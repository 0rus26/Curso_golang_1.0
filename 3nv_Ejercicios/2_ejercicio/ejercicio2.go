package main

import "fmt"

/*
Imprime el rune code point de las letras del alfabeto en mayúsculas tres veces. La salida debe ser como esto:
65
	U+0041 'A'
	U+0041 'A'
U+0041 'A'
66
	U+0042 'B'
	U+0042 'B'
	U+0042 'B'
 … hasta el resto de letras en el alfabeto.

*/
func main() {

	for i := 65; i < 90; i++ {
		fmt.Printf("%v", i)
		for x := 0; x < 3; x++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

}
