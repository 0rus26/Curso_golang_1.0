package main

import "fmt"

/*

●	Crea un strutc persona
●	Crea una función llamada “cambiame” con un *persona como parámetro
○	Cambia el valor almacenado en la dirección de memoria del *persona
■	Importante
●	Para desreferenciar un struct, usa (*valor).campo
○	p1.nombre
○	(*p1).nombre
●	“Como una excepción, si el tipo de x es un tipo puntero con nombre y (*x).c es una expresión válida de selección denotando un campo (pero no un método), x.c es una forma abreviada de (*x).c”.
○	https://golang.org/ref/spec#Selectors
●	Crea un valor de tipo persona
○	Imprime el valor
●	En func main
○	llama “cambiame”
●	En func main
○	Imprime el valor


*/

//este ejercicio sirve para entender que go otorga un espacio en memoria diferente  al del main() ,para hacer el cambio de p.nombre en el espacio de
// ejecucion de la funcion, // si queremos que p1 sea afectado por la función cambiame() debemos apuntar a la dirección
//en memoria de ese dato de tipo persona.

func main() {

	p1 := persona{
		id:     111,
		nombre: "Superman",
	}

	fmt.Println(p1.nombre)
	cambiame(&p1)
	fmt.Println(p1.nombre)
}

type persona struct {
	id     int
	nombre string
}

func cambiame(p *persona) {
	p.nombre = "Spiderman"
}
