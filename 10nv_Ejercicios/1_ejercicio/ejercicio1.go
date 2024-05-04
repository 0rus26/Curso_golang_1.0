package main

/*
Ejercicio Práctico #1
●	Haz que este código funcione:
○	Usando una función literal, también conocida como, función anónima autoejecutable
■	Solución: https://play.golang.org/p/SHr3lpX4so
○	Un canal con búfer
■	Solución: https://play.golang.org/p/Y0Hx6IZc3U

*/

//usar patrones de diseño concurrente cuando usemos canales, por que el main es la unica go rutine que esta ejecutando
//necesitamos usar una go rutines, para que se encargue de mandar el valor al canal.
import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
}
