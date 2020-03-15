package cmd

import (
	"fmt"

	approvecmd "preq/cmd/approve"
	createcmd "preq/cmd/create"
	declinecmd "preq/cmd/decline"
	listcmd "preq/cmd/list"
	opencmd "preq/cmd/open"
	updatecmd "preq/cmd/update"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var rootCmd = &cobra.Command{
	Use:     "preq",
	Short:   "preq command-line utility for pull requests",
	Long:    `Command-line utility for all your pull request needs.`,
	Version: fmt.Sprintf("%v, commit %v, built at %v", version, commit, date),
	Run:     func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	rootCmd.AddCommand(createcmd.New())
	rootCmd.AddCommand(approvecmd.New())
	rootCmd.AddCommand(declinecmd.New())
	rootCmd.AddCommand(listcmd.New())
	rootCmd.AddCommand(opencmd.New())
	rootCmd.AddCommand(updatecmd.New())

	rootCmd.PersistentFlags().StringP("repository", "r", "", "repository in form of owner/repo")
	// TODO: Shorthand names for providers?
	rootCmd.PersistentFlags().StringP("provider", "p", "", "repository host, values - (bitbucket-cloud)")

	rootCmd.Execute()
}
