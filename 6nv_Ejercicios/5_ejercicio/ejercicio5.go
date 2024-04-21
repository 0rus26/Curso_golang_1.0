package main

import (
	"fmt"
	"math"
)

/*
●	Crea un tipo CUADRADO
●	Crea un tipo CÍRCULO
●	Adjunta un método a cada uno que calcule y retorne el ÁREA
○	Área de un círculo= π r 2
○	Área de un cuadrado = L * H
●	Crea un tipo FORMA que defina una interfaz como cualquier cosa que tenga el método ÁREA
●	Crea una func INFO el cuál toma un tipo forma e imprime el área de la misma.
●	Crea un valor de tipo cuadrado
●	Crea un valor de tipo círculo
●	Usa la func info para imprimir el área de un cuadrado
●	Usa la func info para imprimir el área de un círculo



*/

type cuadrado struct {
	lado float32
}

type circulo struct {
	radio float32
}

type forma interface {
	area() float32
}

func informacion(f forma) {
	fmt.Println(f.area())
}

func (c cuadrado) area() float32 {
	return c.lado * c.lado
}

func (c circulo) area() float32 {
	return (c.radio * c.radio) * (math.Pi)
}

func main() {
	cuadrado1 := cuadrado{
		lado: 10.2,
	}

	circulo1 := circulo{
		radio: 10.5,
	}

	//fmt.Println(math.Pi)
	//fmt.Printf("%.4f", (math.Pi))

	fmt.Println("Area cuadrado= ", cuadrado1.area())
	fmt.Println("Area circulo= ", circulo1.area())

	informacion(circulo1)
	informacion(cuadrado1)
}
