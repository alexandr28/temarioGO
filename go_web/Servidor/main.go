package main

import (
	"log"
	"net/http"
)

func main() {
	/*
		http.Handle("/", http.FileServer(http.Dir("public")))
		log.Println("Ejecutando Server en http://localhost:3000")
		log.Println(http.ListenAndServe("localhost:3000", nil))
	*/
	multiplexor := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	multiplexor.Handle("/", fs)
	log.Println("Ejecutando Server en http://localhost:3000")
	log.Println(http.ListenAndServe("localhost:3000", multiplexor))
}
