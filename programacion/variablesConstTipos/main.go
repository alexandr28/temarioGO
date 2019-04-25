package main

import "fmt"

func main() {
	const usuario = "Alexandr Campos"
	fmt.Println("usuario:", usuario)
	/*
		Tipos de Datos
		no se puede realizar operaciones artimeticas con diferentes tipos de datos excepto que se haga casting
	*/
	var a int64 = 12121212
	var b int8 = 88
	//casting
	c := a + int64(b)
	fmt.Println("el resultado es:", c)

	/*
		uso de comodines %s

	*/
	name := "Mariluz"
	lastname := "Garro"
	fmt.Printf("Hola %s %s como estas?\n", name, lastname)
	fmt.Printf("lastname es de tipo: %T\n", lastname)
	fmt.Printf("C es de tipo: %T", c)
}
