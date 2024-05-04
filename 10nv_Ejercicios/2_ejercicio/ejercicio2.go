/*
Ejercicio Práctico #2
●	Haz que este código funcione:
○	https://play.golang.org/p/oB-p3KMiH6
■	Solución:  https://play.golang.org/p/isnJ8hMMKg
○	https://play.golang.org/p/_DBRueImEq
■	Solución: https://play.golang.org/p/mgw750EPp4

*/

package main

import (
	"fmt"
)

func main() {
	//se cambia a bidireccional el canal
	cs := make(chan int)

	go func() {
		cs <- 42
	}()

	fmt.Println(<-cs)
	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
