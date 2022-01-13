package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vahidid/kuknos-cli/controllers"
)

var improvementCmd = &cobra.Command{
	Use:   "improvement",
	Short: "Work with improvements branch",
	Long:  `Use it for improve existing features and maybe refactor or change styles and UI!`,
	Run:   controllers.ImprovementController,
}

func init() {

}
