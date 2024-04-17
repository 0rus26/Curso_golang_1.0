package main

import "fmt"

func main() {
	/*
		camion := struct {
			puertas   int
			color     string
			ruedas4x4 bool
		}{
			puertas:   3,
			color:     "verde",
			ruedas4x4: true,
		}

		fmt.Println(camion.puertas)
	*/

	structura := struct {
		nombre  string
		amigos  map[string]int
		bebidas []string
	}{
		nombre: "Fernando",
		amigos: map[string]int{
			"ana":  111,
			"beto": 222,
			"caro": 333,
		},
		bebidas: []string{"coca", "limonada", "cafe"},
	}

	fmt.Println(structura)
	fmt.Println(structura.amigos)
	fmt.Println(structura.bebidas)

	for k, v := range structura.amigos {
		fmt.Println(k, v)
	}

	for i, v := range structura.bebidas {
		fmt.Println(i, v)
	}

}
