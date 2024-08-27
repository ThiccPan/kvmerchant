package main

import (
	"fmt"
	"log"
	"net"
	"strings"
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
	w := NewWriter(conn)

	for {
		fmt.Println("start")
		resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			log.Println("error:", err)
			conn.Write([]byte("err"))
			break
		}
		log.Println("value:", value)

		if value.typ != "array" {
			fmt.Println("req is not an array")
			w.Write(Value{typ: "string", str: ""})
			continue
		}

		if len(value.array) == 0 {
			fmt.Println("empty array")
			w.Write(Value{typ: "string", str: ""})
			continue
		}

		command := strings.ToUpper(value.array[0].bulk)
		handlerFunc, ok := handler[command]
		args := value.array[1:]
		
		if !ok {
			fmt.Println("invalid command:", command)
			w.Write(Value{typ: "string", str: ""})
			continue
		}

		handlerVal := handlerFunc(args)
		fmt.Println("handlerval:", handlerVal)

		// response to client
		w.Write(handlerVal)
	}
}
