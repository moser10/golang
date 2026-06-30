# golang — serial port reader

Small cross-platform CLI to read bytes from a UART/serial device and print them to stdout.

Originally a single-file demo; refactored into a standard Go layout with tests, flags, and graceful shutdown.

## Features

- Configurable port, baud rate, and output format (`hex` / `raw` / `dec`)
- Graceful exit on `Ctrl+C` (closes the port via `defer`)
- Cross-platform serial I/O via [`go.bug.st/serial`](https://github.com/bugst/go-serial)
- No root required for `go build` (static binary; no `sudo` for compilation)

## Requirements

- Go 1.22+
- Read/write permission on the serial device (OS-dependent)

## Quick start

```bash
git clone https://github.com/moser10/golang.git
cd golang
go run ./cmd/uartread -port /dev/tty.usbserial-1420
```

Or build a binary:

```bash
make build
./uartread -port /dev/tty.usbserial-1420 -baud 460800 -format hex
```

### Windows

```bash
make build-windows   # produces uartread.exe on macOS/Linux
uartread.exe -port COM3 -baud 460800
```

On Windows you can also build natively with `go build -o uartread.exe ./cmd/uartread`.

## CLI

```
uartread -port <device> [-baud 460800] [-format hex|raw|dec]
uartread <device>                 # positional port (backward compatible)
```

| Flag | Default | Description |
|------|---------|-------------|
| `-port` | — | Device path (`/dev/ttyUSB0`, `/dev/tty.usbserial-*`, `COM3`, …) |
| `-baud` | `460800` | Baud rate |
| `-format` | `hex` | `hex` = uppercase hex dump; `raw` = raw bytes; `dec` = space-separated decimals |

## Output example

```
2026/06/30 12:00:00 listening on /dev/tty.usbserial-1420 @ 460800 baud (format=hex); press Ctrl+C to stop
hex[4] A1B2C3D4
hex[8] 0102030405060708
```

## Project layout

```
cmd/uartread/          # program entrypoint (thin main)
internal/config/       # flag parsing + validation
internal/sniffer/      # read loop + formatting
```

`internal/` keeps library code private to this module — idiomatic Go for small tools that may grow later.

## Development

```bash
make test    # unit tests
make lint    # go vet
```

## License

MIT — see [LICENSE](LICENSE).
