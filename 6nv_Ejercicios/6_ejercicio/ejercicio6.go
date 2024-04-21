package main

import (
	"fmt"
)

/*
●	Crea y usa una func anónima
func(definir variables de entrada) {lógica de la función salida}(var entrada)
*/

func main() {
	x := "Joe"
	funcion := func() { fmt.Println("Ejecutando una función anónima") }
	xx := func(name string) string { return fmt.Sprintln("Hola desde función anónima", name) }(x)
	fmt.Println("salida funcion anonima:", xx)
	func() {
		for i := 0; i <= 99; i++ {
			fmt.Println(i + 1)
		}
		fmt.Println("Terminó")
	}()

	funcion()
	fmt.Printf("%T,%v", funcion)
}
