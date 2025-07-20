package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var genApiCmd = &cobra.Command{
	Use:   "gen-api",
	Short: "Generate a secure API template",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🛠️ Generating secure API template...")
		exec.Command("node", "scripts/secure_api_gen.js").Run()
	},
}