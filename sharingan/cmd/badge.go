package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var badgeCmd = &cobra.Command{
	Use:   "badge [target]",
	Short: "Generate a security badge for a scanned target",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := args[0]
		fmt.Println("🏅 Generating badge for:", target)
		exec.Command("ruby", "scripts/badge_gen.rb", target).Run()
	},
}