package cli

import "github.com/spf13/cobra"

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Manage Reality clients",
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
