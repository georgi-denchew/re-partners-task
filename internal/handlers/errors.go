package handlers

type InvalidRequestError struct {
	Message string `json:"message"`
}

func NewInvalidRequestError(message string) error {
	return &InvalidRequestError{
		Message: message,
	}
}

func (e *InvalidRequestError) Error() string {
	return e.Message
}
