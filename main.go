package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	tCol, err := terminalColumnLength()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(strings.Repeat("-", tCol))
}

func terminalColumnLength() (int, error) {
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
