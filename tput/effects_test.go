package tput_test

import (
	"testing"

	"github.com/kotaoue/go-tput/tput"
)

func TestSetaf(t *testing.T) {
	if _, err := tput.Setaf(tput.Red); err != nil {
		t.Errorf("Setaf(%d) returned error: %v", tput.Red, err)
	}
}

func TestSmul(t *testing.T) {
	if _, err := tput.Smul(); err != nil {
		t.Errorf("Smul() returned error: %v", err)
	}
}

func TestBold(t *testing.T) {
	if _, err := tput.Bold(); err != nil {
		t.Errorf("Bold() returned error: %v", err)
	}
}

func TestSgr0(t *testing.T) {
	if _, err := tput.Sgr0(); err != nil {
		t.Errorf("Sgr0() returned error: %v", err)
	}
}
