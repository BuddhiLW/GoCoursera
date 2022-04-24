package main

import (
	"fmt"
	"math"
)

func main() {
	var acc, initVel, initDisp, time float64
	fmt.Println("Enter Acceleration, initial velocity and initial displacement: ")
	fmt.Scan(&acc, &initVel, &initDisp)
	fn := GenDispFn(acc, initVel, initDisp)
	fmt.Println("Enter the value of time: ")
	fmt.Scan(&time)
	fmt.Println(fn(time))
}

func GenDispFn(acc, initVel, initDisp float64) func(time float64) float64 {
	return func(time float64) float64 {
		return initDisp + initVel*time + (0.5)*(acc*math.Pow(time, 2))
	}
}
