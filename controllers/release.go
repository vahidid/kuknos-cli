package controllers

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vahidid/kuknos-cli/helpers"
)

func ReleaseController(cmd *cobra.Command, args []string) {

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Kuknos git tool not initialized! use 'kuknos git init' first")
		return
	}

	// kuknos git release <operation-name> <release-name>

	// Validation first argument is operation name
	helpers.ValidationArgs(args, 1)

	ValidateOperations(args[0])

	// Validation second argument is release name
	helpers.ValidationArgs(args, 2)

	branchName := args[1]
	releaseBranch := fmt.Sprint(viper.Get("release"))
	mainBranch := fmt.Sprint(viper.Get("main"))

	// Check if develop branch is not exists
	if !helpers.ExistsBranch(mainBranch) {
		fmt.Fprintln(os.Stderr, "Develop branch not exists!")
		return
	}

	switch args[0] {
	case "start":
		Start(releaseBranch+"/"+branchName, mainBranch)
	case "finish":
		Finish(releaseBranch+"/"+branchName, mainBranch, "Release")
		helpers.CreateTag(branchName)
		helpers.Push(branchName, "origin")
	case "push":
		helpers.Push(releaseBranch+"/"+branchName, "origin")
	case "pull":
		helpers.Pull(releaseBranch+"/"+branchName, "origin")
	}

}
