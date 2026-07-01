package cli

import (
	"fmt"
	"os"

	"github.com/soulcpd1/xrm2/internal/xray"
	"github.com/spf13/cobra"
)

var clientListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Reality clients",
	Run: func(cmd *cobra.Command, args []string) {

		cfg, err := xray.Read()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		port := xray.RealityPort(cfg)
		clients := xray.RealityClients(cfg)

		fmt.Printf("Reality inbound (port %d)\n", port)
		fmt.Printf("Clients: %d\n\n", len(clients))

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
