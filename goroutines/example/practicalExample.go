package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type UserProfile struct {
	ID       int
	Comments []string
	Likes    int
	Friends  []int
}

type Response struct {
	data any
	err  error
}

func main() {
	start := time.Now()
	userProfile, err := handleGetUserProfile(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userProfile)
	fmt.Println("fetch the user profile took", time.Since(start))
}

func handleGetUserProfile(id int) (*UserProfile, error) {
	var (
		respch = make(chan Response, 3)
		wg     = &sync.WaitGroup{}
	)

	// Add to WaitGroup before starting goroutines
	wg.Add(3)
	go getComments(id, respch, wg)
	go getLikes(id, respch, wg)
	go getFriends(id, respch, wg)

	wg.Wait() // Block until all goroutines finish
	close(respch)

	userProfile := &UserProfile{}

	// Process responses from the channel
	for resp := range respch {
		if resp.err != nil {
			return nil, resp.err
		}
		switch msg := resp.data.(type) {
		case int:
			userProfile.Likes = msg
		case []int:
			userProfile.Friends = msg
		case []string:
			userProfile.Comments = msg
		}
	}

	return userProfile, nil // Return the populated user profile
}

func getComments(id int, respch chan Response, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Microsecond * 200)
	fmt.Println("getComments finished")
	respch <- Response{
		data: []string{"Hey, that was great", "Yeah Buddy", "Ow, I didn't know that"},
		err:  nil,
	}
}

func getLikes(id int, respch chan Response, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Microsecond * 200)
	fmt.Println("getLikes finished")
	respch <- Response{
		data: 200,
		err:  nil,
	}
}

func getFriends(id int, respch chan Response, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Microsecond * 100)
	fmt.Println("getFriends finished")
	respch <- Response{
		data: []int{11, 34, 454},
		err:  nil,
	}
}
