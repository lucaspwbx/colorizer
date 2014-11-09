package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Config struct {
	color Color
	style Style
}

type Color struct {
	fg string
	bg string
}

type Style struct {
	bold      bool
	underline bool
}

const (
	Red       = "31"
	Green     = "32"
	Yellow    = "33"
	Blue      = "34"
	Magenta   = "35"
	Cyan      = "36"
	White     = "37"
	Nc        = "0"
	bgRed     = "41"
	bgGreen   = "42"
	bgYellow  = "43"
	bgBlue    = "44"
	bgMagenta = "45"
	bgCyan    = "46"
	bgWhite   = "47"
	Normal    = "0"
	Bold      = "1"
	Underline = "4"
)

func colorize(text string, c *Color, s *Style) {
	echo := mount(text, c, s)
	cmd := exec.Command("echo", echo)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func mount(text string, c *Color, s *Style) string {
	osSpecific := "\x1B"
	echo := fmt.Sprintf("%s", osSpecific)
	if s.bold {
		echo += fmt.Sprintf("[%s;", Bold)
	} else {
		echo += fmt.Sprintf("[%s;", Normal)
	}
	if s.underline {
		echo += fmt.Sprintf("%s;", Underline)
	}
	echo += fmt.Sprintf("%s;", c.fg)
	echo += fmt.Sprintf("%sm", c.bg)
	echo += fmt.Sprintf("%s", text)
	echo += fmt.Sprintf("%s", "\x1B[0m")
	return echo
}

func main() {
	colorize("rammstein",
		&Color{fg: Red, bg: bgCyan},
		&Style{bold: false, underline: true},
	)
}
