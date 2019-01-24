package main

import (
	"fmt"
	"log"
	"net/http"
)

// funcion q recibe un mensaje personalizado
func messageHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola este es un HandlerFunc</h1>")
}

// esta funcion devuelve un handler del mensaje personalizado
func messageHFCuston(message string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, message)
		},
	)
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	mux.HandleFunc("/saludo", messageHandlerFunc)
	mux.Handle("/otros", messageHFCuston("<h1>Hola amigos</h1>"))
	mux.Handle("/despedida", messageHFCuston("<h1>adios amigos</h1>"))
	log.Println("Ejecutando Server en http://localhost:3000")
	log.Println(http.ListenAndServe("localhost:3000", mux))

}
