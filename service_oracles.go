package defillama

import (
	"context"
	"net/url"
)

// OraclesService provides access to the Oracles endpoints.
type OraclesService struct {
	t *transport
}

// GetOracleMetrics Get oracle data overview
// See: https://api-docs.defillama.com/#tag/oracles/get/api/v2/metrics/oracle
func (s *OraclesService) GetOracleMetrics(ctx context.Context) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /api/v2/metrics/oracle", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetOracleChart Get timeseries chart data for all oracles
// See: https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle
func (s *OraclesService) GetOracleChart(ctx context.Context) ([]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/oracle", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetOracleChartChainBreakdown Get timeseries chart data breakdown by chain
// See: https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle/chain-breakdown
func (s *OraclesService) GetOracleChartChainBreakdown(ctx context.Context) ([]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/oracle/chain-breakdown", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetOracleChartProtocolBreakdown Get timeseries chart data breakdown by protocol
// See: https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle/protocol-breakdown
func (s *OraclesService) GetOracleChartProtocolBreakdown(ctx context.Context) ([]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/oracle/protocol-breakdown", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetOracleProtocolChart Get timeseries chart data by protocol/oracle
// See: https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle/protocol/%7Bprotocol%7D
func (s *OraclesService) GetOracleProtocolChart(ctx context.Context, protocol string) ([]any, error) {
	query := url.Values{}
	path := map[string]any{
		"protocol": protocol}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/oracle/protocol/{protocol}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetOracleProtocolChartChainBreakdown Get chain breakdown timeseries chart data by protocol/oracle
// See: https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle/protocol/%7Bprotocol%7D/chain-breakdown
func (s *OraclesService) GetOracleProtocolChartChainBreakdown(ctx context.Context, protocol string) ([]any, error) {
	query := url.Values{}
	path := map[string]any{
		"protocol": protocol}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/oracle/protocol/{protocol}/chain-breakdown", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetOracleChainChart Get timeseries chart data by chain
// See: https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle/chain/%7Bchain%7D
func (s *OraclesService) GetOracleChainChart(ctx context.Context, chain string) ([]any, error) {
	query := url.Values{}
	path := map[string]any{
		"chain": chain}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/oracle/chain/{chain}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetOracleChainChartProtocolBreakdown Get protocol breakdown timeseries chart data by chain
// See: https://api-docs.defillama.com/#tag/oracles/get/api/v2/chart/oracle/chain/%7Bchain%7D/protocol-breakdown
func (s *OraclesService) GetOracleChainChartProtocolBreakdown(ctx context.Context, chain string) ([]any, error) {
	query := url.Values{}
	path := map[string]any{
		"chain": chain}
	var raw []any
	if err := s.t.get(ctx, "GET /api/v2/chart/oracle/chain/{chain}/protocol-breakdown", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
