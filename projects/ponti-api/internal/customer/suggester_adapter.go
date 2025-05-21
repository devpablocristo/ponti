package customer

import (
	"context"

	pkgsuggester "github.com/alphacodinggroup/ponti-backend/pkg/words-suggestors/pg_trgm-gin"
)

type SuggesterAdapter struct {
	s pkgsuggester.Suggester
}

func NewCustomerSuggester(engine pkgsuggester.Suggester) *SuggesterAdapter {
	return &SuggesterAdapter{engine}
}

func (s *SuggesterAdapter) Suggest(ctx context.Context, query string) ([]string, error) {
	suggestions, err := s.s.Suggest(ctx, query)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(suggestions))
	for _, sug := range suggestions {
		result = append(result, sug.Text)
	}

	return result, nil
}
