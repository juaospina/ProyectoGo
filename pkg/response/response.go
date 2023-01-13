package response

type ResponseRequest struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func OKRequest(message string, data interface{}) *ResponseRequest {
	return &ResponseRequest{Message: message, Data: data}
}

func RequestError(err error) *ResponseRequest {
	return &ResponseRequest{Message: err.Error(), Data: nil}
}