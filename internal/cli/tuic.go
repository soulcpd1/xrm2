package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var tuicCmd = &cobra.Command{
	Use:   "tuic",
	Short: "Manage TUIC",
}

var tuicInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install TUIC",
	Run: func(cmd *cobra.Command, args []string) {

		if err := tuic.Install(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	tuicCmd.AddCommand(tuicInstallCmd)
}
