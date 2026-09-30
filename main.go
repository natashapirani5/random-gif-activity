package main

import (
	"fmt"
	"log"
	"net/http"
)

const page = "<!DOCTYPE html>\n" +
	"<html lang=\"en\">\n" +
	"<head>\n" +
	"  <meta charset=\"utf-8\">\n" +
	"  <title>Hello</title>\n" +
	"  <style>\n" +
	"    body { background-color: #dbeafe; font-family: sans-serif; text-align: center; padding-top: 4rem; }\n" +
	"  </style>\n" +
	"</head>\n" +
	"<body>\n" +
	"  <h1>Hello, visitor!</h1>\n" +
	"</body>\n" +
	"</html>\n"

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, page)
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}