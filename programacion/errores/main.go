package main

import (
	"errors"
	"fmt"
)

func main() {
	r, err := divison(100, 0)
	if err != nil {
		fmt.Println("Error ", err)
		// con el return  se controla la interrupcion de la operación
		return
	}
	fmt.Println("Resultado: ", r)
}

func divison(dividendo, divisor float64) (resultado float64, err error) {
	if divisor == 0 {
		err = errors.New("No se puede dividir por cero")
	} else {
		resultado = dividendo / divisor
	}

	return
}
