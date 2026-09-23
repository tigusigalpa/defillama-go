// Package defillama provides an idiomatic Go client for the DefiLlama Free and
// Pro APIs (https://api-docs.defillama.com/).
//
// All 132 GET operations of the pinned OpenAPI spec (spec/defillama-api.json)
// are exposed through 21 services on the Client. Free endpoints use the
// per-operation origin declared by the spec (api.llama.fi, coins.llama.fi,
// stablecoins.llama.fi, yields.llama.fi). Pro endpoints embed the API key as a
// single URL path segment: https://pro-api.llama.fi/{API_KEY}/<path>.
//
// The client is safe for concurrent use and holds no request-level state.
package defillama

// Version is the package version, used in the default User-Agent.
const Version = "0.1.0"

// Client is the DefiLlama API client. Construct it with New.
type Client struct {
	cfg config
	t   *transport

	tvl                 *TVLService
	prices              *PricesService
	stablecoins         *StablecoinsService
	yields              *YieldsService
	volumes             *VolumesService
	fees                *FeesService
	emissions           *EmissionsService
	ecosystem           *EcosystemService
	bridges             *BridgesService
	etfs                *ETFsService
	narratives          *NarrativesService
	account             *AccountService
	dat                 *DATService
	treasury            *TreasuryService
	oracles             *OraclesService
	forks               *ForksService
	dimensions          *DimensionsService
	financialStatements *FinancialStatementsService
	equities            *EquitiesService
	preIPO              *PreIPOService
	rwa                 *RWAService
}

// New builds a Client. It performs no network I/O; invalid configuration is
// reported as *ConfigError.
func New(opts ...Option) (*Client, error) {
	c := &Client{cfg: defaultConfig()}
	for _, opt := range opts {
		if opt == nil {
			return nil, &ConfigError{Msg: "option must not be nil"}
		}
		if err := opt(&c.cfg); err != nil {
			return nil, err
		}
	}
	if err := c.cfg.validate(); err != nil {
		return nil, err
	}
	c.t = newTransport(c.cfg)

	c.tvl = &TVLService{t: c.t}
	c.prices = &PricesService{t: c.t}
	c.stablecoins = &StablecoinsService{t: c.t}
	c.yields = &YieldsService{t: c.t}
	c.volumes = &VolumesService{t: c.t}
	c.fees = &FeesService{t: c.t}
	c.emissions = &EmissionsService{t: c.t}
	c.ecosystem = &EcosystemService{t: c.t}
	c.bridges = &BridgesService{t: c.t}
	c.etfs = &ETFsService{t: c.t}
	c.narratives = &NarrativesService{t: c.t}
	c.account = &AccountService{t: c.t}
	c.dat = &DATService{t: c.t}
	c.treasury = &TreasuryService{t: c.t}
	c.oracles = &OraclesService{t: c.t}
	c.forks = &ForksService{t: c.t}
	c.dimensions = &DimensionsService{t: c.t}
	c.financialStatements = &FinancialStatementsService{t: c.t}
	c.equities = &EquitiesService{t: c.t}
	c.preIPO = &PreIPOService{t: c.t}
	c.rwa = &RWAService{t: c.t}
	return c, nil
}

// IsPro reports whether an API key is configured, i.e. Pro endpoints are usable.
func (c *Client) IsPro() bool { return c.cfg.apiKey != "" }

// TVL returns the TVL service (Free TVL + Pro TVL charts/metrics).
func (c *Client) TVL() *TVLService { return c.tvl }

// Prices returns the coins/token-liquidity service.
func (c *Client) Prices() *PricesService { return c.prices }

// Stablecoins returns the stablecoins service.
func (c *Client) Stablecoins() *StablecoinsService { return c.stablecoins }

// Yields returns the yields service (pools, earn/borrow v2, LST rates, legacy v1).
func (c *Client) Yields() *YieldsService { return c.yields }

// Volumes returns the volumes service (DEX/options volume, open interest, derivatives).
func (c *Client) Volumes() *VolumesService { return c.volumes }

// Fees returns the fees & revenue service.
func (c *Client) Fees() *FeesService { return c.fees }

// Emissions returns the token unlocks/emissions service.
func (c *Client) Emissions() *EmissionsService { return c.emissions }

// Ecosystem returns the main-page service (categories, entities, forks,
// hacks, oracles, raises, treasuries).
func (c *Client) Ecosystem() *EcosystemService { return c.ecosystem }

// Bridges returns the bridges service.
func (c *Client) Bridges() *BridgesService { return c.bridges }

// ETFs returns the ETFs service.
func (c *Client) ETFs() *ETFsService { return c.etfs }

// Narratives returns the narratives service.
func (c *Client) Narratives() *NarrativesService { return c.narratives }

// Account returns the account/meta service (API key usage).
func (c *Client) Account() *AccountService { return c.account }

// DAT returns the Digital Asset Treasury service.
func (c *Client) DAT() *DATService { return c.dat }

// Treasury returns the protocol treasury service.
func (c *Client) Treasury() *TreasuryService { return c.treasury }

// Oracles returns the oracles service.
func (c *Client) Oracles() *OraclesService { return c.oracles }

// Forks returns the forks service.
func (c *Client) Forks() *ForksService { return c.forks }

// Dimensions returns the dimensions service (19 chart/metric endpoints).
func (c *Client) Dimensions() *DimensionsService { return c.dimensions }

// FinancialStatements returns the financial statements service.
func (c *Client) FinancialStatements() *FinancialStatementsService { return c.financialStatements }

// Equities returns the equities service.
func (c *Client) Equities() *EquitiesService { return c.equities }

// PreIPO returns the pre-IPO companies service.
func (c *Client) PreIPO() *PreIPOService { return c.preIPO }

// RWA returns the real-world assets service.
func (c *Client) RWA() *RWAService { return c.rwa }
