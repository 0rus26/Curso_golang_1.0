package main

import "fmt"

/*

Closure es cuando “encerramos” el scope de una variable en un bloque de código. Para este ejercicio, crea una func el cual “encierra” el scope de una variable:
*/

func main() {
	w := woo()
	fmt.Println(w())

	w()
	fmt.Println(w())
}

func woo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
