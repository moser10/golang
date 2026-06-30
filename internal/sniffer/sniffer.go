package sniffer

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/moser10/golang/internal/config"
)

const defaultReadTimeout = 200 * time.Millisecond

// Reader is the minimal interface needed from a serial port.
type Reader interface {
	Read([]byte) (int, error)
	SetReadTimeout(time.Duration) error
}

// Run reads from r until ctx is canceled or a non-recoverable error occurs.
func Run(ctx context.Context, r Reader, format config.OutputFormat, w io.Writer) error {
	if err := r.SetReadTimeout(defaultReadTimeout); err != nil {
		return err
	}

	buf := make([]byte, 512)

	for {
		if err := ctx.Err(); err != nil {
			return err
		}

		n, err := r.Read(buf)
		if n > 0 {
			FormatFrame(w, buf[:n], format)
		}
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
	}
}
