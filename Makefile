.PHONY: build test lint clean build-windows

BINARY := uartread
MAIN   := ./cmd/uartread

build:
	go build -o $(BINARY) $(MAIN)

test:
	go test ./...

lint:
	go vet ./...

clean:
	rm -f $(BINARY) uartread.exe

# Cross-compile for Windows from macOS/Linux (no CGO required).
build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o uartread.exe $(MAIN)
