package tput_test

import (
	"fmt"

	"github.com/kotaoue/go-tput/tput"
)

func ExampleCols() {
	fmt.Println(tput.Cols())
}

func ExampleHR() {
	tput.HR()
	// Output: ----------------------------------------------------------------
}
