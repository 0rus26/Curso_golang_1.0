package main

import "fmt"

/*
●	Usando un COMPOSITE LITERAL:
●	Crea un SLICE de TIPO int
●	Asigna 10 VALORES
●	Haz Range sobre el slice e imprime los valores.
●	Usando format para imprimir
○	Imprime el TIPO del slice



*/

func main() {

	arreglo := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Printf("%T-%v \n", arreglo, arreglo)

}
