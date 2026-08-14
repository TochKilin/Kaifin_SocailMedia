package utls_http

// Function create Response Object ។
type Response struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

// use forList API to Pagination
func NewResponse(message string, statusCode int, data interface{}) Response {
	return Response{
		Success:    true,
		Message:    message,
		StatusCode: statusCode,
		Data:       data,
	}
}

// createResponse for Pagination។
type ResponseWithPaging struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	Total      int         `json:"total"`
}

// createResponse for Pagination។
func NewResponseWithPaing(message string, statusCode int, data interface{}, page int, perpage int, total int) ResponseWithPaging {
	return ResponseWithPaging{
		Success:    true,
		Message:    message,
		StatusCode: statusCode,
		Data:       data,
		Page:       page,
		PerPage:    perpage,
		Total:      total,
	}
}

// resonse api full project
