package main

import "fmt"

func main() {
	str1 := "appleee"
	str2 := "leppaee"
	b := isAnagram(str1, str2)
	fmt.Printf("Strings: %s, %s. Anagram: %v\n", str1, str2, b)
}

func isAnagram(strOne string, strTwo string) bool {
	if len(strOne) != len(strTwo) {
		return false
	}
	mapOne := make(map[rune]int)
	for _, v := range strOne {
		mapOne[v]++
	}
	mapTwo := make(map[rune]int)
	for _, v := range strTwo {
		mapTwo[v]++
	}

	for k, v := range mapOne {
		existsVal := mapTwo[k]
		if existsVal != v {
			return false
		}
	}
	return true
}