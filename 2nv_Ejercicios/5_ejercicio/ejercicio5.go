package main

import "fmt"

/*
Crea una variable de tipo string usando un string literal no interpretado (raw string literal). Imprímelo.
*/

func main() {
	string_literal := `esto es un string  
	literal`
	fmt.Printf("%T\t%v", string_literal, string_literal)
}
