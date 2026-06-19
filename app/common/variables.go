package common

type ResponseMessage struct {
	IsError bool   `json:"error" example:"true"`
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Error message"`
}

func (r *ResponseMessage) Error() string {
	return r.Message
}

func (r *ResponseMessage) GetStatus() int {
	return r.Code
}

func (r *ResponseMessage) ContentType(ct string) string {
	return ct
}
