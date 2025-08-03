package main

import (
	"fmt"
	"os"
	"unicode"

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
	// m := make(map[string]int)
	// result := traverseDoc(doc, m)
	// for k, v := range result {
	// 	fmt.Printf("Element: %s\tCount: %d\n", k, v)
	// }
	traverser(doc)
}

func traverseDoc(n *html.Node, carte map[string]int) map[string]int {
	// fmt.Printf("Nattr: %s\n", n.Attr)
	if n.Type == html.ElementNode {
		carte[n.Data]++
		// fmt.Printf("Data: %v\tElementNode: %v\n", n.Data, html.ElementNode)
	}

	if n.FirstChild != nil {
		traverseDoc(n.FirstChild, carte)
	}
	if n.NextSibling != nil {
		traverseDoc(n.NextSibling, carte)
	}
	return carte
}

// Exercise 5.3: Write a function to print the contents of all text nodes in an HTML document
// tree. Do not descend into <script> or <style> elements, since their contents are not visible
// in a web browser.
func traverser(n *html.Node) {
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}
	if n.Type == html.TextNode {
		var filteredText []rune
		var containsText bool = false
		i := 0
		for _, v := range n.Data {
			if !unicode.IsSpace(v) {
				containsText = true
				break
			}
			i++
		}
		if containsText {
			for _, v := range n.Data[i:] {
				filteredText = append(filteredText, v)
			}
		}
		if len(filteredText) > 0 {
			fmt.Printf("Text: %v\n", string(filteredText))
		}
	}

	if n.FirstChild != nil {
		traverser(n.FirstChild)
	}
	if n.NextSibling != nil {
		traverser(n.NextSibling)
	}
}