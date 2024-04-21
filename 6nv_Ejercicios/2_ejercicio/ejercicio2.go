package main

import "fmt"

/*

●	Crea una función con el identificador foo que
○	Tome un parámetro variable de tipo int
○	Pásale un valor de tipo []int a la función (haz el despliegue del []int)
○	Retorna la suma de todos los valores de tipo int pasados como argumento.
●	Crea una func con el identificador bar que
○	Tome un parámetro de tipo []int
○	Retorne la suma de todos los valores de tipo int pasados.



*/
func main() {

	xx := []int{1, 2, 3, 4, 5}
	fmt.Println(woo(xx...))
	fmt.Println(far(xx))
	//fmt.Println(far())
}

func woo(xi ...int) string {
	suma := 0
	for i := 0; i <= len(xi)-1; i++ {
		suma += xi[i]

	}
	msg := fmt.Sprint("La suma de xi= ", suma)
	return msg

}
func far(xi []int) string {
	suma := 0
	for i := 0; i <= len(xi)-1; i++ {
		suma += xi[i]

	}
	msg := fmt.Sprint("La suma de xi= ", suma)
	return msg

}
