package main


import (
	"encoding/json"
	"fmt"
)

type Person struct {
	name  string
	addr  string
	phone string
}

func main() {
	p1 := Person{name: "joe", addr: "a st.", phone: "123"}

	barr, err := json.Marshal(p1)
	fmt.Println(err)
	fmt.Println("byte array (barr):", barr, "\t person 1 (p1) in Go:", p1)

	var p2 Person
	err2 := json.Unmarshal(barr, &p2)
	fmt.Println("person 2 (p2) in Go:", p2, "error:", err2)
}
