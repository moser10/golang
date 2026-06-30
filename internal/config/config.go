package config

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

// Config holds CLI options for the serial reader.
type Config struct {
	Port   string
	Baud   int
	Format OutputFormat
}

// OutputFormat selects how received bytes are printed.
type OutputFormat string

const (
	FormatHex OutputFormat = "hex"
	FormatRaw OutputFormat = "raw"
	FormatDec OutputFormat = "dec"
)

// Parse reads flags and the optional positional port argument.
func Parse(args []string) (Config, error) {
	fs := flag.NewFlagSet("uartread", flag.ContinueOnError)
	fs.SetOutput(flag.CommandLine.Output())

	var cfg Config
	var format string
	var showHelp bool
	fs.BoolVar(&showHelp, "h", false, "show help")
	fs.BoolVar(&showHelp, "help", false, "show help")
	fs.StringVar(&cfg.Port, "port", "", "serial device path (e.g. /dev/ttyUSB0, COM3)")
	fs.IntVar(&cfg.Baud, "baud", 460800, "baud rate")
	fs.StringVar(&format, "format", string(FormatHex), "output format: hex, raw, dec")

	if err := fs.Parse(args); err != nil {
		return Config{}, err
	}
	if showHelp {
		return Config{}, flag.ErrHelp
	}

	if cfg.Port == "" && fs.NArg() > 0 {
		cfg.Port = fs.Arg(0)
	}
	if cfg.Port == "" {
		return Config{}, errors.New("serial port is required: use -port or pass the device as the first argument")
	}

	cfg.Format = OutputFormat(strings.ToLower(format))
	switch cfg.Format {
	case FormatHex, FormatRaw, FormatDec:
	default:
		return Config{}, fmt.Errorf("unsupported format %q (use hex, raw, or dec)", format)
	}

	if cfg.Baud <= 0 {
		return Config{}, errors.New("baud rate must be positive")
	}

	return cfg, nil
}
