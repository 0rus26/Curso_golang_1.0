package main

import "fmt"

type persona struct {
	name       string
	last_name  string
	saboresFav []string
}

func main() {

	/*
		Crea tu propio tipo “persona” el cual tendrá un tipo subyacente tipo “struct” de manera que pueda almacenar la siguiente data:
		●	Nombre
		●	Apellido
		●	Sabores de helado favoritos
		Crea dos VALORES de TIPO persona. Imprime los valores, usa range sobre los elementos en el slice el cual almacena los valores de helados favoritos.

	*/

	p1 := persona{
		name:       "Andres",
		last_name:  "Roldan",
		saboresFav: []string{"fresa", "brownie", "Chocolate"},
	}

	p2 := persona{
		name:       "Jorge",
		last_name:  "Fuentes",
		saboresFav: []string{"Vainilla", "Ron pasas", "Chicle"},
	}

	fmt.Println(p1.name, p1.last_name, "Sabores Favoritos:")
	for i, v := range p1.saboresFav {
		fmt.Println("\t", "id ", i, " Valor ", v)

	}

	fmt.Println(p2.name, p2.last_name, "Sabores Favoritos:")
	for i, v := range p2.saboresFav {
		fmt.Println("\t", "id ", i, " Valor ", v)

	}

}
