package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vahidid/kuknos-cli/helpers"
)

var operations = []string{"start", "finish", "push", "pull"}

func ValidateOperations(op string) {
	for _, operation := range operations {
		if op == operation {
			return
		}
	}

	panic(fmt.Sprintf("Operation %s is not valid", "'"+op+"'"))
}

func Start(branchName string, baseBranch string) {
	fmt.Println("Start Branch with name: " + branchName)

	helpers.Checkout(baseBranch)
	helpers.CreateBranch(branchName)
	helpers.Checkout(branchName)

	fmt.Println("You are now on branch: " + branchName)
}

func Finish(branchName, targetBranch, label string) {
	fmt.Println("Finish Branch : " + branchName)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Are you sure you want to finish this branch? (Y/n) ")
	ans, _ := reader.ReadString('\n')
	ans = strings.TrimSpace(ans)

	if ans == "Y" || ans == "y" || ans == "" {
		// Get merge request title
		fmt.Print("Enter merge request title: ")
		title, _ := reader.ReadString('\n')

		// Get merge request description
		fmt.Print("Enter merge request description: ")
		description, _ := reader.ReadString('\n')

		// helpers.Checkout(targetBranch)
		helpers.MergeRequest(branchName, targetBranch, strings.TrimSpace(title), strings.TrimSpace(description), label)
	}

	helpers.Checkout(targetBranch)

	fmt.Println("You are now on branch: " + targetBranch)
}
