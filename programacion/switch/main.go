package main

import "fmt"

func main() {
	//en Go no se usa break  y es mejor omitirlo
	var a uint8 = 4
	switch a {
	case 1:
		// fallthrough -> sigue evaluando, la siguiente expresión
		fallthrough
	case 2:
		fmt.Println("ganaste un viaje")
	case 3:
		fallthrough
	case 4:
		fmt.Println("ganas una camioneta 4x4")
	default:
		fmt.Println("sigue intentando")
	}
	/*
		Otra forma de usar Swich en GO usando condicionales en cada case
	*/
	switch x := 7; {
	case x >= 0 && x <= 5:
		fmt.Println("estas entre semana")
	case x >= 6 && x <= 7:
		fmt.Println("estas en fin de semana")
	default:
		fmt.Println("no es un dia valido")
	}
}
