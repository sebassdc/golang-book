// {7}
package main

import "fmt"

func main() {
	var x map[string]int // Wrong way
	x["key"] = 10
	fmt.Println(x)
}
