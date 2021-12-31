package helpers

import (
	"fmt"
	"os/exec"
)

func Checkout(target string) {
	_, err := exec.Command("git", "checkout", target).Output()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Checkout to %s\t", target)
}

func CreateBranch(name string) {
	_, err := exec.Command("git", "branch", name).Output()

	if err != nil {
		panic(err)
	}

	fmt.Println("New Branch created successfully!")
}

func MergeBranch(source string, target string) {

	Checkout(target)

	_, err := exec.Command("git", "merge", source).Output()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Branch %s successfully merged into %s", source, target)
}
