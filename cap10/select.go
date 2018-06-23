package main

import (
	"fmt"
	"time"
)

func main() {
	print := fmt.Println
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <- c1:
				print(msg1)
			case msg2 := <- c2:
				print(msg2)
			}
		}
	}()
	var input string
	fmt.Scanln(&input)
}