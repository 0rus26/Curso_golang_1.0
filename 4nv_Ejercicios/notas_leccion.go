package main

/**/
func main() {
	/*
		var x [5]int
		x[3] = 42
		fmt.Println(x)
		fmt.Printf("%T\n", x)
		fmt.Println(len(x))
	*/
	//literal compuesto
	/*
		x := []int{2, 3, 4, 42}
		fmt.Println(x)

		for i, v := range x {
			fmt.Println("id: ", i, " Valor: ", v)
		}
		fmt.Println("\n")
		fmt.Println(x[1:3])
	*/
	/*
		x := []int{2, 3, 4, 42}
		for i := 0; i <= len(x)-1; i++ {
			fmt.Println("id: ", i, " Valor: ", x[i])
		}
	*/
	/*x := []int{2, 3, 4, 42}
	//y := 1
	fmt.Println("X:", x)
	y := []int{77, 2, 1, 10}
	//x = append(x, y, 1, 2, 3)
	//como agregar un slice a otro slice -->
	x = append(x, y...)
	z := append(x[:2], x[4:]...)
	fmt.Println("x+y:", x)
	fmt.Println("Z:", z)
	*/
	/*x := make([]string,4,1)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Print ln(cap(x))
	*/
	/*
		miMapa := make(map[string]int)

		// Agregar elementos al mapa
		miMapa["clave1"] = 1
		miMapa["clave2"] = 2
		miMapa["clave3"] = 3

		// Imprimir el mapa
		fmt.Println("Mapa:", miMapa)

		// Acceder a un valor en el mapa
		fmt.Println("Valor de la clave 'clave2':", miMapa["clave2"])

		// Modificar un valor en el mapa
		miMapa["clave2"] = 20

		// Imprimir el mapa actualizado
		fmt.Println("Mapa actualizado:", miMapa)

		// Eliminar un elemento del mapa
		delete(miMapa, "clave3")

		// Imprimir el mapa después de eliminar un elemento
		fmt.Println("Mapa después de eliminar 'clave3':", miMapa)

	*/
	/* map puede tener slices o arrays
	agenda := map[string][]string{
		"Juan":  {"123456", "789012"},
		"María": {"456789"},
		"Pedro": {"987654", "345678", "901234"},
	}
	// Imprimir la agenda
	fmt.Println("Agenda:", agenda)

	// Acceder a los números de teléfono de Juan
	numerosJuan := agenda ["Juan"]
	fmt.Println("Números de teléfono de Juan:", numerosJuan)

	// Agregar un nuevo número de teléfono a María
	agenda["María"] = append(agenda["María"], "012345")

	// Imprimir la agenda actualizada
	fmt.Println("Agenda actualizada:", agenda)
	*/

	/*
		slice_a := []string{"Eduard", "Tua", "Crossfit", "Baseball", "Montañismo"}
		slice_b := []string{"Manolo", "Vargas", "Correr", "Natación", "jugar play"}

		fmt.Printf("Tipo %T, Valor: %v \n", slice_a, slice_a)
		fmt.Printf("Tipo %T, Valor: %v \n", slice_b, slice_b)
		nuevo_slice := [][]string{slice_a, slice_b}
		fmt.Println("\nNuevo slice: ", nuevo_slice)
	*/
	/*
		mapa := map[string]int{
			"Batman": 32,
			"Robin":  27,
		}

		fmt.Println(mapa)
		fmt.Println("Edad:", mapa["Batman"])
		fmt.Println("Edad:", mapa["XX"])

		valor, ok := mapa["XX"]

		fmt.Println("Valor: ", valor, "Estado: ", ok)
	*/
	/*
		mapa := map[string]int{
			"Batman": 32,
			"Robin":  27,
		}

		if valor, ok := mapa["Robin"]; ok {
			fmt.Println("Print desde if:", valor)
		}
	*/

	// Como agregar elementos a un mapa despues de haberlo creado:
	/*
		mapa := map[string]int{
			"Batman": 32,
			"Robin":  27,
		}

		mapa["Spiderman"] = 15

			for llave_superhero, valor := range mapa {
				fmt.Println(llave_superhero, valor)
			}


		mapa["Wolverine"] = 50

		delete(mapa, "Spiderman")
		fmt.Print("mapa", mapa, "\n")
		super_h := "Robin"

		if v, ok := mapa[super_h]; ok {
			fmt.Println("borrado efectivo de la llave", v, "Con valor", mapa[super_h])
			delete(mapa, super_h)
		}
	*/

}
