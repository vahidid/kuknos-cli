package helpers

import "fmt"

func ValidationArgs(args []string, number int) {
	if len(args) < number {
		panic("You have to pass at least " + fmt.Sprint(number) + " parameter as operation")
	}
}
