package main

import "fmt"

func main() {
	/*
			Arrays
			Todos los valores son del mismo tipo de dato
			tamaño fijo

		var nombres [3]string
		nombres[0] = "Alexandra"
		nombres[1] = "Nicolas"
		nombres[2] = "Arianna"
		fmt.Println(nombres)
	*/
	nombres := [3]string{"Alexandra", "Rosa", "Carlos"}
	size := len(nombres)
	fmt.Println("el tamaño del arreglo es: ", size)
	fmt.Println(nombres[0])
	fmt.Println(nombres[0:2])
}
