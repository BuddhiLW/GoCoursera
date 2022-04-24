package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
)

func main() {

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

	// sort the integers
	sort.Ints(nums)
	// print the sorted list
	fmt.Println("sorted slice:", nums)
}
