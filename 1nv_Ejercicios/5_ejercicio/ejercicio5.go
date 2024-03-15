package main

import "fmt"

/*
1.	A nivel de paquete usando la palabra clave “var”, crear una VARIABLE con el IDENTIFICADOR “y”. La variable debería ser del mismo TIPO SUBYACENTE que tu TIPO “x” creado anteriormente
a.	ejemplo:
	type hotdog int
	var x hotdpg
	var y int

2.	en func main
a.	Esto lo debería tener listo
i.	Imprimir el valor de la variable “x”
ii.	Imprimir el tipo de la variable “x”
iii.	Asigna un VALOR a la VARIABLE “x” usando el OPERADOR “=”
iv.	Imprime el valor de la variable “x”
b.	Ahora haces esto
i.	Ahora usa CONVERSIÓN para convertir el TIPO del VALOR almacenado en “x” al TIPO IMPLÍCITO
1.	Usa el operador “=” para ASIGNAR ese valor a “y”
2.	Imprime el valor almacenado en “y”
3.	Imprime el tipo de “y”

*/

type perro int

var x perro
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("La variable x es del tipo %T\n", x)
	x = 5677
	fmt.Printf("x= %v\n", x)
	y = int(x)

	fmt.Println("Y= ", y)
	fmt.Printf("La variable y es del tipo %T\n", y)
	fmt.Printf("y= %v", y)
}
