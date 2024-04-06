package main

import "fmt"

/*
a.	a x asignale 42
b.	a y asignale “James Bond”
c.	a z asignale true

Usa fmt.Sprintf para imprimir todos los VALORES en un solo string. ASIGNA el valor retornado de TIPO string usando el operador de declaración corta a  la VARIABLE con el IDENTIFICADOR “s”
Imprime el valor almacenado por la variable “s”

*/

var x int = 42
var y string = "james Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
