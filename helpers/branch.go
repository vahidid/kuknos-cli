package helpers

import (
	"fmt"
	"strings"
)

func ExistsBranch(name string) bool {
	return strings.Contains(RunCMD("git", "branch"), name)

}

func Checkout(target string) {

	fmt.Println(RunCMD("git", "checkout", target))
}

func CreateBranch(name string) {

	fmt.Println(RunCMD("git", "branch", name))
}

func MergeBranch(source string, target string, mr bool) {

	Checkout(target)
	fmt.Println(RunCMD("git", "merge", source, "-o ", "merge_request.create"))
	fmt.Printf("Branch %s successfully merged into %s", source, target)
}

func MergeRequest(source string, target string) {

	Checkout(target)

	fmt.Println(RunCMD("git", "push", "origin-demo", source, "-o", "merge_request.create"))

	fmt.Printf("Branch %s successfully merged into %s", source, target)
}
