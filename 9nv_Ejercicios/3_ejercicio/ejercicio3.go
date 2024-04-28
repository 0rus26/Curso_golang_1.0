/*

●	Usando gorutinas, crea un programa incremento
○	Haz que una variable guarde el valor del incremento
○	Lanza varias gorutinas
■	cada gorutina deberá
●	Leer el valor del incremento
○	Almacenarlo en una nueva variable
●	Ceder el procesador con runtime.Gosched()
●	Incrementa la nueva variable
●	Escribe el valor en la nueva variable de vuelta a la variable incremento
●	Usa waitgroups para esperar que todas las gorutinas finalicen
●	Lo anterior generará una race condition.
●	Comprueba que es una race condition usando el flag -race


-- EN este ejercicio hay race condition

*/

package main

/*
func main() {
	fmt.Println("Número de CPUs:", runtime.NumCPU())
	fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())
	runtime.GOMAXPROCS(1)
	fmt.Println("Número de CPUs:", runtime.NumCPU())
	var contador int32 = 0

	const gs = 100000
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt32(&contador, 1)
			runtime.Gosched()
			fmt.Println("Valor de Contador:", atomic.LoadInt32(&contador))
			wg.Done()
		}()
		fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Cuenta:", contador)
}
*/

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Número de CPUs:", runtime.NumCPU())
	fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())
	runtime.GOMAXPROCS(1)

	contador := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			fmt.Println(contador)
			wg.Done()
		}()
		//fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Cuenta:", contador)
}
