package generalpurposeconverter

import (
	"fmt"
	"os"
	"strconv"
	//"/Users/alexforest/Desktop/GitHubProjects/GoBook/Ch2/Ch2TempConv/tempconv.go"
	//"../Ch2TempConv"
)

func main() {

	for _, arg := range os.Args[1:] {
		entry, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error reading CMDLineArg: %v", err)
			os.Exit(1)
		}
		//Perform all conversions on Arg
		f := Feet(entry)
		m := Meters(entry)
		fmt.Print("If a length value: \n")
		fmt.Printf("%s = %s, %s = %s\n", f, FtoM(f), m, MtoF(m))
		p := Pounds(entry)
		kg := Kilograms(entry)
		fmt.Print("If a weight value: \n")
		fmt.Printf("%s = %s, %s = %s\n", p, PtoK(p), kg, KtoP(kg))

	}
}