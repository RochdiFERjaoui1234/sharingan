package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var runAllCmd = &cobra.Command{
	Use:   "all [target]",
	Short: "Run full scan, badge, and API generation sequence",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := args[0]
		fmt.Println("🚀 Running full pipeline on", target)
		exec.Command("ruby", "scripts/secrets_scan.rb", target).Run()
		exec.Command("ruby", "scripts/header_tester.rb", target).Run()
		exec.Command("python3", "scripts/deep_scan_api.py", target).Run()
		exec.Command("ruby", "scripts/badge_gen.rb", target).Run()
		exec.Command("node", "scripts/secure_api_gen.js").Run()
	},
}
