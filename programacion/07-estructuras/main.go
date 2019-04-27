package main

import "fmt"

//Persona es un tipo de estructura, son equivalentes a clases
type Persona struct {
	Name    string
	Address string
	Age     uint8
	Phone   int
}

func main() {
	var persona1 Persona
	persona1.Name = "Alexandr"
	persona1.Address = "Enmanuel Kant 372"
	persona1.Age = 33
	persona1.Phone = 989292417
	fmt.Println(persona1.Name)
}
