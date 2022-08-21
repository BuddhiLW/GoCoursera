package main

import (
	"fmt"
)

func main() {
	go fmt.Printf("New routine")
	fmt.Printf("Main routine")
}
