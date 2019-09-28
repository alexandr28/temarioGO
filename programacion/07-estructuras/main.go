package main

import "fmt"

//Persona es un tipo de estructura, son equivalentes a clases
type Persona struct {
	Name    string
	Address string
	Age     uint8
	Phone   int
	Emails  []string
}

func main() {
	var persona1 Persona
	persona1.Name = "Alexandr"
	persona1.Address = "Enmanuel Kant 372"
	persona1.Age = 33
	persona1.Phone = 989292417
	fmt.Println(persona1.Name)

	emails := []string{"maluga@gmail.com", "m.garro@gmail.com"}
	person2 := Persona{
		"Mariluz",
		"calle los girasoles 145",
		28,
		962728383,
		emails,
	}
	fmt.Println(person2.Emails)
}
