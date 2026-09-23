package defillama

import (
	"context"
	"net/url"
)

// EcosystemService provides access to the Ecosystem endpoints.
type EcosystemService struct {
	t *transport
}

// GetCategories Overview of all categories across all protocols
// See: https://api-docs.defillama.com/#tag/main-page/get/api/categories
func (s *EcosystemService) GetCategories(ctx context.Context) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /api/categories", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetForks Overview of all forks across all protocols
// See: https://api-docs.defillama.com/#tag/main-page/get/api/forks
func (s *EcosystemService) GetForks(ctx context.Context) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /api/forks", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetOracles Overview of all oracles across all protocols
// See: https://api-docs.defillama.com/#tag/main-page/get/api/oracles
func (s *EcosystemService) GetOracles(ctx context.Context) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /api/oracles", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetHacks Overview of all hacks on our Hacks dashboard
// See: https://api-docs.defillama.com/#tag/main-page/get/api/hacks
func (s *EcosystemService) GetHacks(ctx context.Context) ([]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw []any
	if err := s.t.get(ctx, "GET /api/hacks", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetRaises Overview of all raises on our Raises dashboard
// See: https://api-docs.defillama.com/#tag/main-page/get/api/raises
func (s *EcosystemService) GetRaises(ctx context.Context) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /api/raises", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetTreasuries List all protocols on our Treasuries dashboard
// See: https://api-docs.defillama.com/#tag/main-page/get/api/treasuries
func (s *EcosystemService) GetTreasuries(ctx context.Context) ([]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw []any
	if err := s.t.get(ctx, "GET /api/treasuries", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetEntities List all entities
// See: https://api-docs.defillama.com/#tag/main-page/get/api/entities
func (s *EcosystemService) GetEntities(ctx context.Context) ([]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw []any
	if err := s.t.get(ctx, "GET /api/entities", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
