package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sharingan",
	Short: "Sharingan - DevSecOps multi-tool CLI",
	Long:  `Sharingan is a CLI toolkit that scans APIs, detects secrets, performs deep vulnerability scans, generates secure API templates, and creates security badges.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(scanCmd)
	rootCmd.AddCommand(badgeCmd)
	rootCmd.AddCommand(genApiCmd)
	rootCmd.AddCommand(runAllCmd)
}
