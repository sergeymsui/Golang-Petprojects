package main

import (
	"fmt"
	"strconv"
)

func calculate(line string) float64 {

	for i, e := range line {

		s := string(e)
		v, _ := strconv.ParseFloat(line[:i], 64)
		u, _ := strconv.ParseFloat(line[i+1:], 64)
		if s == "+" {
			return float64(v + u)
		} else if s == "-" {
			return float64(v - u)
		} else if s == "*" {
			return float64(v * u)
		} else if s == "/" {
			return float64(v / u)
		}
	}

	return 0
}

func main() {
	var line string
	fmt.Scan(&line)

	fmt.Println(calculate((line)))
}
