package controllers

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vahidid/kuknos-cli/helpers"
)

func FeatureController(cmd *cobra.Command, args []string) {

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Kuknos git tool not initialized! use 'kuknos git init' first")
		return
	}

	// kuknos git feature <operation-name> <feature-name>

	// Validation first argument is operation name
	helpers.ValidationArgs(args, 1)

	ValidateOperations(args[0])

	// Validation second argument is feature name
	helpers.ValidationArgs(args, 2)

	branchName := args[1]
	developBranch := fmt.Sprint(viper.Get("developbranch"))

	switch args[0] {
	case "start":
		Start(branchName, developBranch)
	case "finish":
		Finish(branchName, developBranch, true)
	case "push":
		helpers.Checkout(branchName)
	case "pull":
		helpers.Checkout(branchName)
	}

}
