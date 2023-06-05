package request

type Pagination struct {
	Page    uint `json:"page"`
	PerPage uint `json:"page_per"`
}
