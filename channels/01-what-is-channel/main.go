package main

import (
	"fmt"
)

func main() {

	messages := make(chan string, 2)
	fmt.Printf("%d - %T\n", messages, messages)

	go func() { messages <- "ping" }() //channellar her zaman goroutine ile calisir.

	msg := <-messages
	fmt.Println(msg)
	fmt.Printf("%v - %T", msg, msg)

}
