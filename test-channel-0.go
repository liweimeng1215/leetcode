// two go routine print 0~2N-1 one by one
package main

import (
	"fmt"
	"time"
)

var c1 chan int
var c2 chan int
var n int
var N int

func init() {
	c1 = make(chan int)
	c2 = make(chan int)
	n = 0
	N = 10
}

func t1() {
	for i := 0; i < N; i++ {
		<-c2
		fmt.Printf("%d\n", n)
		n++
		c1 <- 1
	}
}

func t2() {
	for i := 0; i < N; i++ {
		<-c1
		fmt.Printf("%d\n", n)
		n++
		c2 <- 1
	}
}

func main() {
	go t1()
	go t2()
	c2 <- 1
	time.Sleep(10 * time.Second)
}
