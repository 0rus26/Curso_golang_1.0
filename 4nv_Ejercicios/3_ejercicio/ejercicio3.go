package main

import "fmt"

/*
Usando el código del ejercicio anterior, usa SLICING para crear los siguientes nuevos slices el cual serán impresos:

*/

func main() {

	arreglo := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(arreglo[:5])
	fmt.Println(arreglo[5:])
	fmt.Println(arreglo[2:7])
	fmt.Println(arreglo[1:6])
}
