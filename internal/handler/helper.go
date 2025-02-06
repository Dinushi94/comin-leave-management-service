package handler

type ErrorResponse struct {
	Error string `json:"error"`
}

type MetaResponse struct {
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalPages int64 `json:"total_pages"`
}

type ListResponse struct {
	Data interface{}  `json:"data"`
	Meta MetaResponse `json:"meta"`
}
