package helpers

type PaginationInfo struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalItems int64 `json:"total_items"`
	TotalPages int   `json:"total_pages"`
}

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func SuccessPaginationResponse(message string, data interface{}, pagination PaginationInfo) Response {
	return Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message string) Response {
	return Response{
		Status:  false,
		Message: message,
	}
}
