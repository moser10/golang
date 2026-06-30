package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go.bug.st/serial"

	"github.com/moser10/golang/internal/config"
	"github.com/moser10/golang/internal/sniffer"
)

func main() {
	log.SetFlags(log.LstdFlags)
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	cfg, err := config.Parse(args)
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			printUsage(os.Stderr)
			return 0
		}
		fmt.Fprintf(os.Stderr, "uartread: %v\n\n", err)
		printUsage(os.Stderr)
		return 2
	}

	mode := &serial.Mode{
		BaudRate: cfg.Baud,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open(cfg.Port, mode)
	if err != nil {
		fmt.Fprintf(os.Stderr, "uartread: open %s: %v\n", cfg.Port, err)
		return 1
	}
	defer port.Close()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log.Printf("listening on %s @ %d baud (format=%s); press Ctrl+C to stop", cfg.Port, cfg.Baud, cfg.Format)

	if err := sniffer.Run(ctx, port, cfg.Format, os.Stdout); err != nil && !errors.Is(err, context.Canceled) {
		fmt.Fprintf(os.Stderr, "uartread: %v\n", err)
		return 1
	}

	log.Println("stopped")
	return 0
}

func printUsage(w *os.File) {
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w, "  uartread -port <device> [-baud 460800] [-format hex|raw|dec]")
	fmt.Fprintln(w, "  uartread <device>                    # positional port, same defaults")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  uartread -port /dev/tty.usbserial-1420")
	fmt.Fprintln(w, "  uartread -port COM3 -baud 115200 -format dec")
}
