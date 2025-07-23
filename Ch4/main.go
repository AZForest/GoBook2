package main

import "fmt"

func main() {
	// s := []int{8, 5, 6, 90, 75}
	// fmt.Printf("%v\n", s)
	// reverse(s)
	// fmt.Printf("Reversing...\n")
	// fmt.Printf("%v\n", s)

	//4.3
	// x := [8]byte{2,5,4,3,90,88,77,6}
	// fmt.Printf("%v\n", x)
	// fmt.Printf("addr: %p\n", &x)
	// reverseP(&x)
	// fmt.Printf("Reversing...\n")
	// fmt.Printf("%v\n", x)
	// fmt.Printf("addr: %p\n", &x)

	//4.4
	direction := [2]string{"right", "left"}
	slice := []byte{1, 2, 3, 4, 5, 77, 46, 35}
	slice = rotate(slice, direction[1]) 
	fmt.Printf("Slice: %v\n", slice)
}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//Exercise 4.3: Rewrite reverse to use an array pointer instead of a slice.
func reverseP(p *[8]byte) {
	for i, j := 0, len(p) -1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}
}

//Exercise 4.4: Write a version of rotate that operates in a single pass.
func rotate(b []byte, direction string) []byte {
	if (direction == "left") {
		//Rotate Left by 1
		//[1, 2, 3, 4, 5] -> [2, 3, 4, 5, 1]
		//Rotate Left by 3
		//[1, 2, 3, 4, 5, 6] -> [4, 5, 6, 1, 2, 3]
		//]
		pH := b[0]
		for i := range b {
			if i != len(b) - 1 {
				b[i] = b[i+1]
			} else {
				b[i] = pH 
			}
		}
		return b
	} else {
		//Rotate Right by 1
		//[1, 2, 3, 4, 5] -> [5, 1, 2, 3, 4]
		nVal := b[len(b) - 1]
		for i := range b {
			if i != len(b) - 1 {
				pH := b[i]
				b[i] = nVal
				nVal = pH
			} else {
				b[i] = nVal
			}
		}
		return b
	}
}