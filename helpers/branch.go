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

func MergeRequest(source, target, title, description, label string) {

	// Checkout(target)
	// commandString := "push " + source + " -o merge_request.create"
	// commandString += " -o merge_request.target=\"" + target + "\""
	// commandString += " -o merge_request.title=\"" + title + "\""
	// commandString += " -o merge_request.description=\"" + description + "\""
	// commandString += " -o merge_request.label=\"" + label + "\""

	// Fetch branches
	fmt.Println(Fetch())

	branches := RunCMD("git", "branch", "-a")
	if !strings.Contains(branches, "remotes/origin/"+target) {
		Push(target, "origin")
	}

	args := []string{
		"push",
		"origin",
		source,
		"-o",
		"merge_request.create",
		"-o",
		"merge_request.target=" + target,
		"-o merge_request.title=\"" + title + "\"",
		"-o merge_request.description=\"" + description + "\"",
		"-o merge_request.label=" + label,
	}

	fmt.Println(RunCMD("git", args...))
	fmt.Printf("Branch %s successfully merged into %s", source, target)
}

func Push(branch, remote string) {

	fmt.Println(RunCMD("git", "push", "--set-upstream", remote, branch))
}

func Pull(branch, remote string) {

	fmt.Println(RunCMD("git", "pull", remote, branch))
}

func Fetch() string {
	return RunCMD("git", "fetch")
}
