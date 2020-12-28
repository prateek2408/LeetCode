package main

import "fmt"

func main() {
	fmt.Println(carvedStone(3))
}

func carvedStone(n int) int {
	if n == 0 {
		return 0
	}
	return n + carvedStone(n-1)
}
