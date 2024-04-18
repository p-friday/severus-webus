package main

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buffer[:n-1]))
	conn.Write(buffer[:n-1])
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:42069")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Listening on port 42069")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}
