package err

import "fmt"

// NewCustom creates a new errCustom
func NewCustom(status int, message string) *errCustom {
	return &errCustom{
		status:  status,
		Message: message,
	}
}

// errCustom es un error personalizado para http
type errCustom struct {
	status  int
	Message string `json:"error"`
}

func (e *errCustom) Error() string {
	return fmt.Sprintf(e.Message)
}

// Status http status code
func (e *errCustom) Status() int {
	return e.status
}
