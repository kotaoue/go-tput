package tput

import (
	"os/exec"
)

func Sgr0() ([]byte, error) {
	return output(exec.Command("tput", "sgr0"))
}
