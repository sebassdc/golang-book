// 1
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age int
}

type ByAge []Person

func (this ByAge) Len() int {
	return len(this)
}
func (this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}
func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main()  {
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}
	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}
