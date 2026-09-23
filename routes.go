// Code generated from spec/defillama-api.json - DO NOT EDIT.

package defillama

var routeRegistry = map[string]route{
	"GET /protocols": {
		Method: "GET", Path: "/protocols", Tier: "free", Server: "https://api.llama.fi",
		ProPath: "/api/protocols",
		Service: "TVL", GoMethod: "GetProtocols", DocsURL: "https://api-docs.defillama.com/#tag/tvl/get/protocols",
		Summary: "List all protocols on defillama along with their tvl",
	},
	"GET /protocol/{protocol}": {
		Method: "GET", Path: "/protocol/{protocol}", Tier: "free", Server: "https://api.llama.fi",
		ProPath: "/api/protocol/{protocol}",
		Service: "TVL", GoMethod: "GetProtocol", DocsURL: "https://api-docs.defillama.com/#tag/tvl/get/protocol/%7Bprotocol%7D",
		Summary: "Get historical TVL of a protocol and breakdowns by token and chain",
		Params: []routeParam{
			{Name: "protocol", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /v2/historicalChainTvl": {
		Method: "GET", Path: "/v2/historicalChainTvl", Tier: "free", Server: "https://api.llama.fi",
		ProPath: "/api/v2/historicalChainTvl",
		Service: "TVL", GoMethod: "GetHistoricalChainTVL", DocsURL: "https://api-docs.defillama.com/#tag/tvl/get/v2/historicalChainTvl",
		Summary: "Get historical TVL (excludes liquid staking and double counted tvl) of DeFi on all chains",
	},
	"GET /v2/historicalChainTvl/{chain}": {
		Method: "GET", Path: "/v2/historicalChainTvl/{chain}", Tier: "free", Server: "https://api.llama.fi",
		ProPath: "/api/v2/historicalChainTvl/{chain}",
		Service: "TVL", GoMethod: "GetHistoricalChainTVLByChain", DocsURL: "https://api-docs.defillama.com/#tag/tvl/get/v2/historicalChainTvl/%7Bchain%7D",
		Summary: "Get historical TVL (excludes liquid staking and double counted tvl) of a chain",
		Params: []routeParam{
			{Name: "chain", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /tvl/{protocol}": {
		Method: "GET", Path: "/tvl/{protocol}", Tier: "free", Server: "https://api.llama.fi",
		ProPath: "/api/tvl/{protocol}",
		Service: "TVL", GoMethod: "GetTVL", DocsURL: "https://api-docs.defillama.com/#tag/tvl/get/tvl/%7Bprotocol%7D",
		Summary: "Simplified endpoint to get current TVL of a protocol",
		Params: []routeParam{
			{Name: "protocol", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /v2/chains": {
		Method: "GET", Path: "/v2/chains", Tier: "free", Server: "https://api.llama.fi",
		ProPath: "/api/v2/chains",
		Service: "TVL", GoMethod: "GetChains", DocsURL: "https://api-docs.defillama.com/#tag/tvl/get/v2/chains",
		Summary: "Get current TVL of all chains",
	},
	"GET /prices/current/{coins}": {
		Method: "GET", Path: "/prices/current/{coins}", Tier: "free", Server: "https://coins.llama.fi",
		ProPath: "/coins/prices/current/{coins}",
		Service: "Prices", GoMethod: "GetCurrentPrices", DocsURL: "https://api-docs.defillama.com/#tag/coins/get/prices/current/%7Bcoins%7D",
		Summary: "Get current prices of tokens by contract address",
		Params: []routeParam{
			{Name: "coins", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /prices/historical/{timestamp}/{coins}": {
		Method: "GET", Path: "/prices/historical/{timestamp}/{coins}", Tier: "free", Server: "https://coins.llama.fi",
		ProPath: "/coins/prices/historical/{timestamp}/{coins}",
		Service: "Prices", GoMethod: "GetHistoricalPrices", DocsURL: "https://api-docs.defillama.com/#tag/coins/get/prices/historical/%7Btimestamp%7D/%7Bcoins%7D",
		Summary: "Get historical prices of tokens by contract address",
		Params: []routeParam{
			{Name: "coins", In: "path", Required: true, Type: "string"},
			{Name: "timestamp", In: "path", Required: true, Type: "number"},
		},
	},
	"GET /batchHistorical": {
		Method: "GET", Path: "/batchHistorical", Tier: "free", Server: "https://coins.llama.fi",
		ProPath: "/coins/batchHistorical",
		Service: "Prices", GoMethod: "GetBatchHistoricalPrices", DocsURL: "https://api-docs.defillama.com/#tag/coins/get/batchHistorical",
		Summary: "Get historical prices for multiple tokens at multiple different timestamps",
		Params: []routeParam{
			{Name: "coins", In: "query", Required: true, Type: "string"},
		},
	},
	"GET /chart/{coins}": {
		Method: "GET", Path: "/chart/{coins}", Tier: "free", Server: "https://coins.llama.fi",
		ProPath: "/coins/chart/{coins}",
		Service: "Prices", GoMethod: "GetChart", DocsURL: "https://api-docs.defillama.com/#tag/coins/get/chart/%7Bcoins%7D",
		Summary: "Get token prices at regular time intervals",
		Params: []routeParam{
			{Name: "coins", In: "path", Required: true, Type: "string"},
			{Name: "start", In: "query", Required: false, Type: "number"},
			{Name: "end", In: "query", Required: false, Type: "number"},
			{Name: "span", In: "query", Required: false, Type: "number"},
			{Name: "period", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /percentage/{coins}": {
		Method: "GET", Path: "/percentage/{coins}", Tier: "free", Server: "https://coins.llama.fi",
		ProPath: "/coins/percentage/{coins}",
		Service: "Prices", GoMethod: "GetPercentageChange", DocsURL: "https://api-docs.defillama.com/#tag/coins/get/percentage/%7Bcoins%7D",
		Summary: "Get percentage change in price over time",
		Params: []routeParam{
			{Name: "coins", In: "path", Required: true, Type: "string"},
			{Name: "timestamp", In: "query", Required: false, Type: "number"},
			{Name: "lookForward", In: "query", Required: false, Type: "boolean"},
			{Name: "period", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /prices/first/{coins}": {
		Method: "GET", Path: "/prices/first/{coins}", Tier: "free", Server: "https://coins.llama.fi",
		ProPath: "/coins/prices/first/{coins}",
		Service: "Prices", GoMethod: "GetFirstPrices", DocsURL: "https://api-docs.defillama.com/#tag/coins/get/prices/first/%7Bcoins%7D",
		Summary: "Get earliest timestamp price record for coins",
		Params: []routeParam{
			{Name: "coins", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /block/{chain}/{timestamp}": {
		Method: "GET", Path: "/block/{chain}/{timestamp}", Tier: "free", Server: "https://coins.llama.fi",
		ProPath: "/coins/block/{chain}/{timestamp}",
		Service: "Prices", GoMethod: "GetBlockAtTimestamp", DocsURL: "https://api-docs.defillama.com/#tag/coins/get/block/%7Bchain%7D/%7Btimestamp%7D",
		Summary: "Get the closest block to a timestamp",
		Params: []routeParam{
			{Name: "chain", In: "path", Required: true, Type: "string"},
			{Name: "timestamp", In: "path", Required: true, Type: "integer"},
		},
	},
	"GET /stablecoins": {
		Method: "GET", Path: "/stablecoins", Tier: "free", Server: "https://stablecoins.llama.fi",
		ProPath: "/stablecoins/stablecoins",
		Service: "Stablecoins", GoMethod: "GetStablecoins", DocsURL: "https://api-docs.defillama.com/#tag/stablecoins/get/stablecoins",
		Summary: "List all stablecoins along with their circulating amounts",
		Params: []routeParam{
			{Name: "includePrices", In: "query", Required: false, Type: "boolean"},
		},
	},
	"GET /stablecoincharts/all": {
		Method: "GET", Path: "/stablecoincharts/all", Tier: "free", Server: "https://stablecoins.llama.fi",
		ProPath: "/stablecoins/stablecoincharts/all",
		Service: "Stablecoins", GoMethod: "GetStablecoinCharts", DocsURL: "https://api-docs.defillama.com/#tag/stablecoins/get/stablecoincharts/all",
		Summary: "Get historical mcap sum of all stablecoins",
		Params: []routeParam{
			{Name: "stablecoin", In: "query", Required: false, Type: "integer"},
		},
	},
	"GET /stablecoincharts/{chain}": {
		Method: "GET", Path: "/stablecoincharts/{chain}", Tier: "free", Server: "https://stablecoins.llama.fi",
		ProPath: "/stablecoins/stablecoincharts/{chain}",
		Service: "Stablecoins", GoMethod: "GetStablecoinChartsByChain", DocsURL: "https://api-docs.defillama.com/#tag/stablecoins/get/stablecoincharts/%7Bchain%7D",
		Summary: "Get historical mcap sum of all stablecoins in a chain",
		Params: []routeParam{
			{Name: "chain", In: "path", Required: true, Type: "string"},
			{Name: "stablecoin", In: "query", Required: false, Type: "integer"},
		},
	},
	"GET /stablecoin/{asset}": {
		Method: "GET", Path: "/stablecoin/{asset}", Tier: "free", Server: "https://stablecoins.llama.fi",
		ProPath: "/stablecoins/stablecoin/{asset}",
		Service: "Stablecoins", GoMethod: "GetStablecoin", DocsURL: "https://api-docs.defillama.com/#tag/stablecoins/get/stablecoin/%7Basset%7D",
		Summary: "Get historical mcap and historical chain distribution of a stablecoin",
		Params: []routeParam{
			{Name: "asset", In: "path", Required: true, Type: "integer"},
		},
	},
	"GET /stablecoinchains": {
		Method: "GET", Path: "/stablecoinchains", Tier: "free", Server: "https://stablecoins.llama.fi",
		ProPath: "/stablecoins/stablecoinchains",
		Service: "Stablecoins", GoMethod: "GetStablecoinChains", DocsURL: "https://api-docs.defillama.com/#tag/stablecoins/get/stablecoinchains",
		Summary: "Get current mcap sum of all stablecoins on each chain",
	},
	"GET /stablecoinprices": {
		Method: "GET", Path: "/stablecoinprices", Tier: "free", Server: "https://stablecoins.llama.fi",
		ProPath: "/stablecoins/stablecoinprices",
		Service: "Stablecoins", GoMethod: "GetStablecoinPrices", DocsURL: "https://api-docs.defillama.com/#tag/stablecoins/get/stablecoinprices",
		Summary: "Get historical prices of all stablecoins",
	},
	"GET /pools": {
		Method: "GET", Path: "/pools", Tier: "free", Server: "https://yields.llama.fi",
		ProPath: "/yields/pools",
		Service: "Yields", GoMethod: "GetPools", DocsURL: "https://api-docs.defillama.com/#tag/yields/get/pools",
		Summary: "Retrieve the latest data for all pools, including enriched information such as predictions",
	},
	"GET /chart/{pool}": {
		Method: "GET", Path: "/chart/{pool}", Tier: "free", Server: "https://yields.llama.fi",
		ProPath: "/yields/chart/{pool}",
		Service: "Yields", GoMethod: "GetPoolChart", DocsURL: "https://api-docs.defillama.com/#tag/yields/get/chart/%7Bpool%7D",
		Summary: "Get historical APY and TVL of a pool",
		Params: []routeParam{
			{Name: "pool", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /overview/dexs": {
		Method: "GET", Path: "/overview/dexs", Tier: "free", Server: "https://api.llama.fi",
		ProPath: "/api/overview/dexs",
		Service: "Volumes", GoMethod: "GetDEXOverview", DocsURL: "https://api-docs.defillama.com/#tag/volumes/get/overview/dexs",
		Summary: "List all dexs along with summaries of their volumes and dataType history data",
		Params: []routeParam{
			{Name: "excludeTotalDataChart", In: "query", Required: true, Type: "boolean"},
			{Name: "excludeTotalDataChartBreakdown", In: "query", Required: true, Type: "boolean"},
		},
	},
	"GET /overview/dexs/{chain}": {
		Method: "GET", Path: "/overview/dexs/{chain}", Tier: "free", Server: "https://api.llama.fi",
		ProPath: "/api/overview/dexs/{chain}",
		Service: "Volumes", GoMethod: "GetDEXOverviewByChain", DocsURL: "https://api-docs.defillama.com/#tag/volumes/get/overview/dexs/%7Bchain%7D",
		Summary: "List all dexs along with summaries of their volumes and dataType history data filtering by chain",
		Params: []routeParam{
			{Name: "chain", In: "path", Required: true, Type: "string"},
			{Name: "excludeTotalDataChart", In: "query", Required: true, Type: "boolean"},
			{Name: "excludeTotalDataChartBreakdown", In: "query", Required: true, Type: "boolean"},
		},
	},
	"GET /summary/dexs/{protocol}": {
		Method: "GET", Path: "/summary/dexs/{protocol}", Tier: "free", Server: "https://api.llama.fi",
		ProPath: "/api/summary/dexs/{protocol}",
		Service: "Volumes", GoMethod: "GetDEXSummary", DocsURL: "https://api-docs.defillama.com/#tag/volumes/get/summary/dexs/%7Bprotocol%7D",
		Summary: "Get summary of dex volume with historical data",
		Params: []routeParam{
			{Name: "protocol", In: "path", Required: true, Type: "string"},
			{Name: "excludeTotalDataChart", In: "query", Required: true, Type: "boolean"},
			{Name: "excludeTotalDataChartBreakdown", In: "query", Required: true, Type: "boolean"},
		},
	},
	"GET /overview/options": {
		Method: "GET", Path: "/overview/options", Tier: "free", Server: "https://api.llama.fi",
		ProPath: "/api/overview/options",
		Service: "Volumes", GoMethod: "GetOptionsOverview", DocsURL: "https://api-docs.defillama.com/#tag/volumes/get/overview/options",
		Summary: "List all options dexs along with summaries of their volumes and dataType history data",
		Params: []routeParam{
			{Name: "excludeTotalDataChart", In: "query", Required: true, Type: "boolean"},
			{Name: "excludeTotalDataChartBreakdown", In: "query", Required: true, Type: "boolean"},
			{Name: "dataType", In: "query", Required: false, Type: "string", Enum: []string{"dailyPremiumVolume", "dailyNotionalVolume"}},
		},
	},
	"GET /overview/options/{chain}": {
		Method: "GET", Path: "/overview/options/{chain}", Tier: "free", Server: "https://api.llama.fi",
		ProPath: "/api/overview/options/{chain}",
		Service: "Volumes", GoMethod: "GetOptionsOverviewByChain", DocsURL: "https://api-docs.defillama.com/#tag/volumes/get/overview/options/%7Bchain%7D",
		Summary: "List all options dexs along with summaries of their volumes and dataType history data filtering by chain",
		Params: []routeParam{
			{Name: "chain", In: "path", Required: true, Type: "string"},
			{Name: "excludeTotalDataChart", In: "query", Required: true, Type: "boolean"},
			{Name: "excludeTotalDataChartBreakdown", In: "query", Required: true, Type: "boolean"},
			{Name: "dataType", In: "query", Required: false, Type: "string", Enum: []string{"dailyPremiumVolume", "dailyNotionalVolume"}},
		},
	},
	"GET /summary/options/{protocol}": {
		Method: "GET", Path: "/summary/options/{protocol}", Tier: "free", Server: "https://api.llama.fi",
		ProPath: "/api/summary/options/{protocol}",
		Service: "Volumes", GoMethod: "GetOptionsSummary", DocsURL: "https://api-docs.defillama.com/#tag/volumes/get/summary/options/%7Bprotocol%7D",
		Summary: "Get summary of options dex volume with historical data",
		Params: []routeParam{
			{Name: "protocol", In: "path", Required: true, Type: "string"},
			{Name: "dataType", In: "query", Required: false, Type: "string", Enum: []string{"dailyPremiumVolume", "dailyNotionalVolume"}},
		},
	},
	"GET /overview/open-interest": {
		Method: "GET", Path: "/overview/open-interest", Tier: "free", Server: "https://api.llama.fi",
		ProPath: "/api/overview/open-interest",
		Service: "Volumes", GoMethod: "GetOpenInterestOverview", DocsURL: "https://api-docs.defillama.com/#tag/perps/get/overview/open-interest",
		Summary: "List all open interest dex exchanges along with summaries of their open interest",
		Params: []routeParam{
			{Name: "excludeTotalDataChart", In: "query", Required: true, Type: "boolean"},
			{Name: "excludeTotalDataChartBreakdown", In: "query", Required: true, Type: "boolean"},
		},
	},
	"GET /overview/fees": {
		Method: "GET", Path: "/overview/fees", Tier: "free", Server: "https://api.llama.fi",
		ProPath: "/api/overview/fees",
		Service: "Fees", GoMethod: "GetOverview", DocsURL: "https://api-docs.defillama.com/#tag/fees-and-revenue/get/overview/fees",
		Summary: "List all protocols along with summaries of their fees and revenue and dataType history data",
		Params: []routeParam{
			{Name: "excludeTotalDataChart", In: "query", Required: true, Type: "boolean"},
			{Name: "excludeTotalDataChartBreakdown", In: "query", Required: true, Type: "boolean"},
			{Name: "dataType", In: "query", Required: false, Type: "string", Enum: []string{"dailyFees", "dailyRevenue", "dailyHoldersRevenue"}},
		},
	},
	"GET /overview/fees/{chain}": {
		Method: "GET", Path: "/overview/fees/{chain}", Tier: "free", Server: "https://api.llama.fi",
		ProPath: "/api/overview/fees/{chain}",
		Service: "Fees", GoMethod: "GetOverviewByChain", DocsURL: "https://api-docs.defillama.com/#tag/fees-and-revenue/get/overview/fees/%7Bchain%7D",
		Summary: "List all protocols along with summaries of their fees and revenue and dataType history data by chain",
		Params: []routeParam{
			{Name: "chain", In: "path", Required: true, Type: "string"},
			{Name: "excludeTotalDataChart", In: "query", Required: true, Type: "boolean"},
			{Name: "excludeTotalDataChartBreakdown", In: "query", Required: true, Type: "boolean"},
			{Name: "dataType", In: "query", Required: false, Type: "string", Enum: []string{"dailyFees", "dailyRevenue", "dailyHoldersRevenue"}},
		},
	},
	"GET /summary/fees/{protocol}": {
		Method: "GET", Path: "/summary/fees/{protocol}", Tier: "free", Server: "https://api.llama.fi",
		ProPath: "/api/summary/fees/{protocol}",
		Service: "Fees", GoMethod: "GetSummary", DocsURL: "https://api-docs.defillama.com/#tag/fees-and-revenue/get/summary/fees/%7Bprotocol%7D",
		Summary: "Get summary of protocol fees and revenue with historical data",
		Params: []routeParam{
			{Name: "protocol", In: "path", Required: true, Type: "string"},
			{Name: "dataType", In: "query", Required: false, Type: "string", Enum: []string{"dailyFees", "dailyRevenue", "dailyHoldersRevenue"}},
		},
	},
	"GET /api/tokenProtocols/{symbol}": {
		Method: "GET", Path: "/api/tokenProtocols/{symbol}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "TVL", GoMethod: "GetTokenProtocols", DocsURL: "https://api-docs.defillama.com/#tag/tvl/get/api/tokenProtocols/%7Bsymbol%7D",
		Summary: "Lists the amount of a certain token within all protocols. Data for the Token Usage page",
		Params: []routeParam{
			{Name: "symbol", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /api/inflows/{protocol}/{timestamp}": {
		Method: "GET", Path: "/api/inflows/{protocol}/{timestamp}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "TVL", GoMethod: "GetInflows", DocsURL: "https://api-docs.defillama.com/#tag/tvl/get/api/inflows/%7Bprotocol%7D/%7Btimestamp%7D",
		Summary: "Lists the amount of inflows and outflows for a protocol at a given date",
		Params: []routeParam{
			{Name: "protocol", In: "path", Required: true, Type: "string"},
			{Name: "timestamp", In: "path", Required: true, Type: "integer"},
		},
	},
	"GET /api/chainAssets": {
		Method: "GET", Path: "/api/chainAssets", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "TVL", GoMethod: "GetChainAssets", DocsURL: "https://api-docs.defillama.com/#tag/tvl/get/api/chainAssets",
		Summary: "Get assets of all chains",
	},
	"GET /api/emissions": {
		Method: "GET", Path: "/api/emissions", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Emissions", GoMethod: "GetEmissions", DocsURL: "https://api-docs.defillama.com/#tag/unlocks/get/api/emissions",
		Summary: "List of all tokens along with basic info for each",
	},
	"GET /api/emission/{protocol}": {
		Method: "GET", Path: "/api/emission/{protocol}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Emissions", GoMethod: "GetEmission", DocsURL: "https://api-docs.defillama.com/#tag/unlocks/get/api/emission/%7Bprotocol%7D",
		Summary: "Unlocks data for a given token/protocol. You can find a list of available slugs to query by querying /emissions and then extracting the key `gecko_id`",
		Params: []routeParam{
			{Name: "protocol", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /api/categories": {
		Method: "GET", Path: "/api/categories", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Ecosystem", GoMethod: "GetCategories", DocsURL: "https://api-docs.defillama.com/#tag/main-page/get/api/categories",
		Summary: "Overview of all categories accross all protocols",
	},
	"GET /api/forks": {
		Method: "GET", Path: "/api/forks", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Ecosystem", GoMethod: "GetForks", DocsURL: "https://api-docs.defillama.com/#tag/main-page/get/api/forks",
		Summary: "Overview of all forks accross all protocols",
	},
	"GET /api/oracles": {
		Method: "GET", Path: "/api/oracles", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Ecosystem", GoMethod: "GetOracles", DocsURL: "https://api-docs.defillama.com/#tag/main-page/get/api/oracles",
		Summary: "Overview of all oracles accross all protocols",
	},
	"GET /api/hacks": {
		Method: "GET", Path: "/api/hacks", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Ecosystem", GoMethod: "GetHacks", DocsURL: "https://api-docs.defillama.com/#tag/main-page/get/api/hacks",
		Summary: "Overview of all hacks on our Hacks dashboard",
	},
	"GET /api/raises": {
		Method: "GET", Path: "/api/raises", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Ecosystem", GoMethod: "GetRaises", DocsURL: "https://api-docs.defillama.com/#tag/main-page/get/api/raises",
		Summary: "Overview of all raises on our Raises dashboard",
	},
	"GET /api/treasuries": {
		Method: "GET", Path: "/api/treasuries", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Ecosystem", GoMethod: "GetTreasuries", DocsURL: "https://api-docs.defillama.com/#tag/main-page/get/api/treasuries",
		Summary: "List all protocols on our Treasuries dashboard",
	},
	"GET /api/entities": {
		Method: "GET", Path: "/api/entities", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Ecosystem", GoMethod: "GetEntities", DocsURL: "https://api-docs.defillama.com/#tag/main-page/get/api/entities",
		Summary: "List all entities",
	},
	"GET /api/historicalLiquidity/{token}": {
		Method: "GET", Path: "/api/historicalLiquidity/{token}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Prices", GoMethod: "GetHistoricalLiquidity", DocsURL: "https://api-docs.defillama.com/#tag/token-liquidity/get/api/historicalLiquidity/%7Btoken%7D",
		Summary: "Provides the name of contracts on a determined chain",
		Params: []routeParam{
			{Name: "token", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /api/overview/derivatives": {
		Method: "GET", Path: "/api/overview/derivatives", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Volumes", GoMethod: "GetDerivativesOverview", DocsURL: "https://api-docs.defillama.com/#tag/perps/get/api/overview/derivatives",
		Summary: "Lists all derivatives along summaries of their volumes filtering by chain",
		Params: []routeParam{
			{Name: "excludeTotalDataChart", In: "query", Required: true, Type: "boolean"},
			{Name: "excludeTotalDataChartBreakdown", In: "query", Required: true, Type: "boolean"},
		},
	},
	"GET /api/summary/derivatives/{protocol}": {
		Method: "GET", Path: "/api/summary/derivatives/{protocol}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Volumes", GoMethod: "GetDerivativesSummary", DocsURL: "https://api-docs.defillama.com/#tag/perps/get/api/summary/derivatives/%7Bprotocol%7D",
		Summary: "Volume Details about a specific perp protocol",
		Params: []routeParam{
			{Name: "protocol", In: "path", Required: true, Type: "string"},
			{Name: "excludeTotalDataChart", In: "query", Required: true, Type: "boolean"},
			{Name: "excludeTotalDataChartBreakdown", In: "query", Required: true, Type: "boolean"},
		},
	},
	"GET /api/v2/metrics/tvl/protocol/{protocol}": {
		Method: "GET", Path: "/api/v2/metrics/tvl/protocol/{protocol}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "TVL", GoMethod: "GetProtocolTVLMetrics", DocsURL: "https://api-docs.defillama.com/#tag/tvl/get/api/v2/metrics/tvl/protocol/%7Bprotocol%7D",
		Summary: "Get aggregate TVL metrics for a protocol",
		Params: []routeParam{
			{Name: "protocol", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /api/v2/chart/tvl/protocol/{protocol}": {
		Method: "GET", Path: "/api/v2/chart/tvl/protocol/{protocol}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "TVL", GoMethod: "GetProtocolTVLChart", DocsURL: "https://api-docs.defillama.com/#tag/tvl/get/api/v2/chart/tvl/protocol/%7Bprotocol%7D",
		Summary: "Get historical TVL chart for a protocol",
		Params: []routeParam{
			{Name: "protocol", In: "path", Required: true, Type: "string"},
			{Name: "key", In: "query", Required: false, Type: "string", Enum: []string{"all", "staking", "borrowed", "vesting", "pool2"}},
		},
	},
	"GET /api/v2/chart/tvl/protocol/{protocol}/chain-breakdown": {
		Method: "GET", Path: "/api/v2/chart/tvl/protocol/{protocol}/chain-breakdown", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "TVL", GoMethod: "GetProtocolTVLChartChainBreakdown", DocsURL: "https://api-docs.defillama.com/#tag/tvl/get/api/v2/chart/tvl/protocol/%7Bprotocol%7D/chain-breakdown",
		Summary: "Get historical TVL chart for a protocol broken down by chain",
		Params: []routeParam{
			{Name: "protocol", In: "path", Required: true, Type: "string"},
			{Name: "key", In: "query", Required: false, Type: "string", Enum: []string{"all", "staking", "borrowed", "vesting", "pool2"}},
		},
	},
	"GET /api/v2/chart/tvl/protocol/{protocol}/token-breakdown": {
		Method: "GET", Path: "/api/v2/chart/tvl/protocol/{protocol}/token-breakdown", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "TVL", GoMethod: "GetProtocolTVLChartTokenBreakdown", DocsURL: "https://api-docs.defillama.com/#tag/tvl/get/api/v2/chart/tvl/protocol/%7Bprotocol%7D/token-breakdown",
		Summary: "Get historical TVL chart for a protocol broken down by token",
		Params: []routeParam{
			{Name: "protocol", In: "path", Required: true, Type: "string"},
			{Name: "key", In: "query", Required: false, Type: "string", Enum: []string{"all", "staking", "borrowed", "vesting", "pool2"}},
			{Name: "currency", In: "query", Required: false, Type: "string", Enum: []string{"usd", "token", "raw"}},
		},
	},
	"GET /api/v2/metrics/treasury/protocol/{protocol}": {
		Method: "GET", Path: "/api/v2/metrics/treasury/protocol/{protocol}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Treasury", GoMethod: "GetTreasuryMetrics", DocsURL: "https://api-docs.defillama.com/#tag/treasury/get/api/v2/metrics/treasury/protocol/%7Bprotocol%7D",
		Summary: "Get aggregate treasury metrics for a protocol",
		Params: []routeParam{
			{Name: "protocol", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /api/v2/chart/treasury/protocol/{protocol}": {
		Method: "GET", Path: "/api/v2/chart/treasury/protocol/{protocol}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Treasury", GoMethod: "GetTreasuryChart", DocsURL: "https://api-docs.defillama.com/#tag/treasury/get/api/v2/chart/treasury/protocol/%7Bprotocol%7D",
		Summary: "Get historical treasury chart for a protocol",
		Params: []routeParam{
			{Name: "protocol", In: "path", Required: true, Type: "string"},
			{Name: "key", In: "query", Required: false, Type: "string", Enum: []string{"OwnTokens", "all"}},
		},
	},
	"GET /api/v2/chart/treasury/protocol/{protocol}/chain-breakdown": {
		Method: "GET", Path: "/api/v2/chart/treasury/protocol/{protocol}/chain-breakdown", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Treasury", GoMethod: "GetTreasuryChartChainBreakdown", DocsURL: "https://api-docs.defillama.com/#tag/treasury/get/api/v2/chart/treasury/protocol/%7Bprotocol%7D/chain-breakdown",
		Summary: "Get historical treasury chart for a protocol broken down by chain",
		Params: []routeParam{
			{Name: "protocol", In: "path", Required: true, Type: "string"},
			{Name: "key", In: "query", Required: false, Type: "string", Enum: []string{"OwnTokens", "all"}},
		},
	},
	"GET /api/v2/chart/treasury/protocol/{protocol}/token-breakdown": {
		Method: "GET", Path: "/api/v2/chart/treasury/protocol/{protocol}/token-breakdown", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Treasury", GoMethod: "GetTreasuryChartTokenBreakdown", DocsURL: "https://api-docs.defillama.com/#tag/treasury/get/api/v2/chart/treasury/protocol/%7Bprotocol%7D/token-breakdown",
		Summary: "Get historical treasury chart for a protocol broken down by token",
		Params: []routeParam{
			{Name: "protocol", In: "path", Required: true, Type: "string"},
			{Name: "key", In: "query", Required: false, Type: "string", Enum: []string{"OwnTokens", "all"}},
			{Name: "currency", In: "query", Required: false, Type: "string", Enum: []string{"usd", "token", "raw"}},
		},
	},
	"GET /api/v2/metrics/oracle": {
		Method: "GET", Path: "/api/v2/metrics/oracle", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Oracles", GoMethod: "GetOracleMetrics", DocsURL: "https://api-docs.defillama.com/#tag/oracles/get/api/v2/metrics/oracle",
		Summary: "Get oracle data overview",
	},
	"GET /api/v2/chart/oracle": {
		Method: "GET", Path: "/api/v2/chart/oracle", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Oracles", GoMethod: "GetOracleChart", DocsURL: "https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle",
		Summary: "Get timeseries chart data for all oracles",
	},
	"GET /api/v2/chart/oracle/chain-breakdown": {
		Method: "GET", Path: "/api/v2/chart/oracle/chain-breakdown", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Oracles", GoMethod: "GetOracleChartChainBreakdown", DocsURL: "https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle/chain-breakdown",
		Summary: "Get timeseries chart data breakdown by chain",
	},
	"GET /api/v2/chart/oracle/protocol-breakdown": {
		Method: "GET", Path: "/api/v2/chart/oracle/protocol-breakdown", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Oracles", GoMethod: "GetOracleChartProtocolBreakdown", DocsURL: "https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle/protocol-breakdown",
		Summary: "Get timeseries chart data breakdown by protocol",
	},
	"GET /api/v2/chart/oracle/protocol/{protocol}": {
		Method: "GET", Path: "/api/v2/chart/oracle/protocol/{protocol}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Oracles", GoMethod: "GetOracleProtocolChart", DocsURL: "https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle/protocol/%7Bprotocol%7D",
		Summary: "Get timeseries chart data by protocol/oracle",
		Params: []routeParam{
			{Name: "protocol", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /api/v2/chart/oracle/protocol/{protocol}/chain-breakdown": {
		Method: "GET", Path: "/api/v2/chart/oracle/protocol/{protocol}/chain-breakdown", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Oracles", GoMethod: "GetOracleProtocolChartChainBreakdown", DocsURL: "https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle/protocol/%7Bprotocol%7D/chain-breakdown",
		Summary: "Get chain breakdown timeseries chart data by protocol/oracle",
		Params: []routeParam{
			{Name: "protocol", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /api/v2/chart/oracle/chain/{chain}": {
		Method: "GET", Path: "/api/v2/chart/oracle/chain/{chain}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Oracles", GoMethod: "GetOracleChainChart", DocsURL: "https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle/chain/%7Bchain%7D",
		Summary: "Get timeseries chart data by chain",
		Params: []routeParam{
			{Name: "chain", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /api/v2/chart/oracle/chain/{chain}/protocol-breakdown": {
		Method: "GET", Path: "/api/v2/chart/oracle/chain/{chain}/protocol-breakdown", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Oracles", GoMethod: "GetOracleChainChartProtocolBreakdown", DocsURL: "https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle/chain/%7Bchain%7D/protocol-breakdown",
		Summary: "Get protocol breakdown timeseries chart data by chain",
		Params: []routeParam{
			{Name: "chain", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /api/v2/metrics/fork": {
		Method: "GET", Path: "/api/v2/metrics/fork", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Forks", GoMethod: "GetForkMetrics", DocsURL: "https://api-docs.defillama.com/#tag/forks/get/api/v2/metrics/fork",
		Summary: "Get fork data overview",
	},
	"GET /api/v2/chart/fork/protocol-breakdown": {
		Method: "GET", Path: "/api/v2/chart/fork/protocol-breakdown", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Forks", GoMethod: "GetForkChartProtocolBreakdown", DocsURL: "https://api-docs.defillama.com/#tag/forks/get/api/v2/chart/fork/protocol-breakdown",
		Summary: "Get timeseries chart data breakdown by protocol",
	},
	"GET /api/v2/chart/fork/protocol/{protocol}": {
		Method: "GET", Path: "/api/v2/chart/fork/protocol/{protocol}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Forks", GoMethod: "GetForkProtocolChart", DocsURL: "https://api-docs.defillama.com/#tag/forks/get/api/v2/chart/fork/protocol/%7Bprotocol%7D",
		Summary: "Get timeseries chart data by protocol",
		Params: []routeParam{
			{Name: "protocol", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /api/v2/metrics/{metric}": {
		Method: "GET", Path: "/api/v2/metrics/{metric}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Dimensions", GoMethod: "GetMetrics", DocsURL: "https://api-docs.defillama.com/#tag/dimensions/get/api/v2/metrics/%7Bmetric%7D",
		Summary: "Get dimension data overview",
		Params: []routeParam{
			{Name: "metric", In: "path", Required: true, Type: "string", Enum: []string{"fees", "liquidations", "dexs", "derivatives", "options", "aggregators", "bridge-aggregators", "open-interest", "normalized-volume", "active-users"}},
			{Name: "dataType", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /api/v2/chart/{metric}": {
		Method: "GET", Path: "/api/v2/chart/{metric}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Dimensions", GoMethod: "GetChart", DocsURL: "https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D",
		Summary: "Get historical timeseries chart data",
		Params: []routeParam{
			{Name: "metric", In: "path", Required: true, Type: "string", Enum: []string{"fees", "liquidations", "dexs", "derivatives", "options", "aggregators", "bridge-aggregators", "open-interest", "normalized-volume", "active-users"}},
			{Name: "dataType", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /api/v2/chart/{metric}/chain-breakdown": {
		Method: "GET", Path: "/api/v2/chart/{metric}/chain-breakdown", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Dimensions", GoMethod: "GetChartChainBreakdown", DocsURL: "https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/chain-breakdown",
		Summary: "Get historical timeseries chart data breakdown by chain",
		Params: []routeParam{
			{Name: "metric", In: "path", Required: true, Type: "string", Enum: []string{"fees", "liquidations", "dexs", "derivatives", "options", "aggregators", "bridge-aggregators", "open-interest", "normalized-volume", "active-users"}},
			{Name: "dataType", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /api/v2/chart/{metric}/protocol-breakdown": {
		Method: "GET", Path: "/api/v2/chart/{metric}/protocol-breakdown", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Dimensions", GoMethod: "GetChartProtocolBreakdown", DocsURL: "https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/protocol-breakdown",
		Summary: "Get historical timeseries chart data breakdown by protocol",
		Params: []routeParam{
			{Name: "metric", In: "path", Required: true, Type: "string", Enum: []string{"fees", "liquidations", "dexs", "derivatives", "options", "aggregators", "bridge-aggregators", "open-interest", "normalized-volume"}},
			{Name: "dataType", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /api/v2/metrics/{metric}/chain/{chain}": {
		Method: "GET", Path: "/api/v2/metrics/{metric}/chain/{chain}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Dimensions", GoMethod: "GetChainMetrics", DocsURL: "https://api-docs.defillama.com/#tag/dimensions/get/api/v2/metrics/%7Bmetric%7D/chain/%7Bchain%7D",
		Summary: "Get chain dimension data overview",
		Params: []routeParam{
			{Name: "metric", In: "path", Required: true, Type: "string", Enum: []string{"fees", "liquidations", "dexs", "derivatives", "options", "aggregators", "bridge-aggregators", "open-interest", "normalized-volume"}},
			{Name: "chain", In: "path", Required: true, Type: "string"},
			{Name: "dataType", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /api/v2/chart/{metric}/chain/{chain}": {
		Method: "GET", Path: "/api/v2/chart/{metric}/chain/{chain}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Dimensions", GoMethod: "GetChainChart", DocsURL: "https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/chain/%7Bchain%7D",
		Summary: "Get chain historical timeseries chart data",
		Params: []routeParam{
			{Name: "metric", In: "path", Required: true, Type: "string", Enum: []string{"fees", "liquidations", "dexs", "derivatives", "options", "aggregators", "bridge-aggregators", "open-interest", "normalized-volume"}},
			{Name: "chain", In: "path", Required: true, Type: "string"},
			{Name: "dataType", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /api/v2/chart/{metric}/chain/{chain}/protocol-breakdown": {
		Method: "GET", Path: "/api/v2/chart/{metric}/chain/{chain}/protocol-breakdown", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Dimensions", GoMethod: "GetChainChartProtocolBreakdown", DocsURL: "https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/chain/%7Bchain%7D/protocol-breakdown",
		Summary: "Get chain timeseries chart data breakdown by protocol",
		Params: []routeParam{
			{Name: "metric", In: "path", Required: true, Type: "string", Enum: []string{"fees", "liquidations", "dexs", "derivatives", "options", "aggregators", "bridge-aggregators", "open-interest", "normalized-volume"}},
			{Name: "chain", In: "path", Required: true, Type: "string"},
			{Name: "dataType", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /api/v2/metrics/{metric}/protocol/{protocol}": {
		Method: "GET", Path: "/api/v2/metrics/{metric}/protocol/{protocol}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Dimensions", GoMethod: "GetProtocolMetrics", DocsURL: "https://api-docs.defillama.com/#tag/dimensions/get/api/v2/metrics/%7Bmetric%7D/protocol/%7Bprotocol%7D",
		Summary: "Get protocol dimension data overview",
		Params: []routeParam{
			{Name: "metric", In: "path", Required: true, Type: "string", Enum: []string{"fees", "liquidations", "dexs", "derivatives", "options", "aggregators", "bridge-aggregators", "open-interest", "normalized-volume"}},
			{Name: "protocol", In: "path", Required: true, Type: "string"},
			{Name: "dataType", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /api/v2/chart/{metric}/protocol/{protocol}": {
		Method: "GET", Path: "/api/v2/chart/{metric}/protocol/{protocol}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Dimensions", GoMethod: "GetProtocolChart", DocsURL: "https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/protocol/%7Bprotocol%7D",
		Summary: "Get protocol historical timeseries chart data",
		Params: []routeParam{
			{Name: "metric", In: "path", Required: true, Type: "string", Enum: []string{"fees", "liquidations", "dexs", "derivatives", "options", "aggregators", "bridge-aggregators", "open-interest", "normalized-volume"}},
			{Name: "protocol", In: "path", Required: true, Type: "string"},
			{Name: "dataType", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /api/v2/chart/{metric}/protocol/{protocol}/chain-breakdown": {
		Method: "GET", Path: "/api/v2/chart/{metric}/protocol/{protocol}/chain-breakdown", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Dimensions", GoMethod: "GetProtocolChartChainBreakdown", DocsURL: "https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/protocol/%7Bprotocol%7D/chain-breakdown",
		Summary: "Get protocol timeseries chart data breakdown by chain",
		Params: []routeParam{
			{Name: "metric", In: "path", Required: true, Type: "string", Enum: []string{"fees", "liquidations", "dexs", "derivatives", "options", "aggregators", "bridge-aggregators", "open-interest", "normalized-volume"}},
			{Name: "protocol", In: "path", Required: true, Type: "string"},
			{Name: "dataType", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /api/v2/chart/{metric}/protocol/{protocol}/version-breakdown": {
		Method: "GET", Path: "/api/v2/chart/{metric}/protocol/{protocol}/version-breakdown", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Dimensions", GoMethod: "GetProtocolChartVersionBreakdown", DocsURL: "https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/protocol/%7Bprotocol%7D/version-breakdown",
		Summary: "Get protocol timeseries chart data breakdown by version",
		Params: []routeParam{
			{Name: "metric", In: "path", Required: true, Type: "string", Enum: []string{"fees", "liquidations", "dexs", "derivatives", "options", "aggregators", "bridge-aggregators", "open-interest", "normalized-volume"}},
			{Name: "protocol", In: "path", Required: true, Type: "string"},
			{Name: "dataType", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /api/v2/chart/{metric}/protocol/{protocol}/label-breakdown": {
		Method: "GET", Path: "/api/v2/chart/{metric}/protocol/{protocol}/label-breakdown", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Dimensions", GoMethod: "GetProtocolChartLabelBreakdown", DocsURL: "https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/protocol/%7Bprotocol%7D/label-breakdown",
		Summary: "Get protocol timeseries chart data breakdown by label",
		Params: []routeParam{
			{Name: "metric", In: "path", Required: true, Type: "string", Enum: []string{"fees", "liquidations", "dexs", "derivatives", "options", "aggregators", "bridge-aggregators", "open-interest", "normalized-volume"}},
			{Name: "protocol", In: "path", Required: true, Type: "string"},
			{Name: "dataType", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /api/v2/metrics/{metric}/category/{category}": {
		Method: "GET", Path: "/api/v2/metrics/{metric}/category/{category}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Dimensions", GoMethod: "GetCategoryMetrics", DocsURL: "https://api-docs.defillama.com/#tag/dimensions/get/api/v2/metrics/%7Bmetric%7D/category/%7Bcategory%7D",
		Summary: "Get category dimension data overview",
		Params: []routeParam{
			{Name: "metric", In: "path", Required: true, Type: "string", Enum: []string{"fees", "liquidations", "dexs", "derivatives", "options", "aggregators", "bridge-aggregators", "open-interest", "normalized-volume"}},
			{Name: "category", In: "path", Required: true, Type: "string"},
			{Name: "dataType", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /api/v2/chart/{metric}/category/{category}": {
		Method: "GET", Path: "/api/v2/chart/{metric}/category/{category}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Dimensions", GoMethod: "GetCategoryChart", DocsURL: "https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/category/%7Bcategory%7D",
		Summary: "Get category historical timeseries chart data",
		Params: []routeParam{
			{Name: "metric", In: "path", Required: true, Type: "string", Enum: []string{"fees", "liquidations", "dexs", "derivatives", "options", "aggregators", "bridge-aggregators", "open-interest", "normalized-volume"}},
			{Name: "category", In: "path", Required: true, Type: "string"},
			{Name: "dataType", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /api/v2/chart/{metric}/category/{category}/chain-breakdown": {
		Method: "GET", Path: "/api/v2/chart/{metric}/category/{category}/chain-breakdown", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Dimensions", GoMethod: "GetCategoryChartChainBreakdown", DocsURL: "https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/category/%7Bcategory%7D/chain-breakdown",
		Summary: "Get category timeseries chart data breakdown by chain",
		Params: []routeParam{
			{Name: "metric", In: "path", Required: true, Type: "string", Enum: []string{"fees", "liquidations", "dexs", "derivatives", "options", "aggregators", "bridge-aggregators", "open-interest", "normalized-volume"}},
			{Name: "category", In: "path", Required: true, Type: "string"},
			{Name: "dataType", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /api/v2/chart/{metric}/category/{category}/protocol-breakdown": {
		Method: "GET", Path: "/api/v2/chart/{metric}/category/{category}/protocol-breakdown", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Dimensions", GoMethod: "GetCategoryChartProtocolBreakdown", DocsURL: "https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/category/%7Bcategory%7D/protocol-breakdown",
		Summary: "Get category timeseries chart data breakdown by protocol",
		Params: []routeParam{
			{Name: "metric", In: "path", Required: true, Type: "string", Enum: []string{"fees", "liquidations", "dexs", "derivatives", "options", "aggregators", "bridge-aggregators", "open-interest", "normalized-volume"}},
			{Name: "category", In: "path", Required: true, Type: "string"},
			{Name: "dataType", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /api/v2/metrics/{metric}/category/{category}/chain/{chain}": {
		Method: "GET", Path: "/api/v2/metrics/{metric}/category/{category}/chain/{chain}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Dimensions", GoMethod: "GetCategoryChainMetrics", DocsURL: "https://api-docs.defillama.com/#tag/dimensions/get/api/v2/metrics/%7Bmetric%7D/category/%7Bcategory%7D/chain/%7Bchain%7D",
		Summary: "Get category chain dimension data overview",
		Params: []routeParam{
			{Name: "metric", In: "path", Required: true, Type: "string", Enum: []string{"fees", "liquidations", "dexs", "derivatives", "options", "aggregators", "bridge-aggregators", "open-interest", "normalized-volume"}},
			{Name: "category", In: "path", Required: true, Type: "string"},
			{Name: "chain", In: "path", Required: true, Type: "string"},
			{Name: "dataType", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /api/v2/chart/{metric}/category/{category}/chain/{chain}": {
		Method: "GET", Path: "/api/v2/chart/{metric}/category/{category}/chain/{chain}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Dimensions", GoMethod: "GetCategoryChainChart", DocsURL: "https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/category/%7Bcategory%7D/chain/%7Bchain%7D",
		Summary: "Get category chain historical timeseries chart data",
		Params: []routeParam{
			{Name: "metric", In: "path", Required: true, Type: "string", Enum: []string{"fees", "liquidations", "dexs", "derivatives", "options", "aggregators", "bridge-aggregators", "open-interest", "normalized-volume"}},
			{Name: "category", In: "path", Required: true, Type: "string"},
			{Name: "chain", In: "path", Required: true, Type: "string"},
			{Name: "dataType", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /api/v2/chart/{metric}/category/{category}/chain/{chain}/protocol-breakdown": {
		Method: "GET", Path: "/api/v2/chart/{metric}/category/{category}/chain/{chain}/protocol-breakdown", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Dimensions", GoMethod: "GetCategoryChainChartProtocolBreakdown", DocsURL: "https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/category/%7Bcategory%7D/chain/%7Bchain%7D/protocol-breakdown",
		Summary: "Get category chain timeseries chart data breakdown by protocol",
		Params: []routeParam{
			{Name: "metric", In: "path", Required: true, Type: "string", Enum: []string{"fees", "liquidations", "dexs", "derivatives", "options", "aggregators", "bridge-aggregators", "open-interest", "normalized-volume"}},
			{Name: "category", In: "path", Required: true, Type: "string"},
			{Name: "chain", In: "path", Required: true, Type: "string"},
			{Name: "dataType", In: "query", Required: false, Type: "string"},
		},
	},
	"GET /api/v2/metrics/financial-statement/protocol/{protocol}": {
		Method: "GET", Path: "/api/v2/metrics/financial-statement/protocol/{protocol}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "FinancialStatements", GoMethod: "GetIncomeStatement", DocsURL: "https://api-docs.defillama.com/#tag/financial-statements/get/api/v2/metrics/financial-statement/protocol/%7Bprotocol%7D",
		Summary: "Get protocol income statement report",
		Params: []routeParam{
			{Name: "protocol", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /stablecoins/stablecoindominance/{chain}": {
		Method: "GET", Path: "/stablecoins/stablecoindominance/{chain}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Stablecoins", GoMethod: "GetStablecoinDominance", DocsURL: "https://api-docs.defillama.com/#tag/stablecoins/get/stablecoins/stablecoindominance/%7Bchain%7D",
		Summary: "Get stablecoin dominance per chain along with the info about the larges coin in a chain",
		Params: []routeParam{
			{Name: "chain", In: "path", Required: true, Type: "string"},
			{Name: "stablecoin", In: "query", Required: false, Type: "integer"},
		},
	},
	"GET /yields/v1/pools": {
		Method: "GET", Path: "/yields/v1/pools", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Yields", GoMethod: "GetLegacyPools", DocsURL: "https://api-docs.defillama.com/#tag/yields/get/yields/v1/pools",
		Summary: "Retrieve the latest data for all pools in the legacy v1 response format",
	},
	"GET /yields/v1/chart/{pool}": {
		Method: "GET", Path: "/yields/v1/chart/{pool}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Yields", GoMethod: "GetLegacyPoolChart", DocsURL: "https://api-docs.defillama.com/#tag/yields/get/yields/v1/chart/%7Bpool%7D",
		Summary: "Get historical APY and TVL of a pool in the legacy v1 response format",
		Params: []routeParam{
			{Name: "pool", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /yields/v1/poolsBorrow": {
		Method: "GET", Path: "/yields/v1/poolsBorrow", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Yields", GoMethod: "GetLegacyBorrowPools", DocsURL: "https://api-docs.defillama.com/#tag/yields/get/yields/v1/poolsBorrow",
		Summary: "Borrow costs APY of assets from lending markets in the legacy v1 response format",
	},
	"GET /yields/v1/chartLendBorrow/{pool}": {
		Method: "GET", Path: "/yields/v1/chartLendBorrow/{pool}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Yields", GoMethod: "GetLegacyLendBorrowChart", DocsURL: "https://api-docs.defillama.com/#tag/yields/get/yields/v1/chartLendBorrow/%7Bpool%7D",
		Summary: "Historical borrow cost APY from a pool on a lending market in the legacy v1 response format",
		Params: []routeParam{
			{Name: "pool", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /yields/v2/earn": {
		Method: "GET", Path: "/yields/v2/earn", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Yields", GoMethod: "GetEarnPools", DocsURL: "https://api-docs.defillama.com/#tag/yields/get/yields/v2/earn",
		Summary: "Retrieve the latest data for all earn pools",
	},
	"GET /yields/v2/earn/query": {
		Method: "GET", Path: "/yields/v2/earn/query", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Yields", GoMethod: "QueryEarnPools", DocsURL: "https://api-docs.defillama.com/#tag/yields/get/yields/v2/earn/query",
		Summary: "Retrieve a filtered, paginated page of earn pools",
		Params: []routeParam{
			{Name: "chain", In: "query", Required: false, Type: "string"},
			{Name: "protocol", In: "query", Required: false, Type: "string"},
			{Name: "stablecoin", In: "query", Required: false, Type: "boolean"},
			{Name: "single_asset_exposure", In: "query", Required: false, Type: "boolean"},
			{Name: "impermanent_loss_risk", In: "query", Required: false, Type: "boolean"},
			{Name: "min_tvl", In: "query", Required: false, Type: "number"},
			{Name: "page", In: "query", Required: false, Type: "integer"},
			{Name: "limit", In: "query", Required: false, Type: "integer"},
		},
	},
	"GET /yields/v2/earn/{id}/history": {
		Method: "GET", Path: "/yields/v2/earn/{id}/history", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Yields", GoMethod: "GetEarnPoolHistory", DocsURL: "https://api-docs.defillama.com/#tag/yields/get/yields/v2/earn/%7Bid%7D/history",
		Summary: "Retrieve daily history for an earn pool",
		Params: []routeParam{
			{Name: "id", In: "path", Required: true, Type: "string"},
			{Name: "range", In: "query", Required: false, Type: "string", Enum: []string{"30d", "90d", "max"}},
		},
	},
	"GET /yields/v2/earn/verified": {
		Method: "GET", Path: "/yields/v2/earn/verified", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Yields", GoMethod: "GetVerifiedEarnPools", DocsURL: "https://api-docs.defillama.com/#tag/yields/get/yields/v2/earn/verified",
		Summary: "Retrieve the latest verified data for all measured earn pools",
	},
	"GET /yields/v2/earn/verified/{id}": {
		Method: "GET", Path: "/yields/v2/earn/verified/{id}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Yields", GoMethod: "GetVerifiedEarnPool", DocsURL: "https://api-docs.defillama.com/#tag/yields/get/yields/v2/earn/verified/%7Bid%7D",
		Summary: "Retrieve the latest verified data for one earn pool",
		Params: []routeParam{
			{Name: "id", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /yields/v2/earn/verified/{id}/history": {
		Method: "GET", Path: "/yields/v2/earn/verified/{id}/history", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Yields", GoMethod: "GetVerifiedEarnPoolHistory", DocsURL: "https://api-docs.defillama.com/#tag/yields/get/yields/v2/earn/verified/%7Bid%7D/history",
		Summary: "Retrieve daily verified history for an earn pool",
		Params: []routeParam{
			{Name: "id", In: "path", Required: true, Type: "string"},
			{Name: "range", In: "query", Required: false, Type: "string", Enum: []string{"30d", "90d", "max"}},
		},
	},
	"GET /yields/v2/borrow/markets": {
		Method: "GET", Path: "/yields/v2/borrow/markets", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Yields", GoMethod: "GetBorrowMarkets", DocsURL: "https://api-docs.defillama.com/#tag/yields/get/yields/v2/borrow/markets",
		Summary: "Retrieve the latest data for all borrow markets",
	},
	"GET /yields/v2/borrow/markets/{id}/history": {
		Method: "GET", Path: "/yields/v2/borrow/markets/{id}/history", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Yields", GoMethod: "GetBorrowMarketHistory", DocsURL: "https://api-docs.defillama.com/#tag/yields/get/yields/v2/borrow/markets/%7Bid%7D/history",
		Summary: "Retrieve daily borrow-side history for a borrow market",
		Params: []routeParam{
			{Name: "id", In: "path", Required: true, Type: "string"},
			{Name: "range", In: "query", Required: false, Type: "string", Enum: []string{"30d", "90d", "max"}},
		},
	},
	"GET /yields/v2/borrow/routes": {
		Method: "GET", Path: "/yields/v2/borrow/routes", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Yields", GoMethod: "GetBorrowRoutes", DocsURL: "https://api-docs.defillama.com/#tag/yields/get/yields/v2/borrow/routes",
		Summary: "Retrieve the latest data for all borrow routes",
	},
	"GET /yields/v2/loops": {
		Method: "GET", Path: "/yields/v2/loops", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Yields", GoMethod: "GetLoopStrategies", DocsURL: "https://api-docs.defillama.com/#tag/yields/get/yields/v2/loops",
		Summary: "Retrieve the latest data for all loop strategies",
	},
	"GET /yields/perps": {
		Method: "GET", Path: "/yields/perps", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Yields", GoMethod: "GetPerpFundingRates", DocsURL: "https://api-docs.defillama.com/#tag/yields/get/yields/perps",
		Summary: "Funding rates and Open Interest of perps across exchanges, including both Decentralized and Centralized",
	},
	"GET /api/lstRates": {
		Method: "GET", Path: "/api/lstRates", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Yields", GoMethod: "GetLSTRates", DocsURL: "https://api-docs.defillama.com/#tag/yields/get/api/lstRates",
		Summary: "Exchange rates and ETH peg of liquid staking tokens",
	},
	"GET /etfs/snapshot": {
		Method: "GET", Path: "/etfs/snapshot", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "ETFs", GoMethod: "GetSnapshot", DocsURL: "https://api-docs.defillama.com/#tag/etfs/get/etfs/snapshot",
		Summary: "Get ETFs and their metrics (aum, flows, fees...)",
	},
	"GET /etfs/flows": {
		Method: "GET", Path: "/etfs/flows", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "ETFs", GoMethod: "GetFlows", DocsURL: "https://api-docs.defillama.com/#tag/etfs/get/etfs/flows",
		Summary: "Historical Flows at the Asset Level",
	},
	"GET /fdv/performance/{period}": {
		Method: "GET", Path: "/fdv/performance/{period}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Narratives", GoMethod: "GetPerformance", DocsURL: "https://api-docs.defillama.com/#tag/narratives/get/fdv/performance/%7Bperiod%7D",
		Summary: "Get chart of narratives based on category performance (with individual coins weighted by mcap)",
		Params: []routeParam{
			{Name: "period", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /bridges/bridges": {
		Method: "GET", Path: "/bridges/bridges", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Bridges", GoMethod: "GetBridges", DocsURL: "https://api-docs.defillama.com/#tag/bridges/get/bridges/bridges",
		Summary: "List all bridges along with summaries of recent bridge volumes.",
		Params: []routeParam{
			{Name: "includeChains", In: "query", Required: false, Type: "boolean"},
		},
	},
	"GET /bridges/bridge/{id}": {
		Method: "GET", Path: "/bridges/bridge/{id}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Bridges", GoMethod: "GetBridge", DocsURL: "https://api-docs.defillama.com/#tag/bridges/get/bridges/bridge/%7Bid%7D",
		Summary: "Get summary of bridge volume and volume breakdown by chain",
		Params: []routeParam{
			{Name: "id", In: "path", Required: true, Type: "integer"},
		},
	},
	"GET /bridges/bridgevolume/{chain}": {
		Method: "GET", Path: "/bridges/bridgevolume/{chain}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Bridges", GoMethod: "GetBridgeVolume", DocsURL: "https://api-docs.defillama.com/#tag/bridges/get/bridges/bridgevolume/%7Bchain%7D",
		Summary: "Get historical volumes for a bridge, chain, or bridge on a particular chain",
		Params: []routeParam{
			{Name: "chain", In: "path", Required: true, Type: "string"},
			{Name: "id", In: "query", Required: false, Type: "integer"},
		},
	},
	"GET /bridges/bridgedaystats/{timestamp}/{chain}": {
		Method: "GET", Path: "/bridges/bridgedaystats/{timestamp}/{chain}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Bridges", GoMethod: "GetBridgeDayStats", DocsURL: "https://api-docs.defillama.com/#tag/bridges/get/bridges/bridgedaystats/%7Btimestamp%7D/%7Bchain%7D",
		Summary: "Get a 24hr token and address volume breakdown for a bridge",
		Params: []routeParam{
			{Name: "timestamp", In: "path", Required: true, Type: "integer"},
			{Name: "chain", In: "path", Required: true, Type: "string"},
			{Name: "id", In: "query", Required: false, Type: "integer"},
		},
	},
	"GET /bridges/transactions/{id}": {
		Method: "GET", Path: "/bridges/transactions/{id}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Bridges", GoMethod: "GetBridgeTransactions", DocsURL: "https://api-docs.defillama.com/#tag/bridges/get/bridges/transactions/%7Bid%7D",
		Summary: "Get all transactions for a bridge within a date range",
		Params: []routeParam{
			{Name: "id", In: "path", Required: true, Type: "integer"},
			{Name: "starttimestamp", In: "query", Required: false, Type: "integer"},
			{Name: "endtimestamp", In: "query", Required: false, Type: "integer"},
			{Name: "sourcechain", In: "query", Required: false, Type: "string"},
			{Name: "address", In: "query", Required: false, Type: "string"},
			{Name: "limit", In: "query", Required: false, Type: "integer"},
		},
	},
	"GET /usage/APIKEY": {
		Method: "GET", Path: "/usage/APIKEY", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Account", GoMethod: "GetUsage", DocsURL: "https://api-docs.defillama.com/#tag/meta/get/usage/APIKEY",
		Summary: "Get amount of credits left in the api key, these reset on the 1st of each month",
	},
	"GET /dat/institutions": {
		Method: "GET", Path: "/dat/institutions", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "DAT", GoMethod: "GetInstitutions", DocsURL: "https://api-docs.defillama.com/#tag/dat/get/dat/institutions",
		Summary: "Get list of all institutions with Digital Asset Treasury data",
	},
	"GET /dat/institutions/{symbol}": {
		Method: "GET", Path: "/dat/institutions/{symbol}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "DAT", GoMethod: "GetInstitution", DocsURL: "https://api-docs.defillama.com/#tag/dat/get/dat/institutions/%7Bsymbol%7D",
		Summary: "Get individual institution Digital Asset Treasury details",
		Params: []routeParam{
			{Name: "symbol", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /equities/v1/companies-list": {
		Method: "GET", Path: "/equities/v1/companies-list", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Equities", GoMethod: "GetCompaniesList", DocsURL: "https://api-docs.defillama.com/#tag/equities/get/equities/v1/companies-list",
		Summary: "Get list of all tracked public companies",
	},
	"GET /equities/v1/statements": {
		Method: "GET", Path: "/equities/v1/statements", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Equities", GoMethod: "GetStatements", DocsURL: "https://api-docs.defillama.com/#tag/equities/get/equities/v1/statements",
		Summary: "Get financial statements for a company",
		Params: []routeParam{
			{Name: "ticker", In: "query", Required: true, Type: "string"},
			{Name: "country", In: "query", Required: true, Type: "string"},
		},
	},
	"GET /equities/v1/dimensions": {
		Method: "GET", Path: "/equities/v1/dimensions", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Equities", GoMethod: "GetDimensions", DocsURL: "https://api-docs.defillama.com/#tag/equities/get/equities/v1/dimensions",
		Summary: "Get financial dimensions for a company",
		Params: []routeParam{
			{Name: "ticker", In: "query", Required: true, Type: "string"},
			{Name: "country", In: "query", Required: true, Type: "string"},
		},
	},
	"GET /equities/v1/price-history": {
		Method: "GET", Path: "/equities/v1/price-history", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Equities", GoMethod: "GetPriceHistory", DocsURL: "https://api-docs.defillama.com/#tag/equities/get/equities/v1/price-history",
		Summary: "Get historical price data for a company",
		Params: []routeParam{
			{Name: "ticker", In: "query", Required: true, Type: "string"},
			{Name: "country", In: "query", Required: true, Type: "string"},
			{Name: "timeframe", In: "query", Required: false, Type: "string", Enum: []string{"1D", "7D", "1W", "1M", "3M", "6M", "YTD", "1Y", "5Y", "MAX"}},
		},
	},
	"GET /equities/v1/ohlcv": {
		Method: "GET", Path: "/equities/v1/ohlcv", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Equities", GoMethod: "GetOHLCV", DocsURL: "https://api-docs.defillama.com/#tag/equities/get/equities/v1/ohlcv",
		Summary: "Get OHLCV candle data for a company",
		Params: []routeParam{
			{Name: "ticker", In: "query", Required: true, Type: "string"},
			{Name: "country", In: "query", Required: true, Type: "string"},
			{Name: "timeframe", In: "query", Required: false, Type: "string", Enum: []string{"1D", "7D", "1W", "1M", "3M", "6M", "YTD", "1Y", "5Y", "MAX"}},
		},
	},
	"GET /equities/v1/summary": {
		Method: "GET", Path: "/equities/v1/summary", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Equities", GoMethod: "GetSummary", DocsURL: "https://api-docs.defillama.com/#tag/equities/get/equities/v1/summary",
		Summary: "Get live market summary for a company",
		Params: []routeParam{
			{Name: "ticker", In: "query", Required: true, Type: "string"},
			{Name: "country", In: "query", Required: true, Type: "string"},
		},
	},
	"GET /equities/v1/filings": {
		Method: "GET", Path: "/equities/v1/filings", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Equities", GoMethod: "GetFilings", DocsURL: "https://api-docs.defillama.com/#tag/equities/get/equities/v1/filings",
		Summary: "Get filings for a company",
		Params: []routeParam{
			{Name: "ticker", In: "query", Required: true, Type: "string"},
			{Name: "country", In: "query", Required: true, Type: "string"},
		},
	},
	"GET /equities/v1/onchain": {
		Method: "GET", Path: "/equities/v1/onchain", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "Equities", GoMethod: "GetOnchainMarkets", DocsURL: "https://api-docs.defillama.com/#tag/equities/get/equities/v1/onchain",
		Summary: "Get on-chain tradeable markets for a company",
		Params: []routeParam{
			{Name: "ticker", In: "query", Required: true, Type: "string"},
			{Name: "country", In: "query", Required: true, Type: "string"},
		},
	},
	"GET /pre-ipo/v1/companies-list": {
		Method: "GET", Path: "/pre-ipo/v1/companies-list", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "PreIPO", GoMethod: "GetCompaniesList", DocsURL: "https://api-docs.defillama.com/#tag/pre-ipo/get/pre-ipo/v1/companies-list",
		Summary: "Get list of all tracked pre-IPO companies",
	},
	"GET /pre-ipo/v1/valuations": {
		Method: "GET", Path: "/pre-ipo/v1/valuations", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "PreIPO", GoMethod: "GetValuations", DocsURL: "https://api-docs.defillama.com/#tag/pre-ipo/get/pre-ipo/v1/valuations",
		Summary: "Get valuation history for a pre-IPO company",
		Params: []routeParam{
			{Name: "company", In: "query", Required: true, Type: "string"},
		},
	},
	"GET /pre-ipo/v1/raises": {
		Method: "GET", Path: "/pre-ipo/v1/raises", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "PreIPO", GoMethod: "GetRaises", DocsURL: "https://api-docs.defillama.com/#tag/pre-ipo/get/pre-ipo/v1/raises",
		Summary: "Get funding rounds for a pre-IPO company",
		Params: []routeParam{
			{Name: "company", In: "query", Required: true, Type: "string"},
		},
	},
	"GET /pre-ipo/v1/summary": {
		Method: "GET", Path: "/pre-ipo/v1/summary", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "PreIPO", GoMethod: "GetSummary", DocsURL: "https://api-docs.defillama.com/#tag/pre-ipo/get/pre-ipo/v1/summary",
		Summary: "Get profile and latest valuation for a pre-IPO company",
		Params: []routeParam{
			{Name: "company", In: "query", Required: true, Type: "string"},
		},
	},
	"GET /pre-ipo/v1/integrations": {
		Method: "GET", Path: "/pre-ipo/v1/integrations", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "PreIPO", GoMethod: "GetIntegrations", DocsURL: "https://api-docs.defillama.com/#tag/pre-ipo/get/pre-ipo/v1/integrations",
		Summary: "Get tradeable markets for a pre-IPO company",
		Params: []routeParam{
			{Name: "company", In: "query", Required: true, Type: "string"},
		},
	},
	"GET /rwa/current": {
		Method: "GET", Path: "/rwa/current", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "RWA", GoMethod: "GetCurrentAssets", DocsURL: "https://api-docs.defillama.com/#tag/rwa/get/rwa/current",
		Summary: "List all current RWA assets",
	},
	"GET /rwa/stats": {
		Method: "GET", Path: "/rwa/stats", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "RWA", GoMethod: "GetStats", DocsURL: "https://api-docs.defillama.com/#tag/rwa/get/rwa/stats",
		Summary: "Get aggregate RWA stats by chain, category, platform, and asset group",
	},
	"GET /rwa/list": {
		Method: "GET", Path: "/rwa/list", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "RWA", GoMethod: "GetList", DocsURL: "https://api-docs.defillama.com/#tag/rwa/get/rwa/list",
		Summary: "List RWA ids and filter values",
	},
	"GET /rwa/chain/{chain}": {
		Method: "GET", Path: "/rwa/chain/{chain}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "RWA", GoMethod: "GetAssetsByChain", DocsURL: "https://api-docs.defillama.com/#tag/rwa/get/rwa/chain/%7Bchain%7D",
		Summary: "List current RWA assets on a chain",
		Params: []routeParam{
			{Name: "chain", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /rwa/chart/chain/{chain}": {
		Method: "GET", Path: "/rwa/chart/chain/{chain}", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "RWA", GoMethod: "GetChainChart", DocsURL: "https://api-docs.defillama.com/#tag/rwa/get/rwa/chart/chain/%7Bchain%7D",
		Summary: "Get historical RWA chart data for a chain",
		Params: []routeParam{
			{Name: "chain", In: "path", Required: true, Type: "string"},
		},
	},
	"GET /rwa/chart/chain-breakdown": {
		Method: "GET", Path: "/rwa/chart/chain-breakdown", Tier: "pro", Server: "https://pro-api.llama.fi",
		Service: "RWA", GoMethod: "GetChartChainBreakdown", DocsURL: "https://api-docs.defillama.com/#tag/rwa/get/rwa/chart/chain-breakdown",
		Summary: "Get historical RWA metric breakdown by chain",
		Params: []routeParam{
			{Name: "key", In: "query", Required: false, Type: "string", Enum: []string{"onChainMcap", "activeMcap", "defiActiveTvl"}},
			{Name: "includeStablecoin", In: "query", Required: false, Type: "boolean"},
			{Name: "includeGovernance", In: "query", Required: false, Type: "boolean"},
		},
	},
}
