package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vahidid/kuknos-cli/controllers"
)

var featureCmd = &cobra.Command{
	Use:   "feature",
	Short: "Work with features branch",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run:   controllers.FeatureController,
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
