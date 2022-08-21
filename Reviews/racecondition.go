package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := []int{}
	for i := 1; i <= 5; i++ {
		go addNegativeNumber(&numbers, i)
		go addPositiveNumber(&numbers, i)
		time.Sleep(100 * time.Millisecond)
	}
	// time.Sleep(time.Second * 1)
	fmt.Println(numbers)
}

func addNegativeNumber(numbers *[]int, number int) {
	*numbers = append(*numbers, number-1)
}

func addPositiveNumber(numbers *[]int, number int) {
	*numbers = append(*numbers, number)
}

/* Race Conditions
Race conditions can occour when multiple functions try to access the same memory concurrently.

When you run this programm multiple times the printed list is going to change.
Its not because we change the code, its te result of a race condition.
*/
