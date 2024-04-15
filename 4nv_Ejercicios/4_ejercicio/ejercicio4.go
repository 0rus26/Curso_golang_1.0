package main

import "fmt"

/*
Usando el código del ejercicio anterior, usa SLICING para crear los siguientes nuevos slices el cual serán impresos:

*/

func main() {

	arreglo := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	arreglo = append(arreglo, (52))
	fmt.Println(arreglo)
	arreglo = append(arreglo, 53, 54, 55)
	fmt.Println(arreglo)
	arreglo2 := []int{56, 57, 58, 59, 60}
	arreglo = append(arreglo, arreglo2...)
	fmt.Println(arreglo)
}
