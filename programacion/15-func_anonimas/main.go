package main

import "fmt"

func main() {
	/*
		anonimo := func() {
			fmt.Println("Dentro de una variable anonima")
		}
		anonimo()*/
	sec := secuencia()
	fmt.Println(sec())
}

func secuencia() func() int32 {
	var x int32
	return func() int32 {
		x++
		return x
	}
}
