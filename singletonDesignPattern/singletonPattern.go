/**
signleton pattern  ensures a class or struct in Go has only one instance and provides a global
point of access to it. This pattern is useful when you want to contril access to resources
or maintain a single instance of a resource througghout the application.

*/

package main

import (
	"fmt"
	"sync"
)

type Logger struct {
	LogFile string
}

var instance *Logger
var once sync.Once

// Get Logger returns the singleton instance of the Logger
func GetLoggerInstance() *Logger {
	once.Do(func() {
		instance = &Logger{LogFile: "app.log"}
		// Additional initialization if needed
	})

	return instance
}

func (l *Logger) Log(message string) {
	fmt.Printf("[%s] %s\n", l.LogFile, message)
}

func main() {
	log := GetLoggerInstance()

	// lgging messages

	log.Log("starting application...")
	log.Log("An error occured")
}
