package cmd

import (
	"fmt"
	"os"

	"github.com/enda-mullally/duckdns/api"
	"github.com/enda-mullally/duckdns/consts"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: fmt.Sprintf("Update your %s domain(s). See: `duckdns update --help` for further details", consts.DuckDNSOrg),
	Long: fmt.Sprintf(`Update your %s domain(s). For example:
  
  duckdns update --domains "a9275" --token "4292f3a3-2842-4ae3-8978-*************"
  
  duckdns update --env`, consts.DuckDNSOrg),

	PreRun: func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetBool("env")

		if env {
			// The user has set the --env flag
			// Manually add the required and optional flags reading from env
			cmd.Flags().Set("domains", os.Getenv(consts.UpdateDomainsEnvName))
			cmd.Flags().Set("token", os.Getenv(consts.UpdateTokenEnvName))
			ip := os.Getenv(consts.UpdateIPEnvName)

			if len(ip) > 0 {
				cmd.Flags().Set("ip", ip)
			}
		}
	},

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

	updateCmd.Flags().Bool("env", false, fmt.Sprintf("If this flag is provided, use environment variables '%s' for domain(s) and '%s' for token. All other flags will be ignored", consts.UpdateDomainsEnvName, consts.UpdateTokenEnvName))
	updateCmd.Flags().String("domains", "", fmt.Sprintf("The list of your %s domain(s) to update, e.g \"mydns01,mydns02\"", consts.DuckDNSOrg))
	updateCmd.Flags().String("token", "", fmt.Sprintf("Your %s token, e.g \"4292f3a3-2842-4ae3-8978-*************\"", consts.DuckDNSOrg))
	updateCmd.Flags().String("ip", "", fmt.Sprintf("The preferred ip address to set for these domain(s). If not provided %s will use your current external ip address", consts.DuckDNSOrg))

	updateCmd.MarkFlagRequired("domains")
	updateCmd.MarkFlagRequired("token")
}
