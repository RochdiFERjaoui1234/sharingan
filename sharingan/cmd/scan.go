package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan [target]",
	Short: "Scan an API or URL for secrets, headers, and deep vulnerabilities",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := args[0]
		fmt.Println("🔍 Scanning for secrets:", target)
		exec.Command("ruby", "scripts/secrets_scan.rb", target).Run()

		fmt.Println("🔐 Testing security headers:", target)
		exec.Command("ruby", "scripts/header_tester.rb", target).Run()

		fmt.Println("🕵️‍♂️ Running deep vulnerability scan:", target)
		exec.Command("python3", "scripts/deep_scan_api.py", target).Run()
	},
}
