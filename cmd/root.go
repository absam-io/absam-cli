package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "absam-cli [command]",
	Short: "absam-cli is a command-line interface for the Absam API",
	Long:  "absam-cli is a command-line interface for the Absam API",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(firewallCmd)
}
