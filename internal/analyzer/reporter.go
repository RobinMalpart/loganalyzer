package analyzer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func ExportResultToJsonfile(outputPath string, results []Result) error {

	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("échec de création des répertoires pour %s : %w", outputPath, err)
	}

	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("échec de la sérialisation des résultats : %w", err)
	}

	if err := os.WriteFile(outputPath, data, 0644); err != nil {
		return fmt.Errorf("échec d'écriture du fichier %s : %w", outputPath, err)
	}

	return nil
}
