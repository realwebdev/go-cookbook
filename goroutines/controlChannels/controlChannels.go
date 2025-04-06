package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{} // 1 byte
	msgch  chan string
} // true or false binary expression for empty channels

func newServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string, 128),
	}

}

func (s *Server) start() {
	fmt.Println("server starting")
	s.loop() // it will block
}

func (s *Server) quit() {
	// close(s.quitch)
	// or
	s.quitch <- struct{}{}
}
func (s *Server) sendMessage(msg string) {
	s.msgch <- msg

}

func (s *Server) loop() {
mainloop:
	for {
		select {
		case <-s.quitch: // continuously check for quit signal or do we have a m
			// message recieved
			fmt.Println("quitting server")
			break mainloop
		// do some stuff when we need to quit
		case msg := <-s.msgch:
			s.handleMessage(msg)
			// default:
			// do some stuff when we have a message
		}
	}
	fmt.Println("server is shutting down gracefully")
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("we recieved a message:", msg)
}
func main() {
	server := newServer()

	go func() {
		time.Sleep(time.Second * 5)
		server.quit()

	}()
	server.start()
	// go server.start()
	// for i := 0; i < 100; i++ {
	// 	server.sendMessage(fmt.Sprintf("handle this number,%d", i))

	// }

	// time.Sleep(time.Second * 5)

}
