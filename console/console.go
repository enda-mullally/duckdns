package console

import "github.com/enda-mullally/duckdns-cli/consts"

// GetBannerText - Returns the application banner text - ascii text from: http://www.patorjk.com/software/taag/#p=display&h=1&f=Standard&t=DuckDNS
func GetBannerText() string {

	var bannerText = ``

	bannerText += ` ____                _     ____   _   _  ____  ` + consts.NewLine
	bannerText += `|  _ \  _   _   ___ | | __|  _ \ | \ | |/ ___| ` + consts.NewLine
	bannerText += `| | | || | | | / __|| |/ /| | | ||  \| |\___ \ ` + consts.NewLine
	bannerText += `| |_| || |_| || (__ |   < | |_| || |\  | ___) |` + consts.NewLine
	bannerText += `|____/  \__,_| \___||_|\_\|____/ |_| \_||____/ ` + consts.NewLine

	bannerText += consts.NewLine

	bannerText += "DuckDNS is a command line tool (cli) for the popular DuckDNS.org service. https://www.duckdns.org/"

	bannerText += consts.NewLine

	return bannerText
}
