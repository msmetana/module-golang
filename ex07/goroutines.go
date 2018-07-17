package goroutines

func Process(c chan string) chan string {
	output := make(chan string)
	done := make(chan bool)

	go func() {
		str := "(" + <-c + ")"
		output <- str
		done <- true
	}()

	go func() {
		<-done
		close(output)
	}()

	return output
}
