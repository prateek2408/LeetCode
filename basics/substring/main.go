package main

import "fmt"

func main() {
	fmt.Println(countSubstring("dragonhorsedragons", "dragon"))
}

func countSubstring(s string, ss string) int {
	if len(s) == 0 || len(s) < len(ss) {
		return 0
	}
	if s[0:len(ss)] == ss {
		return 1 + countSubstring(s[len(ss):], ss)
	}
	return countSubstring(s[1:], ss)
}
