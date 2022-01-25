package tput

import (
	"os/exec"
	"strconv"
)

func Setaf(i int) ([]byte, error) {
	return output(exec.Command("tput", "setaf", strconv.Itoa(i)))
}

func Smul() ([]byte, error) {
	return output(exec.Command("tput", "smul"))
}

func Bold() ([]byte, error) {
	return output(exec.Command("tput", "bold"))
}

func Sgr0() ([]byte, error) {
	return output(exec.Command("tput", "sgr0"))
}
