package main

import "fmt"
import "dev/golang-book/cap8/packages/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
