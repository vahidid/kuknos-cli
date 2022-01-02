package helpers

import (
	"fmt"
)

func Checkout(target string) {
	RunCMD("git", "checkout", target)
	fmt.Printf("Checkout to %s\t", target)
}

func CreateBranch(name string) {

	RunCMD("git", "branch", name)

	fmt.Println("New Branch created successfully!")
}

func MergeBranch(source string, target string, mr bool) {

	Checkout(target)

	RunCMD("git", "merge", source, "-o ", "merge_request.create")

	fmt.Printf("Branch %s successfully merged into %s", source, target)
}

func MergeRequest(source string, target string) {

	Checkout(target)

	RunCMD("git", "push", "origin-demo", source, "-o", "merge_request.create")

	fmt.Printf("Branch %s successfully merged into %s", source, target)
}
