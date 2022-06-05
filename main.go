package main

import (
	"fmt"
	"io"
	"net"
)

func HandleConnection(connection net.Conn) {
	receivedBuffer := make([]byte, 4096)

	for {
		size, err := connection.Read(receivedBuffer)

		if err != nil && err != io.EOF {
			fmt.Println(err)
		}

		if size > 0 {
			data := receivedBuffer[:size]

			fmt.Println(data)
		}

	}
}

func main() {
	server, err := net.Listen("tcp", ":25565")

	if err != nil {
		fmt.Println(err)
	}

	defer server.Close()

	for {
		connection, err := server.Accept()

		if err != nil {
			fmt.Println(err)
		}

		go HandleConnection(connection)
	}
}
