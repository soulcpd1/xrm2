package cli

import (
	"github.com/soulcpd1/xrm2/internal/doctor"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xrm",
	Short: "XRM2 VPN Manager",
}

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Run diagnostics",
	Run: func(cmd *cobra.Command, args []string) {

		results := []doctor.Result{
			doctor.KernelChecker{}.Run(),
			doctor.XrayChecker{}.Run(),
			doctor.ServiceChecker{}.Run(),
			doctor.ListenerChecker{}.Run(),
		}

		doctor.Print(results)
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
	rootCmd.AddCommand(backupCmd)
	rootCmd.AddCommand(clientCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
