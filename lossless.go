package defillama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
)

// LosslessObject stores an open JSON object without converting number lexemes
// to float64. It intentionally exposes accessors rather than a mutable map so
// unknown provider fields remain inspectable without sharing SDK-owned bytes.
type LosslessObject struct {
	fields map[string]json.RawMessage
}

// UnmarshalJSON implements json.Unmarshaler while preserving every member's
// original JSON representation.
func (o *LosslessObject) UnmarshalJSON(data []byte) error {
	if len(bytes.TrimSpace(data)) == 0 || bytes.TrimSpace(data)[0] != '{' {
		return fmt.Errorf("defillama: expected a JSON object")
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	o.fields = make(map[string]json.RawMessage, len(fields))
	for name, value := range fields {
		o.fields[name] = append(json.RawMessage(nil), value...)
	}
	return nil
}

// Has reports whether a member was present, including when its value was null.
func (o LosslessObject) Has(name string) bool {
	_, ok := o.fields[name]
	return ok
}

// Field returns a lossless member value and whether it was present. A present
// null is returned with ok=true and Value.IsNull()==true.
func (o LosslessObject) Field(name string) (LosslessValue, bool) {
	raw, ok := o.fields[name]
	if !ok {
		return LosslessValue{}, false
	}
	return LosslessValue{raw: append(json.RawMessage(nil), raw...)}, true
}

// Keys returns object member names in stable lexical order.
func (o LosslessObject) Keys() []string {
	keys := make([]string, 0, len(o.fields))
	for name := range o.fields {
		keys = append(keys, name)
	}
	sort.Strings(keys)
	return keys
}

// LosslessValue preserves one JSON value exactly and exposes safe, typed
// convenience accessors. It distinguishes an absent object member (Field's
// ok=false), a present null (IsNull), and numeric zero (Number returns "0").
type LosslessValue struct {
	raw json.RawMessage
}

// RawJSON returns a copy of the exact JSON value bytes.
func (v LosslessValue) RawJSON() []byte {
	return append([]byte(nil), v.raw...)
}

// IsNull reports whether the preserved value is the JSON literal null.
func (v LosslessValue) IsNull() bool {
	return bytes.Equal(bytes.TrimSpace(v.raw), []byte("null"))
}

// Number returns the original numeric lexeme without a float64 conversion.
func (v LosslessValue) Number() (json.Number, error) {
	dec := json.NewDecoder(bytes.NewReader(v.raw))
	dec.UseNumber()
	var value any
	if err := dec.Decode(&value); err != nil {
		return "", err
	}
	number, ok := value.(json.Number)
	if !ok {
		return "", fmt.Errorf("defillama: expected JSON number")
	}
	return number, nil
}

// String returns a JSON string value.
func (v LosslessValue) String() (string, error) {
	var value string
	if err := json.Unmarshal(v.raw, &value); err != nil {
		return "", fmt.Errorf("defillama: expected JSON string: %w", err)
	}
	return value, nil
}

// Object returns a nested open JSON object.
func (v LosslessValue) Object() (LosslessObject, error) {
	var object LosslessObject
	if err := json.Unmarshal(v.raw, &object); err != nil {
		return LosslessObject{}, err
	}
	return object, nil
}

// Array returns values from a JSON array while retaining each element exactly.
func (v LosslessValue) Array() ([]LosslessValue, error) {
	var rawValues []json.RawMessage
	if err := json.Unmarshal(v.raw, &rawValues); err != nil {
		return nil, fmt.Errorf("defillama: expected JSON array: %w", err)
	}
	values := make([]LosslessValue, len(rawValues))
	for i, raw := range rawValues {
		values[i] = LosslessValue{raw: append(json.RawMessage(nil), raw...)}
	}
	return values, nil
}

// LosslessProtocols is a provider-native response from GET /protocols. It
// exposes only selected convenience fields; all provider additions remain in
// each entry's Fields object.
type LosslessProtocols struct {
	entries []LosslessProtocol
}

// UnmarshalJSON implements json.Unmarshaler for the protocol array response.
func (p *LosslessProtocols) UnmarshalJSON(data []byte) error {
	values, err := decodeLosslessArray(data)
	if err != nil {
		return err
	}
	p.entries = make([]LosslessProtocol, len(values))
	for i, value := range values {
		object, err := value.Object()
		if err != nil {
			return fmt.Errorf("defillama: protocol %d: %w", i, err)
		}
		p.entries[i] = LosslessProtocol{fields: object}
	}
	return nil
}

// Entries returns the provider-native protocol records.
func (p LosslessProtocols) Entries() []LosslessProtocol {
	return append([]LosslessProtocol(nil), p.entries...)
}

// LosslessProtocol is one provider-native protocol record. Chain and token
// maps remain open through Fields rather than being constrained to an enum.
type LosslessProtocol struct {
	fields LosslessObject
}

// Fields returns the complete open object for this protocol record.
func (p LosslessProtocol) Fields() LosslessObject { return p.fields }

// ID returns the provider-native protocol identifier when present.
func (p LosslessProtocol) ID() (string, bool, error) { return stringField(p.fields, "id") }

// Name returns the provider-native protocol name when present.
func (p LosslessProtocol) Name() (string, bool, error) { return stringField(p.fields, "name") }

// TVL returns the exact provider TVL number lexeme when present and non-null.
func (p LosslessProtocol) TVL() (json.Number, bool, error) { return numberField(p.fields, "tvl") }

// LosslessProtocolDetails is a provider-native response from GET
// /protocol/{protocol}. It intentionally models only selected stable fields;
// Fields retains the full response, including open chain/token maps.
type LosslessProtocolDetails struct {
	fields LosslessObject
}

// UnmarshalJSON implements json.Unmarshaler.
func (d *LosslessProtocolDetails) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &d.fields)
}

// Fields returns the complete open provider object.
func (d LosslessProtocolDetails) Fields() LosslessObject { return d.fields }

// ID returns the provider-native protocol identifier when present.
func (d LosslessProtocolDetails) ID() (string, bool, error) { return stringField(d.fields, "id") }

// Name returns the provider-native protocol name when present.
func (d LosslessProtocolDetails) Name() (string, bool, error) { return stringField(d.fields, "name") }

// TVLHistory returns provider history records exactly as open objects. Provider
// date values are provider timestamps (Unix seconds), not capture times.
func (d LosslessProtocolDetails) TVLHistory() ([]LosslessObject, bool, error) {
	return objectArrayField(d.fields, "tvl")
}

// ChainTVLs returns the open per-chain TVL map when present.
func (d LosslessProtocolDetails) ChainTVLs() (LosslessObject, bool, error) {
	return objectField(d.fields, "chainTvls")
}

// LosslessYieldPools is a provider-native response from GET /pools.
type LosslessYieldPools struct {
	fields LosslessObject
}

// UnmarshalJSON implements json.Unmarshaler.
func (p *LosslessYieldPools) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &p.fields)
}

// Fields returns the complete open response object.
func (p LosslessYieldPools) Fields() LosslessObject { return p.fields }

// Pools returns provider-native pool records as open objects.
func (p LosslessYieldPools) Pools() ([]LosslessYieldPool, bool, error) {
	objects, present, err := objectArrayField(p.fields, "data")
	if err != nil || !present {
		return nil, present, err
	}
	pools := make([]LosslessYieldPool, len(objects))
	for i, object := range objects {
		pools[i] = LosslessYieldPool{fields: object}
	}
	return pools, true, nil
}

// LosslessYieldPool is a provider-native yield pool. APY values are provider
// percentages, not ratios. Predictions and classification fields remain open
// in Fields and are not interpreted as observed APY.
type LosslessYieldPool struct {
	fields LosslessObject
}

// Fields returns the complete open provider object.
func (p LosslessYieldPool) Fields() LosslessObject { return p.fields }

// PoolID returns the provider-native pool identifier, not a universal asset ID.
func (p LosslessYieldPool) PoolID() (string, bool, error) { return stringField(p.fields, "pool") }

// APY returns the exact APY percentage lexeme when present and non-null.
func (p LosslessYieldPool) APY() (json.Number, bool, error) { return numberField(p.fields, "apy") }

// APYBase returns the exact base APY percentage lexeme when present and non-null.
func (p LosslessYieldPool) APYBase() (json.Number, bool, error) {
	return numberField(p.fields, "apyBase")
}

// APYReward returns the raw APY reward member. Use the returned present value
// plus IsNull to distinguish absent, null, and numeric zero.
func (p LosslessYieldPool) APYReward() (LosslessValue, bool) {
	return p.fields.Field("apyReward")
}

// RewardTokens returns the raw rewardTokens member. It preserves the
// distinction between absent, null, and an empty array.
func (p LosslessYieldPool) RewardTokens() (LosslessValue, bool) {
	return p.fields.Field("rewardTokens")
}

// LosslessYieldChart is a provider-native response from GET /chart/{pool}.
// It contains only selected convenience accessors; all extra response and
// point fields remain available in open Fields objects.
type LosslessYieldChart struct {
	fields LosslessObject
}

// UnmarshalJSON implements json.Unmarshaler.
func (c *LosslessYieldChart) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &c.fields)
}

// Fields returns the complete open chart response object.
func (c LosslessYieldChart) Fields() LosslessObject { return c.fields }

// Points returns chart records as open provider objects.
func (c LosslessYieldChart) Points() ([]LosslessYieldChartPoint, bool, error) {
	objects, present, err := objectArrayField(c.fields, "data")
	if err != nil || !present {
		return nil, present, err
	}
	points := make([]LosslessYieldChartPoint, len(objects))
	for i, object := range objects {
		points[i] = LosslessYieldChartPoint{fields: object}
	}
	return points, true, nil
}

// LosslessYieldChartPoint is one provider-native chart record. Timestamp is
// an ISO-8601 provider timestamp, distinct from ResponseReceipt.CapturedAt.
type LosslessYieldChartPoint struct {
	fields LosslessObject
}

// Fields returns the complete open provider object.
func (p LosslessYieldChartPoint) Fields() LosslessObject { return p.fields }

// Timestamp returns the provider's ISO-8601 chart timestamp when present.
func (p LosslessYieldChartPoint) Timestamp() (string, bool, error) {
	return stringField(p.fields, "timestamp")
}

// TVLUSD returns the exact provider TVL number lexeme when present and non-null.
func (p LosslessYieldChartPoint) TVLUSD() (json.Number, bool, error) {
	return numberField(p.fields, "tvlUsd")
}

// APY returns the exact APY percentage lexeme when present and non-null.
func (p LosslessYieldChartPoint) APY() (json.Number, bool, error) {
	return numberField(p.fields, "apy")
}

// APYBase returns the exact base APY percentage lexeme when present and non-null.
func (p LosslessYieldChartPoint) APYBase() (json.Number, bool, error) {
	return numberField(p.fields, "apyBase")
}

// APYReward returns the raw reward APY member, preserving absent/null/zero.
func (p LosslessYieldChartPoint) APYReward() (LosslessValue, bool) {
	return p.fields.Field("apyReward")
}

func decodeLosslessArray(data []byte) ([]LosslessValue, error) {
	var rawValues []json.RawMessage
	if err := json.Unmarshal(data, &rawValues); err != nil {
		return nil, fmt.Errorf("defillama: expected JSON array: %w", err)
	}
	values := make([]LosslessValue, len(rawValues))
	for i, raw := range rawValues {
		values[i] = LosslessValue{raw: append(json.RawMessage(nil), raw...)}
	}
	return values, nil
}

func stringField(object LosslessObject, name string) (string, bool, error) {
	value, present := object.Field(name)
	if !present || value.IsNull() {
		return "", false, nil
	}
	result, err := value.String()
	return result, true, err
}

func numberField(object LosslessObject, name string) (json.Number, bool, error) {
	value, present := object.Field(name)
	if !present || value.IsNull() {
		return "", false, nil
	}
	result, err := value.Number()
	return result, true, err
}

func objectField(object LosslessObject, name string) (LosslessObject, bool, error) {
	value, present := object.Field(name)
	if !present || value.IsNull() {
		return LosslessObject{}, false, nil
	}
	result, err := value.Object()
	return result, true, err
}

func objectArrayField(object LosslessObject, name string) ([]LosslessObject, bool, error) {
	value, present := object.Field(name)
	if !present || value.IsNull() {
		return nil, false, nil
	}
	values, err := value.Array()
	if err != nil {
		return nil, true, err
	}
	objects := make([]LosslessObject, len(values))
	for i, value := range values {
		entry, err := value.Object()
		if err != nil {
			return nil, true, fmt.Errorf("defillama: %s[%d]: %w", name, i, err)
		}
		objects[i] = entry
	}
	return objects, true, nil
}
