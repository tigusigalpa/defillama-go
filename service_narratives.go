package defillama

import (
	"context"
	"net/url"
)

// NarrativesService provides access to the Narratives endpoints.
type NarrativesService struct {
	t *transport
}

// GetPerformance Get chart of narratives based on category performance (with individual coins weighted by mcap)
// See: https://api-docs.defillama.com/#tag/narratives/get/fdv/performance/%7Bperiod%7D
func (s *NarrativesService) GetPerformance(ctx context.Context, period string) ([]any, error) {
	query := url.Values{}
	path := map[string]any{
		"period": period}
	var raw []any
	if err := s.t.get(ctx, "GET /fdv/performance/{period}", path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
