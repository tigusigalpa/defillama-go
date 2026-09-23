package defillama

import (
	"context"
	"net/url"
)

// AccountService provides access to the Account endpoints.
type AccountService struct {
	t *transport
}

// GetUsage Get amount of credits left in the api key, these reset on the 1st of each month
// See: https://api-docs.defillama.com/#tag/meta/get/usage/APIKEY
func (s *AccountService) GetUsage(ctx context.Context) (*APIUsage, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /usage/APIKEY", path, query, &raw); err != nil {
		return nil, err
	}
	return &APIUsage{Raw: raw}, nil
}
