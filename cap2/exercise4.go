package main

import "fmt"

func main() {
	fmt.Println("Farenheit to Celsius")
	fmt.Println("")
	fmt.Println("Ingress a temp in Farenheit")
	var f float64
	fmt.Scanf("%f", &f)
	c := (f - 32) * (5/9.0)
	fmt.Println("The temp in Celsius is", c)
}
