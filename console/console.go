package console

// GetBannerText - Returns the application banner text - ascii text from: http://www.patorjk.com/software/taag/#p=display&h=1&f=Standard&t=DuckDNS
func GetBannerText() string {

	var bannerText = ``

	bannerText += ` ____                _     ____   _   _  ____  ` + "\r\n"
	bannerText += `|  _ \  _   _   ___ | | __|  _ \ | \ | |/ ___| ` + "\r\n"
	bannerText += `| | | || | | | / __|| |/ /| | | ||  \| |\___ \ ` + "\r\n"
	bannerText += `| |_| || |_| || (__ |   < | |_| || |\  | ___) |` + "\r\n"
	bannerText += `|____/  \__,_| \___||_|\_\|____/ |_| \_||____/ ` + "\r\n"

	bannerText += "\r\n"

	bannerText += "DuckDNS is a command line tool (cli) for the popular DuckDNS.org service. https://www.duckdns.org/"

	bannerText += "\r\n"

	return bannerText
}
