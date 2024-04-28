package main

import "fmt"

/*
Method Sets

Este ejercicio te ayudará a reforzar el concepto de method sets:
●	Crea un tipo struct persona
●	Adjunta el método hablar usando un receptor de tipo puntero
○	*persona
●	Crea un tipo interfaz humano
○	Para implícitamente implementar esa interfaz, un tipo humano debe tener el método hablar
●	Crea la función “diAlgo”
○	Haz que tome un humano como parámetro
○	Haz que llame al método hablar
●	Muestra lo siguiente en tu código
○	PUEDES pasar un valor de tipo *persona a diAlgo
○	NO puedes pasar un valor de tipo persona a diAlgo
●	Aquí hay una pista si necesitas ayuda


*/

type persona struct {
	Nombre   string
	Apellido string
}

type humano interface {
	hablar()
}

func (p *persona) hablar() {
	fmt.Println("Persona ", p.Nombre, p.Apellido, "está hablando")

}

func diAlgo(h humano) {
	h.hablar()

}

func main() {

	p1 := persona{Nombre: "Samuel", Apellido: "Martinez"}
	diAlgo(&p1)
	p1.hablar()
	//fmt.Println(p1)
}
