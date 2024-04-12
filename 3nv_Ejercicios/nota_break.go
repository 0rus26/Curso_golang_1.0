package main

import "fmt"

func main() {

	/*i := 0
	for {
		i++
		if i > 100 {
			break
		}
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}


	for i := 32; i <= 122; i++ {
		fmt.Printf("%d-->%#X\t%#U\n", i, i, i)
	}
	*/

	/*
		   Sintaxis para switch case

		switch {
		case 2 == 4:
			fmt.Println("No debería imprimir")
		case 3 == 3:
			fmt.Println("Si debería imprimir")
			fallthrough
		case 4 == 5:
			fmt.Println("No debería imprimir 4 no es igual a 5")
		default:
			fmt.Println("No debería imprimir")
		}
	*/
	/*
		for i := 0; i <= 100; i++ {
			if i%2 == 0 {
				fmt.Println(i)
			}

		}
	*/

	switch "Pera" {

	case "Pera", "Limón":
		fmt.Println("frutas verdes")
	case "Fresa", "Mora":
		fmt.Println("frutas rojas")
	default:
		fmt.Println("caso default")
	}
}
