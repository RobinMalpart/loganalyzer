package cmd

import (
	"encoding/json"
	"fmt"
	"loganalyzer/internal/config"
	"os"

	"github.com/spf13/cobra"
)

var (
	newLogID   string
	newLogPath string
	newLogType string
	configFile string
)

var addLogCmd = &cobra.Command{
	Use:   "add-log",
	Short: "Ajouter un log au fichier de configuration",
	Long:  `Ajoute une nouvelle entrée de log dans le fichier de configuration JSON fourni.`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile(configFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erreur lors de la lecture du fichier de configuration : %v\n", err)
			os.Exit(1)
		}

		var logs []config.LogConfig
		if len(data) > 0 {
			if err := json.Unmarshal(data, &logs); err != nil {
				fmt.Fprintf(os.Stderr, "Erreur de parsing JSON : %v\n", err)
				os.Exit(1)
			}
		}

		newLog := config.LogConfig{
			ID:   newLogID,
			Path: newLogPath,
			Type: newLogType,
		}
		logs = append(logs, newLog)

		updatedData, err := json.MarshalIndent(logs, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erreur lors de la sérialisation JSON : %v\n", err)
			os.Exit(1)
		}
		if err := os.WriteFile(configFile, updatedData, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Erreur lors de l'écriture du fichier : %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Log ajouté avec succès dans le fichier de configuration.")
	},
}

func init() {
	addLogCmd.Flags().StringVar(&newLogID, "id", "", "Identifiant du log")
	addLogCmd.Flags().StringVar(&newLogPath, "path", "", "Chemin du fichier de log")
	addLogCmd.Flags().StringVar(&newLogType, "type", "", "Type de log (ex: access, error)")
	addLogCmd.Flags().StringVar(&configFile, "file", "config.json", "Fichier de configuration cible")

	addLogCmd.MarkFlagRequired("id")
	addLogCmd.MarkFlagRequired("path")
	addLogCmd.MarkFlagRequired("type")
	addLogCmd.MarkFlagRequired("file")

	rootcmd.AddCommand(addLogCmd)
}
