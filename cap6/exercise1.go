package main

import "fmt"

func sum(numbers []int) int {
	sumer := 0
	for _, value := range numbers{
		sumer += value
	}
	return sumer
}

func main() {
	slice := []int{ 98, 93, 77, 82, 83 }
	fmt.Println(sum(slice))
}
