package main

import (
	"fmt"
)

//Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice.
func main() {
	// test := "543252435    5 5 33  "
	// testSlice := []string{"80", "a", " ", "opp", " 6 ", "6", "aaa", "d", "aa", "aa", "zoo", "asheron"}
	// t2 := []string{"aa", "bb", "bb", "cc", "dd", "dd", "dd", "dd"}
	t3 := []string{"dd", "dd", "dd", "aa"}
	// r := elim7(&t2)
	elim7(&t3)
	//fmt.Printf("V: %s\n", r)
	fmt.Printf("%s\n", t3)
}

// func elimDupInPlace(str []string) {
// 	//["aa", "bb", "bb", "cc", "dd", "dd"]
// 	//"dd", "dd", "dd", "aa"]
// 	prev := str[0]
// 	totalLen := len(str) //6
// 	i := 0
// 	for i <= totalLen {
// 		if i == 0 {
// 			i++
// 			continue
// 		}
// 		if str[i] == prev {
// 			if i < totalLen - 1 {
// 				//fmt.Printf("String rn : %s\n", str)
// 				// s := str[:i]
// 				// for _, v := range str[i+1:] {
// 				// 	s = append(s, v)
// 				// }
// 				// str = s
// 				s := str[:i + 1]
// 				fmt.Printf("S val: %s\n", s)
// 				// copy(str[:i], str[i+1:])
// 				// str = str[:len(str) - 1]
// 			} else if i == totalLen - 1 {
// 				fmt.Printf("String rnnn : %s and i: %d\n", str, i)
// 				str = str[:len(str)-1]
// 				fmt.Printf("String rn : %s\n", str)
// 				break
// 			}
// 			totalLen--
// 		} else {
// 			i++
// 		}
// 		prev = str[i]
// 		//i++
// 	}


	func elim(str []string) {
		//["dd", "dd", "dd", "aa"]
		prev := str[0]
		for i := range str {
			// fmt.Printf("i: %d, str: %s, strLen: %d\n", i, str[i], len(str))
			if i == 0 {
				continue
			}
			if str[i] == prev && i < len(str) {
				copy(str[:i],str[i+1:])
				str = str[:len(str) -1]
			}
			// if str[i] == prev && i == len(str) {
			// 	str = str[:len(str) -1]
			// 	break
			// }
			// prev = str[i]
		}
	}

	//modifies in place
	func elim4(str *[]string) {
		s := *str
		for i := range s {
			s[i] += "apple"
		}
	}

	//Not in place But Correct 
	func elim5(str []string) []string {
		var res []string
		for i, v := range str {
			fmt.Printf("res: %s, i: %d\n", res, i)
			if i == 0 || str[i] != str[i - 1] {
				res = append(res, v)
			}
		}
		return res
	}

	//Works in place?
	func elim6(str []string) []string {
		write := 0
		for i, v := range str {
			if i == 0 || str[i] != str[i - 1] {
				str[write] = v
				write++
			}
		}
		return str[:write]
	}

	//In place with a pointer
	func elim7(s *[]string) {
		str := *s
		write := 0
		for i, v := range str {
			if i == 0 || str[i] != str[i - 1] {
				str[write] = v
				write++
			}
		}
		*s = str[:write]
	}
