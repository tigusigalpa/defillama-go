package defillama

import (
	"context"
	"net/url"
)

// DATService provides access to the DAT endpoints.
type DATService struct {
	t *transport
}

// GetInstitutions Get list of all institutions with Digital Asset Treasury data
// See: https://api-docs.defillama.com/#tag/dat/get/dat/institutions
func (s *DATService) GetInstitutions(ctx context.Context) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /dat/institutions", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetInstitution Get individual institution Digital Asset Treasury details
// See: https://api-docs.defillama.com/#tag/dat/get/dat/institutions/%7Bsymbol%7D
func (s *DATService) GetInstitution(ctx context.Context, symbol string) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{
		"symbol": symbol}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /dat/institutions/{symbol}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
