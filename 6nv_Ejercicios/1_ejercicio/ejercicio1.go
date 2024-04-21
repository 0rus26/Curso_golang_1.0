package main

import "fmt"

/*

●	Ejercicio práctico
○	crea una func con el identificador foo que retorne un int
○	crea una func con el identificador bar que retorne un int y un string
○	Llama ambas funciones
○	Imprime sus resultados



*/
func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 1
}

func bar() (int, string) {
	return 1, "Hola"
}
