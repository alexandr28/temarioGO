package main

import "fmt"

func main() {
	/*
		idiomas := make(map[string]string)
		idiomas["es"] = "Español"
		idiomas["fr"] = "Frances"
		idiomas["us"] = "Ingles"
		fmt.Println(idiomas)
	*/
	idiomas := map[string]string{"es": "Español", "en": "Ingles", "fr": "Frances", "ger": "Aleman"}
	fmt.Println(idiomas["fr"])
	fmt.Println(idiomas)
	delete(idiomas, "ger")
	fmt.Println(idiomas)

	numeros := map[int]string{1: "BMW", 2: "Toyota", 8: "Kia", 10: "Hyundai"}
	fmt.Println(numeros)
}
