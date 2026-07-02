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
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {

		engine := xray.NewEngine()

		cfg, err := engine.Load()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		inb := cfg.RealityInbound()
		if inb == nil {
			fmt.Println("no reality inbound found")
			return
		}

		fmt.Println("Reality clients")
		fmt.Printf("Clients: %d\n\n", len(inb.Settings.Clients))

		for i, c := range inb.Settings.Clients {
			email := c.Email
			if email == "" {
				email = "-"
			}
			fmt.Printf("%2d. %-30s %s\n", i+1, email, c.ID)
		}
	},
}

var clientAddCmd = &cobra.Command{
	Use:  "add <email>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		err := xray.ApplyWithEngine(func(cfg *xray.Config) error {
			return cfg.AddClient(xray.ClientMutation{
				Email: args[0],
			})
		})

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("Client added:", args[0])
	},
}

func init() {
	clientCmd.AddCommand(clientListCmd)
	clientCmd.AddCommand(clientAddCmd)
}
