package helpers

import (
	"fmt"
	"os/exec"
)

func Checkout(target string) {
	// cmd, err := exec.Command("git", "checkout", target).Output()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Checkout to %s\t", target)

	RunCMD("git", "checkout", target)
}

func CreateBranch(name string) {

	_, err := exec.Command("git", "branch", name).Output()

	if err != nil {
		panic(err)
	}

	fmt.Println("New Branch created successfully!")
}

func MergeBranch(source string, target string, mr bool) {

	Checkout(target)

	_, err := exec.Command("git", "merge", source, "-o ", "merge_request.create").Output()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Branch %s successfully merged into %s", source, target)
}

func MergeRequest(source string, target string) {

	Checkout(target)

	_, err := exec.Command("git", "push", "origin-demo", source, "-o", "merge_request.create").Output()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Branch %s successfully merged into %s", source, target)
}
