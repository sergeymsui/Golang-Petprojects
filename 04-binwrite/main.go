package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	buf := bytes.NewBuffer([]byte{})

	if err := binary.Write(buf, binary.BigEndian, 1.004); err != nil {
		panic(err)
	}

	var num float64

	if err := binary.Read(buf, binary.BigEndian, &num); err != nil {
		panic(err)
	}

	fmt.Println(num)
}
