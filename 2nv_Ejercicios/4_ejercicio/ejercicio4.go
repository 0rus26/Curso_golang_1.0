package main

import "fmt"

/*
Crea constantes CON TIPO y SIN TIPO. Imprime el valor de las mismas.
*/

var a = 42

func main() {
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#x\t", b, b, b)

}
