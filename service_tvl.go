package defillama

import (
	"context"
	"net/url"
)

// TVLService provides access to the TVL endpoints.
type TVLService struct {
	t *transport
}

// GetProtocols List all protocols on defillama along with their tvl
// See: https://api-docs.defillama.com/#tag/tvl/get/protocols
func (s *TVLService) GetProtocols(ctx context.Context) ([]Protocol, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw []Protocol
	if err := s.t.get(ctx, "GET /protocols", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetProtocol Get historical TVL of a protocol and breakdowns by token and chain
// See: https://api-docs.defillama.com/#tag/tvl/get/protocol/%7Bprotocol%7D
func (s *TVLService) GetProtocol(ctx context.Context, protocol string) (*ProtocolDetails, error) {
	query := url.Values{}
	path := map[string]any{
		"protocol": protocol}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /protocol/{protocol}", path, query, &raw); err != nil {
		return nil, err
	}
	return &ProtocolDetails{Raw: raw}, nil
}

// GetHistoricalChainTVL Get historical TVL (excludes liquid staking and double counted tvl) of DeFi on all chains
// See: https://api-docs.defillama.com/#tag/tvl/get/v2/historicalChainTvl
func (s *TVLService) GetHistoricalChainTVL(ctx context.Context) ([]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw []any
	if err := s.t.get(ctx, "GET /v2/historicalChainTvl", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetHistoricalChainTVLByChain Get historical TVL (excludes liquid staking and double counted tvl) of a chain
// See: https://api-docs.defillama.com/#tag/tvl/get/v2/historicalChainTvl/%7Bchain%7D
func (s *TVLService) GetHistoricalChainTVLByChain(ctx context.Context, chain string) ([]any, error) {
	query := url.Values{}
	path := map[string]any{
		"chain": chain}
	var raw []any
	if err := s.t.get(ctx, "GET /v2/historicalChainTvl/{chain}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetTVL Simplified endpoint to get current TVL of a protocol
// See: https://api-docs.defillama.com/#tag/tvl/get/tvl/%7Bprotocol%7D
func (s *TVLService) GetTVL(ctx context.Context, protocol string) (float64, error) {
	query := url.Values{}
	path := map[string]any{
		"protocol": protocol}
	var raw float64
	if err := s.t.get(ctx, "GET /tvl/{protocol}", path, query, &raw); err != nil {
		return 0, err
	}
	return raw, nil
}

// GetChains Get current TVL of all chains
// See: https://api-docs.defillama.com/#tag/tvl/get/v2/chains
func (s *TVLService) GetChains(ctx context.Context) ([]Chain, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw []Chain
	if err := s.t.get(ctx, "GET /v2/chains", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetTokenProtocols Lists the amount of a certain token within all protocols. Data for the Token Usage page
// See: https://api-docs.defillama.com/#tag/tvl/get/api/tokenProtocols/%7Bsymbol%7D
func (s *TVLService) GetTokenProtocols(ctx context.Context, symbol string) ([]any, error) {
	query := url.Values{}
	path := map[string]any{
		"symbol": symbol}
	var raw []any
	if err := s.t.get(ctx, "GET /api/tokenProtocols/{symbol}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetInflows Lists the amount of inflows and outflows for a protocol at a given date
// See: https://api-docs.defillama.com/#tag/tvl/get/api/inflows/%7Bprotocol%7D/%7Btimestamp%7D
func (s *TVLService) GetInflows(ctx context.Context, protocol string, timestamp int64) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{
		"protocol":  protocol,
		"timestamp": timestamp}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /api/inflows/{protocol}/{timestamp}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetChainAssets Get assets of all chains
// See: https://api-docs.defillama.com/#tag/tvl/get/api/chainAssets
func (s *TVLService) GetChainAssets(ctx context.Context) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /api/chainAssets", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetProtocolTVLMetrics Get aggregate TVL metrics for a protocol
// See: https://api-docs.defillama.com/#tag/tvl/get/api/v2/metrics/tvl/protocol/%7Bprotocol%7D
func (s *TVLService) GetProtocolTVLMetrics(ctx context.Context, protocol string) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{
		"protocol": protocol}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /api/v2/metrics/tvl/protocol/{protocol}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetProtocolTVLChart Get historical TVL chart for a protocol
// See: https://api-docs.defillama.com/#tag/tvl/get/api/v2/chart/tvl/protocol/%7Bprotocol%7D
func (s *TVLService) GetProtocolTVLChart(ctx context.Context, protocol string, opts *TVLChartOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"protocol": protocol}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/tvl/protocol/{protocol}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetProtocolTVLChartChainBreakdown Get historical TVL chart for a protocol broken down by chain
// See: https://api-docs.defillama.com/#tag/tvl/get/api/v2/chart/tvl/protocol/%7Bprotocol%7D/chain-breakdown
func (s *TVLService) GetProtocolTVLChartChainBreakdown(ctx context.Context, protocol string, opts *TVLChartOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"protocol": protocol}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/tvl/protocol/{protocol}/chain-breakdown", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetProtocolTVLChartTokenBreakdown Get historical TVL chart for a protocol broken down by token
// See: https://api-docs.defillama.com/#tag/tvl/get/api/v2/chart/tvl/protocol/%7Bprotocol%7D/token-breakdown
func (s *TVLService) GetProtocolTVLChartTokenBreakdown(ctx context.Context, protocol string, opts *TVLChartOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"protocol": protocol}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/tvl/protocol/{protocol}/token-breakdown", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
