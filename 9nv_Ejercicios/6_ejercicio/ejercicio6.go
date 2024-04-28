/*

Crea un programa que imprima tu OS y ARCH. Usa los siguientes comandos para correrlo
●	go run
●	go build
●	go install


*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("ARCH", runtime.GOARCH)
}
