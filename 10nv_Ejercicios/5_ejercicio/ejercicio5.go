//●	Demuestra el uso del idioma coma ok con este código.

// v, ok:= <- canal recibir del canal, leer una posición del canal
// x eso a la segunda ocación ya no hay data y da 0 y un false
// por que solo se envió un dato al canal

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 50
	}()
	v, ok := <-c
	fmt.Println(v, ok)

	close(c)
	v, ok = <-c
	fmt.Println(v, ok)

}
