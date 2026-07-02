package cli

import (
	"github.com/soulcpd1/xrm2/internal/hysteria"
	"github.com/spf13/cobra"
)

var hysteriaCmd = &cobra.Command{
	Use:   "hysteria",
	Short: "Manage Hysteria2",
}

var hysteriaInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Show Hysteria installation plan",
	Run: func(cmd *cobra.Command, args []string) {
		hysteria.PrintInstallPlan()
	},
}

func init() {
	hysteriaCmd.AddCommand(hysteriaInstallCmd)
}
