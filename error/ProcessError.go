package derrors

import (
	"fmt"
)

type ProcessError struct {
	Errcode string
	Errmsg  string
}

func (e ProcessError) Error() string {
	return fmt.Sprintf("%s:%s", e.Errcode, e.Errmsg)
}

func New(errcode, errmsg string) error {
	return ProcessError{errcode, errmsg}
}
