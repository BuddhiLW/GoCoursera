package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

func useFlags() []float64 {
	flag.Parse()
	args := flag.Args()
	fmt.Println("args:", args)

	nums := make([]float64, len(args))
	// for each argument, insert in position nums[i] the string-converted value num64
	for i, arg := range args {
		num64, err := strconv.ParseFloat(arg, 0)
		if err != nil {
			fmt.Println("You probably didn't enter only integers.")
		}
		nums[i] = float64(num64)
	}
	return nums
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	space := func(t float64) float64 {
		// return (1/2)*a*math.Pow(t, 2) + v0*t + s0
		return ((a * math.Pow(t, 2) / 2) + v0*t + s0)
	}
	return space
}

func main() {
	// fmt.Println("Enter a acceleration (a), initial velocity (v0) and (s0):")
	var params []float64
	params = useFlags()
	fn := GenDisplaceFn(params[0], params[1], params[2])

	var t float64
	fmt.Println("Enter a time (t) to calculate the displacement, with the \n the acceleration (a), initial velocity (v0) and initial displacement (s0)\n you executed the program:")
	_, _ = fmt.Scan(&t)
	fmt.Println("The displacement for t =", t, "s(t) =", fn(t), "\n", "a:", params[0], "v0:", params[1], "s0:", params[2])
}
