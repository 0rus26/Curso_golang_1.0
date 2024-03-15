package main

import "fmt"

/*
1.	Crea tu propio tipo. Haz que el tipo subyacente, raíz o implícito sea un int.
2.	Crea una VARIABLE de tu nuevo TIPO con el IDENTIFICADOR “x” usando la palabra clave “VAR”
3.	En func main
*/

type numero int

var x numero

func main() {
	fmt.Println(x)
	fmt.Printf("La variable x es del tipo %T\n", x)
	x = 42
	fmt.Printf("x= %v", x)
}
