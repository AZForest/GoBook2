package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "This 🔥 Hello 🌍"
	b := []byte(s)
	// for _, v := range b {
	// 	fmt.Printf("%d ", v)
	// }
	// fmt.Printf("Before: %s\n", string(b))
	// res := reverseBytes(b)
	// fmt.Printf("After: %s\n", string(res))
	// fmt.Println()
	// r := []rune(s)
	// for _, v := range r {
	// 	fmt.Printf("%d ", v)
	// }
	// fmt.Println()
	fmt.Printf("Before: %s\n", string(b))
	res2 := reverseBytesR(b)
	fmt.Printf("After: %s\n", string(res2))
}

// Exercise 4.7: Modify reverse to reverse the characters of a []byte slice that represents a
// UTF-8-encoded string, in place. Can you do it without allocating new memory?
func reverseBytes(bytes []byte) []byte {
	for i, j := 0, len(bytes) - 1; i< j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return bytes
}


//Allocates new memory but works, only solution on page
func reverseBytesR(bytes []byte) []byte {
	r := []rune(string(bytes))
	for i, j := 0, len(r) - 1; i< j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return []byte(string(r))
}

func reverseBytesRII(bytes []byte) []byte {
	i := iterator(bytes)
	for i < len(bytes) {
		size := iterator(bytes[i:])
		i += size
	}
	return bytes
}

func iterator(bytes []byte) int {
	r, size := utf8.DecodeRune(bytes)
	fmt.Printf("R: %v, S: %d\n", r, size)
	return size
}

func wrapper(s []byte) []byte {
	i := 0
	j := len(s) - 1
	for i < j {
		fSize, sSize := tourner(i, j, s)
		s[i], s[j] = s[j], s[i]
		i += fSize
		j -= sSize
	}
	return s
}

func tourner(i, j int, s []byte) (int, int) {
	runeFirst, sizeFirst := utf8.DecodeRune(s[i:])
	runeLast, sizeLast := utf8.DecodeLastRune(s[:j+1])
	fmt.Printf("Fi: %d, s1: %d\nSi: %d, s2: %d\n", runeFirst, sizeFirst, runeLast, sizeLast)
	return sizeFirst, sizeLast;
}


// reverse reverses a slice of ints in place.
// func reverse(s []int) {
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		s[i], s[j] = s[j], s[i]
// 		}
// }