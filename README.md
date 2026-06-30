# Serial Port Reader — Go

跨平台串口读取工具，将 UART 收到的字节输出到终端（hex / raw / dec）。

## 本地运行

需要 Go 1.22+。

```bash
git clone https://github.com/moser10/serial-port-reader-golang.git
cd serial-port-reader-golang
go run ./cmd/uartread -port /dev/tty.usbserial-1420
```

或编译：

```bash
make build
./uartread -port COM3 -baud 460800 -format hex
```

Windows 交叉编译：`make build-windows` → `uartread.exe`

## 参数

```
uartread -port <设备> [-baud 460800] [-format hex|raw|dec]
uartread <设备>    # 位置参数，兼容旧用法
```

## License

MIT
