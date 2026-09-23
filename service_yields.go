package defillama

import (
	"context"
	"net/url"
)

// YieldsService provides access to the Yields endpoints.
type YieldsService struct {
	t *transport
}

// GetPools Retrieve the latest data for all pools, including enriched information such as predictions
// See: https://api-docs.defillama.com/#tag/yields/get/pools
func (s *YieldsService) GetPools(ctx context.Context) ([]YieldPool, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw struct {
		Data []YieldPool `json:"data"`
	}
	if err := s.t.get(ctx, "GET /pools", path, query, &raw); err != nil {
		return nil, err
	}
	return raw.Data, nil
}

// GetPoolChart Get historical APY and TVL of a pool
// See: https://api-docs.defillama.com/#tag/yields/get/chart/%7Bpool%7D
func (s *YieldsService) GetPoolChart(ctx context.Context, pool string) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{
		"pool": pool}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /chart/{pool}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetLegacyPools Retrieve the latest data for all pools in the legacy v1 response format
// See: https://api-docs.defillama.com/#tag/yields/get/yields/v1/pools
func (s *YieldsService) GetLegacyPools(ctx context.Context) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /yields/v1/pools", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetLegacyPoolChart Get historical APY and TVL of a pool in the legacy v1 response format
// See: https://api-docs.defillama.com/#tag/yields/get/yields/v1/chart/%7Bpool%7D
func (s *YieldsService) GetLegacyPoolChart(ctx context.Context, pool string) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{
		"pool": pool}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /yields/v1/chart/{pool}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetLegacyBorrowPools Borrow costs APY of assets from lending markets in the legacy v1 response format
// See: https://api-docs.defillama.com/#tag/yields/get/yields/v1/poolsBorrow
func (s *YieldsService) GetLegacyBorrowPools(ctx context.Context) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /yields/v1/poolsBorrow", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetLegacyLendBorrowChart Historical borrow cost APY from a pool on a lending market in the legacy v1 response format
// See: https://api-docs.defillama.com/#tag/yields/get/yields/v1/chartLendBorrow/%7Bpool%7D
func (s *YieldsService) GetLegacyLendBorrowChart(ctx context.Context, pool string) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{
		"pool": pool}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /yields/v1/chartLendBorrow/{pool}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetEarnPools Retrieve the latest data for all earn pools
// See: https://api-docs.defillama.com/#tag/yields/get/yields/v2/earn
func (s *YieldsService) GetEarnPools(ctx context.Context) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /yields/v2/earn", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// QueryEarnPools Retrieve a filtered, paginated page of earn pools
// See: https://api-docs.defillama.com/#tag/yields/get/yields/v2/earn/query
func (s *YieldsService) QueryEarnPools(ctx context.Context, opts EarnPoolsQuery) (map[string]any, error) {
	query := url.Values{}
	if err := opts.validate(); err != nil {
		return nil, err
	}
	opts.apply(query)
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /yields/v2/earn/query", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetEarnPoolHistory Retrieve daily history for an earn pool
// See: https://api-docs.defillama.com/#tag/yields/get/yields/v2/earn/%7Bid%7D/history
func (s *YieldsService) GetEarnPoolHistory(ctx context.Context, id string, opts *HistoryRangeOptions) (map[string]any, error) {
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
	if err := s.t.get(ctx, "GET /yields/v2/earn/{id}/history", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetVerifiedEarnPools Retrieve the latest verified data for all measured earn pools
// See: https://api-docs.defillama.com/#tag/yields/get/yields/v2/earn/verified
func (s *YieldsService) GetVerifiedEarnPools(ctx context.Context) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /yields/v2/earn/verified", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetVerifiedEarnPool Retrieve the latest verified data for one earn pool
// See: https://api-docs.defillama.com/#tag/yields/get/yields/v2/earn/verified/%7Bid%7D
func (s *YieldsService) GetVerifiedEarnPool(ctx context.Context, id string) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{
		"id": id}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /yields/v2/earn/verified/{id}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetVerifiedEarnPoolHistory Retrieve daily verified history for an earn pool
// See: https://api-docs.defillama.com/#tag/yields/get/yields/v2/earn/verified/%7Bid%7D/history
func (s *YieldsService) GetVerifiedEarnPoolHistory(ctx context.Context, id string, opts *HistoryRangeOptions) (map[string]any, error) {
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
	if err := s.t.get(ctx, "GET /yields/v2/earn/verified/{id}/history", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetBorrowMarkets Retrieve the latest data for all borrow markets
// See: https://api-docs.defillama.com/#tag/yields/get/yields/v2/borrow/markets
func (s *YieldsService) GetBorrowMarkets(ctx context.Context) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /yields/v2/borrow/markets", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetBorrowMarketHistory Retrieve daily borrow-side history for a borrow market
// See: https://api-docs.defillama.com/#tag/yields/get/yields/v2/borrow/markets/%7Bid%7D/history
func (s *YieldsService) GetBorrowMarketHistory(ctx context.Context, id string, opts *HistoryRangeOptions) (map[string]any, error) {
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
	if err := s.t.get(ctx, "GET /yields/v2/borrow/markets/{id}/history", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetBorrowRoutes Retrieve the latest data for all borrow routes
// See: https://api-docs.defillama.com/#tag/yields/get/yields/v2/borrow/routes
func (s *YieldsService) GetBorrowRoutes(ctx context.Context) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /yields/v2/borrow/routes", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetLoopStrategies Retrieve the latest data for all loop strategies
// See: https://api-docs.defillama.com/#tag/yields/get/yields/v2/loops
func (s *YieldsService) GetLoopStrategies(ctx context.Context) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /yields/v2/loops", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetPerpFundingRates Funding rates and Open Interest of perps across exchanges, including both Decentralized and Centralized
// See: https://api-docs.defillama.com/#tag/yields/get/yields/perps
func (s *YieldsService) GetPerpFundingRates(ctx context.Context) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /yields/perps", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetLSTRates Exchange rates and ETH peg of liquid staking tokens
// See: https://api-docs.defillama.com/#tag/yields/get/api/lstRates
func (s *YieldsService) GetLSTRates(ctx context.Context) ([]any, error) {
	query := url.Values{}
	path := map[string]any{}
	var raw []any
	if err := s.t.get(ctx, "GET /api/lstRates", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
