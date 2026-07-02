package cli

import (
	"fmt"

	"github.com/soulcpd1/xrm2/internal/registry"
	"github.com/spf13/cobra"
)

var routingCmd = &cobra.Command{
	Use:   "routing",
	Short: "Routing engine",
	Run: func(cmd *cobra.Command, args []string) {

		r := registry.Default()

		fmt.Printf("Registered engines: %d\n", r.Count())

		for _, name := range r.Names() {
			fmt.Println(" -", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(routingCmd)
}
