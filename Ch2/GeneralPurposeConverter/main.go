package main

import (
	gpcconv "GeneralPurposeConverter/GPCConv"
	//"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	testIfArgs := os.Args[1:]
	if (len(testIfArgs) > 0) {
		for _, arg := range os.Args[1:] {
			entry, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stdout, "Error reading CMDLineArg: %v", err)
				os.Exit(1)
			}
			processNum(entry)	
		}
	} else {
		i := 0
		for i < 3 {
			fmt.Printf("Please enter number to be converted: \n")
			var input string
			fmt.Scanln(&input)

			num, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Printf("Error when trying to convert %s to a float64\n", input)
				break
			}
			processNum(num)
			i++
		}
	}
}

func processNum(n float64) {
	f := gpcconv.Feet(n)
	m := gpcconv.Meters(n)
	fmt.Print("If a length value: \n")
	fmt.Printf("%s = %s, %s = %s\n", f, gpcconv.FtoM(f), m, gpcconv.MtoF(m))
	p := gpcconv.Pounds(n)
	kg := gpcconv.Kilograms(n)
	fmt.Print("If a weight value: \n")
	fmt.Printf("%s = %s, %s = %s\n", p, gpcconv.PtoK(p), kg, gpcconv.KtoP(kg))
	fmt.Print("\n")
}