package main

import "fmt"

func main() {
	// forma clasica de condiconal
	if 5 > 10 {
		fmt.Println("es verdadero")
	}
	fmt.Println("fin del ciclo")

	// condicionales anidadas

	if edad, nombre := 17, "Alexandr"; edad < 14 {
		fmt.Println(nombre, "es un niño")
	} else if edad < 18 {
		fmt.Println(nombre, "es un adolecente")
	} else {
		fmt.Println(nombre, "es un adulto")
	}

}
