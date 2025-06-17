package cmd

import (
	"fmt"
	analyzer "loganalyzer/internal/analyzer"
	"loganalyzer/internal/config"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

var (
	inputFilePath  string
	outputFilePath string
)

var rootcmd = &cobra.Command{
	Use:   "loganalyser",
	Short: "Loganalyser is a tool to analyze log files.",
	Long: `Loganalyser is a command-line tool that analyzes log files.
It can be used to extract useful information from logs and generate reports.`,
	Run: func(cmd *cobra.Command, args []string) {

		if inputFilePath == "" {
			fmt.Fprintln(os.Stderr, "Error: input file path is required")
			os.Exit(1)
		}

		logs, err := config.LoadConfig(inputFilePath)

		if len(logs) == 0 {
			fmt.Fprintln(os.Stderr, "No logs to analyze in the input file.")
			os.Exit(1)
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading logs from file %s: %v\n", inputFilePath, err)
			os.Exit(1)
		}

		results := analyzer.RunAll(logs)

		var wg sync.WaitGroup
		wg.Add(len(results))
		for _, res := range results {
			go func(r analyzer.Result) {
				defer wg.Done()
				if r.Status == "FAILED" {
					fmt.Printf("Log %s: %s - %s\n", r.LogID, r.Message, r.ErrorDetails)
				} else {
					fmt.Printf("Log %s: OK\n", r.LogID)
				}
			}(res)
		}
		wg.Wait()
	},
}

func Execute() {
	if err := rootcmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, "Error executing command: %v\n", err)
		os.Exit(1)
	}
}
