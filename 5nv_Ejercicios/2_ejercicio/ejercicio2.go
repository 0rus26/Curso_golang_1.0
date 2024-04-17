package main

import "fmt"

type persona struct {
	name       string
	last_name  string
	saboresFav []string
}

func main() {

	/*
		Usa el código del ejemplo anterior y almacena los valores de tipo persona en un mapa con la llave apellido.
		Accede a cada valor en el mapa. Imprime los valores. Imprime también los valores haciendo range sobre el slice.
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

	p3 := persona{
		name:       "Elena",
		last_name:  "Garcia",
		saboresFav: []string{"Mani", "Chicharron"},
	}
	/*

		fmt.Println(p1.name, p1.last_name, "Sabores Favoritos:")
		for i, v := range p1.saboresFav {
			fmt.Println("\t", "id ", i, " Valor ", v)

		}

		fmt.Println(p2.name, p2.last_name, "Sabores Favoritos:")
		for i, v := range p2.saboresFav {
			fmt.Println("\t", "id ", i, " Valor ", v)

		}
	*/
	mapa := map[string]persona{
		p1.last_name: p1,
		p2.last_name: p2,
		p3.last_name: p3,
	}
	i := 1
	for _, v := range mapa {

		fmt.Println("Usuario:", i, "Nombre: ", v.name, "Apellido: ", v.last_name)
		for i, v := range v.saboresFav {
			fmt.Println("\tId:", i, "Valor: ", v)

		}
		fmt.Println("		+_____+")
		i++

	}

}
