package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	msg := "ping"

	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Println(err)
		return
	}

	respMsg := make([]byte, 1024)
	_, err = conn.Read(respMsg)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(respMsg))	
}