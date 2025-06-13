package analyzer

import "fmt"

type ErrFileNotFound struct {
	Path string
}

func (e ErrFileNotFound) Error() string {
	return fmt.Sprintf("fichier non trouvé: %s", e.Path)
}

type ErrParsing struct {
	Path string
}

func (e ErrParsing) Error() string {
	return fmt.Sprintf("erreur de parsing pour le fichier: %s", e.Path)
}
