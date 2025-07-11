// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	
	//Read in Response
	sz, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error: %s", err)
		return
	}
	numBytes := len(sz)
	fmt.Printf("Num Bytes: %d\n", numBytes)

	//Write To File
	fileName := truncName(url) + ".txt"
	fmt.Printf("FileName: %s\n", fileName)
	erreur := os.WriteFile(fileName, sz, 0644)
	if erreur != nil {
		ch <- fmt.Sprintf("Error: %s", erreur)
		return
	}
	fmt.Print("Worked!\n")

	//Count bytes Only
	// nbytes, err := io.Copy(io.Discard, resp.Body)
	// resp.Body.Close() // don't leak resources
	// if err != nil {
	// 	ch <- fmt.Sprintf("while reading %s: %v", url, err)
	// 	return
	// }
	
	resp.Body.Close()
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, numBytes, url)
}

func truncName(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] != '.' {
			continue
		}
		s = s[i+1:]
	}
	for i := 0; i < len(s); i++ {
		if s[i] != '.' {
			continue
		}
		s = s[:i]
	}
	return s
}

//!-

