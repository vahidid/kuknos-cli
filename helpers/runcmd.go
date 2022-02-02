package helpers

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/pterm/pterm"
)

func RunCMD(name string, arg ...string) string {
	cmd := exec.Command(name, arg...)

	pterm.Info.Println("Running command:", name, arg)
	// fmt.Println("Running: ", name, arg)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		// fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		pterm.Error.Println(stderr.String())
		// return ""
		os.Exit(1)
		// panic(stderr.String())
	}
	// fmt.Println(out.String())
	return out.String()
}
