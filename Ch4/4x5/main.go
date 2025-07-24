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
	totalLen := len(str) //6
	i := 0
	for i <= totalLen {
		if i == 0 {
			i++
			continue
		}
		if str[i] == prev {
			if i < totalLen - 1 {
				//fmt.Printf("String rn : %s\n", str)
				s := str[:i]
				for _, v := range str[i+1:] {
					s = append(s, v)
				}
				str = s
			} else if i == totalLen - 1 {
				fmt.Printf("String rnnn : %s and i: %d\n", str, i)
				str = str[:len(str)-1]
				break
			}
			totalLen--
		}
		prev = str[i]
		i++
	}
}