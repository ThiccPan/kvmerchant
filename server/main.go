package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	readMsg()
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Println(err)
		return
	}

	conn, err := l.Accept()
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		fmt.Println("now reading")
		_, err := conn.Read(buf)
		fmt.Println("done reading")
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("error reading:", err)
		} else {
			println(string(buf))
		}
		conn.Write([]byte("pong"))
	}
}

func readMsg() {
	input := "$5\r\nAhmed\r\n"
	reader := bufio.NewReader(strings.NewReader(input))
	b, _ := reader.ReadByte()
	if b != '$' {
		fmt.Println("invalid data type")
		return
	}
	lenMsg, _ := reader.ReadByte()
	size, _ := strconv.ParseInt(string(lenMsg), 10, 64)
	key := make([]byte, size)
	reader.ReadByte()
	reader.ReadByte()
	reader.Read(key)
	fmt.Println(string(key))
}