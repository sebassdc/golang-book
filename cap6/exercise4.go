package main
import "fmt"

func makeOddGenerator() func()uint {
	x := uint(1)
	return func() (ret uint){
		ret = x
		x += 2
		return
	}
}

func main() {
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}
