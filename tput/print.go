package tput

import "fmt"

const (
	TextColor int = iota
	UnderLine
	BoldText
)

type Option struct {
	Attribute int
	Color     int
}

func Printf(opts []*Option, format string, a ...interface{}) (int, error) {
	for _, opt := range opts {
		switch opt.Attribute {
		case TextColor:
			Setaf(opt.Color)
		case UnderLine:
			Smul()
		case BoldText:
			Bold()
		}
	}
	defer Sgr0()
	return fmt.Printf(format, a...)
}
