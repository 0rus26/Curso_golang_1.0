/*
●	Comenzando con este código, extrae los valores del canal usando una declaración select
*/

package main

import (
	"fmt"
)

func main() {
	salir := make(chan int)
	c := gen(salir)

	recibir(c, salir)

	fmt.Println("A punto de finalizar.")
}

func gen(salir chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		salir <- 1
		close(c)
	}()

	return c
}

func recibir(c, s <-chan int) {
	for {

		select {
		case v := <-c:
			fmt.Println(v)
		case <-s:
			return
		}
	}
}
