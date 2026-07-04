package cli

import (
	"fmt"
	"os"

	"github.com/soulcpd1/xrm2/internal/hysteria"
	"github.com/spf13/cobra"
)

var hysteriaCmd = &cobra.Command{
	Use:   "hysteria",
	Short: "Manage Hysteria2",
}

var hysteriaInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install Hysteria2",
	Run: func(cmd *cobra.Command, args []string) {

		if err := hysteria.Install(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	hysteriaCmd.AddCommand(hysteriaInstallCmd)
	hysteriaCmd.AddCommand(hysteriaServeCmd)
}

var hysteriaServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run XRM integrated Hysteria runtime",
	Run: func(cmd *cobra.Command, args []string) {

		rt := runtime.New()
		if err := rt.Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
