// Code generated from spec/defillama-api.json - DO NOT EDIT.

package defillama

import (
	"fmt"
	"net/url"
	"strconv"
)

// OverviewOptions holds optional/required query parameters. Nil pointer fields are omitted.
type OverviewOptions struct {
	ExcludeTotalDataChart          bool    // query: excludeTotalDataChart
	ExcludeTotalDataChartBreakdown bool    // query: excludeTotalDataChartBreakdown
	DataType                       *string // query: dataType
}

func (o *OverviewOptions) validate() error {
	return nil
}

func (o *OverviewOptions) apply(q url.Values) {
	q.Set("excludeTotalDataChart", boolStr(o.ExcludeTotalDataChart))
	q.Set("excludeTotalDataChartBreakdown", boolStr(o.ExcludeTotalDataChartBreakdown))
	if o.DataType != nil {
		q.Set("dataType", *o.DataType)
	}
}

// ChartOptions holds optional/required query parameters. Nil pointer fields are omitted.
type ChartOptions struct {
	Start  *float64 // query: start
	End    *float64 // query: end
	Span   *float64 // query: span
	Period *string  // query: period
}

func (o *ChartOptions) validate() error {
	return nil
}

func (o *ChartOptions) apply(q url.Values) {
	if o.Start != nil {
		q.Set("start", strconv.FormatFloat(*o.Start, 'f', -1, 64))
	}
	if o.End != nil {
		q.Set("end", strconv.FormatFloat(*o.End, 'f', -1, 64))
	}
	if o.Span != nil {
		q.Set("span", strconv.FormatFloat(*o.Span, 'f', -1, 64))
	}
	if o.Period != nil {
		q.Set("period", *o.Period)
	}
}

// PercentageOptions holds optional/required query parameters. Nil pointer fields are omitted.
type PercentageOptions struct {
	Timestamp   *float64 // query: timestamp
	LookForward *bool    // query: lookForward
	Period      *string  // query: period
}

func (o *PercentageOptions) validate() error {
	return nil
}

func (o *PercentageOptions) apply(q url.Values) {
	if o.Timestamp != nil {
		q.Set("timestamp", strconv.FormatFloat(*o.Timestamp, 'f', -1, 64))
	}
	if o.LookForward != nil {
		q.Set("lookForward", boolStr(*o.LookForward))
	}
	if o.Period != nil {
		q.Set("period", *o.Period)
	}
}

// StablecoinsListOptions holds optional/required query parameters. Nil pointer fields are omitted.
type StablecoinsListOptions struct {
	IncludePrices *bool // query: includePrices
}

func (o *StablecoinsListOptions) validate() error {
	return nil
}

func (o *StablecoinsListOptions) apply(q url.Values) {
	if o.IncludePrices != nil {
		q.Set("includePrices", boolStr(*o.IncludePrices))
	}
}

// StablecoinChartOptions holds optional/required query parameters. Nil pointer fields are omitted.
type StablecoinChartOptions struct {
	Stablecoin *int64 // query: stablecoin
}

func (o *StablecoinChartOptions) validate() error {
	return nil
}

func (o *StablecoinChartOptions) apply(q url.Values) {
	if o.Stablecoin != nil {
		q.Set("stablecoin", strconv.FormatInt(*o.Stablecoin, 10))
	}
}

// DimensionOptions holds optional/required query parameters. Nil pointer fields are omitted.
type DimensionOptions struct {
	DataType *string // query: dataType
}

func (o *DimensionOptions) validate() error {
	return nil
}

func (o *DimensionOptions) apply(q url.Values) {
	if o.DataType != nil {
		q.Set("dataType", *o.DataType)
	}
}

// TVLChartOptions holds optional/required query parameters. Nil pointer fields are omitted.
type TVLChartOptions struct {
	Key      *string // query: key
	Currency *string // query: currency
}

func (o *TVLChartOptions) validate() error {
	if o.Key != nil && !inStrings(*o.Key, []string{"all", "staking", "borrowed", "vesting", "pool2"}) {
		return fmt.Errorf("defillama: invalid key value %q", *o.Key)
	}
	if o.Currency != nil && !inStrings(*o.Currency, []string{"usd", "token", "raw"}) {
		return fmt.Errorf("defillama: invalid currency value %q", *o.Currency)
	}
	return nil
}

func (o *TVLChartOptions) apply(q url.Values) {
	if o.Key != nil {
		q.Set("key", *o.Key)
	}
	if o.Currency != nil {
		q.Set("currency", *o.Currency)
	}
}

// TreasuryChartOptions holds optional/required query parameters. Nil pointer fields are omitted.
type TreasuryChartOptions struct {
	Key      *string // query: key
	Currency *string // query: currency
}

func (o *TreasuryChartOptions) validate() error {
	if o.Key != nil && !inStrings(*o.Key, []string{"OwnTokens", "all"}) {
		return fmt.Errorf("defillama: invalid key value %q", *o.Key)
	}
	if o.Currency != nil && !inStrings(*o.Currency, []string{"usd", "token", "raw"}) {
		return fmt.Errorf("defillama: invalid currency value %q", *o.Currency)
	}
	return nil
}

func (o *TreasuryChartOptions) apply(q url.Values) {
	if o.Key != nil {
		q.Set("key", *o.Key)
	}
	if o.Currency != nil {
		q.Set("currency", *o.Currency)
	}
}

// EarnPoolsQuery holds optional/required query parameters. Nil pointer fields are omitted.
type EarnPoolsQuery struct {
	Chain               *string  // query: chain
	Protocol            *string  // query: protocol
	Stablecoin          *bool    // query: stablecoin
	SingleAssetExposure *bool    // query: single_asset_exposure
	ImpermanentLossRisk *bool    // query: impermanent_loss_risk
	MinTVL              *float64 // query: min_tvl
	Page                *int64   // query: page
	Limit               *int64   // query: limit
}

func (o *EarnPoolsQuery) validate() error {
	return nil
}

func (o *EarnPoolsQuery) apply(q url.Values) {
	if o.Chain != nil {
		q.Set("chain", *o.Chain)
	}
	if o.Protocol != nil {
		q.Set("protocol", *o.Protocol)
	}
	if o.Stablecoin != nil {
		q.Set("stablecoin", boolStr(*o.Stablecoin))
	}
	if o.SingleAssetExposure != nil {
		q.Set("single_asset_exposure", boolStr(*o.SingleAssetExposure))
	}
	if o.ImpermanentLossRisk != nil {
		q.Set("impermanent_loss_risk", boolStr(*o.ImpermanentLossRisk))
	}
	if o.MinTVL != nil {
		q.Set("min_tvl", strconv.FormatFloat(*o.MinTVL, 'f', -1, 64))
	}
	if o.Page != nil {
		q.Set("page", strconv.FormatInt(*o.Page, 10))
	}
	if o.Limit != nil {
		q.Set("limit", strconv.FormatInt(*o.Limit, 10))
	}
}

// HistoryRangeOptions holds optional/required query parameters. Nil pointer fields are omitted.
type HistoryRangeOptions struct {
	Range *string // query: range
}

func (o *HistoryRangeOptions) validate() error {
	if o.Range != nil && !inStrings(*o.Range, []string{"30d", "90d", "max"}) {
		return fmt.Errorf("defillama: invalid range value %q", *o.Range)
	}
	return nil
}

func (o *HistoryRangeOptions) apply(q url.Values) {
	if o.Range != nil {
		q.Set("range", *o.Range)
	}
}

// BridgesOptions holds optional/required query parameters. Nil pointer fields are omitted.
type BridgesOptions struct {
	IncludeChains *bool // query: includeChains
}

func (o *BridgesOptions) validate() error {
	return nil
}

func (o *BridgesOptions) apply(q url.Values) {
	if o.IncludeChains != nil {
		q.Set("includeChains", boolStr(*o.IncludeChains))
	}
}

// BridgeVolumeOptions holds optional/required query parameters. Nil pointer fields are omitted.
type BridgeVolumeOptions struct {
	ID *int64 // query: id
}

func (o *BridgeVolumeOptions) validate() error {
	return nil
}

func (o *BridgeVolumeOptions) apply(q url.Values) {
	if o.ID != nil {
		q.Set("id", strconv.FormatInt(*o.ID, 10))
	}
}

// BridgeTransactionsOptions holds optional/required query parameters. Nil pointer fields are omitted.
type BridgeTransactionsOptions struct {
	Starttimestamp *int64  // query: starttimestamp
	Endtimestamp   *int64  // query: endtimestamp
	Sourcechain    *string // query: sourcechain
	Address        *string // query: address
	Limit          *int64  // query: limit
}

func (o *BridgeTransactionsOptions) validate() error {
	if o.Limit != nil && (*o.Limit < 1 || *o.Limit > 6000) {
		return fmt.Errorf("defillama: limit must be between 1 and 6000")
	}
	return nil
}

func (o *BridgeTransactionsOptions) apply(q url.Values) {
	if o.Starttimestamp != nil {
		q.Set("starttimestamp", strconv.FormatInt(*o.Starttimestamp, 10))
	}
	if o.Endtimestamp != nil {
		q.Set("endtimestamp", strconv.FormatInt(*o.Endtimestamp, 10))
	}
	if o.Sourcechain != nil {
		q.Set("sourcechain", *o.Sourcechain)
	}
	if o.Address != nil {
		q.Set("address", *o.Address)
	}
	if o.Limit != nil {
		q.Set("limit", strconv.FormatInt(*o.Limit, 10))
	}
}

// EquitiesQuery holds optional/required query parameters. Nil pointer fields are omitted.
type EquitiesQuery struct {
	Ticker    string  // query: ticker
	Country   string  // query: country
	Timeframe *string // query: timeframe
}

func (o *EquitiesQuery) validate() error {
	if o.Timeframe != nil && !inStrings(*o.Timeframe, []string{"1D", "7D", "1W", "1M", "3M", "6M", "YTD", "1Y", "5Y", "MAX"}) {
		return fmt.Errorf("defillama: invalid timeframe value %q", *o.Timeframe)
	}
	return nil
}

func (o *EquitiesQuery) apply(q url.Values) {
	q.Set("ticker", o.Ticker)
	q.Set("country", o.Country)
	if o.Timeframe != nil {
		q.Set("timeframe", *o.Timeframe)
	}
}

// PreIPOQuery holds optional/required query parameters. Nil pointer fields are omitted.
type PreIPOQuery struct {
	Company string // query: company
}

func (o *PreIPOQuery) validate() error {
	return nil
}

func (o *PreIPOQuery) apply(q url.Values) {
	q.Set("company", o.Company)
}

// RWAChartOptions holds optional/required query parameters. Nil pointer fields are omitted.
type RWAChartOptions struct {
	Key               *string // query: key
	IncludeStablecoin *bool   // query: includeStablecoin
	IncludeGovernance *bool   // query: includeGovernance
}

func (o *RWAChartOptions) validate() error {
	if o.Key != nil && !inStrings(*o.Key, []string{"onChainMcap", "activeMcap", "defiActiveTvl"}) {
		return fmt.Errorf("defillama: invalid key value %q", *o.Key)
	}
	return nil
}

func (o *RWAChartOptions) apply(q url.Values) {
	if o.Key != nil {
		q.Set("key", *o.Key)
	}
	if o.IncludeStablecoin != nil {
		q.Set("includeStablecoin", boolStr(*o.IncludeStablecoin))
	}
	if o.IncludeGovernance != nil {
		q.Set("includeGovernance", boolStr(*o.IncludeGovernance))
	}
}
