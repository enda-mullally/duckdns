package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/enda-mullally/duckdns-cli/console"
	"github.com/enda-mullally/duckdns-cli/consts"
)

// Update - call the duckdns.org update endpoint. Basic HTTP request with url params.
func Update(domains string, token string, ip string) error {
	fmt.Print(console.GetBannerText())

	fmt.Println("Domains:", domains)
	fmt.Println("Token:", token)

	if len(ip) > 0 {
		fmt.Println("IP:", ip)
	}

	fmt.Println()
	fmt.Print("Updating...")

	url := fmt.Sprintf("https://www.duckdns.org/update?domains=%s&token=%s&ip=%s", domains, token, ip)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-Application-Name", "duckdns-cli")
	req.Header.Add("X-Application-Source", "https://github.com/enda-mullally/duckdns-cli")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

	fmt.Println()

	if string(body) == "OK" {
		fmt.Println("Update successful.")
	} else {
		fmt.Println("Failed")

		return fmt.Errorf("Update failed! %s didn't like your parameters", consts.DuckDNSOrg)
	}

	return nil
}
