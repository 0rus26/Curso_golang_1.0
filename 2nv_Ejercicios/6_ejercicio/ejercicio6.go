package main

import "fmt"

/*
Usando iota, crea 4 constantes para los PRÓXIMOS 4 años. Imprime los valores de las constantes.

*/

const (
	a = 2024 + iota
	b = 2024 + iota
	c = 2024 + iota
	d = 2024 + iota
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
