package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("server running")
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Println(err)
		return
	}

	// listen for conn
	conn, err := l.Accept()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("new conn:", conn)
	defer conn.Close()

	for {
		resp := NewResp(conn)
		val, err := resp.Read()
		if err != nil {
			log.Println("error:", err)
			conn.Write([]byte("err"))
			break
		}
		conn.Write([]byte("pong"))
		log.Println("value:", val)
	}
}
