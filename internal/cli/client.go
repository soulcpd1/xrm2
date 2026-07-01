package cli

import (
	"fmt"
	"os"

	"github.com/soulcpd1/xrm2/internal/xray"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Manage Reality clients",
}

var clientListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Reality clients",
	Run: func(cmd *cobra.Command, args []string) {

		cfg, err := xray.Read()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		clients := xray.RealityClients(cfg)

		fmt.Println("Reality clients")
		fmt.Println("---------------")

		if len(clients) == 0 {
			fmt.Println("No clients")
			return
		}

		for i, c := range clients {
			email := c.Email
			if email == "" {
				email = "-"
			}

			fmt.Printf("%2d. %-30s %s\n", i+1, email, c.ID)
		}
	},
}

func init() {
	clientCmd.AddCommand(clientListCmd)
}
