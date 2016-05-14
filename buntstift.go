package buntstift

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/tj/go-spin"
)

// Output defines the standard output of the print functions. By default
// os.Stdout from color is used.
var Output = color.Output

// Options struct
type Options struct {
	NoColor bool
	NoUtf8  bool
}

// Buntstift struct
type Buntstift struct {
	options Options
	icons   map[string]string
}

func (b *Buntstift) unsetColor() {
	if b.options.NoColor {
		return
	}
	color.Unset()
}

func (b *Buntstift) printf(Color *color.Color, format string, a ...interface{}) (n int, err error) {
	Color.Set()
	defer b.unsetColor()
	return fmt.Fprintf(Output, format, a...)
}

func (b *Buntstift) colorize(values ...color.Attribute) *color.Color {
	Color := color.New(values...)
	if b.options.NoColor {
		Color.DisableColor()
	}
	return Color
}

func (b *Buntstift) spin(stop chan bool) {
	spinner := spin.New()

loop:
	for {
		select {
		case <-stop:
			fmt.Fprintf(os.Stderr, "\r")
			break loop

		default:
			fmt.Fprintf(os.Stderr, "\r%s", spinner.Next())
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func (b *Buntstift) getTerminalSize() (int, int) {
	var width, height int
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, _ := cmd.Output()
	fmt.Sscan(string(output), &height, &width)
	return width, height
}

/*
	Public Api
*/

// New Buntstift returns Buntstift instance
func New(params ...interface{}) *Buntstift {
	var b *Buntstift

	if len(params) > 0 {
		param, ok := params[0].(Options)
		if !ok {
			panic("Wrong parameter type, must be Buntstift.Options")
		}
		b = new(Buntstift)
		b.options = param
	} else {
		b = &Buntstift{options: Options{}}
	}

	b.icons = unicode
	if b.options.NoUtf8 {
		b.icons = ascii
	}

	return b
}

// Error prints a red cross mark and text
func (b *Buntstift) Error(text string) {
	Color := b.colorize(color.FgRed, color.Bold)
	b.printf(Color, "%v %v\n", b.icons["crossMark"], text)
}

// Info prints white text
func (b *Buntstift) Info(text string) {
	Color := b.colorize(color.FgWhite)
	b.printf(Color, "  %v\n", text)
}

// Line prints a white dashed line
func (b *Buntstift) Line() {
	width, _ := b.getTerminalSize()
	Color := b.colorize(color.FgWhite)
	b.printf(Color, "%v", strings.Repeat("-", width))
}

// List prints a white list entry with optional indentation paramter
func (b *Buntstift) List(text string, indentLevel ...int) {
	indent := 0
	if len(indentLevel) > 0 {
		indent = indentLevel[0]
	}
	Color := b.colorize(color.FgWhite)
	b.printf(Color, "%v%v %v\n", strings.Repeat(" ", indent*2), b.icons["multiplicationDot"], text)
}

// NewLine prints new line
func (b *Buntstift) NewLine() {
	fmt.Fprintf(Output, "\r \n")
}

// Success prints a green check mark and text
func (b *Buntstift) Success(text string) {
	Color := b.colorize(color.FgGreen, color.Bold)
	b.printf(Color, "%v %v\n", b.icons["checkMark"], text)
}

// Warn prints a yellow arrow and text
func (b *Buntstift) Warn(text string) {
	Color := b.colorize(color.FgYellow, color.Bold)
	b.printf(Color, "%v %v\n", b.icons["rightPointingPointer"], text)
}

// WaitFor prints a white spinner until stop <- true
func (b *Buntstift) WaitFor(worker func(stop chan bool)) {
	stop := make(chan bool)
	go b.spin(stop)
	worker(stop)
}
