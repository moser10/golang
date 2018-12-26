package main

import (
	// "encoding/hex"
	// "flag"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"log"
	"os"
)

func main() {
	pname := "/dev/tty.usbserial-1420"
	if len(os.Args) > 1 {
		pname = os.Args[1]
	}

	//set up options
	options := serial.OpenOptions{
		PortName:        pname,
		BaudRate:        460800,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	//open port

	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	buf := make([]byte, 512)

	for true {
		n, err := port.Read(buf)
		if n > 0 {
			log.Printf("十六进制%X", buf[:n], n)
			// log.Printf("十进制%d", buf[:n])
			// log.Printf("字节%q", buf[:n])
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	//close later
	// defer port.Close()
	//write 4 Bytes
	// b := []byte{0x00, 0x01, 0x02, 0x03}
	// n, err := port.Write(b)
	// if err != nil {
	// 	log.Fatalf("port.Write:%v", err)
	// }
	fmt.Println("-----------------")
}
