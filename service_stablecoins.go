package defillama

import (
	"context"
	"net/url"
)

// StablecoinsService provides access to the Stablecoins endpoints.
type StablecoinsService struct {
	t *transport
}

// GetStablecoins List all stablecoins along with their circulating amounts
// See: https://api-docs.defillama.com/#tag/stablecoins/get/stablecoins
func (s *StablecoinsService) GetStablecoins(ctx context.Context, opts *StablecoinsListOptions) ([]Stablecoin, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{}
	var raw struct {
		PeggedAssets []Stablecoin `json:"peggedAssets"`
	}
	if err := s.t.get(ctx, "GET /stablecoins", path, query, &raw); err != nil {
		return nil, err
	}
	return raw.PeggedAssets, nil
}

// GetStablecoinCharts Get historical mcap sum of all stablecoins
// See: https://api-docs.defillama.com/#tag/stablecoins/get/stablecoincharts/all
func (s *StablecoinsService) GetStablecoinCharts(ctx context.Context, opts *StablecoinChartOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{}
	var raw []any
	if err := s.t.get(ctx, "GET /stablecoincharts/all", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetStablecoinChartsByChain Get historical mcap sum of all stablecoins in a chain
// See: https://api-docs.defillama.com/#tag/stablecoins/get/stablecoincharts/%7Bchain%7D
func (s *StablecoinsService) GetStablecoinChartsByChain(ctx context.Context, chain string, opts *StablecoinChartOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"chain": chain}
	var raw []any
	if err := s.t.get(ctx, "GET /stablecoincharts/{chain}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetStablecoin Get historical mcap and historical chain distribution of a stablecoin
// See: https://api-docs.defillama.com/#tag/stablecoins/get/stablecoin/%7Basset%7D
func (s *StablecoinsService) GetStablecoin(ctx context.Context, asset int64) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{
		"asset": asset}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /stablecoin/{asset}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetStablecoinChains Get current mcap sum of all stablecoins on each chain
// See: https://api-docs.defillama.com/#tag/stablecoins/get/stablecoinchains
func (s *StablecoinsService) GetStablecoinChains(ctx context.Context) ([]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw []any
	if err := s.t.get(ctx, "GET /stablecoinchains", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetStablecoinPrices Get historical prices of all stablecoins
// See: https://api-docs.defillama.com/#tag/stablecoins/get/stablecoinprices
func (s *StablecoinsService) GetStablecoinPrices(ctx context.Context) ([]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw []any
	if err := s.t.get(ctx, "GET /stablecoinprices", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetStablecoinDominance Get stablecoin dominance per chain along with the info about the larges coin in a chain
// See: https://api-docs.defillama.com/#tag/stablecoins/get/stablecoins/stablecoindominance/%7Bchain%7D
func (s *StablecoinsService) GetStablecoinDominance(ctx context.Context, chain string, opts *StablecoinChartOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"chain": chain}
	var raw []any
	if err := s.t.get(ctx, "GET /stablecoins/stablecoindominance/{chain}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
