package main

import (
	"io"
	"os"
)

func main() {
	io.WriteString(os.Stdout, "LOL\n")

	io.WriteString(os.Stderr, "LOL 2\n")

	buf := []byte{0xAF, 0xFF, 0xFE}
	for i := 0; i < 200; i++ {
		if _, e := os.Stdout.Write(buf); e != nil {
			panic(e)
		}
	}
}
