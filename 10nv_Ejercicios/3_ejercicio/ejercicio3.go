/*
Ejercicio Práctico #3
●	Comenzando con este código, extrae los valores del canal usando un ciclo for range
solución: https://play.golang.org/p/U2iGzRTtbxg
*/
package main

import (
	"fmt"
)

func main() {
	c := gen()
	recibir(c)

	fmt.Println("A punto de finalizar.")
}

// se debe definir el canal para recibir desde el canal
func recibir(c <-chan int) {

	for i := range c {
		fmt.Println(i)
	}

}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
