// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

// package tempconv
package ch2tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//−273.15°C
//500degC =  226.85K
//30degC = 303.15
func CtoK(c Celsius) Kelvin { return Kelvin(c + 273.15)}

func KtoC(k Kelvin) Celsius { return Celsius(k - 273.15)}

func FtoK(f Fahrenheit) Kelvin { return Kelvin(FToC(f) + 273.15) }

func KtoF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(Celsius(k - 273.15)))}

//!-