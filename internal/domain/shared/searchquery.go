package domain

type SearchQuery struct {
	Page      int
	PerPage   int
	Terms     string
	Sort      string
	Direction string
}
