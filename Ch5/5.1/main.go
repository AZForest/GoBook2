package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// Also contains exercise 5.4

// Exercise 5.1: Change the findlinks program to traverse the n.FirstChild linked list using
// recursive calls to visit instead of a loop.

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink1: %v\n", err)
		os.Exit(1)
	}
	//5.1
	counter := 0
	for _, link := range visit(nil, doc) {
		counter++
		fmt.Println(link)
	}
	fmt.Printf("Total Links: %d\n", counter)

	//5.4 main 
	m := visiter(nil, doc)
	for k, v := range m {
		fmt.Printf("Key: %s\tVal: %v\n", k, v)
		// for _, val := range v {
		// 	fmt.Printf("%s\n", val)
		// }
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
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	// for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 	links = visit(links, c)
	// 	}
	return links
}

// Extend the visit function s o that it extracts other kinds of links from the doc-
// ument, such as images, scripts, and style sheets.
// visit appends to links each link found in n and returns the result.
func visiter(tracker map[string]string, n *html.Node) map[string]string {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			for _, a := range n.Attr {
				if a.Key == "href" {
					tracker["links"] += a.Val
				}
			}
		case "script":
			tracker["script"] += n.Data
		case "image":
			tracker["image"] += n.Data
		case "style":
			tracker["style"] += n.Data
		default:
			fmt.Printf("No matches.\n")
		}
		}
	 
	if n.FirstChild != nil {
		visiter(tracker, n.FirstChild)
	}
	if n.NextSibling != nil {
		visiter(tracker, n.NextSibling)
	}
	// for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 	links = visit(links, c)
	// 	}
	return tracker

}