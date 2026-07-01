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

var clientShowCmd = &cobra.Command{
	Use:   "show <email|uuid>",
	Short: "Show client information",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		cfg, err := xray.Read()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		key := args[0]

		client := xray.FindClientByEmail(cfg, key)
		if client == nil {
			client = xray.FindClientByUUID(cfg, key)
		}

		if client == nil {
			fmt.Println("Client not found")
			os.Exit(1)
		}

		fmt.Println("Reality client")
		fmt.Println("--------------")
		fmt.Printf("Email : %s\n", client.Email)
		fmt.Printf("UUID  : %s\n", client.ID)
		fmt.Printf("Port  : %d\n", xray.RealityPort(cfg))
	},
}

func init() {
	clientCmd.AddCommand(clientListCmd)
	clientCmd.AddCommand(clientShowCmd)
}
