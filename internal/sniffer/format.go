package sniffer

import (
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"github.com/moser10/golang/internal/config"
)

// FormatFrame writes one received chunk using the selected output style.
func FormatFrame(w io.Writer, data []byte, format config.OutputFormat) {
	switch format {
	case config.FormatRaw:
		_, _ = w.Write(data)
		if len(data) == 0 || data[len(data)-1] != '\n' {
			_, _ = io.WriteString(w, "\n")
		}
	case config.FormatDec:
		parts := make([]string, len(data))
		for i, b := range data {
			parts[i] = fmt.Sprintf("%d", b)
		}
		fmt.Fprintf(w, "dec[%d] %s\n", len(data), strings.Join(parts, " "))
	default:
		fmt.Fprintf(w, "hex[%d] %s\n", len(data), strings.ToUpper(hex.EncodeToString(data)))
	}
}
