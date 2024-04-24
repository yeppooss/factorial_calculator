package rp

type Response struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message,omitempty"`
}

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

func Ok(msg interface{}) Response {
	return Response{Status: StatusOk, Message: msg}
}

func Error(msg string) Response {
	return Response{
		Status:  StatusError,
		Message: msg}
}
