package internal

import (
	"fmt"
	"io"

	"go.uber.org/atomic"
)

var (
	termSizeWidth  = atomic.Int32{}
	termSizeHeight = atomic.Int32{}
)

func updateTerminalSize() error {
	termSizeWidth.Store(800)
	termSizeHeight.Store(800)
	return nil
}

func moveCursorUp(w io.Writer, n int) {
	_, _ = fmt.Fprintf(w, "\033[%dA", n)
}

func moveCursorDown(w io.Writer, n int) {
	_, _ = fmt.Fprintf(w, "\033[%dB", n)
}

func moveCursorToLineStart(w io.Writer) {
	_, _ = fmt.Fprintf(w, "\r")
}

func clearLine(w io.Writer) {
	_, _ = fmt.Fprintf(w, "\033[2K")
}

func init() {
	_ = updateTerminalSize()
}
