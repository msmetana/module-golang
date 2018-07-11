package letter

import "sync"

var mutex = &sync.Mutex{}

func Frequency(s string) map[rune]int {
	m := map[rune]int{}

	for _, v := range s {
		m[v] += 1
	}

	return m
}

func Concurrent(s string, m *map[rune]int, c chan byte) {
	for _, v := range s {
		mutex.Lock()
		(*m)[v] += 1
		mutex.Unlock()
	}
	c <- 1
}

func ConcurrentFrequency(s []string) map[rune]int {
	m := map[rune]int{}
	done := make(chan byte)

	for _, v := range s {
		go Concurrent(v, &m, done)
	}

	for range s {
		<-done
	}

	return m
}
