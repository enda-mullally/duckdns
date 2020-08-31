package cmd

import (
	"fmt"

	"github.com/enda-mullally/duckdns-cli/consts"
	"github.com/spf13/cobra"
)

var versionTemplate = "Version: " + consts.AppVersion

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v", "ver"},
	Short:   "Displays the current version of this tool",
	Hidden:  true,
	Long:    "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(versionTemplate)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
