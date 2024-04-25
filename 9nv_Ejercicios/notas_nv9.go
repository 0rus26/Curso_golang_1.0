package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var contador int64
	const gs int = 100
	fmt.Println("Número de CPUs: ", runtime.NumCPU())
	fmt.Println("Número de Rutinas: ", runtime.NumGoroutine())
	runtime.GOMAXPROCS(6)
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("Contador", atomic.LoadInt64(&contador))
			wg.Done() // Cambiado de Add(-1) a Done()
		}()
		fmt.Println("Número de Rutinas: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Contador: ", contador)
}

/*
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	var contador int64
	const gs int = 100
	fmt.Println("Número de CPUs: ", runtime.NumCPU())
	fmt.Println("Número de Rutinas: ", runtime.NumGoroutine())
	runtime.GOMAXPROCS(1)

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {

			mu.Lock() // Se adquiere el bloqueo antes de modificar el contador
			v := contador
			v++
			runtime.Gosched()
			contador = v
			mu.Unlock() // Se libera el bloqueo después de modificar el contador
			wg.Done()   // Cambiado de Add(-1) a Done()
		}()
		fmt.Println("Número de Rutinas: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Contador: ", contador)
}

*/
