package generalpurposeconverter

import (
	"fmt"
	"os"
	"strconv"
	// "/GoBook/Ch2/Ch2TempConv/tempconv.go"
)

func main() {

	for _, arg := range os.Args[1:] {
		entry, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error reading CMDLineArg: %v", err)
			os.Exit(1)
		}
		//Perform all conversions on Arg
		fmt.Print(entry) //just to avoid compiler error

	}
}