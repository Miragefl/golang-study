package main

import (
	"fmt"
	"io"
	"os"
)

func readOps() {
	file, _ := os.Open("text2.txt")
	for {
		buf := make([]byte, 3)
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Println("read byte length", n)
		fmt.Println("read content", string(buf))
	}
	file.Close()

}

func main() {
	readOps()
}
