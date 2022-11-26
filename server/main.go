package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go handle(c)
	}
}

func handle(c net.Conn) {
	log.Println("Accepting connections on:", c.RemoteAddr().String())

	for {
		data, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		if strings.TrimSpace(data) == "quit" {
			break
		}
	}

	c.Close()
}
