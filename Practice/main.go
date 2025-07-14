package main

import (
	"fmt"
)

func main() {
	eq1 := 10 % 3
	fmt.Printf("10 mod 3 == %d\n", eq1)
	eq2 := 7 % 100
	fmt.Printf("7 mod 100 == %d\n", eq2)
	eq3 := 0 % 100000
	fmt.Printf("0 mod 10000 == %d\n", eq3)
}