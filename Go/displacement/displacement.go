package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v_o, s_o float64) func(t float64) float64 {
	return func(t float64) float64 {
		return .5*a*math.Pow(t, 2) + v_o*t + s_o
	}
}

func main() {
	var (
		a   float64 = 10.0
		v_o float64 = 2.0
		s_o float64 = 1.0
		t   float64 = 3
	)

	// Asking the inital values
	fmt.Printf(":::::::::: Displacement Calculator ::::::::::\n")
	fmt.Printf("Write acceleration: ")
	fmt.Scanf("%f", &a)
	fmt.Printf("Write inital velocity: ")
	fmt.Scanf("%f", &v_o)
	fmt.Printf("Write initial displacement: ")
	fmt.Scanf("%f", &s_o)

	// Generation of displacement function
	fn := GenDisplaceFn(a, v_o, s_o)

	// Asking for time
	fmt.Printf("Write time: ")
	fmt.Scanf("%f", &t)

	fmt.Printf("===== Displacement =====\n")
	fmt.Printf("Displacement %v\n", fn(t))
}
