package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, _ := net.Listen("tcp", ":1721")
	fmt.Println("Hello, World!", listener)
	for {
		conn, err := listener.Accept()
		fmt.Println("Connected")
		if err != nil {
			log.Fatal(err)
		}
		go do(conn)
	}

}

func do(conn net.Conn) {
	var byteBuf []byte
	_, err := conn.Read(byteBuf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Processing the connection")
	time.Sleep(8 * time.Second)
	_, _ = conn.Write([]byte("Http/1.1 200 Ok\r\n\r\nHello, World!\r\n"))

	conn.Close()
}
