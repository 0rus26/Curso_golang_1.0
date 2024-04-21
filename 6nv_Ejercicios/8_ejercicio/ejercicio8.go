package main

import "fmt"

/*

●	Crea una func el cual retorna una func
●	Asigna la func retornada a una variable
●	Llama la func retornada

*/

func main() {
	f := woo()
	fmt.Println(f())
}

func woo() func() int {
	return func() int {
		return 1
	}
}
