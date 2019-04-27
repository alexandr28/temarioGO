package main

import "../interfaces/animales"
func main()  {
	var p animales.Perro
	var g animales.Gato
	p.Nombre="Galgo"
	g.Nombre="Alexa"
	//AdoptarPerro(p)
	//AdoptarGato(g)
	AdoptarMascota(p)
	AdoptarMascota(g)
}

func AdoptarMascota(m animales.Mascota)  {
	m.Comunicarse()
}
/*
func AdoptarPerro(p animales.Perro)  {
	p.Comunicarse()
}
func AdoptarGato(g animales.Gato)  {
	g.Comunicarse()
}*/
