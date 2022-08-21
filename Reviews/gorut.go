package main

import (
	"fmt"
	"sync"
)

func incrBy2(x *int, wg * sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("initial value of x is", *x)
	*x = *x + 2
	fmt.Println("value of x is", *x, "Increamented by 2")

}

func incrBy3(x *int, wg * sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("initial value of x is", *x)
	*x = *x + 3
	fmt.Println("value of x is", *x, "Increamented by 3")
}

func main() {
	var wg sync.WaitGroup
	x := 1
	wg.Add(1)
	go incrBy2(&x, &wg)
	wg.Add(1)
	go incrBy3(&x, &wg)
	wg.Wait()
}
/*
Above two go routines incrBy2 and incrBy3 will try to increment the sam evariable 'x'.
As the processor will run interleaved instructions,
it is never guaranteed which one function will execute first.
Based on the function which is executed first the value incremented/printed by next function to be executed will vary.
This is know as race condition. when two threads/process try to access and modify a single shared resource.

Output:
C:\Users\ag\go\src\awesomeProject>go run gorut.go
initial value of x is 1
value of x is 4 Increamented by 3
initial value of x is 1
value of x is 6 Increamented by 2


Note: Here when printing initial value both read the same initial value.
and done the increment on the value incremented by the first function in execution.
 */
