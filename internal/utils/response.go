package utils

type SuccessResponse struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Type    string `json:"type,omitempty"`
	Message string `json:"message"`
}

func NewApiSuccessResponse(data interface{}) SuccessResponse {
	return SuccessResponse{
		Data: data,
	}
}

func NewApiErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Message: message,
	}
}
