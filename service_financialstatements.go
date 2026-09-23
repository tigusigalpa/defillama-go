package defillama

import (
	"context"
	"net/url"
)

// FinancialStatementsService provides access to the FinancialStatements endpoints.
type FinancialStatementsService struct {
	t *transport
}

// GetIncomeStatement Get protocol income statement report
// See: https://api-docs.defillama.com/#tag/financial-statements/get/api/v2/metrics/financial-statement/protocol/%7Bprotocol%7D
func (s *FinancialStatementsService) GetIncomeStatement(ctx context.Context, protocol string) (map[string]any, error) {
	query := url.Values{}
	path := map[string]any{
		"protocol": protocol}
	var raw map[string]any
	if err := s.t.get(ctx, "GET /api/v2/metrics/financial-statement/protocol/{protocol}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
