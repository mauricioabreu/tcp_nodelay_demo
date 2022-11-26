package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	conn.SetNoDelay(false) // true is the default value


	for x := 0; x < 10; x++ {
		b, err := conn.Write([]byte("test\n"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d bytes were written\n", b)
	}
}
