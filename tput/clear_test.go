package tput_test

import (
	"testing"

	"github.com/kotaoue/go-tput/tput"
)

func TestEl(t *testing.T) {
	if _, err := tput.El(); err != nil {
		t.Skipf("el not supported by this terminal: %v", err)
	}
}

func TestEl1(t *testing.T) {
	if _, err := tput.El1(); err != nil {
		t.Skipf("el1 not supported by this terminal: %v", err)
	}
}

func TestEd(t *testing.T) {
	if _, err := tput.Ed(); err != nil {
		t.Skipf("ed not supported by this terminal: %v", err)
	}
}

func TestClear(t *testing.T) {
	if _, err := tput.Clear(); err != nil {
		t.Skipf("clear not supported by this terminal: %v", err)
	}
}
