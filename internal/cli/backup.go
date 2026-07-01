package cli

import (
	"fmt"
	"os"

	"github.com/soulcpd1/xrm2/internal/backup"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup operations",
}

var backupCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Xray config backup",
	Run: func(cmd *cobra.Command, args []string) {

		msg, err := backup.Create()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(msg)
	},
}

func init() {
	backupCmd.AddCommand(backupCreateCmd)
}
