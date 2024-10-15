// https://youtu.be/LxJLuW5aUDQ?t=426

package main

import (
	"log"
	"main/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	err := s.Start()

	if err != nil {
		log.Fatal(err)
	}
}
