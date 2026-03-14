package tput_test

import (
	"testing"

	"github.com/kotaoue/go-tput/tput"
)

func TestColorConstants(t *testing.T) {
	tests := []struct {
		name string
		got  int
		want int
	}{
		{"Black", tput.Black, 0},
		{"Red", tput.Red, 1},
		{"Green", tput.Green, 2},
		{"Yellow", tput.Yellow, 3},
		{"Blue", tput.Blue, 4},
		{"Magenta", tput.Magenta, 5},
		{"Cyan", tput.Cyan, 6},
		{"White", tput.White, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("tput.%s = %d, want %d", tt.name, tt.got, tt.want)
			}
		})
	}
}
