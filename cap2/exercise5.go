package main

import "fmt"

func main() {
	fmt.Println("Feets to meters")
	fmt.Println("")
	fmt.Println("Ingress feets: ")
	var feets float64
	fmt.Scanf("%f", &feets)
	meters := feets * 0.3048
	fmt.Println(meters, " meters")
}
