package analyzer

import (
	"fmt"
)

type ErrFileNotFound struct {
	Path string
	Err  error
}

func (e *ErrFileNotFound) Error() string {
	return fmt.Sprintf("open : %s", e.Path)
}

func (e *ErrFileNotFound) Unwrap() error {
	return e.Err
}

type ErrFileEmpty struct {
	Path string
	Err  error
}

func (e *ErrFileEmpty) Error() string {
	return fmt.Sprintf("empty log file : %s", e.Path)
}

func (e *ErrFileEmpty) Unwrap() error {
	return e.Err
}
