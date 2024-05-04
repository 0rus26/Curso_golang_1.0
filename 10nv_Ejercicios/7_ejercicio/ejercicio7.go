/*

●	Escribe un programa que:
○	Lance 10 gorutinas
■	Cada gorutina agrega 10 números a un canal
○	Extrae los números del canal e imprímelos



package main

import (
	"fmt"
	"math/rand"
)

func main() {
	canal_numeros := make(chan int)

	enviar(canal_numeros)

	leer_canal(canal_numeros)
	//close(canal_numeros)
	fmt.Println("Finalizando")
}

func enviar(c chan<- int) {
	num_gr := 10
	for i := 1; i <= num_gr; i++ {
		go func() {
			func() {
				for valor := range operacion() {
					c <- valor
				}
			}()

		}()

	}

}

func operacion() []int {
	var c []int
	for i := 0; i <= 10; i++ {
		c = append(c, i*rand.Intn(100))

	}
	return c
}

func leer_canal(c <-chan int) {

	for x := 0; x <= 100; x++ {

		fmt.Println(x, ":leyendo canal-> ", <-c)

	}
}

*/
//otra forma de resolver

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 10
	y := 10
	c := gen(x, y)
	leer_canal(x, y, c)
}

func gen(a, b int) <-chan int {
	canal := make(chan int)

	for i := 0; i <= a; i++ {
		go func() {
			for i := 0; i <= b; i++ {
				canal <- i * rand.Intn(100)
			}
		}()
	}

	return canal
}

func leer_canal(x, y int, canal <-chan int) {

	for i := 0; i < x*y; i++ {
		fmt.Println("id:", i, " valor canal:", <-canal)
	}
}
