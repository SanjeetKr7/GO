/*
Print A B C in order 10 times
ex-
A
B
C
A
B
C
A
B
C
A
B
C
...
*/
package main

import (
	"fmt"
	"time"
)

func printA(A chan int, B chan int) {
	for i := 1; i <= 10; i++ {
		<-A
		fmt.Println("A")
		B <- 1
	}
}

func printB(B chan int, C chan int) {
	for i := 1; i <= 10; i++ {
		<-B
		fmt.Println("B")
		C <- 1
	}
}

func printC(C chan int, A chan int) {
	for i := 1; i <= 10; i++ {
		<-C
		fmt.Println("C")
		A <- 1
	}
}

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go printA(a, b)
	go printB(b, c)
	go printC(c, a)
	a <- 1
	time.Sleep(time.Duration(2) * time.Second)
}
