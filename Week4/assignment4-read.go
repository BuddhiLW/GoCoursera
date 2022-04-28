package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type T struct {
	fname string
	lname string
}

func main() {
	// Ask for file name (fn)
	fmt.Println("Enter file name, e.g.: 'names.txt'")
	var fn string
	_, _ = fmt.Scan(&fn)

	// Open file
	file, err := os.Open(fn)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Create a slice, in which the T constructs will be populated.
	var slice []T

	// Counter for each line
	i := 0

	// Read each line and iterate on it
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		strline := scanner.Text()
		// words := strings.Split(strline, " ")
		// var words [20]string
		words := strings.Fields(strline)
		// fmt.Println(words, len(words))
		fmt.Println("First name:", words[0], "Last name:", words[1])
		slice = append(slice, T{fname: words[0], lname: words[1]})
		fmt.Println("Slice so far:", slice)
		i = i + 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// return slice with T-type elements
	fmt.Println(slice)

	// Iterate and return first and last names
	fmt.Println("List of first and last names in the file:")

	for _, value := range slice {
		fmt.Println("First name:", value.fname, "Last name:", value.lname)
	}
