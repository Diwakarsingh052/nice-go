package web

// ErrorResponse is the struct used for API responses from failures in the API.
// this is for json
type ErrorResponseJson struct {
	Error string
}

// Error struct is used by core of our app to generate error messages that we want to send to end user
type Error struct {
	Err    error
	Status int
}

func (err *Error) Error() string {
	return err.Err.Error()
}

// NewRequestError would be called by other layers to generate user friend msg without disclosing internal details
func NewRequestError(err error, status int) error {
	return &Error{
		Err:    err,
		Status: status,
	}

}
