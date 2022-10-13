package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)

	go PrintX()
	wg.Wait()
	fmt.Println()
	printY()
	// time.Sleep(time.Second)

}

func PrintX() { // buyuk yazmamin sebebi dısaridan bu fonk. erisilebilmesi ama su an icin gerek yok.
	for i := 0; i < 1000000; i++ {
		fmt.Print("X")
	}
	wg.Done()
}

func printY() {
	for i := 0; i < 10; i++ {
		fmt.Print("Y")
	}

}
