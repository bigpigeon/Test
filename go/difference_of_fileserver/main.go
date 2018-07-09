package main

import (
	"log"
	"net/http"
)

func listen1() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func listen2() {
	fs := http.FileServer(http.Dir(".")) // root at the root directory.
	http.Handle("/static/", fs)          //leave off the StripPrefix call.

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func main() {
	//listen1()
	listen2()
}
