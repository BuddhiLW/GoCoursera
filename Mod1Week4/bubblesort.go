package main

import (
	"flag"
	"fmt"
	"strconv"
)

func Swap(sli []int, i int) bool {
	if sli[i] > sli[i+1] {
		sli[i], sli[i+1] = sli[i+1], sli[i]
		return true
	} else {
		return false
	}
}

func BubbleSort(sli []int) []int {
	count_sort := 1
	for count_sort > 0 {
		count_sort = 0
		for i := 0; i < len(sli)-1; i++ { // sort two-by-two, all 2-tuples
			if Swap(sli, i) {
				count_sort++
			}
		}
		if count_sort != 0 {
			BubbleSort(sli)
		}
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
