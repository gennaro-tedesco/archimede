package cmd

import (
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "archimede",
	Short: "Fetch directory info",
	Long:  "Unobtrusive directory information fetcher",
	Run: func(cmd *cobra.Command, args []string) {
		fileFormat, _ := cmd.Flags().GetString("file")
		textColour, _ := cmd.Flags().GetString("colour")
		separator, _ := cmd.Flags().GetString("separator")
		displayInfo(fileFormat, textColour, separator)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringP("file", "f", "long", "short/long file format")
	rootCmd.Flags().StringP("colour", "c", "cyan", "text colour")
	rootCmd.Flags().StringP("separator", "s", ":", "key-value separator character")
	rootCmd.SetHelpTemplate(getRootHelp())
}

func getRootHelp() string {
	return `
help page
`
}
