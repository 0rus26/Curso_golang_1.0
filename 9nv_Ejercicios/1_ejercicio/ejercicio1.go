package main

import "fmt"

func main() {

	go func() {
		fmt.Println("Hola desde una goroutine1")
	}()

	go func() {
		fmt.Println("Hola desde una goroutine2")
	}()
}
