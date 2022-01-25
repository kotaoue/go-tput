package tput

import (
	"os/exec"
)

func El() ([]byte, error) {
	return output(exec.Command("tput", "el"))
}

func El1() ([]byte, error) {
	return output(exec.Command("tput", "el1"))
}

func Ed() ([]byte, error) {
	return output(exec.Command("tput", "ed"))
}

func Clear() ([]byte, error) {
	return output(exec.Command("tput", "clear"))
}
