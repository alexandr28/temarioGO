package main

import "fmt"

func main()  {
	imprenta:=[]string{"php","js","java","git","html","css","Asp.net","C#","Mysql","oracle","sqlServer","python","GO"}
	libreria:=[]string{}
	//canales
	fotocopiadora:=make(chan string)
	fotocopiado:=make(chan string)

	go func() {
		for _,libro:= range imprenta {
			fotocopiadora <-libro
		}
	}()

	// se carga la fotocopiadora
	go func() {
		for  {
			// entrega el contenido de la fotocopiadora a la variable libro
			libro := <-fotocopiadora
			fotocopiado <- libro
			//agregar el libro a la libreria
			libreria = append(libreria,libro)
			if len(imprenta) == len(libreria) {
				// cerrar el canal fotocopiado
				close(fotocopiado)
			}
		}
	}()
	// se deja en la bodega destino
	fmt.Println("** Listado de Libros **")
}
