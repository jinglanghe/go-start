package cmd

import (
	"github.com/jinglanghe/go-start/cmd/sub_commands"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "go-start",
	Short: "go-start",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute the current command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(sub_commands.ApiServer)
	rootCmd.AddCommand(sub_commands.VersionCmd)
}
