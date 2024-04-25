package main

import (
	"encoding/json"
	"fmt"
)

/*
Comenzando con este código,
https://go.dev/play/p/qXkZG4ISFTX
marshal el slice []usuario a JSON. Hay una pequeña curva en este ejercicio
recuerda preguntarte qué necesitas para EXPORTAR un valor de un paquete
*/

type usuario struct {
	Nombre string
	Edad   int
}

func main() {
	u1 := usuario{
		Nombre: "Eduar",
		Edad:   32,
	}

	u2 := usuario{
		Nombre: "Juan",
		Edad:   27,
	}

	u3 := usuario{
		Nombre: "Che",
		Edad:   54,
	}

	usuarios := []usuario{u1, u2, u3}
	fmt.Println("Usuarios:", usuarios)

	bs, err := json.Marshal(usuarios)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
