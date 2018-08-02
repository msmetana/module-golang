package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	maxWaiting = 3
	nCustomers = 10
)

var (
	lobby = make(chan chan int, maxWaiting)
	wg    = new(sync.WaitGroup)
	flag  = 1
	mutex = &sync.Mutex{}
)

func barber() {
	for {
		select {
		case ch := <-lobby:
			flag = 0
			fmt.Println("Barber cuts the hair of Customer", <-ch)
			time.Sleep(1000)
			fmt.Println("Barber finishes cutting hair.")
			ch <- 0
			flag = 1
		}
	}
}

func customer(id int) {
	defer wg.Done()
	ch := make(chan int)

	mutex.Lock()
	if len(lobby) == 0 && flag == 1 {
		fmt.Println("Barber sleep")
		fmt.Println("Customer", id, "enters the barbershop and woke up Barber.")
	} else {
		fmt.Println("Customer", id, "enters the barbershop")
	}
	mutex.Unlock()

	select {
	case lobby <- ch:
		ch <- id
		<-ch
	default:
		fmt.Println("Customer", id, "go away.")
		time.Sleep(1 * time.Second)
		wg.Add(1)
		customer(id)
	}
}

func main() {
	wg.Add(nCustomers)
	go barber()

	for i := 0; i < nCustomers; i++ {
		go customer(i)
	}

	wg.Wait()
}
