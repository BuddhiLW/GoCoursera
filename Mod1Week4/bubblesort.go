package main

import (
	"flag"
	"fmt"
	"strconv"
)

func Swap(first, next, count int) (int, int, int) {
	if first > next {
		first, next = next, first
		count = count + 1
	}
	return first, next, count
}

func BubbleSort(sli []int) []int {
	count_sort := 0
	for i := 0; i < len(sli)-1; i++ { // sort two-by-two, all 2-tuples
		sli[i], sli[i+1], count_sort = Swap(sli[i], sli[i+1], count_sort)
	}
	if count_sort != 0 {
		BubbleSort(sli)
	}
	return sli
}

func useFlags() []int {
	flag.Parse()
	args := flag.Args()
	fmt.Println("args:", args)

	nums := make([]int, len(args))
	// for each argument, insert in position nums[i] the string-converted value num64
	for i, arg := range args {
		num64, err := strconv.ParseInt(arg, 0, 0)
		if err != nil {
			fmt.Println("You probably didn't enter only integers.")
		}
		nums[i] = int(num64)
	}
	return nums
}

func main() {
	unsorted := useFlags()
	sorted := BubbleSort(unsorted)
	fmt.Println(sorted)
}
