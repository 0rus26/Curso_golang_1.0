package main

import "fmt"

type vehiculo struct {
	puertas int
	color   string
}

type camion struct {
	vehiculo
	ruedas4x4 bool
}

type toyota struct {
	vehiculo
	lujoso bool
}

func main() {
	/*
		Usa el código del ejemplo anterior y almacena los valores de tipo persona en un mapa con la llave apellido.
		Accede a cada valor en el mapa. Imprime los valores. Imprime también los valores haciendo range sobre el slice.
	*/

	c := camion{
		vehiculo: vehiculo{
			puertas: 2,
			color:   "verde",
		},
		ruedas4x4: true,
	}

	hiunday := toyota{
		vehiculo: vehiculo{
			puertas: 4,
			color:   "blanca"},
		lujoso: false,
	}

	fmt.Println(c.puertas)
	fmt.Println(hiunday.puertas)

}
