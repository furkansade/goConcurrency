package main

import (
	"fmt"
	"time"
)

var p = fmt.Println

func AfterDuration(d time.Duration) <-chan struct{} {
	c := make(chan struct{}, 1)
	go func() {
		time.Sleep(d)
		c <- struct{}{}
	}()
	return c

}

func main() {
	p("Hi!")
	<-AfterDuration(time.Second)
	p("Hello!")
	<-AfterDuration(time.Second)
	p("Bye!")

	<-time.After(time.Second) //AfterDuration'i kendimiz yazdik ama time.After fonks. kullanabiliriz.
	p("Last!")
}
