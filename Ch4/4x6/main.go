package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "   He\nllo 🌍	"
	b := []byte(s)
	res := removeUniSpaces(b)
	fmt.Println(string(res))
}

func removeUniSpaces(str []byte) []byte {
	write := 0
	for _, v := range str {
		// Lines 20 - 26 Not needed, just more an expressive version, only the if statement at bottom needed
		// r := rune(v)
		// fmt.Printf("i: %d, v: %v\n", i, r)
		// if !unicode.IsSpace(r) {
		// 	fmt.Println("True!")
		// 	str[write] = v
		// 	write++
		// }
		if !unicode.IsSpace(rune(v)) {
			str[write] = v
			write++
		}
	}
	return str[:write]
}