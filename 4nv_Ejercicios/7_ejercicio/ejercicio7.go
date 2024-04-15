package main

import "fmt"

/*
Crear un slice de slice de string ([][]string). Almacena los siguientes valores en un slice multi-dimensional:

"Batman", "Jefe", "Vestido de negro"
"Robin", "Ayudante", "Vestido de colores"


*/

func main() {
	slice_multiD := [][]string{{"Batman", "Jefe", "Vestido de negro"}, {"Robin", "Ayudante", "Vestido de colores"}}
	fmt.Println(slice_multiD[0][0])
	for i, v := range slice_multiD {
		fmt.Println("Id:", i, "Valor", v)
		for i2, v2 := range slice_multiD[i] {
			fmt.Println("Id:", i2, "Valor", v2)
		}
	}
}
