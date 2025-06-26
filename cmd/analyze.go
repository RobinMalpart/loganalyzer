package cmd

import (
	"fmt"
	"loganalyzer/internal/analyzer"
	"loganalyzer/internal/config"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	inputFilePath  string
	outputFilePath string
)

var analyseCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze log files",
	Long:  `Analyze log files to extract useful information and generate reports.`,
	Run: func(cmd *cobra.Command, args []string) {
		if inputFilePath == "" {
			fmt.Fprintln(os.Stderr, "Erreur: le chemin du fichier de configuration d'entrée doit être spécifié avec --input")
			os.Exit(1)
		}

		logs, err := config.LoadConfig(inputFilePath)

		if len(logs) == 0 {
			fmt.Fprintln(os.Stderr, "Aucun aucun log à analyser. Veuillez vérifier le fichier de configuration.")
			os.Exit(1)
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erreur lors du chargement de la configuration depuis le fichier %s: %v\n", inputFilePath, err)
			os.Exit(1)
		}

		//bonus2
		datePrefix := time.Now().Format("02012006")
		outputFilePath = fmt.Sprintf("%s_%s", datePrefix, outputFilePath)

		results := analyzer.Analyse(logs)

		reportingError := analyzer.ExportResultToJsonfile(outputFilePath, results)

		if reportingError != nil {
			fmt.Fprintf(os.Stderr, "Erreur pendant l'exportation des résultats vers le fichier %s: %v\n", outputFilePath, reportingError)
			os.Exit(1)
		} else {
			fmt.Printf("Analyse des logs terminée avec succès. Résultats exportés vers %s\n", outputFilePath)
		}
	},
}

func init() {
	analyseCmd.Flags().StringVarP(&inputFilePath, "input", "i", "", "Chemin du fichier de configuration d'entrée")
	analyseCmd.Flags().StringVarP(&outputFilePath, "output", "o", "result.json", "Chemin du fichier de sortie pour les résultats de l'analyse")
	analyseCmd.MarkFlagRequired("input")

	rootcmd.AddCommand(analyseCmd)
}
