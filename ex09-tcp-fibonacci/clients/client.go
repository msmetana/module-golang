package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"os"
	"time"
)

type message struct {
	Fib  *big.Int
	Time time.Duration
}

func main() {
	go Input()
	ListeningFromServer()
}

func Input() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		go SendToServer("127.0.0.1:1337", scanner.Text())
		if scanner.Text() == "exit" {
			break
		}
	}
}

func SendToServer(ip string, str string) {
	conn, err := net.Dial("tcp", ip)
	defer conn.Close()

	if err != nil {
		fmt.Println(err)
	}
	encoder := json.NewEncoder(conn)
	encoder.Encode(str)
}

func ListeningFromServer() {
	ln, err := net.Listen("tcp", ":1338")

	if err != nil {
		fmt.Println(err)
	}

	for {
		var num message
		conn, _ := ln.Accept()
		dec := json.NewDecoder(conn)
		err := dec.Decode(&num)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(num.Time, num.Fib)
	}
}
