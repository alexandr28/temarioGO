package main

import "fmt"

func main() {
	wy()
}

func wy() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%T\n", r)
			fmt.Println("Recuperado  en wy", r)
		}
	}()
	fmt.Println("llamando WY")
	wx(8)
	fmt.Println("Retorna normalmente desde WX")
}

func wx(i int) {
	if i > 3 {
		fmt.Println("error errror")
		panic("El numero no puede ser mayor que 3")
	}
	defer fmt.Println("Defer en la función WY")
	fmt.Println("Imprimiendo en WY", i)
}
