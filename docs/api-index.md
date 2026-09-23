# DefiLlama API — endpoint index

Generated from `spec/defillama-api.json` (132 GET operations). Each row maps an official documentation page to the SDK route and public method.

| # | Tier | Tag | Operation | Go method | Documentation |
|---|------|-----|-----------|----------|---------------|
| 1 | free | TVL | `GET /protocols` | `TVLService.GetProtocols()` | [ref](https://api-docs.defillama.com/#tag/tvl/get/protocols) |
| 2 | free | TVL | `GET /protocol/{protocol}` | `TVLService.GetProtocol()` | [ref](https://api-docs.defillama.com/#tag/tvl/get/protocol/%7Bprotocol%7D) |
| 3 | free | TVL | `GET /v2/historicalChainTvl` | `TVLService.GetHistoricalChainTVL()` | [ref](https://api-docs.defillama.com/#tag/tvl/get/v2/historicalChainTvl) |
| 4 | free | TVL | `GET /v2/historicalChainTvl/{chain}` | `TVLService.GetHistoricalChainTVLByChain()` | [ref](https://api-docs.defillama.com/#tag/tvl/get/v2/historicalChainTvl/%7Bchain%7D) |
| 5 | free | TVL | `GET /tvl/{protocol}` | `TVLService.GetTVL()` | [ref](https://api-docs.defillama.com/#tag/tvl/get/tvl/%7Bprotocol%7D) |
| 6 | free | TVL | `GET /v2/chains` | `TVLService.GetChains()` | [ref](https://api-docs.defillama.com/#tag/tvl/get/v2/chains) |
| 7 | free | coins | `GET /prices/current/{coins}` | `PricesService.GetCurrentPrices()` | [ref](https://api-docs.defillama.com/#tag/coins/get/prices/current/%7Bcoins%7D) |
| 8 | free | coins | `GET /prices/historical/{timestamp}/{coins}` | `PricesService.GetHistoricalPrices()` | [ref](https://api-docs.defillama.com/#tag/coins/get/prices/historical/%7Btimestamp%7D/%7Bcoins%7D) |
| 9 | free | coins | `GET /batchHistorical` | `PricesService.GetBatchHistoricalPrices()` | [ref](https://api-docs.defillama.com/#tag/coins/get/batchHistorical) |
| 10 | free | coins | `GET /chart/{coins}` | `PricesService.GetChart()` | [ref](https://api-docs.defillama.com/#tag/coins/get/chart/%7Bcoins%7D) |
| 11 | free | coins | `GET /percentage/{coins}` | `PricesService.GetPercentageChange()` | [ref](https://api-docs.defillama.com/#tag/coins/get/percentage/%7Bcoins%7D) |
| 12 | free | coins | `GET /prices/first/{coins}` | `PricesService.GetFirstPrices()` | [ref](https://api-docs.defillama.com/#tag/coins/get/prices/first/%7Bcoins%7D) |
| 13 | free | coins | `GET /block/{chain}/{timestamp}` | `PricesService.GetBlockAtTimestamp()` | [ref](https://api-docs.defillama.com/#tag/coins/get/block/%7Bchain%7D/%7Btimestamp%7D) |
| 14 | free | stablecoins | `GET /stablecoins` | `StablecoinsService.GetStablecoins()` | [ref](https://api-docs.defillama.com/#tag/stablecoins/get/stablecoins) |
| 15 | free | stablecoins | `GET /stablecoincharts/all` | `StablecoinsService.GetStablecoinCharts()` | [ref](https://api-docs.defillama.com/#tag/stablecoins/get/stablecoincharts/all) |
| 16 | free | stablecoins | `GET /stablecoincharts/{chain}` | `StablecoinsService.GetStablecoinChartsByChain()` | [ref](https://api-docs.defillama.com/#tag/stablecoins/get/stablecoincharts/%7Bchain%7D) |
| 17 | free | stablecoins | `GET /stablecoin/{asset}` | `StablecoinsService.GetStablecoin()` | [ref](https://api-docs.defillama.com/#tag/stablecoins/get/stablecoin/%7Basset%7D) |
| 18 | free | stablecoins | `GET /stablecoinchains` | `StablecoinsService.GetStablecoinChains()` | [ref](https://api-docs.defillama.com/#tag/stablecoins/get/stablecoinchains) |
| 19 | free | stablecoins | `GET /stablecoinprices` | `StablecoinsService.GetStablecoinPrices()` | [ref](https://api-docs.defillama.com/#tag/stablecoins/get/stablecoinprices) |
| 20 | free | yields | `GET /pools` | `YieldsService.GetPools()` | [ref](https://api-docs.defillama.com/#tag/yields/get/pools) |
| 21 | free | yields | `GET /chart/{pool}` | `YieldsService.GetPoolChart()` | [ref](https://api-docs.defillama.com/#tag/yields/get/chart/%7Bpool%7D) |
| 22 | free | volumes | `GET /overview/dexs` | `VolumesService.GetDEXOverview()` | [ref](https://api-docs.defillama.com/#tag/volumes/get/overview/dexs) |
| 23 | free | volumes | `GET /overview/dexs/{chain}` | `VolumesService.GetDEXOverviewByChain()` | [ref](https://api-docs.defillama.com/#tag/volumes/get/overview/dexs/%7Bchain%7D) |
| 24 | free | volumes | `GET /summary/dexs/{protocol}` | `VolumesService.GetDEXSummary()` | [ref](https://api-docs.defillama.com/#tag/volumes/get/summary/dexs/%7Bprotocol%7D) |
| 25 | free | volumes | `GET /overview/options` | `VolumesService.GetOptionsOverview()` | [ref](https://api-docs.defillama.com/#tag/volumes/get/overview/options) |
| 26 | free | volumes | `GET /overview/options/{chain}` | `VolumesService.GetOptionsOverviewByChain()` | [ref](https://api-docs.defillama.com/#tag/volumes/get/overview/options/%7Bchain%7D) |
| 27 | free | volumes | `GET /summary/options/{protocol}` | `VolumesService.GetOptionsSummary()` | [ref](https://api-docs.defillama.com/#tag/volumes/get/summary/options/%7Bprotocol%7D) |
| 28 | free | perps | `GET /overview/open-interest` | `VolumesService.GetOpenInterestOverview()` | [ref](https://api-docs.defillama.com/#tag/perps/get/overview/open-interest) |
| 29 | free | fees and revenue | `GET /overview/fees` | `FeesService.GetOverview()` | [ref](https://api-docs.defillama.com/#tag/fees-and-revenue/get/overview/fees) |
| 30 | free | fees and revenue | `GET /overview/fees/{chain}` | `FeesService.GetOverviewByChain()` | [ref](https://api-docs.defillama.com/#tag/fees-and-revenue/get/overview/fees/%7Bchain%7D) |
| 31 | free | fees and revenue | `GET /summary/fees/{protocol}` | `FeesService.GetSummary()` | [ref](https://api-docs.defillama.com/#tag/fees-and-revenue/get/summary/fees/%7Bprotocol%7D) |
| 32 | pro | TVL | `GET /api/tokenProtocols/{symbol}` | `TVLService.GetTokenProtocols()` | [ref](https://api-docs.defillama.com/#tag/tvl/get/api/tokenProtocols/%7Bsymbol%7D) |
| 33 | pro | TVL | `GET /api/inflows/{protocol}/{timestamp}` | `TVLService.GetInflows()` | [ref](https://api-docs.defillama.com/#tag/tvl/get/api/inflows/%7Bprotocol%7D/%7Btimestamp%7D) |
| 34 | pro | TVL | `GET /api/chainAssets` | `TVLService.GetChainAssets()` | [ref](https://api-docs.defillama.com/#tag/tvl/get/api/chainAssets) |
| 35 | pro | Unlocks | `GET /api/emissions` | `EmissionsService.GetEmissions()` | [ref](https://api-docs.defillama.com/#tag/unlocks/get/api/emissions) |
| 36 | pro | Unlocks | `GET /api/emission/{protocol}` | `EmissionsService.GetEmission()` | [ref](https://api-docs.defillama.com/#tag/unlocks/get/api/emission/%7Bprotocol%7D) |
| 37 | pro | main page | `GET /api/categories` | `EcosystemService.GetCategories()` | [ref](https://api-docs.defillama.com/#tag/main-page/get/api/categories) |
| 38 | pro | main page | `GET /api/forks` | `EcosystemService.GetForks()` | [ref](https://api-docs.defillama.com/#tag/main-page/get/api/forks) |
| 39 | pro | main page | `GET /api/oracles` | `EcosystemService.GetOracles()` | [ref](https://api-docs.defillama.com/#tag/main-page/get/api/oracles) |
| 40 | pro | main page | `GET /api/hacks` | `EcosystemService.GetHacks()` | [ref](https://api-docs.defillama.com/#tag/main-page/get/api/hacks) |
| 41 | pro | main page | `GET /api/raises` | `EcosystemService.GetRaises()` | [ref](https://api-docs.defillama.com/#tag/main-page/get/api/raises) |
| 42 | pro | main page | `GET /api/treasuries` | `EcosystemService.GetTreasuries()` | [ref](https://api-docs.defillama.com/#tag/main-page/get/api/treasuries) |
| 43 | pro | main page | `GET /api/entities` | `EcosystemService.GetEntities()` | [ref](https://api-docs.defillama.com/#tag/main-page/get/api/entities) |
| 44 | pro | token liquidity | `GET /api/historicalLiquidity/{token}` | `PricesService.GetHistoricalLiquidity()` | [ref](https://api-docs.defillama.com/#tag/token-liquidity/get/api/historicalLiquidity/%7Btoken%7D) |
| 45 | pro | perps | `GET /api/overview/derivatives` | `VolumesService.GetDerivativesOverview()` | [ref](https://api-docs.defillama.com/#tag/perps/get/api/overview/derivatives) |
| 46 | pro | perps | `GET /api/summary/derivatives/{protocol}` | `VolumesService.GetDerivativesSummary()` | [ref](https://api-docs.defillama.com/#tag/perps/get/api/summary/derivatives/%7Bprotocol%7D) |
| 47 | pro | TVL | `GET /api/v2/metrics/tvl/protocol/{protocol}` | `TVLService.GetProtocolTVLMetrics()` | [ref](https://api-docs.defillama.com/#tag/tvl/get/api/v2/metrics/tvl/protocol/%7Bprotocol%7D) |
| 48 | pro | TVL | `GET /api/v2/chart/tvl/protocol/{protocol}` | `TVLService.GetProtocolTVLChart()` | [ref](https://api-docs.defillama.com/#tag/tvl/get/api/v2/chart/tvl/protocol/%7Bprotocol%7D) |
| 49 | pro | TVL | `GET /api/v2/chart/tvl/protocol/{protocol}/chain-breakdown` | `TVLService.GetProtocolTVLChartChainBreakdown()` | [ref](https://api-docs.defillama.com/#tag/tvl/get/api/v2/chart/tvl/protocol/%7Bprotocol%7D/chain-breakdown) |
| 50 | pro | TVL | `GET /api/v2/chart/tvl/protocol/{protocol}/token-breakdown` | `TVLService.GetProtocolTVLChartTokenBreakdown()` | [ref](https://api-docs.defillama.com/#tag/tvl/get/api/v2/chart/tvl/protocol/%7Bprotocol%7D/token-breakdown) |
| 51 | pro | Treasury | `GET /api/v2/metrics/treasury/protocol/{protocol}` | `TreasuryService.GetTreasuryMetrics()` | [ref](https://api-docs.defillama.com/#tag/treasury/get/api/v2/metrics/treasury/protocol/%7Bprotocol%7D) |
| 52 | pro | Treasury | `GET /api/v2/chart/treasury/protocol/{protocol}` | `TreasuryService.GetTreasuryChart()` | [ref](https://api-docs.defillama.com/#tag/treasury/get/api/v2/chart/treasury/protocol/%7Bprotocol%7D) |
| 53 | pro | Treasury | `GET /api/v2/chart/treasury/protocol/{protocol}/chain-breakdown` | `TreasuryService.GetTreasuryChartChainBreakdown()` | [ref](https://api-docs.defillama.com/#tag/treasury/get/api/v2/chart/treasury/protocol/%7Bprotocol%7D/chain-breakdown) |
| 54 | pro | Treasury | `GET /api/v2/chart/treasury/protocol/{protocol}/token-breakdown` | `TreasuryService.GetTreasuryChartTokenBreakdown()` | [ref](https://api-docs.defillama.com/#tag/treasury/get/api/v2/chart/treasury/protocol/%7Bprotocol%7D/token-breakdown) |
| 55 | pro | Oracles | `GET /api/v2/metrics/oracle` | `OraclesService.GetOracleMetrics()` | [ref](https://api-docs.defillama.com/#tag/oracles/get/api/v2/metrics/oracle) |
| 56 | pro | Oracles | `GET /api/v2/chart/oracle` | `OraclesService.GetOracleChart()` | [ref](https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle) |
| 57 | pro | Oracles | `GET /api/v2/chart/oracle/chain-breakdown` | `OraclesService.GetOracleChartChainBreakdown()` | [ref](https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle/chain-breakdown) |
| 58 | pro | Oracles | `GET /api/v2/chart/oracle/protocol-breakdown` | `OraclesService.GetOracleChartProtocolBreakdown()` | [ref](https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle/protocol-breakdown) |
| 59 | pro | Oracles | `GET /api/v2/chart/oracle/protocol/{protocol}` | `OraclesService.GetOracleProtocolChart()` | [ref](https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle/protocol/%7Bprotocol%7D) |
| 60 | pro | Oracles | `GET /api/v2/chart/oracle/protocol/{protocol}/chain-breakdown` | `OraclesService.GetOracleProtocolChartChainBreakdown()` | [ref](https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle/protocol/%7Bprotocol%7D/chain-breakdown) |
| 61 | pro | Oracles | `GET /api/v2/chart/oracle/chain/{chain}` | `OraclesService.GetOracleChainChart()` | [ref](https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle/chain/%7Bchain%7D) |
| 62 | pro | Oracles | `GET /api/v2/chart/oracle/chain/{chain}/protocol-breakdown` | `OraclesService.GetOracleChainChartProtocolBreakdown()` | [ref](https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle/chain/%7Bchain%7D/protocol-breakdown) |
| 63 | pro | Forks | `GET /api/v2/metrics/fork` | `ForksService.GetForkMetrics()` | [ref](https://api-docs.defillama.com/#tag/forks/get/api/v2/metrics/fork) |
| 64 | pro | Forks | `GET /api/v2/chart/fork/protocol-breakdown` | `ForksService.GetForkChartProtocolBreakdown()` | [ref](https://api-docs.defillama.com/#tag/forks/get/api/v2/chart/fork/protocol-breakdown) |
| 65 | pro | Forks | `GET /api/v2/chart/fork/protocol/{protocol}` | `ForksService.GetForkProtocolChart()` | [ref](https://api-docs.defillama.com/#tag/forks/get/api/v2/chart/fork/protocol/%7Bprotocol%7D) |
| 66 | pro | Dimensions | `GET /api/v2/metrics/{metric}` | `DimensionsService.GetMetrics()` | [ref](https://api-docs.defillama.com/#tag/dimensions/get/api/v2/metrics/%7Bmetric%7D) |
| 67 | pro | Dimensions | `GET /api/v2/chart/{metric}` | `DimensionsService.GetChart()` | [ref](https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D) |
| 68 | pro | Dimensions | `GET /api/v2/chart/{metric}/chain-breakdown` | `DimensionsService.GetChartChainBreakdown()` | [ref](https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/chain-breakdown) |
| 69 | pro | Dimensions | `GET /api/v2/chart/{metric}/protocol-breakdown` | `DimensionsService.GetChartProtocolBreakdown()` | [ref](https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/protocol-breakdown) |
| 70 | pro | Dimensions | `GET /api/v2/metrics/{metric}/chain/{chain}` | `DimensionsService.GetChainMetrics()` | [ref](https://api-docs.defillama.com/#tag/dimensions/get/api/v2/metrics/%7Bmetric%7D/chain/%7Bchain%7D) |
| 71 | pro | Dimensions | `GET /api/v2/chart/{metric}/chain/{chain}` | `DimensionsService.GetChainChart()` | [ref](https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/chain/%7Bchain%7D) |
| 72 | pro | Dimensions | `GET /api/v2/chart/{metric}/chain/{chain}/protocol-breakdown` | `DimensionsService.GetChainChartProtocolBreakdown()` | [ref](https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/chain/%7Bchain%7D/protocol-breakdown) |
| 73 | pro | Dimensions | `GET /api/v2/metrics/{metric}/protocol/{protocol}` | `DimensionsService.GetProtocolMetrics()` | [ref](https://api-docs.defillama.com/#tag/dimensions/get/api/v2/metrics/%7Bmetric%7D/protocol/%7Bprotocol%7D) |
| 74 | pro | Dimensions | `GET /api/v2/chart/{metric}/protocol/{protocol}` | `DimensionsService.GetProtocolChart()` | [ref](https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/protocol/%7Bprotocol%7D) |
| 75 | pro | Dimensions | `GET /api/v2/chart/{metric}/protocol/{protocol}/chain-breakdown` | `DimensionsService.GetProtocolChartChainBreakdown()` | [ref](https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/protocol/%7Bprotocol%7D/chain-breakdown) |
| 76 | pro | Dimensions | `GET /api/v2/chart/{metric}/protocol/{protocol}/version-breakdown` | `DimensionsService.GetProtocolChartVersionBreakdown()` | [ref](https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/protocol/%7Bprotocol%7D/version-breakdown) |
| 77 | pro | Dimensions | `GET /api/v2/chart/{metric}/protocol/{protocol}/label-breakdown` | `DimensionsService.GetProtocolChartLabelBreakdown()` | [ref](https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/protocol/%7Bprotocol%7D/label-breakdown) |
| 78 | pro | Dimensions | `GET /api/v2/metrics/{metric}/category/{category}` | `DimensionsService.GetCategoryMetrics()` | [ref](https://api-docs.defillama.com/#tag/dimensions/get/api/v2/metrics/%7Bmetric%7D/category/%7Bcategory%7D) |
| 79 | pro | Dimensions | `GET /api/v2/chart/{metric}/category/{category}` | `DimensionsService.GetCategoryChart()` | [ref](https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/category/%7Bcategory%7D) |
| 80 | pro | Dimensions | `GET /api/v2/chart/{metric}/category/{category}/chain-breakdown` | `DimensionsService.GetCategoryChartChainBreakdown()` | [ref](https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/category/%7Bcategory%7D/chain-breakdown) |
| 81 | pro | Dimensions | `GET /api/v2/chart/{metric}/category/{category}/protocol-breakdown` | `DimensionsService.GetCategoryChartProtocolBreakdown()` | [ref](https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/category/%7Bcategory%7D/protocol-breakdown) |
| 82 | pro | Dimensions | `GET /api/v2/metrics/{metric}/category/{category}/chain/{chain}` | `DimensionsService.GetCategoryChainMetrics()` | [ref](https://api-docs.defillama.com/#tag/dimensions/get/api/v2/metrics/%7Bmetric%7D/category/%7Bcategory%7D/chain/%7Bchain%7D) |
| 83 | pro | Dimensions | `GET /api/v2/chart/{metric}/category/{category}/chain/{chain}` | `DimensionsService.GetCategoryChainChart()` | [ref](https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/category/%7Bcategory%7D/chain/%7Bchain%7D) |
| 84 | pro | Dimensions | `GET /api/v2/chart/{metric}/category/{category}/chain/{chain}/protocol-breakdown` | `DimensionsService.GetCategoryChainChartProtocolBreakdown()` | [ref](https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/category/%7Bcategory%7D/chain/%7Bchain%7D/protocol-breakdown) |
| 85 | pro | Financial Statements | `GET /api/v2/metrics/financial-statement/protocol/{protocol}` | `FinancialStatementsService.GetIncomeStatement()` | [ref](https://api-docs.defillama.com/#tag/financial-statements/get/api/v2/metrics/financial-statement/protocol/%7Bprotocol%7D) |
| 86 | pro | stablecoins | `GET /stablecoins/stablecoindominance/{chain}` | `StablecoinsService.GetStablecoinDominance()` | [ref](https://api-docs.defillama.com/#tag/stablecoins/get/stablecoins/stablecoindominance/%7Bchain%7D) |
| 87 | pro | yields | `GET /yields/v1/pools` | `YieldsService.GetLegacyPools()` | [ref](https://api-docs.defillama.com/#tag/yields/get/yields/v1/pools) |
| 88 | pro | yields | `GET /yields/v1/chart/{pool}` | `YieldsService.GetLegacyPoolChart()` | [ref](https://api-docs.defillama.com/#tag/yields/get/yields/v1/chart/%7Bpool%7D) |
| 89 | pro | yields | `GET /yields/v1/poolsBorrow` | `YieldsService.GetLegacyBorrowPools()` | [ref](https://api-docs.defillama.com/#tag/yields/get/yields/v1/poolsBorrow) |
| 90 | pro | yields | `GET /yields/v1/chartLendBorrow/{pool}` | `YieldsService.GetLegacyLendBorrowChart()` | [ref](https://api-docs.defillama.com/#tag/yields/get/yields/v1/chartLendBorrow/%7Bpool%7D) |
| 91 | pro | yields | `GET /yields/v2/earn` | `YieldsService.GetEarnPools()` | [ref](https://api-docs.defillama.com/#tag/yields/get/yields/v2/earn) |
| 92 | pro | yields | `GET /yields/v2/earn/query` | `YieldsService.QueryEarnPools()` | [ref](https://api-docs.defillama.com/#tag/yields/get/yields/v2/earn/query) |
| 93 | pro | yields | `GET /yields/v2/earn/{id}/history` | `YieldsService.GetEarnPoolHistory()` | [ref](https://api-docs.defillama.com/#tag/yields/get/yields/v2/earn/%7Bid%7D/history) |
| 94 | pro | yields | `GET /yields/v2/earn/verified` | `YieldsService.GetVerifiedEarnPools()` | [ref](https://api-docs.defillama.com/#tag/yields/get/yields/v2/earn/verified) |
| 95 | pro | yields | `GET /yields/v2/earn/verified/{id}` | `YieldsService.GetVerifiedEarnPool()` | [ref](https://api-docs.defillama.com/#tag/yields/get/yields/v2/earn/verified/%7Bid%7D) |
| 96 | pro | yields | `GET /yields/v2/earn/verified/{id}/history` | `YieldsService.GetVerifiedEarnPoolHistory()` | [ref](https://api-docs.defillama.com/#tag/yields/get/yields/v2/earn/verified/%7Bid%7D/history) |
| 97 | pro | yields | `GET /yields/v2/borrow/markets` | `YieldsService.GetBorrowMarkets()` | [ref](https://api-docs.defillama.com/#tag/yields/get/yields/v2/borrow/markets) |
| 98 | pro | yields | `GET /yields/v2/borrow/markets/{id}/history` | `YieldsService.GetBorrowMarketHistory()` | [ref](https://api-docs.defillama.com/#tag/yields/get/yields/v2/borrow/markets/%7Bid%7D/history) |
| 99 | pro | yields | `GET /yields/v2/borrow/routes` | `YieldsService.GetBorrowRoutes()` | [ref](https://api-docs.defillama.com/#tag/yields/get/yields/v2/borrow/routes) |
| 100 | pro | yields | `GET /yields/v2/loops` | `YieldsService.GetLoopStrategies()` | [ref](https://api-docs.defillama.com/#tag/yields/get/yields/v2/loops) |
| 101 | pro | yields | `GET /yields/perps` | `YieldsService.GetPerpFundingRates()` | [ref](https://api-docs.defillama.com/#tag/yields/get/yields/perps) |
| 102 | pro | yields | `GET /api/lstRates` | `YieldsService.GetLSTRates()` | [ref](https://api-docs.defillama.com/#tag/yields/get/api/lstRates) |
| 103 | pro | ETFs | `GET /etfs/snapshot` | `ETFsService.GetSnapshot()` | [ref](https://api-docs.defillama.com/#tag/etfs/get/etfs/snapshot) |
| 104 | pro | ETFs | `GET /etfs/flows` | `ETFsService.GetFlows()` | [ref](https://api-docs.defillama.com/#tag/etfs/get/etfs/flows) |
| 105 | pro | narratives | `GET /fdv/performance/{period}` | `NarrativesService.GetPerformance()` | [ref](https://api-docs.defillama.com/#tag/narratives/get/fdv/performance/%7Bperiod%7D) |
| 106 | pro | bridges | `GET /bridges/bridges` | `BridgesService.GetBridges()` | [ref](https://api-docs.defillama.com/#tag/bridges/get/bridges/bridges) |
| 107 | pro | bridges | `GET /bridges/bridge/{id}` | `BridgesService.GetBridge()` | [ref](https://api-docs.defillama.com/#tag/bridges/get/bridges/bridge/%7Bid%7D) |
| 108 | pro | bridges | `GET /bridges/bridgevolume/{chain}` | `BridgesService.GetBridgeVolume()` | [ref](https://api-docs.defillama.com/#tag/bridges/get/bridges/bridgevolume/%7Bchain%7D) |
| 109 | pro | bridges | `GET /bridges/bridgedaystats/{timestamp}/{chain}` | `BridgesService.GetBridgeDayStats()` | [ref](https://api-docs.defillama.com/#tag/bridges/get/bridges/bridgedaystats/%7Btimestamp%7D/%7Bchain%7D) |
| 110 | pro | bridges | `GET /bridges/transactions/{id}` | `BridgesService.GetBridgeTransactions()` | [ref](https://api-docs.defillama.com/#tag/bridges/get/bridges/transactions/%7Bid%7D) |
| 111 | pro | meta | `GET /usage/APIKEY` | `AccountService.GetUsage()` | [ref](https://api-docs.defillama.com/#tag/meta/get/usage/APIKEY) |
| 112 | pro | DAT | `GET /dat/institutions` | `DATService.GetInstitutions()` | [ref](https://api-docs.defillama.com/#tag/dat/get/dat/institutions) |
| 113 | pro | DAT | `GET /dat/institutions/{symbol}` | `DATService.GetInstitution()` | [ref](https://api-docs.defillama.com/#tag/dat/get/dat/institutions/%7Bsymbol%7D) |
| 114 | pro | Equities | `GET /equities/v1/companies-list` | `EquitiesService.GetCompaniesList()` | [ref](https://api-docs.defillama.com/#tag/equities/get/equities/v1/companies-list) |
| 115 | pro | Equities | `GET /equities/v1/statements` | `EquitiesService.GetStatements()` | [ref](https://api-docs.defillama.com/#tag/equities/get/equities/v1/statements) |
| 116 | pro | Equities | `GET /equities/v1/dimensions` | `EquitiesService.GetDimensions()` | [ref](https://api-docs.defillama.com/#tag/equities/get/equities/v1/dimensions) |
| 117 | pro | Equities | `GET /equities/v1/price-history` | `EquitiesService.GetPriceHistory()` | [ref](https://api-docs.defillama.com/#tag/equities/get/equities/v1/price-history) |
| 118 | pro | Equities | `GET /equities/v1/ohlcv` | `EquitiesService.GetOHLCV()` | [ref](https://api-docs.defillama.com/#tag/equities/get/equities/v1/ohlcv) |
| 119 | pro | Equities | `GET /equities/v1/summary` | `EquitiesService.GetSummary()` | [ref](https://api-docs.defillama.com/#tag/equities/get/equities/v1/summary) |
| 120 | pro | Equities | `GET /equities/v1/filings` | `EquitiesService.GetFilings()` | [ref](https://api-docs.defillama.com/#tag/equities/get/equities/v1/filings) |
| 121 | pro | Equities | `GET /equities/v1/onchain` | `EquitiesService.GetOnchainMarkets()` | [ref](https://api-docs.defillama.com/#tag/equities/get/equities/v1/onchain) |
| 122 | pro | Pre-IPO | `GET /pre-ipo/v1/companies-list` | `PreIPOService.GetCompaniesList()` | [ref](https://api-docs.defillama.com/#tag/pre-ipo/get/pre-ipo/v1/companies-list) |
| 123 | pro | Pre-IPO | `GET /pre-ipo/v1/valuations` | `PreIPOService.GetValuations()` | [ref](https://api-docs.defillama.com/#tag/pre-ipo/get/pre-ipo/v1/valuations) |
| 124 | pro | Pre-IPO | `GET /pre-ipo/v1/raises` | `PreIPOService.GetRaises()` | [ref](https://api-docs.defillama.com/#tag/pre-ipo/get/pre-ipo/v1/raises) |
| 125 | pro | Pre-IPO | `GET /pre-ipo/v1/summary` | `PreIPOService.GetSummary()` | [ref](https://api-docs.defillama.com/#tag/pre-ipo/get/pre-ipo/v1/summary) |
| 126 | pro | Pre-IPO | `GET /pre-ipo/v1/integrations` | `PreIPOService.GetIntegrations()` | [ref](https://api-docs.defillama.com/#tag/pre-ipo/get/pre-ipo/v1/integrations) |
| 127 | pro | RWA | `GET /rwa/current` | `RWAService.GetCurrentAssets()` | [ref](https://api-docs.defillama.com/#tag/rwa/get/rwa/current) |
| 128 | pro | RWA | `GET /rwa/stats` | `RWAService.GetStats()` | [ref](https://api-docs.defillama.com/#tag/rwa/get/rwa/stats) |
| 129 | pro | RWA | `GET /rwa/list` | `RWAService.GetList()` | [ref](https://api-docs.defillama.com/#tag/rwa/get/rwa/list) |
| 130 | pro | RWA | `GET /rwa/chain/{chain}` | `RWAService.GetAssetsByChain()` | [ref](https://api-docs.defillama.com/#tag/rwa/get/rwa/chain/%7Bchain%7D) |
| 131 | pro | RWA | `GET /rwa/chart/chain/{chain}` | `RWAService.GetChainChart()` | [ref](https://api-docs.defillama.com/#tag/rwa/get/rwa/chart/chain/%7Bchain%7D) |
| 132 | pro | RWA | `GET /rwa/chart/chain-breakdown` | `RWAService.GetChartChainBreakdown()` | [ref](https://api-docs.defillama.com/#tag/rwa/get/rwa/chart/chain-breakdown) |
