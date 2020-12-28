package main

import "fmt"

func main() {
	li := []int{1, 2, 3, 4, 5, 6, 7, 8}
	li = append(li[:3], li[4:]...)
	fmt.Println(li)
}
