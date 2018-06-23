package main
import (
	"fmt"
	"bytes"
)

func main()  {
	var buf bytes.Buffer
	buf.Write([]byte("test"))
	fmt.Println(buf)
}
