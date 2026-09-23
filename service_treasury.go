package defillama

import (
	"context"
	"net/url"
)

// TreasuryService provides access to the Treasury endpoints.
type TreasuryService struct {
	t *transport
}

// GetTreasuryMetrics Get aggregate treasury metrics for a protocol
// See: https://api-docs.defillama.com/#tag/treasury/get/api/v2/metrics/treasury/protocol/%7Bprotocol%7D
func (s *TreasuryService) GetTreasuryMetrics(ctx context.Context, protocol string) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{
		"protocol": protocol}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /api/v2/metrics/treasury/protocol/{protocol}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetTreasuryChart Get historical treasury chart for a protocol
// See: https://api-docs.defillama.com/#tag/treasury/get/api/v2/chart/treasury/protocol/%7Bprotocol%7D
func (s *TreasuryService) GetTreasuryChart(ctx context.Context, protocol string, opts *TreasuryChartOptions) ([]any, error) {
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
	if err := s.t.get(ctx, "GET /api/v2/chart/treasury/protocol/{protocol}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetTreasuryChartChainBreakdown Get historical treasury chart for a protocol broken down by chain
// See: https://api-docs.defillama.com/#tag/treasury/get/api/v2/chart/treasury/protocol/%7Bprotocol%7D/chain-breakdown
func (s *TreasuryService) GetTreasuryChartChainBreakdown(ctx context.Context, protocol string, opts *TreasuryChartOptions) ([]any, error) {
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
	if err := s.t.get(ctx, "GET /api/v2/chart/treasury/protocol/{protocol}/chain-breakdown", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetTreasuryChartTokenBreakdown Get historical treasury chart for a protocol broken down by token
// See: https://api-docs.defillama.com/#tag/treasury/get/api/v2/chart/treasury/protocol/%7Bprotocol%7D/token-breakdown
func (s *TreasuryService) GetTreasuryChartTokenBreakdown(ctx context.Context, protocol string, opts *TreasuryChartOptions) ([]any, error) {
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
	if err := s.t.get(ctx, "GET /api/v2/chart/treasury/protocol/{protocol}/token-breakdown", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
