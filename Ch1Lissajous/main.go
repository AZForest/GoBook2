// package ch1lissajous

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

//!-main
// Packages not needed by version in book.

//!+main


var palette = []color.Color{
	color.Black,
	color.RGBA{0xff, 0x7e, 0x03, 0xe0}, //#ff7e03e0
	color.RGBA{0x59, 0xff, 0xa3, 0x3b}, //#59ffa33b
	color.RGBA{0xf7, 0x47, 0x2b, 0x1}, //#f7472b
	color.RGBA{0x2b, 0xd7, 0xf7, 0xd9}, //#2bd7f7d9
	color.RGBA{0xf7, 0x2b, 0xe8, 0xb5}, //#f72be8b5
	color.RGBA{0xff, 0xff, 0x58, 0xde}, //#ffff58de
	color.White}

// const (
// 	whiteIndex = 0 // first color in palette
// 	blackIndex = 1 // next color in palette
// )

const (
	blackIndex = 0 // first color in palette
	greenIndex = 1 // next color in palette
	redIndex = 2
	blueIndex = 3
	purpleIndex = 4
	yellowIndex = 5
	orangeIndex = 6 
	whiteIndex = 7
)

var timesIncremmented = 0;
var currentIndex uint8 = 0;

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	var cycles = 5
	var size = 100
	var delay = 8

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			// fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
			// fmt.Printf("%s\n", r.Host)

			if err := r.ParseForm(); err != nil {
				log.Print(err)
			}
			for k, v := range r.Form {
				if k == "cycles" {
					cyc, err := strconv.Atoi(v[0])
					if err != nil {
						fmt.Printf("Error while reading URL param, %s", err)
					} else {
						// fmt.Printf("Cyc val: %d\n", cyc)
						cycles = cyc
					}
				}
				if k == "size" {
					s, err := strconv.Atoi(v[0])
					if err != nil {
						fmt.Printf("Error while reading URL param, %s", err)
					} else {
						// fmt.Printf("Size val: %d\n", s)
						size = s
					}
				}
				if k == "delay" {
					dl, err := strconv.Atoi(v[0])
					if err != nil {
						fmt.Printf("Error while reading URL param, %s", err)
					} else {
						// fmt.Printf("Size val: %d\n", dl)
						delay = dl
					}
				}
				// fmt.Printf("Form[%q] = %q\n", k, v)
			}

			lissajous(w, cycles, size, delay)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout, cycles, size, delay)
}

func lissajous(out io.Writer, c int, s int, d int) {
	const (
		// cycles  = 5    // 5def - number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		// size    = 100   // 100def - image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		// delay   = 8     // delay between frames in 10ms units
	)
	cycles := c
	size := s
	delay := d

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			// Another implentation
			if timesIncremmented % 100000 == 0 {
				updateIndex()
			}
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				// blueIndex,
				currentIndex,
				)
			timesIncremmented++
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func updateIndex() {
	currentIndex = uint8(rand.Intn(7) + 1)
}
// var palette = []color.Color{color.White, color.Black}
//!-main