package main

import "fmt"

func toFarenheight(celcius float32) float32 {
	return (celcius * (9 / 5)) + 32
}

fmt.Println(toFarenheight(32))
