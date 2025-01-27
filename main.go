package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	go func() {
		for val := range ch {
			fmt.Println(val)
		}
		done <- true
	}()
	<-done
}
