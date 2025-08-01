package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// Exercise 5.1: Change the findlinks program to traverse the n.FirstChild linked list using
// recursive calls to visit instead of a loop.

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	c := n.FirstChild
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	if c != nil {
		links = visit(links, c)
	}
	return links
}