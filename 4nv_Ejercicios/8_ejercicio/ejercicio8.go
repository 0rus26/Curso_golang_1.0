package main

import "fmt"

/*
Crea un mapa con una llave TIPO string el cual representa el “nombre_apellido” de una persona y un valor de TIPO []string el cual almacena sus cosas favoritas. Almacena tres registros en tu mapa. Imprime todos sus valores con su índice de posición en el slice.

`eduar_tua`, `computadoras`, `montaña`, `playa`
`carlos_ramon`, `leer`, `comprar`, `música`
`juan_bimba`, `helado`, `pintar`, `bailar`

*/

func main() {
	//miMapa := make(map[string]int)
	//mapa := make(map[string][]string)
	mapa := map[string][]string{
		"eduar_tua":    {`computadoras`, `montaña`, `playa`},
		"carlos_ramon": {`leer`, `comprar`, `música`},
		"juan_bimba":   {`helado`, `pintar`, `bailar`},
	}
	for llave, valor := range mapa {
		fmt.Println("LLave: ", llave, "\nValor: ", valor)
		for i, v := range mapa[llave] {
			fmt.Println("\tId: ", i, "\tValor: ", v)
		}
	}

	//miMapa["clave1"] = 1
	//mapa["eduar_tua"] =  {`computadoras`, `montaña`, `playa`}

	/*for i, v := range slice_multiD {
		fmt.Println("Id:", i, "Valor", v)
		for i2, v2 := range slice_multiD[i] {
			fmt.Println("Id:", i2, "Valor", v2)
		}
	}

	agenda := map[string][]string{
		"Juan":  {"123456", "789012"},
		"María": {"456789"},
		"Pedro": {"987654", "345678", "901234"},
	}
	*/
}
