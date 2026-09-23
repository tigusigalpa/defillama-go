package defillama

import (
	"context"
	"net/url"
)

// ETFsService provides access to the ETFs endpoints.
type ETFsService struct {
	t *transport
}

// GetSnapshot Get ETFs and their metrics (aum, flows, fees...)
// See: https://api-docs.defillama.com/#tag/etfs/get/etfs/snapshot
func (s *ETFsService) GetSnapshot(ctx context.Context) ([]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw []any
	if err := s.t.get(ctx, "GET /etfs/snapshot", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetFlows Historical Flows at the Asset Level
// See: https://api-docs.defillama.com/#tag/etfs/get/etfs/flows
func (s *ETFsService) GetFlows(ctx context.Context) ([]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw []any
	if err := s.t.get(ctx, "GET /etfs/flows", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
