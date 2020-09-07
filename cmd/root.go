package cmd

import (
	"fmt"
	"os"

	"github.com/enda-mullally/duckdns/console"
	"github.com/enda-mullally/duckdns/consts"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "duckdns",
	Short: "DuckDNS cli",
	Long:  console.GetFullBannerText(),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	cobra.MousetrapHelpText = "" // Disable the windows explorer mousetrap message.

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Version = consts.AppVersion
	rootCmd.Flags().Bool("version", false, "Displays the current version of this tool")
	rootCmd.SetVersionTemplate(versionTemplate)
}
