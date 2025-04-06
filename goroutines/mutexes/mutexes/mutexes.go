package main

import (
	"sync/atomic"
)

// state => set a state OR update state OR Delete state

type State struct {
	// mu    sync.Mutex
	count int32
}

func (s *State) setState(i int) {
	// s.mu.Lock()
	// defer s.mu.Unlock()

	// s.count = i

	atomic.AddInt32(&s.count, int32(i))
}

func main() {
}
