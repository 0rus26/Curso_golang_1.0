package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 10, 15, 11, 16, 22, 102, 100, 1}

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
