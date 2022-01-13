package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vahidid/kuknos-cli/controllers"
)

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Work with release branch",
	Long:  `Release branches uses for publishing new version of your application! and so documenting or clear logs and clean up comments or anything like these!`,
	Run:   controllers.ReleaseController,
}
