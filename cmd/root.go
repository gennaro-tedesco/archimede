package cmd

import (
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "dirinfo",
	Short: "Fetch directory info",
	Long:  "Unobtrusive directory information fetcher",
	Run: func(cmd *cobra.Command, args []string) {
		displayInfo()
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.SetHelpTemplate(getRootHelp())
}

func getRootHelp() string {
	return `
help page
`
}
