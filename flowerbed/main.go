package main

import "fmt"

func main() {
	v := canPlaceFlowers([]int{1,0,0,0,0,1}, 2)
	fmt.Printf("%v\n", v)
}

func canPlaceFlowers(flowerbed []int, n int) bool {
    plt := 0
    fbl := len(flowerbed)
    if n >=0 && n <= fbl{
        for i := 0; i < fbl; i++{
            if flowerbed[i] == 0{
                if flowerbed[i+1] == 0 && i == 0{
                    plt++
                } else if flowerbed[i+1] == 0 && flowerbed[i-1] == 0{
                    plt++
                } else if i == fbl-1{
                    plt++
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
