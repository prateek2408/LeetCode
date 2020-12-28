package main

import "fmt"

func main() {
	a := make(map[int]int)
	a[0] = 1
	a[1] = 9
	a[2] = 20
	a[3] = 55
	a[4] = 66
	a[5] = 90
	for _, v := range a {
		v %= 60
		fmt.Printf("Hello")
	}
}
