# go-tput

[![Go](https://github.com/kotaoue/go-tput/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/kotaoue/go-tput/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/kotaoue/go-tput/branch/main/graph/badge.svg)](https://codecov.io/gh/kotaoue/go-tput)
[![Go Report Card](https://goreportcard.com/badge/github.com/kotaoue/go-tput)](https://goreportcard.com/report/github.com/kotaoue/go-tput)
[![License](https://img.shields.io/github/license/kotaoue/go-tput)](https://github.com/kotaoue/go-tput/blob/main/LICENSE)

Execute [tput](https://man7.org/linux/man-pages/man1/tput.1.html) commands from Go.  
`go-tput` provides a thin wrapper around the `tput` terminal utility, letting you apply colors, text effects, and terminal control sequences directly from your Go programs.

## Installation

```bash
go get github.com/kotaoue/go-tput
```

## Import
```go
import "github.com/kotaoue/go-tput/tput"
```

## Usage

### Terminal Info

| Function | Description |
| ---- | ---- |
| `tput.Cols() (int, error)` | Return the number of columns in the terminal |
| `tput.HR() error` | Print a horizontal rule that spans the terminal width |

### Colors

Use `tput.Setaf(i int)` to set the foreground color. The following color constants are provided:

| Constant | Value | Color |
| ---- | ---- | ---- |
| `tput.Black` | 0 | Black |
| `tput.Red` | 1 | Red |
| `tput.Green` | 2 | Green |
| `tput.Yellow` | 3 | Yellow |
| `tput.Blue` | 4 | Blue |
| `tput.Magenta` | 5 | Magenta |
| `tput.Cyan` | 6 | Cyan |
| `tput.White` | 7 | White |

### Text Effects

| Function | Description |
| ---- | ---- |
| `tput.Setaf(i int) ([]byte, error)` | Set foreground color (use color constants above) |
| `tput.Smul() ([]byte, error)` | Start underlined text |
| `tput.Bold() ([]byte, error)` | Start bold text |
| `tput.Sgr0() ([]byte, error)` | Turn off all attributes |

### Clear / Erase

| Function | Description |
| ---- | ---- |
| `tput.El() ([]byte, error)` | Erase to end of line |
| `tput.El1() ([]byte, error)` | Erase to beginning of line |
| `tput.Ed() ([]byte, error)` | Erase to end of screen |
| `tput.Clear() ([]byte, error)` | Clear the screen |

### Printf helper

`tput.Printf` applies one or more display options before printing formatted output, and resets all attributes afterwards.

```go
type Option struct {
    Attribute int  // TextColor, UnderLine, or BoldText
    Color     int  // used when Attribute == TextColor
}
```

| Attribute constant | Description |
| ---- | ---- |
| `tput.TextColor` | Set foreground color (requires `Color` field) |
| `tput.UnderLine` | Underline the text |
| `tput.BoldText` | Bold the text |

## Examples

### Build a styled CLI status report

This example shows a typical use case: printing a terminal report with colored status indicators, bold section headers, and horizontal rules.

```go
package main

import (
    "fmt"

    "github.com/kotaoue/go-tput"
)

func main() {
    tput.HR()
    tput.Printf([]*tput.Option{{Attribute: tput.BoldText}}, "Build Report\n")
    tput.HR()

    tput.Printf([]*tput.Option{{Attribute: tput.TextColor, Color: tput.Green}}, "[PASS] ")
    fmt.Println("All tests passed")

    tput.Printf([]*tput.Option{{Attribute: tput.TextColor, Color: tput.Yellow}}, "[WARN] ")
    fmt.Println("Deprecated API usage detected")

    tput.Printf([]*tput.Option{
        {Attribute: tput.TextColor, Color: tput.Red},
        {Attribute: tput.BoldText},
    }, "[FAIL] ")
    fmt.Println("Build failed: missing dependency")

    tput.HR()
}
```

Output:

![Build report showing bold header, green [PASS] for all tests passed, yellow [WARN] for deprecated API usage, and bold red [FAIL] for missing dependency](docs/images/example.png)

### Print a colored message

```go
package main

import (
    "github.com/kotaoue/go-tput"
)

func main() {
    tput.Printf(
        []*tput.Option{
            {Attribute: tput.TextColor, Color: tput.Red},
        },
        "Hello, %s!\n", "world",
    )
}
```

### Print bold underlined text

```go
package main

import (
    "github.com/kotaoue/go-tput"
)

func main() {
    tput.Printf(
        []*tput.Option{
            {Attribute: tput.BoldText},
            {Attribute: tput.UnderLine},
        },
        "Important notice\n",
    )
}
```

### Use low-level color / effect functions directly

```go
package main

import (
    "fmt"
    "log"

    "github.com/kotaoue/go-tput"
)

func main() {
    if _, err := tput.Setaf(tput.Green); err != nil {
        log.Fatal(err)
    }
    if _, err := tput.Bold(); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Success!")
    tput.Sgr0()
}
```
