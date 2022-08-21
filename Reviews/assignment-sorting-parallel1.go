package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

//Sorts a slice using bubble sort algorithm
func bubbleSort(id int, slice []int, wg *sync.WaitGroup) {
	fmt.Printf("Go routine %d sorting: %v\n", id, slice)
	defer wg.Done()
	n := len(slice)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

//Splits a slice into n approximately equal size slices
func splitSlice(slice []int, numberOfChunks int) [][]int {
	if len(slice) == 0 {
		return nil
	}
	if numberOfChunks <= 0 {
		return nil
	}

	if numberOfChunks == 1 {
		return [][]int{slice}
	}

	result := make([][]int, numberOfChunks)

	if numberOfChunks > len(slice) {
		for i := 0; i < len(slice); i++ {
			result[i] = []int{slice[i]}
		}
		return result
	}

	for i := 0; i < numberOfChunks; i++ {
		min := (i * len(slice) / numberOfChunks)
		max := ((i + 1) * len(slice)) / numberOfChunks

		result[i] = slice[min:max]
	}

	return result
}

func main() {
	fmt.Println("Enter a series of integers in comma separated format")
	fmt.Print(">")

	//Read list of numbers from standard input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	arr := strings.Split(input, ",")

	numbers := make([]int, len(arr))
	for i, val := range arr {
		num, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			log.Fatal("invalid input, please enter integers separated by comma")
		}
		numbers[i] = num
	}

	//Split the array into 4
	partitions := 4
	chunks := splitSlice(numbers, partitions)

	//Sort each partition in a different goroutine
	var wg sync.WaitGroup
	wg.Add(partitions)

	for i := 0; i < partitions; i++ {
		go bubbleSort(i, chunks[i], &wg)
	}

	wg.Wait()

	//Merge the sorted slices into one
	result := []int{}
	for i := 0; i < len(chunks); i++ {
		result = merge(result, chunks[i])
	}
	fmt.Printf("Sorted list: %v\n", result)
}

//Merges two sorted arrays into a sorted array
func merge(arr1, arr2 []int) []int {
	result := make([]int, len(arr1)+len(arr2))
	var i, j, k int

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result[k] = arr1[i]
			i++
		} else {
			result[k] = arr2[j]
			j++
		}
		k++
	}

	for i < len(arr1) {
		result[k] = arr1[i]
		i++
		k++
	}

	for j < len(arr2) {
		result[k] = arr2[j]
		j++
		k++
	}

	return result
}
