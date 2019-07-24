package main

import (
	"fmt"
	"net"
	"os"
)

const (
	_type = "tcp"
	_host = "localhost"
	_port = "5000"
)

func handeRequest(conn net.Conn) {
	buffer := make([]byte, 1024)

	reqLen, err := conn.Read(buffer)

	s := string(buffer[:reqLen])

	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}

	fmt.Println(s)
	conn.Write([]byte("Message received."))

	conn.Close()
}

func main() {
	l, err := net.Listen(_type, _host+":"+_port)

	if err != nil {
		fmt.Println("Error Listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()

	fmt.Println("Listening to" + _host + ":" + _port)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go handeRequest(conn)
	}

}
