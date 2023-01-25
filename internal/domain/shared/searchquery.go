package domain

type SearchQuery struct {
	Page      int
	PerPage   int
	Terms     *string
	Sort      *string
	Direction *string
}

func NewSearchQuery(aPage int, aPerPage int, aTerms *string, aSort *string, aDirection *string) *SearchQuery {
	return &SearchQuery{
		Page:      aPage,
		PerPage:   aPerPage,
		Terms:     aTerms,
		Sort:      aSort,
		Direction: aDirection,
	}
}
