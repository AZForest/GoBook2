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
	// counter := 0
	// for _, link := range visit(nil, doc) {
	// 	counter++
	// 	fmt.Println(link)
	// }
	// fmt.Printf("Total Links: %d\n", counter)

	//5.4 main 
	newMap := map[string][]string{
		"links": {""},
		"script": {""},
		"image": {""},
		"style": {""},
	}
	m := visiter(newMap, doc)
	for k, v := range m {
		fmt.Printf("Key: %s\tVal: %v\n", k, v)
	}
}

// 5.1
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

// 5.4
// Extend the visit function s o that it extracts other kinds of links from the doc-
// ument, such as images, scripts, and style sheets.
// visit appends to links each link found in n and returns the result.
func visiter(tracker map[string][]string, n *html.Node) map[string][]string {
	// fmt.Printf("%v ", n.Data)
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			for _, a := range n.Attr {
				if a.Key == "href" {
					tracker["links"] = append(tracker["links"], a.Val)
				}
			}
		case "script":
			for _, s := range n.Attr {
				tracker["script"] = append(tracker["script"], s.Val)
			}
		case "img":
			for _, i := range n.Attr {
				if i.Key == "src" {
					tracker["image"] = append(tracker["image"], i.Val)
				}
			}
		case "link":
			for _, i := range n.Attr {
				if i.Key == "rel" && i.Val == "stylesheet" {
					for _, a := range n.Attr {
						if a.Key == "href" {
							tracker["style"] = append(tracker["style"], a.Val)
						}
					}	
				}
			}
		default:
			//fmt.Printf("No matches.\n")
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