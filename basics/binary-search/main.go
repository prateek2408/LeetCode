package main

import (
	"fmt"
	"math"
)

func main() {
	li := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}
	fmt.Println(binarySearchRecursive(li, 19, 0, len(li)-1))
}

func binarySearchRecursive(pl []int, p, min, max int) int {
	if min > max {
		return -1
	}
	mid := int(math.Floor(float64(max-min)/2)) + min
	if pl[mid] == p {
		return mid
	} else if pl[mid] < p {
		min = mid + 1
	} else {
		max = mid - 1
	}
	return binarySearchRecursive(pl, p, min, max)
}
