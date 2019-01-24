package main

import "fmt"

func main() {
	saludarVarios(18, "Ruben", "Alexandr", "Junnior", "Sofia")
}

func saludarVarios(edad uint8, nombres ...string) {
	for _, v := range nombres {
		fmt.Println("Hola ", v, " tienes ", edad, " años")
	}
}
