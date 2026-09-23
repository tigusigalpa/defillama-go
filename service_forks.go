package defillama

import (
	"context"
	"net/url"
)

// ForksService provides access to the Forks endpoints.
type ForksService struct {
	t *transport
}

// GetForkMetrics Get fork data overview
// See: https://api-docs.defillama.com/#tag/forks/get/api/v2/metrics/fork
func (s *ForksService) GetForkMetrics(ctx context.Context) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /api/v2/metrics/fork", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetForkChartProtocolBreakdown Get timeseries chart data breakdown by protocol
// See: https://api-docs.defillama.com/#tag/forks/get/api/v2/chart/fork/protocol-breakdown
func (s *ForksService) GetForkChartProtocolBreakdown(ctx context.Context) ([]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/fork/protocol-breakdown", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetForkProtocolChart Get timeseries chart data by protocol
// See: https://api-docs.defillama.com/#tag/forks/get/api/v2/chart/fork/protocol/%7Bprotocol%7D
func (s *ForksService) GetForkProtocolChart(ctx context.Context, protocol string) ([]any, error) {
	query := url.Values{}
	path := map[string]any{
		"protocol": protocol}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/fork/protocol/{protocol}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
