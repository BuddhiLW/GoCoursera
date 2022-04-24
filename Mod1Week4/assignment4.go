package main


import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Address  string `json:"address"`
}

func main() {
	var x string
	fmt.Println("Enter 'Name':")
	_, _ = fmt.Scan(&x)

	var y string
	fmt.Println("Enter 'Address':")
	_, _ = fmt.Scan(&y)

	p1 := Person{Name: x, Address: y}

	barr, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(barr))
}
