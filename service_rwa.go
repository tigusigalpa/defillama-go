package defillama

import (
	"context"
	"net/url"
)

// RWAService provides access to the RWA endpoints.
type RWAService struct {
	t *transport
}

// GetCurrentAssets List all current RWA assets
// See: https://api-docs.defillama.com/#tag/rwa/get/rwa/current
func (s *RWAService) GetCurrentAssets(ctx context.Context) ([]RWAAsset, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw []RWAAsset
	if err := s.t.get(ctx, "GET /rwa/current", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetStats Get aggregate RWA stats by chain, category, platform, and asset group
// See: https://api-docs.defillama.com/#tag/rwa/get/rwa/stats
func (s *RWAService) GetStats(ctx context.Context) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /rwa/stats", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetList List RWA ids and filter values
// See: https://api-docs.defillama.com/#tag/rwa/get/rwa/list
func (s *RWAService) GetList(ctx context.Context) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /rwa/list", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetAssetsByChain List current RWA assets on a chain
// See: https://api-docs.defillama.com/#tag/rwa/get/rwa/chain/%7Bchain%7D
func (s *RWAService) GetAssetsByChain(ctx context.Context, chain string) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{
		"chain": chain}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /rwa/chain/{chain}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetChainChart Get historical RWA chart data for a chain
// See: https://api-docs.defillama.com/#tag/rwa/get/rwa/chart/chain/%7Bchain%7D
func (s *RWAService) GetChainChart(ctx context.Context, chain string) ([]any, error) {
	query := url.Values{}
	path := map[string]any{
		"chain": chain}
	var raw []any
	if err := s.t.get(ctx, "GET /rwa/chart/chain/{chain}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetChartChainBreakdown Get historical RWA metric breakdown by chain
// See: https://api-docs.defillama.com/#tag/rwa/get/rwa/chart/chain-breakdown
func (s *RWAService) GetChartChainBreakdown(ctx context.Context, opts *RWAChartOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{}
	var raw []any
	if err := s.t.get(ctx, "GET /rwa/chart/chain-breakdown", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
