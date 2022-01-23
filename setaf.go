package tput

import (
	"os/exec"
	"strconv"
)

func Setaf(i int) ([]byte, error) {
	return output(exec.Command("tput", "setaf", strconv.Itoa(i)))
}
