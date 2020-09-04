package cmd

import (
	"fmt"

	"github.com/enda-mullally/duckdns-cli/api"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update duckdns domain(s). See: `duckdns update --help` for further details",
	Long: `Update duckdns domain(s). For example:
  
  duckdns update -domains "dns01,dns02" -token "4292f3a3-2842-4ae3-8978-*************"`,

	RunE: func(cmd *cobra.Command, args []string) error {
		domains, err := cmd.Flags().GetString("domains")

		if err != nil {
			return err
		}

		if len(domains) == 0 {
			return fmt.Errorf("domains cannot be empty")
		}

		token, err := cmd.Flags().GetString("token")

		if err != nil {
			return err
		}

		if len(token) == 0 {
			return fmt.Errorf("token cannot be empty")
		}

		ip, err := cmd.Flags().GetString("ip")

		if err != nil {
			return err
		}

		return api.Update(domains, token, ip)
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
