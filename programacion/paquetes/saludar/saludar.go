package saludar

import "fmt"
// para utilizar desde otro paquete  con mayuscula al iniciar
var Example string
var exam string
// se pone en mayusculas por q va hacer una función exportada
func Saludar(nombre string) {
	exam="como estas"
	fmt.Println("Hola ", nombre)
}
