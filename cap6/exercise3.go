package main
import "fmt"
import "math"

func findGreater(args ...float64) float64 {
	max := float64(math.MinInt64)
	for _, value := range args {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	fmt.Println(findGreater(2, 3, 4, 56 , 21, 23, 54))
}
