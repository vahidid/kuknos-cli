package helpers

import "fmt"

func CreateTag(tagName string) {
	fmt.Println(RunCMD("git", "tag", tagName))
}
