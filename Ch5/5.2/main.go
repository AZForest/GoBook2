package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// Exercise 5.2: Write a function to populate a mapping from element names—p, div, span, and
// so on—to the number of elements with that name in an HTML document tree.
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}
	// fmt.Printf("Result: %v\n", doc)
	m := make(map[string]int)
	result := traverseDoc(doc, m)
	fmt.Printf("Result: %v\n", result)

}

func traverseDoc(n *html.Node, carte map[string]int) map[string]int {
	// fmt.Printf("Nattr: %s\n", n.Attr)
	if (n.Type == html.ElementNode) {
		fmt.Printf("Type: %v\tElementNode: %v\n", n.Type, html.ElementNode)
	}

	if n.FirstChild != nil {
		carte = traverseDoc(n.FirstChild, carte)
	}
	if n.NextSibling != nil {
		carte = traverseDoc(n.NextSibling, carte)
	}
	return carte
}