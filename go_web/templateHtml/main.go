package main

import (
	"html/template"
	"log"
	"net/http"
)

type Persona struct {
	Nombre string
	Edad uint8
}

func renderizar(w http.ResponseWriter, r *http.Request)  {
	p := &Persona{"Alexandr",33}
	t, err := template.ParseFiles("views/index.html")
	//t,err := template.Must(template.ParseGlob("./views/index.html"))
	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
	t.Execute(w,p)
}

func main()  {
	mux:= http.NewServeMux()
	mux.HandleFunc("/",renderizar)
	log.Println("ejecutando servidor en http://localhost:3000")
	log.Println(http.ListenAndServe("3000",mux))
	log.Println("ejecucion terminada")
}
