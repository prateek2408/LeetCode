package main

import "fmt"

func main() {
	fmt.Println(singleDigitFrequency(10000101, 0))
}

func singleDigitFrequency(n int, t int) int {
	if n == 0 {
		return 0
	}
	if n%10 == t {
		return 1 + singleDigitFrequency(int(n/10), t)
	}
	return singleDigitFrequency(int(n/10), t)
}
