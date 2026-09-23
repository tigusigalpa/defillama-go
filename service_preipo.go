package defillama

import (
	"context"
	"net/url"
)

// PreIPOService provides access to the PreIPO endpoints.
type PreIPOService struct {
	t *transport
}

// GetCompaniesList Get list of all tracked pre-IPO companies
// See: https://api-docs.defillama.com/#tag/pre-ipo/get/pre-ipo/v1/companies-list
func (s *PreIPOService) GetCompaniesList(ctx context.Context) ([]PreIPOCompany, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw struct {
		UpdatedAt string          `json:"updatedAt"`
		Data      []PreIPOCompany `json:"data"`
	}
	if err := s.t.get(ctx, "GET /pre-ipo/v1/companies-list", path, query, &raw); err != nil {
		return nil, err
	}
	return raw.Data, nil
}

// GetValuations Get valuation history for a pre-IPO company
// See: https://api-docs.defillama.com/#tag/pre-ipo/get/pre-ipo/v1/valuations
func (s *PreIPOService) GetValuations(ctx context.Context, opts PreIPOQuery) (map[string]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /pre-ipo/v1/valuations", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetRaises Get funding rounds for a pre-IPO company
// See: https://api-docs.defillama.com/#tag/pre-ipo/get/pre-ipo/v1/raises
func (s *PreIPOService) GetRaises(ctx context.Context, opts PreIPOQuery) (map[string]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /pre-ipo/v1/raises", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetSummary Get profile and latest valuation for a pre-IPO company
// See: https://api-docs.defillama.com/#tag/pre-ipo/get/pre-ipo/v1/summary
func (s *PreIPOService) GetSummary(ctx context.Context, opts PreIPOQuery) (map[string]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /pre-ipo/v1/summary", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetIntegrations Get tradeable markets for a pre-IPO company
// See: https://api-docs.defillama.com/#tag/pre-ipo/get/pre-ipo/v1/integrations
func (s *PreIPOService) GetIntegrations(ctx context.Context, opts PreIPOQuery) (map[string]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /pre-ipo/v1/integrations", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
