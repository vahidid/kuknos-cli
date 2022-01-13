package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vahidid/kuknos-cli/controllers"
)

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run:   controllers.ReleaseController,
}
