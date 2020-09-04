package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/enda-mullally/duckdns-cli/console"

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

		return doUpdate(domains, token, ip)
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

func doUpdate(domains string, token string, ip string) error {
	fmt.Print(console.GetBannerText())

	fmt.Println("Update started...")
	fmt.Println("Domains:", domains)
	fmt.Println("Token:", token)

	if ip != "" {
		fmt.Println("IP:", ip)
	}

	fmt.Println()

	url := "https://www.duckdns.org/update?domains=" + domains + "&token=" + token + "&ip=" + ip

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-Application-Name", "duckdns-cli")
	req.Header.Add("X-Application-Source", "https://github.com/enda-mullally/duckdns-cli")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body) // perhaps checking for a StatusCode of 200 is safer

	if string(body) == "OK" {
		fmt.Println("Update successful.")
	} else {
		return fmt.Errorf("Update failed! Please check your params")
	}

	return nil
}
