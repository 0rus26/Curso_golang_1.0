package main

import "fmt"

/*
Usando los siguientes operadores, escribe expresiones y asigna sus valores a variables:
a.	==
b.	<=
c.	>=
d.	!=
e.	<
f.	>
Imprime los valores de las variables.

*/

func main() {
	a := (42 == 42)
	b := 10 <= 3
	c := 1 >= 2
	d := 5 != 4
	e := 1 < 6
	f := 5 > 1
	fmt.Printf("%v\t%t\t%t\t%t\t%t\t%t\t", a, b, c, d, e, f)
}
