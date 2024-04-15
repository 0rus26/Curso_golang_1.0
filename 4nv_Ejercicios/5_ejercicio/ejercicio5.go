package main

import "fmt"

/*
	Para BORRAR de un slice, usamos APPEND en conjunto con SLICING(dividir). Para este ejercicio sigue los siguientes pasos:
●	Comienza con un slice
○	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
●	Usa APPEND & SLICING para obtener los siguientes valores el cual se los debes asignar a una variable “y” y luego imprimir:
○	[42, 43, 44, 48, 49, 50, 51]


*/

func main() {

	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := append(slice[:3], slice[6:]...)
	fmt.Println(y)
	//slice2 := []int{56, 57, 58, 59, 60}
	//slice = append(slice, slice2...)
	//fmt.Println(slice)

}
