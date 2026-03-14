package tput_test

import (
	"testing"

	"github.com/kotaoue/go-tput/tput"
)

func TestPrintf(t *testing.T) {
	opts := []*tput.Option{
		{Attribute: tput.TextColor, Color: tput.Cyan},
		{Attribute: tput.UnderLine},
		{Attribute: tput.BoldText},
	}
	if _, err := tput.Printf(opts, "%s\n", "hello"); err != nil {
		t.Errorf("Printf() returned error: %v", err)
	}
}
