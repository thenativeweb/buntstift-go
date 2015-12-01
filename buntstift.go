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

// Output defines the standard Color of the print functions. By default
// os.Stdout from color is used.
var Output = color.Output

// Options ...
type Options struct {
	NoColor bool
	NoUtf8  bool
}

// Buntstift ...
type Buntstift struct {
	options Options
	icons   map[string]string
}

// ListOptions ...
type ListOptions struct {
	Indent int
}

// New Buntstift
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

// Success ...
func (b *Buntstift) Success(text string) {
	Color := b.colorize(color.FgGreen, color.Bold)
	b.printf(Color, "%v %v\n", b.icons["checkMark"], text)
}

// Error ...
func (b *Buntstift) Error(text string) {
	Color := b.colorize(color.FgRed, color.Bold)
	b.printf(Color, "%v %v\n", b.icons["crossMark"], text)
}

// Warn ...
func (b *Buntstift) Warn(text string) {
	Color := b.colorize(color.FgYellow, color.Bold)
	b.printf(Color, "%v %v\n", b.icons["rightPointingPointer"], text)
}

// Info ...
func (b *Buntstift) Info(text string) {
	Color := b.colorize(color.FgWhite)
	b.printf(Color, "  %v\n", text)
}

// List ...
func (b *Buntstift) List(text string, optionalOptions ...ListOptions) {
	var options = ListOptions{}
	if len(optionalOptions) > 0 {
		options = optionalOptions[0]
	}
	Color := b.colorize(color.FgWhite)
	b.printf(Color, "%v %v %v\n", strings.Repeat(" ", options.Indent*2), b.icons["multiplicationDot"], text)
}

// NewLine ...
func (b *Buntstift) NewLine() {
	fmt.Fprintf(Output, "\r \n")
}

// Line ...
func (b *Buntstift) Line() {
	width, _ := b.getTerminalSize()
	Color := b.colorize(color.FgWhite)
	b.printf(Color, "%v", strings.Repeat("-", width))
}

// WaitFor ...
func (b *Buntstift) WaitFor(worker func(stop chan bool)) {
	stop := make(chan bool)
	// done := make(chan bool)

	go b.spin(stop)
	worker(stop)

	// stop <- true
	// <-done
}

func (b *Buntstift) spin(stop chan bool) {
	spinner := spin.New()

loop:
	for {
		select {
		case <-stop:
			fmt.Fprintf(Output, "\r")
			break loop

		default:
			fmt.Printf("\r%s", spinner.Next())
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
