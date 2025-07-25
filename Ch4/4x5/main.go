package main

import (
	"fmt"
)

//Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice.
func main() {
	// test := "543252435    5 5 33  "
	// testSlice := []string{"80", "a", " ", "opp", " 6 ", "6", "aaa", "d", "aa", "aa", "zoo", "asheron"}
	//t2 := []string{"aa", "bb", "bb", "cc", "dd", "dd"}
	t3 := []string{"dd", "dd", "dd", "aa"}
	elim3(&t3)
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

	func elim3(s *[]string) {
		str := *s
		//["dd", "dd", "dd", "aa"]
		write := 1
		for read := 1; read < len(str); read++ {
			if (str[read] != str[read - 1]) {
				str[write] = str[read]
				write++
			}
		}
		// str = str[:write]
	}
