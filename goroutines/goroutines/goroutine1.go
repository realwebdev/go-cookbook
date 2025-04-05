package main

import (
	"fmt"
	"time"
)

// 1 unbuffered
// buffered channel
func main() {
	// go fetchResource()
	// fmt.Println("some result")
	resultch := make(chan string)
	go func() {
		result := <-resultch

		fmt.Println(result)
	}() // if not unblocked it will be full. this is the way of unblocking

	resultch <- "foo"

	// a channel is always block if its full even buffered or unbuffered
}

// go routine is basically a function that is getting scheduled by go scheduller

/* go lang has a scheduler in run time. each time you build a binary each time you build your own program
its also going to compiled in golang runtime. in golang runtime you have scheduler. it take all your
go routine and it will make sure that if go routine has no task to perform. it going to schedule another
go routine. go routine do not run in parallel. there is no such thing as parallel.
Its imp to know.
*/

func fetchResource() string {
	time.Sleep(time.Second * 2)
	return "some result"
}
