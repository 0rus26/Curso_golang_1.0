package main

import "fmt"

/*
Crea constantes CON TIPO y SIN TIPO. Imprime el valor de las mismas.
*/

const (
	a     = 42
	b int = 43
)

func main() {
	fmt.Printf("%v\t%T\t%v\t%T", a, a, b, b)

}
