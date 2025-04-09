package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type result struct {
	userId string
	err    error
}

func fetchUserId() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*100)
	defer cancel()

	resultch := make(chan result, 1)

	go func() {
		res, err := thirdPartyHTTPCall()
		resultch <- result{
			userId: res,
			err:    err,
		}
	}()

	select {
	// Done()
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-resultch:
		return res.userId, res.err
	}
}
func thirdPartyHTTPCall() (string, error) {
	time.Sleep(time.Millisecond * 500)
	return "2", nil
}

func main() {
	startTime := time.Now()
	userId, err := fetchUserId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the response took %v -> %+v\n", time.Since(startTime), userId)

}
