package service

type PaginationResponse struct {
	PageSize uint `json:"page_size"`
	Page     uint `json:"page"`
	Pages    uint `json:"pages"`
}
