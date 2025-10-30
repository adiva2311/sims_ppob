package dto

type ApiResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Pagination struct {
	Limit  int         `json:"limit"`
	Offset int         `json:"offset"`
	Record interface{} `json:"record,omitempty"`
}
