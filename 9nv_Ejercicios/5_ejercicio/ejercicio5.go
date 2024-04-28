/*

Arregla la race condition que creaste en el ejercicio anterior usando un mutex
●	Tiene sentido eliminar runtime.Gosched()

*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Número de CPUs:", runtime.NumCPU())
	fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())
	runtime.GOMAXPROCS(1)

	var contador int32 = 0

	const gs = 100
	var wg sync.WaitGroup

	// creamos la variable del tipo sync.Mutex para poder usar los métodos del paquete sync Lock y Unlock

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {

			atomic.AddInt32(&contador, 1)
			fmt.Println("Valor de Contador:", atomic.LoadInt32(&contador))

			wg.Done()
		}()
		//fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Cuenta:", contador)
}
