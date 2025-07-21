package main

import (
	"bytes"
	"fmt"
	"strings"
)
func main() {
	i := 0
	for i < 3 {
		fmt.Printf("Enter any integer or decimal without commas: \n")
		var input string
		fmt.Scanln(&input)
		result := commaWrapper(input)
		fmt.Printf("End result: %s\n", result)
		fmt.Printf("\n")
		i++
	}	
}

//Wrapper Function
func commaWrapper(str string) string {
	s, i := findDecimal(str)
	if len(s) == 0 {
		return commaBuffer(str)
	}
	return commaBuffer(str[:i]) + s
}

//Returns a string with commas placed in correct order using a stack implementation
func commaBuffer(str string) string {
	var buf bytes.Buffer
	count := 0

	var stack []byte
	for i := len(str) - 1; i >= 0; i-- {
		if count % 3 == 0 && i != len(str) - 1 {
			stack = append(stack, ',')
		}
		stack = append(stack, str[i])
		count++
	}
	for i := len(stack) - 1; i >= 0; i-- {
		buf.WriteByte(stack[i])
	}
	return buf.String()
}

//Returns string after the decimal point, and the index of decimal
func findDecimal(str string) (string, int) {
	containsDecimal := strings.LastIndex(str, ".")
	if containsDecimal == -1 {
		return "", 0
	}
	return str[containsDecimal:], containsDecimal
}

