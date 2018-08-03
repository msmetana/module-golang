package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"strconv"
	"time"
)

var m = make(map[int64]*big.Int)

type message struct {
	Fib  *big.Int
	Time time.Duration
}

func main() {
	StartListening()
}

func StartListening() {
	ln, err := net.Listen("tcp", ":1337")

	if err != nil {
		fmt.Println(err)
	}

	for {
		var s string
		conn, _ := ln.Accept()
		d := json.NewDecoder(conn)
		err := d.Decode(&s)

		if err != nil {
			fmt.Println(err)
		}

		number, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			fmt.Println(err)
		}

		go SendNumber(number)
	}
}

func SendNumber(str int64) {
	conn, err := net.Dial("tcp", "127.0.0.1:1338")
	ret := message{}

	if err != nil {
		fmt.Println(err)
	}

	t0 := time.Now()
	ret.Fib = fibonacci(str)
	t1 := time.Now()
	ret.Time = t1.Sub(t0)

	encoder := json.NewEncoder(conn)
	encoder.Encode(ret)
	conn.Close()
}

func fibonacci(n int64) *big.Int {
	val, err := m[n]
	if err {
		return val
	}

	fst, _ := fib(n)
	m[n] = fst
	return fst
}

func fib(n int64) (*big.Int, *big.Int) {
	if n == 0 {
		return big.NewInt(0), big.NewInt(1)
	}
	a, b := fib(n / 2)
	c := Mul(a, Sub(Mul(b, big.NewInt(2)), a))
	d := Add(Mul(a, a), Mul(b, b))
	if n%2 == 0 {
		return c, d
	} else {
		return d, Add(c, d)
	}
}

func Mul(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mul(x, y)
}
func Sub(x, y *big.Int) *big.Int {
	return big.NewInt(0).Sub(x, y)
}
func Add(x, y *big.Int) *big.Int {
	return big.NewInt(0).Add(x, y)
}
