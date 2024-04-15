package main

import "fmt"

/*
●	Usando un COMPOSITE LITERAL:
●	Crea un ARREGLO el cual tenga 5 VALORES de TIPO int
●	Asigna VALORES a cada posición dada por los índices.
●	Itera con Range sobre el arreglo e imprime los valores.
●	Usando el paquete fmt
○	Imprime el TIPO del arreglo


*/

// la diferencia es la definición del arreglo con [X] y el slice va vacio []

func main() {
	var arreglo [5]int
	arreglo = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T-%v \n", arreglo, arreglo)
	for id, v := range arreglo {
		fmt.Printf("ID: %v , Valor: %v \n", id, v)
	}

}
