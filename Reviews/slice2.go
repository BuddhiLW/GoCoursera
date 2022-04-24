package main

import (
	"fmt"
	"sort"
	"strconv"
)

func printSlice(sli []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(sli), cap(sli), sli)
}

func main() {
	sli := make([]int, 0, 3)

	for {
		fmt.Printf("Please enter integers or quit(X):")

		var x string
		_, _ = fmt.Scan(&x)

		if x == "X" {
			break
		} else {
			if s, err := strconv.Atoi(x); err == nil {
				sli = append(sli, s)
				sort.Ints(sli)
				printSlice(sli)
			} else {
				fmt.Printf("\nUnrecognized! Please enter again!\n")
			}
		}
	}
}
