package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Dataset struct {
	Data   string `json:"data"`
	Number int    `json:"number"`
}

func redirect(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "https://ya.ru", http.StatusMovedPermanently)
}

func JSONHandler(res http.ResponseWriter, req *http.Request) {

	d := Dataset{
		Data:   "data_1",
		Number: 1,
	}

	resp, err := json.Marshal(d)

	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	res.Header().Set("content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(resp)

	log.Println("This is JSON handler")
}

func apiPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello world"))
}

func mainPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("This is main page"))
}

func init() {
	// fmt.Println("Hello world")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainPage)
	mux.HandleFunc("/api", apiPage)
	mux.HandleFunc("/json", JSONHandler)
	mux.HandleFunc("/redirect", redirect)

	mux.HandleFunc("/ls", http.FileServer(http.Dir(".")))

	mux.HandleFunc("/file", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "./static/test.txt")
	})

	mux.HandleFunc("/golang", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "./cmd/main.go")
	})

	http.FileServer(http.Dir("./static"))

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}
}
