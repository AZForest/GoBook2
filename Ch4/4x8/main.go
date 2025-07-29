package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	// fourEight()
	fourNine()
}

// Exercise 4.8: Modify charcount to count letters, digits, and so on in their Unicode categories,
// using functions like unicode.IsLetter.
func preFourEight() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

// Exercise 4.8: Modify charcount to count letters, digits, and so on in their Unicode categories,
// using functions like unicode.IsLetter.
func fourEight() {
	counts := make(map[string]int)
	in := bufio.NewReader(os.Stdin)
	for {
		v, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if unicode.IsLetter(v){
			counts["letter"]++
			continue
		}
		if unicode.IsDigit(v) {
			counts["digit"]++
			continue
		}
		if unicode.IsSpace(v) {
			counts["space"]++
			continue
		}
		if unicode.IsSymbol(v) {
			counts["symbol"]++
			continue
		}
		if (unicode.IsPunct(v)) {
			counts["punct"]++
			continue
		}
	}
	fmt.Printf("Final Counts: \n")
	for k, v := range counts {
		fmt.Printf("%s\t%d\n", k, v)
	}
}

// Exercise 4.9: Write a program wordfreq to report the frequency of each word in an input text
// file. Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into
// words instead of lines.
func fourNine() {
	wordCounts := make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	
	for in.Scan() {
		word := in.Text()
		motPropre := cleanWord(word)
		if motPropre == "" {
			continue
		}
		wordCounts[motPropre]++
	}

	if err := in.Err(); err != nil {
		fmt.Printf("Error %s\n", err)
	}

	for k, v := range wordCounts {
		fmt.Printf("Word: %s   Count: %d\n", k, v)
	}
}

func cleanWord(str string) string {
	return strings.Map(func (r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return r
		}
		return -1
	}, str)
}
