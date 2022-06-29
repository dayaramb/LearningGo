package main

import (
	"fmt"
	"sort"
)

func checkAnagram(s1 string, s2 string) bool {
	sort.Strings(s1)
	sort.Strings(s2)
	if s1 == s2 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("the anagram")
	fmt.Println(checkAnagram("daya", "aya"))
}
