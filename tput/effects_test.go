package tput_test

import (
	"testing"

	"github.com/kotaoue/go-tput/tput"
)

func TestSetaf(t *testing.T) {
	if _, err := tput.Setaf(tput.Red); err != nil {
		t.Skipf("setaf not supported by this terminal: %v", err)
	}
}

func TestSmul(t *testing.T) {
	if _, err := tput.Smul(); err != nil {
		t.Skipf("smul not supported by this terminal: %v", err)
	}
}

func TestBold(t *testing.T) {
	if _, err := tput.Bold(); err != nil {
		t.Skipf("bold not supported by this terminal: %v", err)
	}
}

func TestSgr0(t *testing.T) {
	if _, err := tput.Sgr0(); err != nil {
		t.Skipf("sgr0 not supported by this terminal: %v", err)
	}
}
