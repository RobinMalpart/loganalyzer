package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootcmd = &cobra.Command{
	Use:   "loganalyser",
	Short: "Loganalyser is a tool to analyze log files.",
	Long: `Loganalyser is a command-line tool that analyzes log files.
It can be used to extract useful information from logs and generate reports.`,
}

func Execute() {
	if err := rootcmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, "Erreur lors de l'exécution de la commande : ", err)
		os.Exit(1)
	}
}
