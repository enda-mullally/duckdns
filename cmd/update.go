package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update duckdns domain(s). See: `duckdns update --help` for further details",
	Long: `Update duckdns domain(s). For example:
  
  duckdns update -domains "dns01,dns02" -token 4292f3a3-2842-4ae3-8978-5c1236721f**`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Update starting...")
		fmt.Println("OK")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().String("domains", "", "The list of your DuckDNS.org domain(s) to update, e.g \"mydns01,mydns02\"")
	updateCmd.Flags().String("token", "", "Your DuckDNS.org token, e.g 4292f3a3-2842-4ae3-8978-5c1236721f**")
	updateCmd.Flags().String("ip", "", "The preferred ip address to set for these domain(s). If not provided DuckDNS.org will automatically use your current external ip address.")

	updateCmd.MarkFlagRequired("domains")
	updateCmd.MarkFlagRequired("token")
}
