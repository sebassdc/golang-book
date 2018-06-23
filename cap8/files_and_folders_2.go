// {2}
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("files_and_folders.go")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}
