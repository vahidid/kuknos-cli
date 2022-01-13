package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vahidid/kuknos-cli/controllers"
)

var improvementCmd = &cobra.Command{
	Use:   "improvement",
	Short: "Work with improvements branch",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run:   controllers.ImprovementController,
}

func init() {

}
