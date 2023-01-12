package main

import (
	"fmt"
)

func main() {

	pingChannel := make(chan string)
	pongChannel := make(chan string)

	go func() {
		pingChannel <- "ping"
	}()
	go func() {
		pongChannel <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case ping := <-pingChannel:
			fmt.Println(ping)

		case pong := <-pongChannel:
			fmt.Println(pong)
		}
	}
}
