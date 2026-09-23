package defillama

import (
	"context"
	"net/url"
)

// BridgesService provides access to the Bridges endpoints.
type BridgesService struct {
	t *transport
}

// GetBridges List all bridges along with summaries of recent bridge volumes.
// See: https://api-docs.defillama.com/#tag/bridges/get/bridges/bridges
func (s *BridgesService) GetBridges(ctx context.Context, opts *BridgesOptions) (map[string]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /bridges/bridges", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetBridge Get summary of bridge volume and volume breakdown by chain
// See: https://api-docs.defillama.com/#tag/bridges/get/bridges/bridge/%7Bid%7D
func (s *BridgesService) GetBridge(ctx context.Context, id int64) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{
		"id": id}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /bridges/bridge/{id}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetBridgeVolume Get historical volumes for a bridge, chain, or bridge on a particular chain
// See: https://api-docs.defillama.com/#tag/bridges/get/bridges/bridgevolume/%7Bchain%7D
func (s *BridgesService) GetBridgeVolume(ctx context.Context, chain string, opts *BridgeVolumeOptions) ([]any, error) {
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
	if err := s.t.get(ctx, "GET /bridges/bridgevolume/{chain}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetBridgeDayStats Get a 24hr token and address volume breakdown for a bridge
// See: https://api-docs.defillama.com/#tag/bridges/get/bridges/bridgedaystats/%7Btimestamp%7D/%7Bchain%7D
func (s *BridgesService) GetBridgeDayStats(ctx context.Context, timestamp int64, chain string, opts *BridgeVolumeOptions) (map[string]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"timestamp": timestamp,
		"chain":     chain}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /bridges/bridgedaystats/{timestamp}/{chain}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetBridgeTransactions Get all transactions for a bridge within a date range
// See: https://api-docs.defillama.com/#tag/bridges/get/bridges/transactions/%7Bid%7D
func (s *BridgesService) GetBridgeTransactions(ctx context.Context, id int64, opts *BridgeTransactionsOptions) (map[string]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"id": id}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /bridges/transactions/{id}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
