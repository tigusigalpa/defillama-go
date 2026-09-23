package defillama

import (
	"context"
	"net/url"
)

// EquitiesService provides access to the Equities endpoints.
type EquitiesService struct {
	t *transport
}

// GetCompaniesList Get list of all tracked public companies
// See: https://api-docs.defillama.com/#tag/equities/get/equities/v1/companies-list
func (s *EquitiesService) GetCompaniesList(ctx context.Context) ([]EquityCompany, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw []EquityCompany
	if err := s.t.get(ctx, "GET /equities/v1/companies-list", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetStatements Get financial statements for a company
// See: https://api-docs.defillama.com/#tag/equities/get/equities/v1/statements
func (s *EquitiesService) GetStatements(ctx context.Context, opts EquitiesQuery) (map[string]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /equities/v1/statements", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetDimensions Get financial dimensions for a company
// See: https://api-docs.defillama.com/#tag/equities/get/equities/v1/dimensions
func (s *EquitiesService) GetDimensions(ctx context.Context, opts EquitiesQuery) (map[string]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /equities/v1/dimensions", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetPriceHistory Get historical price data for a company
// See: https://api-docs.defillama.com/#tag/equities/get/equities/v1/price-history
func (s *EquitiesService) GetPriceHistory(ctx context.Context, opts EquitiesQuery) ([]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{}
	var raw []any
	if err := s.t.get(ctx, "GET /equities/v1/price-history", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetOHLCV Get OHLCV candle data for a company
// See: https://api-docs.defillama.com/#tag/equities/get/equities/v1/ohlcv
func (s *EquitiesService) GetOHLCV(ctx context.Context, opts EquitiesQuery) ([]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{}
	var raw []any
	if err := s.t.get(ctx, "GET /equities/v1/ohlcv", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetSummary Get live market summary for a company
// See: https://api-docs.defillama.com/#tag/equities/get/equities/v1/summary
func (s *EquitiesService) GetSummary(ctx context.Context, opts EquitiesQuery) (map[string]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /equities/v1/summary", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetFilings Get filings for a company
// See: https://api-docs.defillama.com/#tag/equities/get/equities/v1/filings
func (s *EquitiesService) GetFilings(ctx context.Context, opts EquitiesQuery) ([]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{}
	var raw []any
	if err := s.t.get(ctx, "GET /equities/v1/filings", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetOnchainMarkets Get on-chain tradeable markets for a company
// See: https://api-docs.defillama.com/#tag/equities/get/equities/v1/onchain
func (s *EquitiesService) GetOnchainMarkets(ctx context.Context, opts EquitiesQuery) (map[string]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /equities/v1/onchain", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
