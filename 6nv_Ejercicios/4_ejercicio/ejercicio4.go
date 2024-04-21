package main

import (
	"fmt"
)

/*

●	Usa la palabra clave “defer” para mostrar que una función diferida corre después que la función que la contiene finaliza o retorna.


*/

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func (p persona) presentar() {
	fmt.Println("Hola mi nombre es: ", p.nombre, p.apellido, "y tengo ", p.edad, " años de edad.")
}

func main() {
	p1 := persona{
		nombre:   "Juanma",
		apellido: "Santos Petardos",
		edad:     100,
	}

	p1.presentar()
}
