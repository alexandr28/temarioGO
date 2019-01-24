package main

import "fmt"

func main() {
	//saludar("Alexander", 33)
	n := []int8{52, 63, 47, 19, 83, 120, 33, 44, 73, 78}
	maximo, minimo := maxmin(n)
	fmt.Println("Maximo: ", maximo)
	fmt.Println("Minimo: ", minimo)
}

func saludar(nombre string, edad uint8) {
	//fmt.Println("Hola como estas", nombre, "tu edad es", edad)
	fmt.Printf("hola %s tienes %d años de edad \n", nombre, edad)
}
func maxmin(numeros []int8) (max int8, min int8) {
	for _, v := range numeros {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return
}
