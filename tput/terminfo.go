package tput

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func Cols() (int, error) {
	s, err := exec.Command("tput", "cols").Output()
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(strings.Trim(string(s), "\n"))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func HR() error {
	tCol, _ := Cols()

	if tCol == 0 {
		tCol = 64
	}

	_, err := fmt.Println(strings.Repeat("-", tCol))
	return err
}
