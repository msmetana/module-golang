package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"
)

//var Nodes []struct

func main() {
	//strIp := [...]string{"10.30.8.105:1337", "10.30.8.131:1337", "10.30.8.47:1337"}
	strIp := [...]string{"127.0.0.1:8006"}
	done := make(chan string)
	go StartListening()

	// imitation receive message by leader
	for _, ip := range strIp {
		go Send(ip)
	}

	<-done
}

func StartListening() {
	ln, err := net.Listen("tcp", ":8006")

	if err != nil {
		fmt.Println(err)
	}

	for {
		c1 := make(chan net.Conn, 1)

		go func() {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println(err)
			}
			c1 <- conn
		}()

		select {
		case conn := <-c1:
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("\nMessage Received: ", string(message))
		case <-time.After(220 * time.Millisecond):
			fmt.Println("\nThe message didn`t receive")
			fmt.Println("I want to be a leader and now I send message")
			return
		}
	}
}
func Send(ip string) {
	for x := 0; x < 5; x++ {
		time.Sleep(200 * time.Millisecond)
		conn, err := net.Dial("tcp", ip)

		if err != nil {
			fmt.Println(err)
			continue
		}

		tosend, _ := json.Marshal("x :" + strconv.Itoa(x) + " " + conn.LocalAddr().String())
		conn.Write(tosend)
		conn.Close()
	}

}
