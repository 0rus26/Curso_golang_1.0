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

	born_year := 1995
	i := born_year
	now := 2024
	for {
		if i > now {
			break
		}

		fmt.Printf("Año->(%v)\n", i)
		i++
	}
}
