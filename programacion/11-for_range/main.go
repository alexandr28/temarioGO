package main

import "fmt"

func main() {
	/*
			numeros := []string{"uno", "dos", "tres", "cuatro"}
			for i, v := range numeros {
				fmt.Println(i, v)
			}
		paises := map[string]string{"pe": "Perú", "col": "Colombia", "ar": "Argentina", "br": "Brasil", "cl": "Chile", "ec": "Ecuador", "pr": "Paraguay"}
		for i, v := range paises {
			fmt.Println(i, v)
		}

		frase := "Hola Mundo"
		for _, letra := range frase {
			fmt.Println(string(letra))
		}
	*/
	// con el _ (guion bajo) estamos ignorando un valor u omitir el valor mas no el indice
	for _, entero := range []int{15, 16, 17, 18, 19, 20} {
		fmt.Println(entero)
	}
}
