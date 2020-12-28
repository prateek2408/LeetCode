package main

import "fmt"

func main() {
	m := []int{30, 20, 150, 100, 40}
	fmt.Println(numPairsDivisibleBy60(m))
	fmt.Println(30 % 60)
}

func numPairsDivisibleBy60(time []int) int {
	m := make(map[int]int)
	np := 0
	for _, v := range time {
		fmt.Println(v)
		v %= 60
		fmt.Println(v)
		if n, ok := m[(60-v)%60]; ok {
			m[v]++
			np += n
		} else {
			m[v]++
		}
	}
	fmt.Println(m)
	return np
}
