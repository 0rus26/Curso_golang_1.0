package main

import (
	"fmt"
)

/*
●	Crea y usa una func anónima
func(definir variables de entrada) {lógica de la función salida}(var entrada)
*/

var x int

var metodo1 func() = func() { fmt.Println("Función desde afuera del main") }

func main() {
	//x := "Joe"
	funcion := func() { fmt.Println("Ejecutando una función anónima") }
	//xx := func(name string) string { return fmt.Sprintln("Hola desde función anónima", name) }(x)
	//xx2 := func(name string) { fmt.Println("Hola desde función anónima", name) }
	//fmt.Println("salida funcion anonima:", xx)
	/*
		func() {
			for i := 0; i <= 99; i++ {
				fmt.Println(i + 1)
			}
			fmt.Println("Terminó")
		}()
	*/
	//fmt.Println("Hola ---->")
	funcion()
	fmt.Printf("%T,%v\n funcion: ", funcion, funcion)
	//fmt.Printf("%T,%v\n xx: ", xx, xx)
	//fmt.Printf("%T,%v\n xx2: ", xx2, xx2)
	metodo1()
	fmt.Printf("Metodo1: %T,%v", metodo1, metodo1)

	//fmt.Println("", metodo1)
}
