package errs

type APIError struct {
	Msg        string
	StatusCode int
}

func NewAPIError(statusCode int, msg string) *APIError {
	return &APIError{StatusCode: statusCode, Msg: msg}
}

func (e *APIError) Error() string {
	return e.Msg
}
