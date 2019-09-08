package main

import "fmt"

func channelDemo() {
	message := make(chan string)
	go func() {
		message <- "Hello"
	}()
	msg := <-message
	fmt.Println(msg)
}
