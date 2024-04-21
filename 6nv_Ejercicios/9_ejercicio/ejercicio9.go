package main

import "fmt"

/*

Un “callback” es cuando le pasamos una función a otra función como argumento. Para este ejercicio,
●	Pasa una func a otra función como argumento

func sumarimpares(f func(x ...int) int, vi ...int) int {
*/

func main() {
	x := []int{100, 2, 3, 4}
	//foo(foo2, x)

	foo2 := func(x []int) int {
		long := len(x)
		if long == 0 {
			return 0
		} else if long == 1 {
			return x[0]
		} else {
			return x[0] + x[long-1]
		}
	}

	fmt.Println("foo: ", foo(foo2, x))
}

func foo(f func(y []int) int, x []int) int {
	YY := f(x)
	YY++
	return YY
}
