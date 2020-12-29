package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(power_optimized(12, 7))
	fmt.Println(power(12, 7))
}

func power(x int64, n int64) int64 {
	if n == 0 {
		return 1
	}
	return x * power(x, n-1)
}

func power_optimized(x int64, n int64) int64 {
	if n == 0 {
		return 1
	}
	if n%2 != 0 {
		cv := power_optimized(x, int64(math.Floor(float64(n/2))))
		return x * cv * cv
	}
	cv := power_optimized(x, int64(math.Floor(float64(n/2))))
	return cv * cv
}
