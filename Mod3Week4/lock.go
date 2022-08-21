package main

import (
	"fmt"
	"sync"
)

var i int = 0
var mut sync.Mutex

func inc(i int, t string, p chan int) {
	p <- i + 1
	fmt.Println(t)
}

func main() {
	var sum int = 0
	a := 0
	b := 0
	c := 0
	// b := 0
	p1 := make(chan int)
	p2 := make(chan int)
	p3 := make(chan int)
	// p2 := make(chan int)
	for i := 0; i < 5; i++ {
		// fmt.Println("sum, bef:", sum)
		mut.Lock()
		go inc(sum, "inc, process: 1 --", p1)
		mut.Unlock()
		mut.Lock()
		go inc(sum, "inc, process: 2 --", p2)
		mut.Unlock()
		mut.Lock()
		go inc(sum, "inc, process: 3 --", p3)
		mut.Unlock()
		a = <-p1
		b = <-p2
		c = <-p3
		sum = a + b + c - 2*sum
		fmt.Println("sum:", sum)
		// fmt.Println("a", a)
		// mut.Unlock()
		// mut.Lock()
		// fmt.Println("b", b)
		// mut.Unlock()
		// sum = a + b - sum
		// fmt.Println("sum, aft:", sum)
	}
	fmt.Println(sum)
}
