package main

import "fmt"

/*
Crea un programa que use una declaración switch con la expresión de switch especificada como una variable de TIPO string
y el IDENTIFICADOR “deporteFav”.
*/

var deporteFav string = "voleibol"

func main() {
	switch deporteFav {
	case "micro":
		fmt.Println("No es mi favorito")
	case "baloncesto":
		fmt.Println("No me gusta ")
	case "voleibol":
		fmt.Printf("%v es mi deporte favorito ", deporteFav)
	default:
		fmt.Println("No debería imprimir")
	}
}
