package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// Imprimimos información útil del entorno de ejecucón CPUS

	fmt.Printf("CPUS: %v\n", runtime.NumCPU())
	fmt.Printf("GOroutines inicio: %v\n", runtime.NumGoroutine())
	runtime.GOMAXPROCS(1)
	//Numero de gourutines que se deben ejecutar
	wg.Add(2)
	go func() {
		fmt.Println("Hola desde una goroutine1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hola desde una goroutine2")
		wg.Done()
	}()

	fmt.Printf("CPUS: %v\n", runtime.NumCPU())
	fmt.Printf("GOroutines medio: %v\n", runtime.NumGoroutine())
	wg.Wait() // espera a que todas las go routines hayan sido ejecutadas
	fmt.Printf("GOroutines final: %v\n", runtime.NumGoroutine())
	fmt.Printf("CPUS: %v\n", runtime.NumCPU())
	fmt.Println("Terminando ejecución")
}
