package helpers

import (
	"bytes"
	"fmt"
	"os/exec"
)

func RunCMD(name string, arg ...string) string {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		// fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		panic(stderr.String())
	}
	fmt.Println(out.String())
	return out.String()
}
