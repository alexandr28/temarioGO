package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
)
// Producto contiene los datos de un articulo
type Producto struct {
	gorm.Model // ID, CreateAt, UpdateAt, DeleteAt
	Codigo string
	Nombre string
	Precio uint8
}
func main()  {
	db,err:=gorm.Open("mysql","root:alexandr2124@/gorm?charset=utf8&parseTime=True&loc=Local")
	if err!= nil {
		panic("error en la conexion a la base de datos")
	}
	defer db.Close()
	fmt.Println("se conecto a la BD")
	/*
	db.CreateTable(&Producto{})
	fmt.Println("se creo la tabla prodcuto")
	db.Create(&Producto{
		Codigo:"112464asd12334",
		Nombre:"Mesa",
		Precio:124,
	})

	db.Create(&Producto{
		Codigo:"11334",
		Nombre:"muebles",
		Precio:400,
	})*/
	var p Producto
	db.First(&p,2)
	fmt.Println("el precio del producto consultado es: ",p.Precio)

	p.Precio=100
	db.Save(&p)
	fmt.Println("El nuevo precio del producto es: ",p.Precio)
}
