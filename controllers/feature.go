package controllers

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/vahidid/kuknos-cli/helpers"
)

func FeatureController(cmd *cobra.Command, args []string) {
	if len(args) <= 1 {
		fmt.Println("You have to pass at least one parameter as branch name")
	}

	branchName := args[1]

	branchCommand := exec.Command("git", "branch")
	output, err := branchCommand.Output()
	if err != nil {
		log.Fatal(err)
		return
	}

	branches := helpers.SplitStringByByteArray(output, " ")

	k, exists := helpers.FindInStringSlice(branches, branchName)

	if exists {
		helpers.MergeBranch(branches[k], "develop")
		fmt.Printf("Successfully closed %s branch", branches[k])
	} else {
		helpers.CreateBranch(branchName)
		fmt.Printf("Successfully created %s branch", branchName)
	}

}
