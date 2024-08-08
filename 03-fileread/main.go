package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("./file.txt")

	if err != nil {
		panic(err)
	}

	c, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	f.Close()

	fmt.Printf("Content: %s\n", c)
}
