package analyzer

import (
	"errors"
	"fmt"
)

var (
	ErrFileNotFoundSentinel = errors.New("Fichier introuvable")
	ErrFileEmptySentinel    = errors.New("Fichier vide")
)

type ErrFileNotFound struct {
	Path string
	Err  error
}

func (e *ErrFileNotFound) Error() string {
	return fmt.Sprintf("open %s: no such file or directory", e.Path)
}

func (e *ErrFileNotFound) Unwrap() error {
	return e.Err
}

type ErrFileEmpty struct {
	Path string
	Err  error
}

func (e *ErrFileEmpty) Error() string {
	return fmt.Sprintf("open %s file empty", e.Path)
}

func (e *ErrFileEmpty) Unwrap() error {
	return e.Err
}
