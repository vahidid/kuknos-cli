package controllers

import (
	"fmt"

	"github.com/vahidid/kuknos-cli/helpers"
)

var operations = []string{"start", "finish", "push", "pull"}

func ValidateOperations(op string) {
	for _, operation := range operations {
		if op == operation {
			break
		}
		panic(fmt.Sprintf("Operation %s is not valid", "'"+op+"'"))
	}
}

func Start(branchName string, baseBranch string) {
	fmt.Println("Start Branch with name: " + branchName)

	helpers.Checkout(baseBranch)
	helpers.CreateBranch(branchName)
	helpers.Checkout(branchName)

	fmt.Println("You are now on branch: " + branchName)
}

func Finish(branchName string, targetBranch string) {
	fmt.Println("Finish Branch with name: " + branchName)

	helpers.Checkout(targetBranch)
	helpers.MergeBranch(branchName, targetBranch)

	fmt.Println("You are now on branch: master")
}
