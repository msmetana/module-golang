package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func Worker(index int, jobs <-chan float64, wg *sync.WaitGroup) {
	flag := false
	for t := range jobs {
		if !flag {
			fmt.Printf("worker:%d spawning\n", index)
			flag = true
		}
		fmt.Printf("worker:%d sleep:%.1f\n", index, t)
		s := int(1000 * t)
		time.Sleep(time.Duration(s) * time.Millisecond)
	}
	if flag {
		fmt.Printf("worker:%d stopping\n", index)
	}
	wg.Done()
}

func Run(poolSize int) {
	var wg sync.WaitGroup
	max_workers := 100
	ch := make(chan float64, max_workers)

	read := bufio.NewScanner(os.Stdin)

	for read.Scan() {
		input := string(read.Bytes())
		s, _ := strconv.ParseFloat(input, 64)
		ch <- s
	}
	close(ch)

	for w := 1; w <= poolSize; w++ {
		wg.Add(1)
		go Worker(w, ch, &wg)
		time.Sleep(1 * time.Millisecond)
	}
	wg.Wait()
}
