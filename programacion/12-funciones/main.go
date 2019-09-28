package main

import "fmt"

func main() {
	//saludar("Alexander", 33)
	var n1, n2 int8
	n1 = 18
	n2 = 16
	r := suma(n1, n2)
	var edad uint8
	edad = 17
	fmt.Println("es un ", tipoEdad(edad))
	fmt.Println("la suma es: ", r)
	n := []int8{52, 63, 47, 19, 83, 120, 33, 44, 73, 78}
	maximo, minimo := maxmin(n)
	fmt.Println("Maximo: ", maximo)
	fmt.Println("Minimo: ", minimo)
}

func saludar(nombre string, edad uint8) {
	//fmt.Println("Hola como estas", nombre, "tu edad es", edad)
	fmt.Printf("hola %s tienes %d años de edad \n", nombre, edad)
}

// func return
func suma(a, b int8) int8 {
	return a + b
}

func tipoEdad(edad uint8) string {
	var tipo string
	switch {
	case edad < 12:
		tipo = "children"
	case edad < 18:
		tipo = "Adolecente"
	default:
		tipo = "Adulto"
	}
	return tipo
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
