package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vahidid/kuknos-cli/controllers"
)

var bugfixCmd = &cobra.Command{
	Use:   "bugfix",
	Short: "Work with bugfix branch",
	Long:  `Create a new bugfix branch for debugging sprints!`,
	Run:   controllers.BugfixController,
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// featureCmd.Flags().BoolP("mr", "", false, "create merge request")
}
