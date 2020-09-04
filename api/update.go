package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/enda-mullally/duckdns-cli/console"
)

// Update - call the duckdns.org update endpoint. Basic HTTP request with url params.
func Update(domains string, token string, ip string) error {
	fmt.Print(console.GetBannerText())

	fmt.Println("Domains:", domains)
	fmt.Println("Token:", token)

	if ip != "" {
		fmt.Println("IP:", ip)
	}

	fmt.Println()
	fmt.Print("Updating...")

	url := "https://www.duckdns.org/update?domains=" + domains + "&token=" + token + "&ip=" + ip

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

		return fmt.Errorf("Update failed! DuckDNS.org didn't like your parameters")
	}

	return nil
}
