package checker

import "fmt"

type UnrechabkeURLError struct {
	url string
	Err error
}

func (e *UnrechabkeURLError) Error() string {
	return fmt.Sprintf("URK inaccessible : %s (%v)", e.url, e.Err)
}

func (e *UnrechabkeURLError) Unwrap() error {
	return e.Err
}
