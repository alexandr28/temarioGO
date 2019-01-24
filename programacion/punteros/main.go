package main

import "fmt"

func main() {
	a:=3
	fmt.Println("Antes de duplicar a =",a)
	// enviamos el puntero
	duplicar(&a)
	fmt.Println("despues de duplicar a =",a)
}
// el asterisco indica el valor enmemoria del puntero
func duplicar(x *int) {
	// *X = *x * 2
	*x *=2
	fmt.Println("dentro de la función duplicar x vale ",*x)
}
