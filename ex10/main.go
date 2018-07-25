package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	//"strconv"
	"time"
)

var Nodes []string

func main() {

	//strIp := [...]string{"10.30.8.105:1337", "10.30.8.131:1337", "10.30.8.47:1337", "10.30.8.47:1337"}
	strIp := [...]string{"127.0.0.1:8866"}
	done := make(chan string)
	go StartListening()
	for _, i := range strIp {
		go Send(i)
	}
	<-done
}

func StartListening() {
	ln, err := net.Listen("tcp", ":8866")

	if err != nil {
		fmt.Println(err)
	}
	for {
		//timer1 := time.NewTimer(2 * time.Second)

		conn, err := ln.Accept()

		if err != nil {
			fmt.Println(err)
		}

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("\nMessage Received  : ", string(message))

	}
}
func Send(str string) {
	for x := 0; x < 30; x++ {
		conn, err := net.Dial("tcp", str)
		time.Sleep(2 * time.Second)

		if err != nil {
			fmt.Println(err)
			continue
		}

		tosend, _ := json.Marshal("Max " + conn.LocalAddr().String())
		conn.Write(tosend)
		conn.Close()
		//time.Sleep(2 * time.Second)
	}
	//fmt.Printf("\n End!!")
}
