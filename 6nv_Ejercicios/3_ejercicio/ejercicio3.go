package main

import "fmt"

/*

●	Usa la palabra clave “defer” para mostrar que una función diferida corre después que la función que la contiene finaliza o retorna.


*/
func main() {

	defer woo()
	far()
	//fmt.Println(far())
}

func woo() {

	defer func() {
		fmt.Println("Desde woo segunda función")
	}()
	fmt.Println("Desde woo primera función")

}
func far() {
	fmt.Println("Desde far")

}
