package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

func (c circle) display() {
	fmt.Println("A Circle!")
}

func area(c circle, myChannel2 chan float64) {
	result := math.Pi * c.r * c.r
	myChannel2 <- result
}

func merhaba1(isim string) string {
	return "Merhaba " + isim
}

func merhaba2(chan1 chan string, isim string) {
	chan1 <- "Merhaba " + isim
}

func main() {
	isim := "Sade" // merhaba fonk. isim ile maindeki isim ayni degil.

	fmt.Println(merhaba1(isim)) // return olan fonk. goroutine uygulanmaz.
	fmt.Println("---------------------")
	myChannel := make(chan string)
	go merhaba2(myChannel, isim)
	fmt.Println("myChannel adresi:", myChannel)
	fmt.Println(<-myChannel)
	fmt.Println("---------------------")
	c1 := circle{5}
	chan1 := make(chan float64)
	go area(c1, chan1)
	fmt.Printf("%.2f\n", <-chan1)
	c1.display()

}
