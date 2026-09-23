package defillama

import (
	"context"
	"net/url"
)

// EmissionsService provides access to the Emissions endpoints.
type EmissionsService struct {
	t *transport
}

// GetEmissions List of all tokens along with basic info for each
// See: https://api-docs.defillama.com/#tag/unlocks/get/api/emissions
func (s *EmissionsService) GetEmissions(ctx context.Context) ([]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw []any
	if err := s.t.get(ctx, "GET /api/emissions", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetEmission Unlocks data for a given token/protocol. You can find a list of available slugs to query by querying /emissions and then extracting the key `gecko_id`
// See: https://api-docs.defillama.com/#tag/unlocks/get/api/emission/%7Bprotocol%7D
func (s *EmissionsService) GetEmission(ctx context.Context, protocol string) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{
		"protocol": protocol}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /api/emission/{protocol}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
