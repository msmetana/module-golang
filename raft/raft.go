package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	//"math/rand"
	"net"
	"strconv"
	"time"
)

var (
	//strIps = [...]string{"10.30.8.105:8888", "10.30.8.131:8888", "10.30.8.47:8888", "10.30.8.66:8888"}
	strIp = [...]string{"127.0.0.1:8006"}
	f     follower
	c     candidate
	l     leader
)

func main() {
	done := make(chan string)

	f.work()
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
			f.get = message
			f.data = append(f.data, f.get)
			f.send()
		case <-time.After(220 * time.Millisecond):
			fmt.Println("\nTimeout. The message didn`t receive")
			fmt.Println("I want to be a leader and now I send message")
			c.voice = 0
			go c.work()
			return
		}
	}
}

func Send(ip string) {
	for x := 0; x < 500; x++ {
		time.Sleep(200 * time.Millisecond)
		if x > 5 {
			time.Sleep(20 * time.Millisecond)
		}
		conn, err := net.Dial("tcp", ip)

		if err != nil {
			fmt.Println(err)
			continue
		}

		tosend, _ := json.Marshal("x : " + strconv.Itoa(x) + " " + conn.LocalAddr().String())
		conn.Write(tosend)
		conn.Close()
	}
}
