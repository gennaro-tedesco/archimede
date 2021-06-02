package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "archimede",
	Args:  cobra.NoArgs,
	Short: "Fetch directory info",
	Long:  "Unobtrusive directory information fetcher",
	Run: func(cmd *cobra.Command, args []string) {
		short, _ := cmd.Flags().GetBool("short")
		git, _ := cmd.Flags().GetBool("git")
		excludeDir, _ := cmd.Flags().GetString("exclude-dir")
		excludeFile, _ := cmd.Flags().GetString("exclude-file")
		textColour, _ := cmd.Flags().GetString("colour")
		delimiter, _ := cmd.Flags().GetString("delimiter")
		displayInfo(
			short,
			git,
			excludeDir,
			excludeFile,
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
	rootCmd.Flags().StringP("exclude-dir", "e", "", "directory to exclude from counts and stats")
	rootCmd.Flags().StringP("exclude-file", "v", "", "file extension to exclude from counts and stats")
	rootCmd.Flags().StringP("colour", "c", "cyan", "text colour")
	rootCmd.Flags().StringP("delimiter", "d", " ", "key-value delimiter character")
	rootCmd.SetHelpTemplate(getRootHelp())
}

func getRootHelp() string {
	return `
archimede: unobtrusive project information fetcher

Description:
  archimede fetches project information of the directory
  it is invoked from. Available information

  - Path:  the current working directory path
  - Files: count of regular + hidden files, breakdown per extension
		   (for the breakdown to be shown more than two extension types
		   must be present)
  - Dirs:  count of directories with breakdown in top level, nested (sub-folders)
  		   and sub-nested (sub-sub-folders and more)
  - Git:   if in a git repository, show the git branch and status in terms
  		   of number of files in the staged index and modified
  - Space: the disk usage of the entire directory

Usage:
  archimede [flag]

Flags:
  -h, --help
	open this help page

  -s, --short
	whether to display output in pretty long format (default)
	or short one. For files and directories the short format
	only displays numerical counts without descriptors; for
	git status it only displays the current branch name.
	Notice that the default long format makes use of terminal
	unicode characters that may not render perfectly with all
	terminals or fonts: if so use -s.

	type boolean: archimede -s

  -g, --git
	whether to include the /.git folder in calculating total
	file counts and distribution. Defaults to false.
	Equivalent to invoking archimede -e .git

	type boolean: archimede -s

  -e, --exclude-dir
	directory to explicitly exclude from file counts and
	statistics. Defaults to an empty string "" which is
	checked against; provide a non empty string to exclude
	a directory explicitly.

	type string: archimede -e test

  -v, --exclude-file
	file extension to explicitly exclude from file counts
	and statistics. Defaults to an empty string "" which is
	checked against; provide a non empty file extension to
	exclude a file type explicitly. Notice that the extension
	must contain the dot(.)

	type string: archimede -v .md

  -c, --colour
	colour to display the text. One of "blue", "red", "cyan",
	"magenta", "yellow", "green", "black", "white".

	type string: archimede -c magenta

  -d, --delimiter
	delimiter separator between key and value, in the output.
	Defaults to the empty string (no separator, just space).

	type string: archimede -d ":"

Help commands:
  version     print current version
`
}
