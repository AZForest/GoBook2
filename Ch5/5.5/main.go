package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	words, imgs, err := CountWordsAndImages("https://www.espn.com")
	if err != nil {
		fmt.Printf("Error thrown: %e", err)
	}
	fmt.Printf("Words: %d\tImages: %d\n", words, imgs)
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}
// Exercise 5.5: Implement countWordsAndImages. (See Exercise 4.9 for word-splitting.)
func countWordsAndImages(n *html.Node) (words, images int) {
	if (n.Type == html.ElementNode && n.Data == "img") {
		images++
	} else if (n.Type == html.TextNode) {
		// fmt.Println(n.Data)
		in := bufio.NewScanner(strings.NewReader(n.Data))
		in.Split(bufio.ScanWords)
		for in.Scan() {
			words++
		}
		// arr := strings.Split(n.Data, " ")
		// words += len(arr)
	}



	for c := n.FirstChild ; c != nil ; c = c.NextSibling {
		cw, ci := countWordsAndImages(c)
		words += cw
		images += ci
	}
	// if n.FirstChild != nil {
	// 	childWords, childImages := countWordsAndImages(n.FirstChild)
	// 	words += childWords
	// 	images += childImages
	// }
	// if n.NextSibling != nil {
	// 	sibWords, sibImages := countWordsAndImages(n.NextSibling)
	// 	words += sibWords
	// 	images += sibImages

	// }
	return words, images
 }
