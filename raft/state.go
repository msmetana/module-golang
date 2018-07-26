package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"time"
)

type Block struct {
	PrevHash string
	CurHash  string
	data     []string
}

type Message struct {
	Len    int
	Sender string
	Mcode  string
	Blocks []Block
}

type state interface {
	send()
	work()
}

type follower struct {
	get  string
	data []string
}

type candidate struct {
	voice int
}

type leader struct {
}

func (f follower) send() {
	fmt.Printf("\nfollower append\n")
	fmt.Println(f.data)
}

func (f follower) work() {
	fmt.Println("follower start listening")
	go StartListening()
}

func (c candidate) send() {

	for _, ipad := range strIp {
		go func(ip string) {
			conn, err := net.Dial("tcp", ip)
			if err != nil {
				fmt.Println(err)
			}
			tosend, _ := json.Marshal("127.0.0.1:8006")
			conn.Write(tosend)
			conn.Close()
		}(ipad)
	}
	go StartListening()

}

func (c candidate) work() {
	fmt.Println("candidat start work")
	r := rand.Intn(200)
	time.Sleep(time.Duration(r) * time.Microsecond)

	//c.send()
}

func (l leader) send() {

}

func (l leader) work() {

}
