package defillama

import (
	"context"
	"net/url"
)

// VolumesService provides access to the Volumes endpoints.
type VolumesService struct {
	t *transport
}

// GetDEXOverview List all dexs along with summaries of their volumes and dataType history data
// See: https://api-docs.defillama.com/#tag/volumes/get/overview/dexs
func (s *VolumesService) GetDEXOverview(ctx context.Context, opts OverviewOptions) (map[string]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /overview/dexs", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetDEXOverviewByChain List all dexs along with summaries of their volumes and dataType history data filtering by chain
// See: https://api-docs.defillama.com/#tag/volumes/get/overview/dexs/%7Bchain%7D
func (s *VolumesService) GetDEXOverviewByChain(ctx context.Context, chain string, opts OverviewOptions) (map[string]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{
		"chain": chain}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /overview/dexs/{chain}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetDEXSummary Get summary of dex volume with historical data
// See: https://api-docs.defillama.com/#tag/volumes/get/summary/dexs/%7Bprotocol%7D
func (s *VolumesService) GetDEXSummary(ctx context.Context, protocol string, opts OverviewOptions) (map[string]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{
		"protocol": protocol}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /summary/dexs/{protocol}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetOptionsOverview List all options dexs along with summaries of their volumes and dataType history data
// See: https://api-docs.defillama.com/#tag/volumes/get/overview/options
func (s *VolumesService) GetOptionsOverview(ctx context.Context, opts OverviewOptions) (map[string]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /overview/options", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetOptionsOverviewByChain List all options dexs along with summaries of their volumes and dataType history data filtering by chain
// See: https://api-docs.defillama.com/#tag/volumes/get/overview/options/%7Bchain%7D
func (s *VolumesService) GetOptionsOverviewByChain(ctx context.Context, chain string, opts OverviewOptions) (map[string]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{
		"chain": chain}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /overview/options/{chain}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetOptionsSummary Get summary of options dex volume with historical data
// See: https://api-docs.defillama.com/#tag/volumes/get/summary/options/%7Bprotocol%7D
func (s *VolumesService) GetOptionsSummary(ctx context.Context, protocol string, opts *DimensionOptions) (map[string]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"protocol": protocol}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /summary/options/{protocol}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetOpenInterestOverview List all open interest dex exchanges along with summaries of their open interest
// See: https://api-docs.defillama.com/#tag/perps/get/overview/open-interest
func (s *VolumesService) GetOpenInterestOverview(ctx context.Context, opts OverviewOptions) (map[string]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /overview/open-interest", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetDerivativesOverview Lists all derivatives along summaries of their volumes filtering by chain
// See: https://api-docs.defillama.com/#tag/perps/get/api/overview/derivatives
func (s *VolumesService) GetDerivativesOverview(ctx context.Context, opts OverviewOptions) (map[string]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /api/overview/derivatives", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetDerivativesSummary Volume Details about a specific perp protocol
// See: https://api-docs.defillama.com/#tag/perps/get/api/summary/derivatives/%7Bprotocol%7D
func (s *VolumesService) GetDerivativesSummary(ctx context.Context, protocol string, opts OverviewOptions) (map[string]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{
		"protocol": protocol}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /api/summary/derivatives/{protocol}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
