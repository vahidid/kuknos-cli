package controllers

import (
	"fmt"

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
	fmt.Println("Start Branch with name: " + baseBranch + "/" + branchName)

	helpers.Checkout(baseBranch)
	helpers.CreateBranch(baseBranch + "/" + branchName)
	helpers.Checkout(baseBranch + "/" + branchName)

	fmt.Println("You are now on branch: " + baseBranch + branchName)
}

func Finish(branchName string, targetBranch string, mr bool) {
	fmt.Println("Finish Branch with name: " + branchName)

	helpers.Checkout(targetBranch)
	helpers.MergeRequest(branchName, targetBranch)

	fmt.Println("You are now on branch: master")
}
