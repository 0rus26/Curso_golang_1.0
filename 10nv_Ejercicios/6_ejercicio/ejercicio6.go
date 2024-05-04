/*

●	Escribe un programa que:
○	Ponga 100 números en un canal
○	Extraiga los números del canal y los imprima

*/

package main

import "fmt"

func main() {
	canal_numeros := make(chan int)

	enviar(canal_numeros)

	for valores := range canal_numeros {
		fmt.Println("leyendo canal-> ", valores)
	}
	//leer_canal(canal_numeros)
	//close(canal_numeros)
	fmt.Println("Finalizando")
}

func enviar(c chan<- int) {
	go func() {
		for i := 0; i <= 10; i++ {
			c <- i
		}
		close(c)
	}()

}

func leer_canal(c <-chan int) {

}
