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

const (
	version = "0.1.0"
)

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

// New Buntstift
func New(params ...interface{}) *Buntstift {
	var b *Buntstift

	if len(params) > 0 {
		param, ok := params[0].(Options)
		if ok {
			b = new(Buntstift)
			b.options = param
		} else {
			panic("Wrong parameter type, must be Buntstift.Options")
		}
	} else {
		b = &Buntstift{options: Options{}}
	}

	if b.options.NoUtf8 {
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

// Success ...
func (b *Buntstift) Success(text string) {
	output := b.colorize(color.FgGreen, color.Bold)
	output.Printf(b.icons["checkMark"]+" %v\n", text)
}

// Error
func (b *Buntstift) Error(text string) {
	output := b.colorize(color.FgRed, color.Bold)
	output.Printf("✗ %v\n", text)
}

// Warn ...
func (b *Buntstift) Warn(text string) {
	output := b.colorize(color.FgYellow, color.Bold)
	output.Printf(b.icons["rightPointingPointer"]+" %v\n", text)
}

// Info ...
func (b *Buntstift) Info(text string) {
	output := b.colorize(color.FgWhite)
	output.Printf("  %v\n", text)
}

// List ...
func (b *Buntstift) List(text string) {
	b.ListIndent(0, text)
}

// ListIndent ...
func (b *Buntstift) ListIndent(level int, text string) {
	output := b.colorize(color.FgWhite)
	output.Printf("%v"+b.icons["multiplicationDot"]+" %v\n", strings.Repeat(" ", level*2), text)
}

// Line ...
func (b *Buntstift) Line() {
	w, _ := b.getTerminalSize()
	output := b.colorize(color.FgWhite)
	output.Println(strings.Repeat("-", w))
}

// WaitFor ...
func (b *Buntstift) WaitFor(worker func()) {
	stop := make(chan bool, 1)
	done := make(chan bool, 1)

	go b.spin(stop, done)
	worker()

	stop <- true
	<-done
}

func (b *Buntstift) spin(stop, done chan bool) {
	s := spin.New()

loop:
	for {
		select {
		case <-stop:
			fmt.Printf("\r")
			done <- true
			break loop

		default:
			fmt.Printf("\r%s", s.Next())
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func (b *Buntstift) getTerminalSize() (int, int) {
	var w, h int
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	d, _ := cmd.Output()
	fmt.Sscan(string(d), &h, &w)
	return w, h
}
