package main

import (
	"fmt"
)

//Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice.
func main() {
	// test := "543252435    5 5 33  "
	// testSlice := []string{"80", "a", " ", "opp", " 6 ", "6", "aaa", "d", "aa", "aa", "zoo", "asheron"}
	t2 := []string{"aa", "bb", "bb", "cc", "dd", "dd"}
	elimDupInPlace(t2)
	fmt.Printf("%s\n", t2)
}

func elimDupInPlace(str []string) {
	//["aa", "bb", "bb", "cc", "dd", "dd"]
	prev := str[0]
	for i := range str {
		fmt.Printf("cVal: %s, preVal: %s\n", str[i], prev)
		fmt.Printf("String rn : %s\n", str)
		if i == 0 {
			continue
		}
		if i == len(str) - 1 {
			if (str[i] == prev) {
				str = str[:i]
				str = str[:len(str) - 1]
			}
			break
		}
		if str[i] == prev && i < len(str) - 1 {
			copy(str[:i], str[i+1:])
			str = str[:len(str) - 1]
			//fmt.Printf("String rn : %s\n", str)
		}
		prev = str[i]
	}
}