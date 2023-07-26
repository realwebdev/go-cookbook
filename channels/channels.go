package main

import "fmt"

func main() {
	// Create a channel.
	ch := make(chan int)

	// Start a goroutine to send data to the channel.
	go func() {
		i := 0
		for {
			ch <- i
			i++
		}
	}()

	// Start a goroutine to receive data from the channel.
	go func() {
		for {
			v := <-ch
			fmt.Println(v)
		}
	}()

	// Wait for the goroutines to finish.
	<-ch
	<-ch
}
