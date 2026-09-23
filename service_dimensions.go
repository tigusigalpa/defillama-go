package defillama

import (
	"context"
	"net/url"
)

// DimensionsService provides access to the Dimensions endpoints.
type DimensionsService struct {
	t *transport
}

// GetMetrics Get dimension data overview
// See: https://api-docs.defillama.com/#tag/dimensions/get/api/v2/metrics/%7Bmetric%7D
func (s *DimensionsService) GetMetrics(ctx context.Context, metric string, opts *DimensionOptions) (map[string]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"metric": metric}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /api/v2/metrics/{metric}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetChart Get historical timeseries chart data
// See: https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D
func (s *DimensionsService) GetChart(ctx context.Context, metric string, opts *DimensionOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"metric": metric}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/{metric}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetChartChainBreakdown Get historical timeseries chart data breakdown by chain
// See: https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/chain-breakdown
func (s *DimensionsService) GetChartChainBreakdown(ctx context.Context, metric string, opts *DimensionOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"metric": metric}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/{metric}/chain-breakdown", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetChartProtocolBreakdown Get historical timeseries chart data breakdown by protocol
// See: https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/protocol-breakdown
func (s *DimensionsService) GetChartProtocolBreakdown(ctx context.Context, metric string, opts *DimensionOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"metric": metric}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/{metric}/protocol-breakdown", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetChainMetrics Get chain dimension data overview
// See: https://api-docs.defillama.com/#tag/dimensions/get/api/v2/metrics/%7Bmetric%7D/chain/%7Bchain%7D
func (s *DimensionsService) GetChainMetrics(ctx context.Context, metric string, chain string, opts *DimensionOptions) (map[string]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"metric": metric,
		"chain":  chain}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /api/v2/metrics/{metric}/chain/{chain}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetChainChart Get chain historical timeseries chart data
// See: https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/chain/%7Bchain%7D
func (s *DimensionsService) GetChainChart(ctx context.Context, metric string, chain string, opts *DimensionOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"metric": metric,
		"chain":  chain}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/{metric}/chain/{chain}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetChainChartProtocolBreakdown Get chain timeseries chart data breakdown by protocol
// See: https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/chain/%7Bchain%7D/protocol-breakdown
func (s *DimensionsService) GetChainChartProtocolBreakdown(ctx context.Context, metric string, chain string, opts *DimensionOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"metric": metric,
		"chain":  chain}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/{metric}/chain/{chain}/protocol-breakdown", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetProtocolMetrics Get protocol dimension data overview
// See: https://api-docs.defillama.com/#tag/dimensions/get/api/v2/metrics/%7Bmetric%7D/protocol/%7Bprotocol%7D
func (s *DimensionsService) GetProtocolMetrics(ctx context.Context, metric string, protocol string, opts *DimensionOptions) (map[string]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"metric":   metric,
		"protocol": protocol}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /api/v2/metrics/{metric}/protocol/{protocol}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetProtocolChart Get protocol historical timeseries chart data
// See: https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/protocol/%7Bprotocol%7D
func (s *DimensionsService) GetProtocolChart(ctx context.Context, metric string, protocol string, opts *DimensionOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"metric":   metric,
		"protocol": protocol}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/{metric}/protocol/{protocol}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetProtocolChartChainBreakdown Get protocol timeseries chart data breakdown by chain
// See: https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/protocol/%7Bprotocol%7D/chain-breakdown
func (s *DimensionsService) GetProtocolChartChainBreakdown(ctx context.Context, metric string, protocol string, opts *DimensionOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"metric":   metric,
		"protocol": protocol}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/{metric}/protocol/{protocol}/chain-breakdown", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetProtocolChartVersionBreakdown Get protocol timeseries chart data breakdown by version
// See: https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/protocol/%7Bprotocol%7D/version-breakdown
func (s *DimensionsService) GetProtocolChartVersionBreakdown(ctx context.Context, metric string, protocol string, opts *DimensionOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"metric":   metric,
		"protocol": protocol}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/{metric}/protocol/{protocol}/version-breakdown", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetProtocolChartLabelBreakdown Get protocol timeseries chart data breakdown by label
// See: https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/protocol/%7Bprotocol%7D/label-breakdown
func (s *DimensionsService) GetProtocolChartLabelBreakdown(ctx context.Context, metric string, protocol string, opts *DimensionOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"metric":   metric,
		"protocol": protocol}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/{metric}/protocol/{protocol}/label-breakdown", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetCategoryMetrics Get category dimension data overview
// See: https://api-docs.defillama.com/#tag/dimensions/get/api/v2/metrics/%7Bmetric%7D/category/%7Bcategory%7D
func (s *DimensionsService) GetCategoryMetrics(ctx context.Context, metric string, category string, opts *DimensionOptions) (map[string]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"metric":   metric,
		"category": category}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /api/v2/metrics/{metric}/category/{category}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetCategoryChart Get category historical timeseries chart data
// See: https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/category/%7Bcategory%7D
func (s *DimensionsService) GetCategoryChart(ctx context.Context, metric string, category string, opts *DimensionOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"metric":   metric,
		"category": category}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/{metric}/category/{category}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetCategoryChartChainBreakdown Get category timeseries chart data breakdown by chain
// See: https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/category/%7Bcategory%7D/chain-breakdown
func (s *DimensionsService) GetCategoryChartChainBreakdown(ctx context.Context, metric string, category string, opts *DimensionOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"metric":   metric,
		"category": category}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/{metric}/category/{category}/chain-breakdown", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetCategoryChartProtocolBreakdown Get category timeseries chart data breakdown by protocol
// See: https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/category/%7Bcategory%7D/protocol-breakdown
func (s *DimensionsService) GetCategoryChartProtocolBreakdown(ctx context.Context, metric string, category string, opts *DimensionOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"metric":   metric,
		"category": category}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/{metric}/category/{category}/protocol-breakdown", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetCategoryChainMetrics Get category chain dimension data overview
// See: https://api-docs.defillama.com/#tag/dimensions/get/api/v2/metrics/%7Bmetric%7D/category/%7Bcategory%7D/chain/%7Bchain%7D
func (s *DimensionsService) GetCategoryChainMetrics(ctx context.Context, metric string, category string, chain string, opts *DimensionOptions) (map[string]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"metric":   metric,
		"category": category,
		"chain":    chain}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /api/v2/metrics/{metric}/category/{category}/chain/{chain}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetCategoryChainChart Get category chain historical timeseries chart data
// See: https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/category/%7Bcategory%7D/chain/%7Bchain%7D
func (s *DimensionsService) GetCategoryChainChart(ctx context.Context, metric string, category string, chain string, opts *DimensionOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"metric":   metric,
		"category": category,
		"chain":    chain}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/{metric}/category/{category}/chain/{chain}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetCategoryChainChartProtocolBreakdown Get category chain timeseries chart data breakdown by protocol
// See: https://api-docs.defillama.com/#tag/dimensions/get/api/v2/chart/%7Bmetric%7D/category/%7Bcategory%7D/chain/%7Bchain%7D/protocol-breakdown
func (s *DimensionsService) GetCategoryChainChartProtocolBreakdown(ctx context.Context, metric string, category string, chain string, opts *DimensionOptions) ([]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"metric":   metric,
		"category": category,
		"chain":    chain}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/{metric}/category/{category}/chain/{chain}/protocol-breakdown", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
