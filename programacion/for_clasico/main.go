package main

import "fmt"

func main() {
	/*
			for i := 0; i < 5; i++ {
				var num int
				num = 1 + i
				fmt.Println(num)
			}
		for i := 8; i > 0; i-- {
			if i == 4 {
				break
			}
			fmt.Println(i)
		}*/

	matriz := [3][3]int{}
	valor := 0
	for fila := 0; fila < 3; fila++ {
		for columna := 0; columna < 3; columna++ {
			valor++
			matriz[fila][columna] = valor
		}
	}
	fmt.Println(matriz)
}
