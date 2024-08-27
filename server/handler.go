package main

import (
	"fmt"
	"sync"
)

var SETs = map[string]string{}
var SETsMu = sync.RWMutex{}

var HSETs = map[string]map[string]string{}
var HSETsMu = sync.RWMutex{}

var handler = map[string]func(args []Value) Value{
	"PING":    ping,
	"SET":     set,
	"GET":     get,
	"HGET":    hget,
	"HSET":    hset,
	"HGETALL": hgetall,
}

func ping(args []Value) Value {
	if len(args) == 0 {
		return Value{typ: "string", str: "PONG"}
	}
	return Value{typ: "string", str: args[1].bulk}
}

func set(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'set' command"}
	}

	SETsMu.Lock()
	SETs[args[0].bulk] = args[1].bulk
	SETsMu.Unlock()
	return Value{typ: "string", str: args[1].bulk}
}

func get(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'get' command"}
	}

	SETsMu.Lock()
	value, ok := SETs[args[0].bulk]
	SETsMu.Unlock()
	if !ok {
		return Value{typ: "null"}
	}
	fmt.Println("get vaalue")

	return Value{typ: "bulk", bulk: value}
}

func hset(args []Value) Value {
	if len(args) != 3 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'hset' command"}
	}

	hash := args[0].bulk
	key := args[1].bulk
	value := args[2].bulk

	HSETsMu.Lock()
	_, ok := HSETs[hash]
	if !ok {
		HSETs[hash] = map[string]string{}
	}
	HSETs[hash][key] = value
	HSETsMu.Unlock()

	return Value{typ: "string", str: args[1].bulk}
}

func hget(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'hget' command"}
	}

	hash := args[0].bulk
	key := args[1].bulk

	HSETsMu.Lock()
	value, ok := HSETs[hash][key]
	HSETsMu.Unlock()
	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "bulk", bulk: value}
}

func hgetall(args []Value) Value {
	fmt.Println("args is:", args)
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'hgetall' command"}
	}

	hash := args[0].bulk

	HSETsMu.Lock()
	value, ok := HSETs[hash]
	HSETsMu.Unlock()
	if !ok {
		return Value{typ: "null"}
	}

	result := Value{typ: "array"}
	for k, v := range value {
		result.array = append(result.array, Value{typ: "bulk", bulk: k})
		result.array = append(result.array, Value{typ: "bulk", bulk: v})
	}

	return result
}
