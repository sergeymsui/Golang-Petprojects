package main

import (
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	reader, writer := io.Pipe()

	cmd := exec.Command("echo", "This is echo")
	cmd.Stdout = writer

	f, err := os.OpenFile("file.txt", os.O_CREATE|os.O_RDWR, 0777)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	go func() {
		defer reader.Close()

		if _, err := io.Copy(f, reader); err != nil {
			log.Fatal(err)
		}

	}()

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
