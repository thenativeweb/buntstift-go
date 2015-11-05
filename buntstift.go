package buntstift

import (
	"github.com/fatih/color"
	"strings"
	"fmt"
	"os/exec"
	"os"
)

// Params to set options
type Options struct {
	NoColor bool
}

type Buntstift struct {
	options Options
	icons   map[string]string
}

func Newbuntstift(params ...interface{}) *Buntstift {
	var b *Buntstift

	if len(params) > 0 {
		param, ok := params[0].(Options);
		if ok {
			b = &Buntstift{options: param}
		} else {
			panic("Wrong parameter type, must be Buntstift.Options")
		}
	} else {
		b = &Buntstift{options: Options{}}
	}

	if !isUtf8() {
		b.icons = ascii
	} else {
		b.icons = unicode
	}

	return b
}


func (b *Buntstift) colorize(values ...color.Attribute) *color.Color {
	c := color.New(values...)
	if b.options.NoColor {
		c.DisableColor()
	}

	return c
}

func (b *Buntstift) Success(text string) {
	output := b.colorize(color.FgGreen, color.Bold)
	output.Printf(b.icons["checkMark"] + " %v\n", text)
}

func (b *Buntstift) Error(text string) {
	output := b.colorize(color.FgRed, color.Bold)
	output.Printf("✗ %v\n", text)
}

func (b *Buntstift) Warn(text string) {
	output := b.colorize(color.FgYellow, color.Bold)
	output.Printf(b.icons["rightPointingPointer"] + " %v\n", text)
}

func (b *Buntstift) Info(text string) {
	output := b.colorize(color.FgWhite)
	output.Printf("  %v\n", text)
}

func (b *Buntstift) List(text string) {
	b.ListIndent(0, text)
}

func (b *Buntstift) ListIndent(level int, text string) {
	output := b.colorize(color.FgWhite)
	output.Printf("%v" + b.icons["multiplicationDot"] + " %v\n", strings.Repeat(" ", level * 2), text)
}

func (b *Buntstift) Line() {
	w, _ := b.getTerminalSize()
	output := b.colorize(color.FgWhite)
	output.Println(strings.Repeat("-", w))
}

func (b *Buntstift) getTerminalSize() (int, int) {
	var w, h int
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	d, _ := cmd.Output()
	fmt.Sscan(string(d), &h, &w)
	return w, h
}

func isUtf8() bool {
	env := os.Environ()
	for _, value := range env {
		if strings.Contains(value, "UTF-8") {
			return true
			break;
		}
	}
	return false
}