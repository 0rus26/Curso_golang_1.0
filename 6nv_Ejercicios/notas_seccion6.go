package main

/*
func main() {
		foo()
		bar("James.")
		swoo := woo("Penny")
		fmt.Println(swoo)
		x, y := saludar("Eduard", "Snow")
		fmt.Println(x)
		fmt.Println(y)
	}

	func foo() {
		fmt.Println("Hola desde Foo")
	}
	func bar(s string) {
		fmt.Println("Hola desde bar ", s)
	}
	func woo(s string) string {
		return fmt.Sprint("Hola desde Woo, ", s)
	}
	func saludar(n, a string) (string, bool) {
		p := fmt.Sprint(n, " ", a, ` dice "Hola" `)
		q := true
		return p, q
	estoy muy regalado ----->>>>>>>>>>>>>>> no +
}
*/

// Parámetros variados

//func (r receptor) IDENTIFICADOR(parametros){retorno(s)}{código}
// el parámetro x... dentro de una función, debe ir siempre al final de todas las entradas como parámetros

/*


clase de funciones multiples entradas y salidas

func main() {
	parametro := " XXX"
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s, p := foo(parametro, xi...)

	fmt.Println(s, "Promedio= ", p)
}

func foo(s string, x ...int) (string, float32) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println("capacidad []:", cap(x), "len:", len(x))
	suma := 0
	for i := 0; i <= len(x)-1; i++ {
		suma = suma + x[i]
	}
	suma_total := fmt.Sprint("Suma total= ", suma, "parametro", s)
	promedio := float32(suma) / float32(len(x))
	return suma_total, promedio
}
*/

/*
Lección de defer,  permite ejecutar una función desúes de que retorne la función main. por defecto
func main() {
	defer foo()
	woo()
}

func foo() {
	fmt.Println("foo")
}

func woo() {
	fmt.Println("woo")
}

*/

/*
 */

/*
type persona struct {
	nombre   string
	apellido string
}

type agente struct {
	persona
	licencia bool
}

type humano interface {
	presentarse()
}

func main() {
	ag1 := agente{
		persona: persona{
			nombre:   "007",
			apellido: "Bond",
		},
		licencia: true,
	}

	p1 := persona{
		nombre:   "Pepito",
		apellido: "Perez",
	}

	ag1.presentarse()
	p1.presentarse()

	bar(ag1)
	bar(p1)
}

func (agente agente) presentarse() {
	fmt.Println("Se presenta el agente ", agente.nombre, agente.apellido)
}

func (persona persona) presentarse() {
	fmt.Println("Se presenta la persona ", persona.nombre, persona.apellido)
}

func bar(h humano) {

	switch h.(type) {
	case persona:
		fmt.Println("Impresión desde bar--Persona", h.(persona).nombre)
	case agente:
		fmt.Println("Impresión desde bar--Agente", h.(agente).nombre)
	}
	//fmt.Println("Fui pasado a la función bar", h)
}
*/

/*
función anónima sin nombre
func main() {
	func() { fmt.Println("Ejecutando función anónima") }()
	func(name string) { fmt.Println("Hola desde función anónima", name) }("Jorge")
}
*/

/*
funciones con expresiones

func main() {
	f := func(alguien string) {
		fmt.Println("Evaluando función con expresión", alguien)
	}

	f("Fuckencio")
}

*/

/*

como un f(g(x))

// retornar una función desde otra función

func main() {
	//x := foo()
	//fmt.Println(x)
	//xfunc := woo()
	//fmt.Printf("xfunc: %T \n", xfunc)
	//valor := xfunc()
	fmt.Printf("Tipo: %T Valor: %v", woo(), woo()())
}

func foo() string {
	saludo := "Hola mundo"
	return saludo
}

func woo() func() int { return func() int { return 1 } }

*/

/*Como llamar una función desde otra función callbacks
func main() {
	x := []int{1, 2, 3, 4}
	total := sumar(x...)
	fmt.Println("\nTotal: ", total)

	total_pares := sumarpares(sumar, x...)
	fmt.Println("\nTotal_pares es : ", total_pares)

	total_impares := sumarimpares(sumar, x...)
	fmt.Println("\nTotal_impares es : ", total_impares)
}

func sumar(x ...int) int {
	fmt.Printf("%T,%v", x, x)
	total := 0
	for i := 0; i <= len(x)-1; i++ {
		total += x[i]
	}
	return total
}

func sumarpares(f func(x ...int) int, vi ...int) int {
	var y []int
	for _, v := range vi {
		if v%2 == 0 {
			y = append(y, v)
		}

	}
	return sumar(y...)
}

func sumarimpares(f func(x ...int) int, vi ...int) int {
	var y []int
	for _, v := range vi {
		if v%2 == 1 {
			y = append(y, v)
		}

	}
	return sumar(y...)
}
*/

/*Como reducir el scope de una variable dentro del codigo con llaves {}
var x int

func main() {

	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)

	{
		y := 50
		y++
		fmt.Println(y)
	}
	//fmt.Println(y)

}

func foo() {
	fmt.Println(x)
	x++
}
*/

/* manejar una variable dentro del scope de una función
func main() {

	a := incremental()
	b := incremental()
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func incremental() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

*/

/*
func main() {
	n := 35
	f1 := factorial(n)
	f2 := ciclo_f(n)
	fmt.Println("factorial:", f1)
	fmt.Println("Ciclo factorial: ", f2)
}

func factorial(n int) int {

	if n == 0 {
		return 1
	}
	return n * factorial((n - 1))
}

func ciclo_f(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}

	return total
}

*/
