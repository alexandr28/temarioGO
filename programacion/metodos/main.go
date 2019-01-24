package main

import "fmt"

type Persona struct {
	name string
	age int8
}

type Numero int
func (p Persona) saludar(){
	fmt.Println("Hola soy una persona")
}

func (n Numero) presentarse() {
	fmt.Println("hola soy un numero")
}

func (p *Persona) asignarEdad(i int8) {
	if i>=0{
		p.age=i
	}else{
		fmt.Println("edad no validad.. nose permiten edades negativas")
	}
}

func main()  {
	var persona Persona
	var numero Numero
	persona.saludar()
	numero.presentarse()
	persona.asignarEdad(35)
	fmt.Println( persona)
}
