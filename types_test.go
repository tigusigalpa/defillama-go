package defillama

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestTypedModelsRetainUnknownFields(t *testing.T) {
	tests := []struct {
		name    string
		payload string
		target  any
		field   string
		want    any
	}{
		{"Protocol", `{"name":"Aave","tvl":12.5,"extra":7}`, &Protocol{}, "Name", "Aave"},
		{"Chain", `{"name":"Ethereum","tvl":11,"extra":7}`, &Chain{}, "Name", "Ethereum"},
		{"CoinPrice", `{"price":42.5,"symbol":"ETH","extra":7}`, &CoinPrice{}, "Price", 42.5},
		{"Block", `{"height":42,"timestamp":1700,"extra":7}`, &Block{}, "Height", int64(42)},
		{"Stablecoin", `{"name":"USDC","price":1,"extra":7}`, &Stablecoin{}, "Name", "USDC"},
		{"YieldPool", `{"pool":"pool-1","apy":3.5,"extra":7}`, &YieldPool{}, "Pool", "pool-1"},
		{"EquityCompany", `{"ticker":"NVDA","companyName":"Nvidia","extra":7}`, &EquityCompany{}, "Ticker", "NVDA"},
		{"PreIPOCompany", `{"name":"SpaceX","extra":7}`, &PreIPOCompany{}, "Name", "SpaceX"},
		{"RWAAsset", `{"assetName":"Gold","extra":7}`, &RWAAsset{}, "AssetName", "Gold"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := json.Unmarshal([]byte(tt.payload), tt.target); err != nil {
				t.Fatal(err)
			}
			model := reflect.ValueOf(tt.target).Elem()
			if got := model.FieldByName(tt.field).Interface(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s = %v, want %v", tt.field, got, tt.want)
			}
			raw := model.FieldByName("Raw").Interface().(map[string]any)
			if got := raw["extra"]; got != float64(7) {
				t.Errorf("Raw[extra] = %v, want 7", got)
			}
			fresh := reflect.New(model.Type()).Interface()
			if err := json.Unmarshal([]byte(`{"broken":`), fresh); err == nil {
				t.Error("malformed JSON was accepted")
			}
		})
	}
}

func TestProtocolDetailsAccessors(t *testing.T) {
	d := &ProtocolDetails{Raw: map[string]any{
		"name":      "Aave",
		"slug":      "aave",
		"tvl":       []any{float64(10)},
		"chainTvls": map[string]any{"Ethereum": float64(10)},
	}}
	if d.Name() != "Aave" || d.Slug() != "aave" {
		t.Errorf("name/slug = %q/%q", d.Name(), d.Slug())
	}
	if got := d.TVLHistory(); len(got) != 1 || got[0] != float64(10) {
		t.Errorf("TVLHistory = %v", got)
	}
	if got := d.ChainTVLs(); got["Ethereum"] != float64(10) {
		t.Errorf("ChainTVLs = %v", got)
	}

	d.Raw = map[string]any{"name": 123, "slug": false, "tvl": "wrong", "chainTvls": []any{}}
	if d.Name() != "" || d.Slug() != "" || d.TVLHistory() != nil || d.ChainTVLs() != nil {
		t.Error("accessors did not safely handle unexpected field types")
	}
}

func TestAPIUsageCreditsLeft(t *testing.T) {
	for _, key := range []string{"credits_left", "creditsLeft", "remaining"} {
		u := &APIUsage{Raw: map[string]any{key: float64(42)}}
		if got, ok := u.CreditsLeft(); !ok || got != 42 {
			t.Errorf("%s: got (%v, %v), want (42, true)", key, got, ok)
		}
	}
	for _, raw := range []map[string]any{{"remaining": "42"}, {}} {
		u := &APIUsage{Raw: raw}
		if got, ok := u.CreditsLeft(); ok || got != 0 {
			t.Errorf("%v: got (%v, %v), want (0, false)", raw, got, ok)
		}
	}
}

func TestPtrPreservesValue(t *testing.T) {
	if got := Ptr(42); got == nil || *got != 42 {
		t.Errorf("Ptr(42) = %v", got)
	}
}
