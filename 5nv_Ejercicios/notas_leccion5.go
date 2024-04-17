package main

import "fmt"

/*
type persona struct {
	nombre       string
	apellido     string
	d_nacimiento int
}

type employees struct {
	persona
	labor string
}
*/

func main() {
	/*	p1 := persona{nombre: "Juan",
			apellido:     "Guerrero",
			d_nacimiento: 1500}

		fmt.Println(p1)
		fmt.Println("Nombre: ", p1.nombre)
		fmt.Println("Apellido: ", p1.apellido)
		fmt.Println("Año nacimiento: ", p1.d_nacimiento)

		e1 := employees{persona: p1, labor: "secretario"}
		fmt.Println("Empleado: ", e1.nombre, "Edad: ", time.Now().Year()-e1.d_nacimiento, "años")
	*/

	// struct anónimo

	p1 := struct {
		nombre       string
		apellido     string
		d_nacimiento int
	}{
		nombre:       "vermelio",
		apellido:     "lopez",
		d_nacimiento: 1600,
	}

	fmt.Println(p1)

}
