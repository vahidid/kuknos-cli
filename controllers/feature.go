package controllers

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/vahidid/kuknos-cli/helpers"
)

func FeatureController(cmd *cobra.Command, args []string) {

	fmt.Println(args)
	if len(args) < 1 {
		fmt.Println("You have to pass at least one parameter as operation")
		return
	}
	fmt.Println("You have passed operation: ", "'"+args[0]+"'")
	ValidateOperations(args[0])

	if len(args) < 2 {
		fmt.Println("You have to pass at least one parameter as branch name after operation")
		return
	}

	branchName := args[1]

	branchCommand := exec.Command("git", "branch")
	output, err := branchCommand.Output()
	if err != nil {
		log.Fatal(err)
		return
	}

	branches := helpers.SplitStringByByteArray(output, " ")

	exists := helpers.FindInStringSlice(branches, branchName)
	fmt.Println(branches[3])
	fmt.Println(exists)

	// if exists {
	// 	helpers.MergeBranch(branchName, "develop")
	// 	fmt.Printf("Successfully closed %s branch", branchName)
	// } else {
	// 	helpers.CreateBranch(branchName)
	// 	fmt.Printf("Successfully created %s branch", branchName)
	// }

}
