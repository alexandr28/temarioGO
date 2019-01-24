package main

import (
	"../paquetes/saludar"
	"fmt"
)
func main() {
	nombre := "Alexandr"
	saludar.Saludar(nombre)
	saludar.Example="Prueba 1 2 3"
	fmt.Println(saludar.Example)
}
