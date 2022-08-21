package main

// The goal of this activity is to explore the use of threads by creating
// a program for sorting integers that uses four goroutines to create
// four sub-arrays and then merge the arrays into a single array.

// Review criteria
// Students will receive five points if the program sorts the
// integers and five additional points if there are four goroutines
// that each prints out a set of array elements that it is storing.
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Inputs:
// https://stackoverflow.com/questions/46143974/convert-slice-of-string-input-from-console-to-slice-of-numbers

func input() []float64 {
	scanner := bufio.NewScanner(os.Stdin)
	// var result [][]float64
	var numbers []float64
	var txt string
	for scanner.Scan() {
		txt = scanner.Text()
		if len(txt) > 0 {
			values := strings.Split(txt, ",")
			// var row []float64
			for _, v_text := range values {
				// Convert string-value representing a float to float-value.
				v_float, err := strconv.ParseFloat(strings.Trim(v_text, " "), 64)
				if err != nil {
					panic(fmt.Sprintf("Incorrect value for float64 '%v'", v_float))
				}
				numbers = append(numbers, v_float)
			}
			// result = append(result, row)
		}
	}
	fmt.Printf("Result: %v\n", numbers)
	return numbers
}

func sortVec(vec []float64, v chan []float64) {
	fmt.Println("Vec for sorting: ", vec)
	sort.Float64s(vec)
	fmt.Println("After sort:", vec)
	v <- vec
}

func main() {
	numbers := input()
	quotient, remainder := len(numbers)/4, len(numbers)%4

	v := make(chan []float64)

	vec1 := numbers[0:(quotient + remainder)]
	go sortVec(vec1, v)
	a := <-v
	// fmt.Println("a:", a, a[0])

	vec2 := numbers[(1*quotient)+remainder : (2*quotient)+remainder]
	go sortVec(vec2, v)
	b := <-v
	// fmt.Println("b:", b, b[0])

	vec3 := numbers[(2*quotient)+remainder : (3*quotient)+remainder]
	go sortVec(vec3, v)
	c := <-v
	// fmt.Println("c:", c, c[0])

	vec4 := numbers[(3*quotient)+remainder : (4*quotient)+remainder]
	go sortVec(vec4, v)
	d := <-v
	// fmt.Println("d:", d, d[0], reflect.TypeOf(d[0]))

	map_vecs := map[float64][]float64{
		a[0]: a,
		b[0]: b,
		c[0]: c,
		d[0]: d,
	}
	// fmt.Println("map_vecs:", map_vecs)

	firsts := []float64{float64(a[0]), float64(b[0]), float64(c[0]), float64(d[0])}
	// fmt.Println("first values:", firsts)
	sort.Float64s(firsts)
	// fmt.Println("first values sorted:", firsts)

	var sorted []float64
	for i := 0; i < len(firsts); i++ {
		sorted = append(sorted, map_vecs[firsts[i]]...)
	}
	sort.Float64s(sorted)
	fmt.Println("sorted:", sorted)
}
