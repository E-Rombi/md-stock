package domain

type Pagination[T any] struct {
	CurrentPage int
	PerPage     int
	Total       int64
	Items       []T
}

func NewPagination[T any](aCurrentPage int, aPerPage int, aTotal int64, anItems []T) *Pagination[T] {
	return &Pagination[T]{
		CurrentPage: aCurrentPage,
		PerPage:     aPerPage,
		Total:       aTotal,
		Items:       anItems,
	}
}
