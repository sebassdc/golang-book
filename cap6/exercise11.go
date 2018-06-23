package main
import "fmt"

func swap(x, y *int) {
	*x = *x + *y
	*y = *x - *y
	*x = *x - *y
}

func main() {
	x := 1
	y := 2
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}
