package sniffer

import (
	"bytes"
	"strings"
	"testing"

	"github.com/moser10/serial-port-reader-golang/internal/config"
)

func TestFormatFrameHex(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer
	FormatFrame(&buf, []byte{0xA1, 0x02, 0xFF}, config.FormatHex)
	if buf.String() != "hex[3] A102FF\n" {
		t.Fatalf("got %q", buf.String())
	}
}

func TestFormatFrameDec(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer
	FormatFrame(&buf, []byte{1, 2}, config.FormatDec)
	want := "dec[2] 1 2"
	if strings.TrimSpace(buf.String()) != want {
		t.Fatalf("got %q want %q", buf.String(), want)
	}
}
