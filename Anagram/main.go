package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// sort the string and check for equality.
func checkAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	st1 := strings.Split(s1, "")
	sort.Strings(st1)
	s1 = strings.Join(st1, "")
	st2 := strings.Split(s2, "")
	sort.Strings(st2)
	s2 = strings.Join(st2, "")
	if s1 == s2 {
		return true
	} else {
		return false

	}
}

// If two strings are anagram then they will have even number of characters in the map.
func checkAnagram2(s1 string, s2 string) bool {
	freq := make(map[string]int)
	l1 := strings.Split(s1, "")
	for _, c := range l1 {
		count := freq[c]
		if count == 0 {
			freq[c] = 1
		} else {
			freq[c] += 1
		}
	}
	l2 := strings.Split(s2, "")
	for _, c := range l2 {
		count := freq[c]
		if count == 0 {
			freq[c] = 1
		} else {
			freq[c] += 1
		}
	}
	var isAnagram bool = true
	for _, f := range freq {
		if f%2 != 0 {
			isAnagram = false
		}
	}
	return isAnagram
}

func getSize(name string) {
	file, err := os.Open(name)
	file.Seek()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 20)
	stat, err := os.Stat(name)
	if err != nil {
		fmt.Printf("stat failed %v", err)
	}
	start := stat.Size() - 20
	_, err = file.ReadAt(buf, start)
	fmt.Printf("Size is:%v", stat.Size())
	if err == nil {
		fmt.Printf("%s", buf)
	}

}
func main() {
	fmt.Println(checkAnagram2("daya", "ayad"))
	fmt.Println(checkAnagram("act", "cat"))
	getSize("myfile.txt")

}
