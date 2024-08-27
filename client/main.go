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

	msg := "$5\r\nPesan\r\n"
	msg2 := "*3\r\n$3\r\nSET\r\n$5\r\nhello\r\n$5\r\nworld\r\n"
	msg3 := "*2\r\n$3\r\nGET\r\n$5\r\nhello\r\n"
	msg4 := fmt.Sprint(
		"*4\r\n",
		"$4\r\nHSET\r\n",
		"$5\r\nusers\r\n",
		"$2\r\nu1\r\n",
		"$5\r\nuser1\r\n",
	)
	fmt.Println(msg4)
	msg5 := fmt.Sprint(
		"*4\r\n",
		"$4\r\nHSET\r\n",
		"$5\r\nusers\r\n",
		"$2\r\nu2\r\n",
		"$5\r\nuser2\r\n",
	)
	msg6 := fmt.Sprint(
		"*2\r\n",
		"$7\r\nHGETALL\r\n",
		"$5\r\nusers\r\n",
	)

	sendMsg(conn, msg)
	sendMsg(conn, msg2)
	sendMsg(conn, msg3)
	sendMsg(conn, msg4)
	sendMsg(conn, msg5)
	sendMsg(conn, msg6)
}

func sendMsg(conn net.Conn, msg string) {
	_, err := conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("error writing to conn:", err)
		return
	}

	respMsg := make([]byte, 1024)
	_, err = conn.Read(respMsg)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("resp msg:", string(respMsg))
}
