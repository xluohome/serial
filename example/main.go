package main

import (
	"log"

	"github.com/xluohome/serial"
)

func main() {
	c := &serial.Config{Name: "COM9", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	txbuf := []byte{0xAA, 0x01, 0x0f, 0x00, 0x00, 0xBA}

	n, err := s.Write(txbuf)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%X\n", buf[:n])
}
