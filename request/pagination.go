package request

type PaginationRequest struct {
	Page     int `schema:"page"`
	PageSize int `schema:"page_size"`
}
