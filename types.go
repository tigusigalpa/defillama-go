package defillama

import "encoding/json"

// Every model below keeps the complete JSON payload in Raw, so unknown or
// newly added fields are never lost. Monetary/percentage fields are float64
// transport values; convert to a decimal type in application code where exact
// precision is required.

// Protocol is an entry of GET /protocols.
type Protocol struct {
	ID        string             `json:"id"`
	Name      string             `json:"name"`
	Symbol    string             `json:"symbol"`
	Category  string             `json:"category"`
	Chains    []string           `json:"chains"`
	TVL       float64            `json:"tvl"`
	ChainTVLs map[string]float64 `json:"chainTvls"`
	Change1d  float64            `json:"change_1d"`
	Change7d  float64            `json:"change_7d"`
	Raw       map[string]any     `json:"-"`
}

// UnmarshalJSON decodes the typed fields and retains the raw payload.
func (p *Protocol) UnmarshalJSON(b []byte) error {
	type alias Protocol
	if err := json.Unmarshal(b, (*alias)(p)); err != nil {
		return err
	}
	return json.Unmarshal(b, &p.Raw)
}

// ProtocolDetails is the response of GET /protocol/{protocol}. The shape
// evolves, so the decoded payload is exposed via Raw with typed accessors
// for the stable fields.
type ProtocolDetails struct {
	Raw map[string]any
}

// Name returns the protocol name, if present.
func (d *ProtocolDetails) Name() string { return str(d.Raw["name"]) }

// Slug returns the protocol slug, if present.
func (d *ProtocolDetails) Slug() string { return str(d.Raw["slug"]) }

// TVLHistory returns the raw historical TVL series.
func (d *ProtocolDetails) TVLHistory() []any {
	if v, ok := d.Raw["tvl"].([]any); ok {
		return v
	}
	return nil
}

// ChainTVLs returns the per-chain TVL breakdown.
func (d *ProtocolDetails) ChainTVLs() map[string]any {
	if v, ok := d.Raw["chainTvls"].(map[string]any); ok {
		return v
	}
	return nil
}

// Chain is an entry of GET /v2/chains.
type Chain struct {
	GeckoID     string         `json:"gecko_id"`
	TVL         float64        `json:"tvl"`
	TokenSymbol string         `json:"tokenSymbol"`
	CmcID       string         `json:"cmcId"`
	Name        string         `json:"name"`
	ChainID     float64        `json:"chainId"`
	Raw         map[string]any `json:"-"`
}

// UnmarshalJSON decodes the typed fields and retains the raw payload.
func (c *Chain) UnmarshalJSON(b []byte) error {
	type alias Chain
	if err := json.Unmarshal(b, (*alias)(c)); err != nil {
		return err
	}
	return json.Unmarshal(b, &c.Raw)
}

// CoinPrice is a single record from the coins API.
type CoinPrice struct {
	Price      float64        `json:"price"`
	Symbol     string         `json:"symbol"`
	Timestamp  int64          `json:"timestamp"`
	Decimals   float64        `json:"decimals"`
	Confidence float64        `json:"confidence"`
	Raw        map[string]any `json:"-"`
}

// UnmarshalJSON decodes the typed fields and retains the raw payload.
func (p *CoinPrice) UnmarshalJSON(b []byte) error {
	type alias CoinPrice
	if err := json.Unmarshal(b, (*alias)(p)); err != nil {
		return err
	}
	return json.Unmarshal(b, &p.Raw)
}

// Block is the response of GET /block/{chain}/{timestamp}.
type Block struct {
	Height    int64          `json:"height"`
	Timestamp int64          `json:"timestamp"`
	Raw       map[string]any `json:"-"`
}

// UnmarshalJSON decodes the typed fields and retains the raw payload.
func (b *Block) UnmarshalJSON(data []byte) error {
	type alias Block
	if err := json.Unmarshal(data, (*alias)(b)); err != nil {
		return err
	}
	return json.Unmarshal(data, &b.Raw)
}

// Stablecoin is an entry of GET /stablecoins (peggedAssets).
type Stablecoin struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Symbol       string         `json:"symbol"`
	PegType      string         `json:"pegType"`
	PegMechanism string         `json:"pegMechanism"`
	Circulating  map[string]any `json:"circulating"`
	Chains       []string       `json:"chains"`
	Price        float64        `json:"price"`
	Raw          map[string]any `json:"-"`
}

// UnmarshalJSON decodes the typed fields and retains the raw payload.
func (s *Stablecoin) UnmarshalJSON(b []byte) error {
	type alias Stablecoin
	if err := json.Unmarshal(b, (*alias)(s)); err != nil {
		return err
	}
	return json.Unmarshal(b, &s.Raw)
}

// YieldPool is an entry of GET yields.llama.fi/pools (data array).
type YieldPool struct {
	Pool         string         `json:"pool"`
	Chain        string         `json:"chain"`
	Project      string         `json:"project"`
	Symbol       string         `json:"symbol"`
	TVLUsd       float64        `json:"tvlUsd"`
	APY          float64        `json:"apy"`
	APYBase      float64        `json:"apyBase"`
	APYReward    *float64       `json:"apyReward"`
	Stablecoin   bool           `json:"stablecoin"`
	IlRisk       string         `json:"ilRisk"`
	Exposure     string         `json:"exposure"`
	RewardTokens []string       `json:"rewardTokens"`
	Predictions  map[string]any `json:"predictions"`
	Raw          map[string]any `json:"-"`
}

// UnmarshalJSON decodes the typed fields and retains the raw payload.
func (p *YieldPool) UnmarshalJSON(b []byte) error {
	type alias YieldPool
	if err := json.Unmarshal(b, (*alias)(p)); err != nil {
		return err
	}
	return json.Unmarshal(b, &p.Raw)
}

// APIUsage is the response of GET /usage/APIKEY — the Pro key credit balance.
// The payload shape is not fully documented, so fields are accessed via Raw
// and the helpers below.
type APIUsage struct {
	Raw map[string]any
}

// CreditsLeft extracts the remaining credit count when present.
func (u *APIUsage) CreditsLeft() (float64, bool) {
	for _, k := range []string{"credits_left", "creditsLeft", "remaining"} {
		if v, ok := u.Raw[k]; ok {
			if f, ok := v.(float64); ok {
				return f, true
			}
		}
	}
	return 0, false
}

// EquityCompany is an entry of GET /equities/v1/companies-list.
type EquityCompany struct {
	Ticker      string         `json:"ticker"`
	CompanyName string         `json:"companyName"`
	Country     string         `json:"country"`
	CountryName string         `json:"countryName"`
	Raw         map[string]any `json:"-"`
}

// UnmarshalJSON decodes the typed fields and retains the raw payload.
func (e *EquityCompany) UnmarshalJSON(b []byte) error {
	type alias EquityCompany
	if err := json.Unmarshal(b, (*alias)(e)); err != nil {
		return err
	}
	return json.Unmarshal(b, &e.Raw)
}

// PreIPOCompany is an entry of GET /pre-ipo/v1/companies-list (data array).
type PreIPOCompany struct {
	ID                       string         `json:"id"`
	Name                     string         `json:"name"`
	Description              string         `json:"description"`
	Website                  string         `json:"website"`
	Sector                   string         `json:"sector"`
	LatestEstimatedValuation *float64       `json:"latestEstimatedValuation"`
	LatestRaise              *float64       `json:"latestRaise"`
	LatestFundingValuation   *float64       `json:"latestFundingValuation"`
	LatestFundingDate        string         `json:"latestFundingDate"`
	TotalRaised              *float64       `json:"totalRaised"`
	Raw                      map[string]any `json:"-"`
}

// UnmarshalJSON decodes the typed fields and retains the raw payload.
func (p *PreIPOCompany) UnmarshalJSON(b []byte) error {
	type alias PreIPOCompany
	if err := json.Unmarshal(b, (*alias)(p)); err != nil {
		return err
	}
	return json.Unmarshal(b, &p.Raw)
}

// RWAAsset is an entry of GET /rwa/current and /rwa/chain/{chain}.
type RWAAsset struct {
	ID          string             `json:"id"`
	Ticker      string             `json:"ticker"`
	AssetName   string             `json:"assetName"`
	AssetGroup  string             `json:"assetGroup"`
	Issuer      string             `json:"issuer"`
	Chain       []string           `json:"chain"`
	Category    []string           `json:"category"`
	OnChainMcap map[string]float64 `json:"onChainMcap"`
	Price       *float64           `json:"price"`
	Stablecoin  *bool              `json:"stablecoin"`
	Governance  *bool              `json:"governance"`
	Raw         map[string]any     `json:"-"`
}

// UnmarshalJSON decodes the typed fields and retains the raw payload.
func (a *RWAAsset) UnmarshalJSON(b []byte) error {
	type alias RWAAsset
	if err := json.Unmarshal(b, (*alias)(a)); err != nil {
		return err
	}
	return json.Unmarshal(b, &a.Raw)
}

// str safely converts a decoded value to string.
func str(v any) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
