package main

import "fmt"

func main() {
	v := canPlaceFlowers([]int{0}, 1)
	fmt.Printf("%v\n", v)
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	plt := 0
	fbl := len(flowerbed)
	fmt.Printf("Size of flowerbed=%v\n", fbl)
	if n >= 0 && n <= fbl {
		if flowerbed[0] == 0 && fbl == 1 {
			plt++
			flowerbed[0] = 1
		}
		for i := 0; i < fbl; i++ {
			if flowerbed[i] == 0 {
				if i == fbl-1 && flowerbed[i] == 0 {
					if flowerbed[i-1] == 0 {
						plt++
						flowerbed[i] = 1
					}
				} else if i == 0 && flowerbed[i+1] == 0 {
					plt++
					flowerbed[i] = 1
				} else if flowerbed[i+1] == 0 && flowerbed[i-1] == 0 {
					plt++
					flowerbed[i] = 1
				}
			}
		}
		if n <= plt {
			fmt.Printf("Can be planted\n")
			return true
		} else {
			fmt.Printf("Cannot be planted\n")
			return false
		}
	} else {
		fmt.Printf("Invalid value of n %v\n", n)
		return false
	}
}
