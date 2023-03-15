package domain

type Pagination[T any] struct {
	CurrentPage int   `json:"currentPage"`
	PerPage     int   `json:"perPage"`
	Total       int64 `json:"total"`
	Items       []T   `json:"items"`
}

func NewPagination[T any](aCurrentPage int, aPerPage int, aTotal int64, someItems []T) *Pagination[T] {
	return &Pagination[T]{
		CurrentPage: aCurrentPage,
		PerPage:     aPerPage,
		Total:       aTotal,
		Items:       someItems,
	}
}
