package cmd

import (
	"fmt"
	"os"

	"github.com/RobinMalpart/loganizer/internal/config"
	"github.com/spf13/cobra"
)

var (
	configPath string
	outputPath string
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze logs from a config file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Loading config...")

		logs, err := config.LoadConfig(configPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Erreur lecture config:", err)
			os.Exit(1)
		}

		fmt.Printf("Config chargée : %d logs\n", len(logs))
		for _, log := range logs {
			fmt.Printf("- ID: %s, Path: %s, Type: %s\n", log.ID, log.Path, log.Type)
		}

		// La suite viendra ici (appel à analyzer.RunAll)
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "", "Path to config JSON file")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "report.json", "Path to output JSON report file")
	analyzeCmd.MarkFlagRequired("config")
}
