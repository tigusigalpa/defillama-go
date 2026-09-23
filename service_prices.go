package defillama

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// PricesService provides access to the Prices endpoints.
type PricesService struct {
	t *transport
}

// GetCurrentPrices Get current prices of tokens by contract address
// See: https://api-docs.defillama.com/#tag/coins/get/prices/current/%7Bcoins%7D
func (s *PricesService) GetCurrentPrices(ctx context.Context, coins []string) (map[string]CoinPrice, error) {
	query := url.Values{}
	path := map[string]any{
		"coins": coins}
	var raw struct {
		Coins map[string]CoinPrice `json:"coins"`
	}
	if err := s.t.get(ctx, "GET /prices/current/{coins}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw.Coins, nil
}

// GetHistoricalPrices Get historical prices of tokens by contract address
// See: https://api-docs.defillama.com/#tag/coins/get/prices/historical/%7Btimestamp%7D/%7Bcoins%7D
func (s *PricesService) GetHistoricalPrices(ctx context.Context, coins []string, timestamp int64) (map[string]CoinPrice, error) {
	query := url.Values{}
	path := map[string]any{
		"coins":     coins,
		"timestamp": timestamp}
	var raw struct {
		Coins map[string]CoinPrice `json:"coins"`
	}
	if err := s.t.get(ctx, "GET /prices/historical/{timestamp}/{coins}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw.Coins, nil
}

// GetBatchHistoricalPrices Get historical prices for multiple tokens at multiple different timestamps
// See: https://api-docs.defillama.com/#tag/coins/get/batchHistorical
func (s *PricesService) GetBatchHistoricalPrices(ctx context.Context, coins map[string][]int64) (map[string]any, error) {
	query := url.Values{}
	enc, err := json.Marshal(coins)
	if err != nil {
		return nil, fmt.Errorf("defillama: encode coins: %w", err)
	}
	query.Set("coins", string(enc))
	path := map[string]any{}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /batchHistorical", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetChart Get token prices at regular time intervals
// See: https://api-docs.defillama.com/#tag/coins/get/chart/%7Bcoins%7D
func (s *PricesService) GetChart(ctx context.Context, coins []string, opts *ChartOptions) (map[string]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"coins": coins}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /chart/{coins}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetPercentageChange Get percentage change in price over time
// See: https://api-docs.defillama.com/#tag/coins/get/percentage/%7Bcoins%7D
func (s *PricesService) GetPercentageChange(ctx context.Context, coins []string, opts *PercentageOptions) (map[string]any, error) {
	query := url.Values{}
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, err
		}
		opts.apply(query)
	}
	path := map[string]any{
		"coins": coins}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /percentage/{coins}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// GetFirstPrices Get earliest timestamp price record for coins
// See: https://api-docs.defillama.com/#tag/coins/get/prices/first/%7Bcoins%7D
func (s *PricesService) GetFirstPrices(ctx context.Context, coins []string) (map[string]CoinPrice, error) {
	query := url.Values{}
	path := map[string]any{
		"coins": coins}
	var raw struct {
		Coins map[string]CoinPrice `json:"coins"`
	}
	if err := s.t.get(ctx, "GET /prices/first/{coins}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw.Coins, nil
}

// GetBlockAtTimestamp Get the closest block to a timestamp
// See: https://api-docs.defillama.com/#tag/coins/get/block/%7Bchain%7D/%7Btimestamp%7D
func (s *PricesService) GetBlockAtTimestamp(ctx context.Context, chain string, timestamp int64) (*Block, error) {
	query := url.Values{}
	path := map[string]any{
		"chain":     chain,
		"timestamp": timestamp}
	var raw Block
	if err := s.t.get(ctx, "GET /block/{chain}/{timestamp}", path, query, &raw); err != nil {
		return nil, err
	}
	return &raw, nil
}

// GetHistoricalLiquidity Provides the name of contracts on a determined chain
// See: https://api-docs.defillama.com/#tag/token-liquidity/get/api/historicalLiquidity/%7Btoken%7D
func (s *PricesService) GetHistoricalLiquidity(ctx context.Context, token string) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{
		"token": token}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /api/historicalLiquidity/{token}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
