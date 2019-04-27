package main

import (
	"fmt"
	"log"
	"net/http"
)

type messageHandler struct{ message string }

func (m messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	m1 := &messageHandler{message: "<h1>Hola desde un Handler</h1>"}
	m2 := &messageHandler{message: "<h1>Hola desde UCV</h1>"}
	m3 := &messageHandler{message: "<h1>Hola desde UPAO</h1>"}
	mux.Handle("/saludar", m1)
	mux.Handle("/ucv", m2)
	mux.Handle("/upao", m3)
	log.Println("Ejecutando Server en http://localhost:3000")
	log.Println(http.ListenAndServe("localhost:3000", mux))

}
