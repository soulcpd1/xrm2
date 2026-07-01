package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xrm",
	Short: "XRM2 VPN Manager",
}

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check system",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("XRM2 v0.1")
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
