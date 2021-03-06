package main

import "fmt"

func main() {
	/**
	Slices => Apuntador a un array, tamaño que no es fijo, capacidad
	make(tipo de dato del slice, tamaño inicial,capacidad inicial)
	*/
	var nombres []string
	fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Arianna")
	fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Adriana")
	fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Mariluz")
	fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Luciana")
	fmt.Println(nombres)
	fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))

	// con make permite construir slice dando 3 valores
	//make(tipo de dato, tamaño inicial, capacidad inicial(opcional))
	alumnos := make([]string, 0)
	alumnos = append(alumnos, "Jose")
	alumnos = append(alumnos, "Felipe")
	fmt.Println(alumnos)

	// haciendo slice sin inicializar y de forma rapida
	names := []string{"Alexandr", "Mariluz", "Andrea", "Ana Claudia"}
	fmt.Println(names)
}
