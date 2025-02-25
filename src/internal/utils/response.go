package utils

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(message string, data interface{}) Response {
	return Response{
		Status:  "success2",
		Message: message,
		Data:    data,
	}
}

func NewErrorResponse(message string) Response {
	return Response{
		Status:  "error",
		Message: message,
	}
}
