package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "archimede",
	Short: "Fetch directory info",
	Long:  "Unobtrusive directory information fetcher",
	Run: func(cmd *cobra.Command, args []string) {
		short, _ := cmd.Flags().GetBool("short")
		git, _ := cmd.Flags().GetBool("git")
		textColour, _ := cmd.Flags().GetString("colour")
		delimiter, _ := cmd.Flags().GetString("delimiter")
		displayInfo(
			short,
			git,
			textColour,
			delimiter)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().BoolP("short", "s", false, "print short format: default false")
	rootCmd.Flags().BoolP("git", "g", false, "include .git folder in files stats: default false")
	rootCmd.Flags().StringP("colour", "c", "cyan", "text colour")
	rootCmd.Flags().StringP("delimiter", "d", " ", "key-value delimiter character")
	rootCmd.SetHelpTemplate(getRootHelp())
}

func getRootHelp() string {
	return `
help page
`
}
