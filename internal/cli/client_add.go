package cli

import (
	"fmt"
	"os"

	"github.com/soulcpd1/xrm2/internal/xray"
	"github.com/spf13/cobra"
)

var clientAddCmd = &cobra.Command{
	Use:   "add <email>",
	Short: "Add Reality client",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		email := args[0]

		err := xray.Apply(func(cfg *xray.Config) error {
			return cfg.AddClient(email)
		})

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("Client added:", email)
	},
}

func init() {
	clientCmd.AddCommand(clientAddCmd)
}
