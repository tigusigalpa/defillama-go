package defillama

import (
	"context"
	"net/url"
)

// FeesService provides access to the Fees endpoints.
type FeesService struct {
	t *transport
}

// GetOverview List all protocols along with summaries of their fees and revenue and dataType history data
// See: https://api-docs.defillama.com/#tag/fees-and-revenue/get/overview/fees
func (s *FeesService) GetOverview(ctx context.Context, opts OverviewOptions) (map[string]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /overview/fees", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetOverviewByChain List all protocols along with summaries of their fees and revenue and dataType history data by chain
// See: https://api-docs.defillama.com/#tag/fees-and-revenue/get/overview/fees/%7Bchain%7D
func (s *FeesService) GetOverviewByChain(ctx context.Context, chain string, opts OverviewOptions) (map[string]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{
		"chain": chain}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /overview/fees/{chain}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetSummary Get summary of protocol fees and revenue with historical data
// See: https://api-docs.defillama.com/#tag/fees-and-revenue/get/summary/fees/%7Bprotocol%7D
func (s *FeesService) GetSummary(ctx context.Context, protocol string, opts *DimensionOptions) (map[string]any, error) {
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
	if err := s.t.get(ctx, "GET /summary/fees/{protocol}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
