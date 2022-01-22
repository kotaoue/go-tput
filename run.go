package tput

import (
	"fmt"
	"os/exec"
)

func output(cmd *exec.Cmd) ([]byte, error) {
	out, err := cmd.Output()
	fmt.Printf("%s", out)
	return out, err
}
