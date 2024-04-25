package main

/*
ordena el []Usuario por
●	Edad
●	Apellido
También ordena el []string “Dichos” para cada Usuario
●	Imprime todo de una manera agradable


*/

import (
	"fmt"
	"sort"
)

type Usuario struct {
	Nombre   string
	Apellido string
	Edad     int
	Dichos   []string
}

type Ordenar_Usuario_Edad []Usuario

func (e Ordenar_Usuario_Edad) Len() int           { return len(e) }
func (e Ordenar_Usuario_Edad) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e Ordenar_Usuario_Edad) Less(i, j int) bool { return e[i].Edad < e[j].Edad }

type Ordenar_Usuario_Nombre []Usuario

func (e Ordenar_Usuario_Nombre) Less(i, j int) bool { return e[i].Nombre < e[j].Nombre }
func (e Ordenar_Usuario_Nombre) Len() int           { return len(e) }
func (e Ordenar_Usuario_Nombre) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

type Ordenar_Usuario_Dicho []Usuario

func (e Ordenar_Usuario_Dicho) Less(i, j int) bool { return e[i].Dichos[i] < e[j].Dichos[j] }
func (e Ordenar_Usuario_Dicho) Len() int           { return len(e) }
func (e Ordenar_Usuario_Dicho) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

func (u Usuario) String() string {
	return fmt.Sprintf("N:%s ,A: %s, E: %d, D: %v \n", u.Nombre, u.Apellido, u.Edad, u.Dichos)
}

func main() {
	u1 := Usuario{
		Nombre:   "Eduar",
		Apellido: "Tua",
		Edad:     32,
		Dichos: []string{
			"Cachicamo diciéndole a morrocoy conchudo",
			"La mona, aunque se vista de seda, mona se queda",
			"Poco a poco se anda lejos",
		},
	}

	u2 := Usuario{
		Nombre:   "Carlos",
		Apellido: "Pérez",
		Edad:     27,
		Dichos: []string{
			"Camarón que se duerme se lo lleva la corriente",
			"A ponerse las alpargatas que lo que viene es joropo",
			"No gastes pólvora en zamuro",
		},
	}

	u3 := Usuario{
		Nombre:   "Che",
		Apellido: "López",
		Edad:     54,
		Dichos: []string{
			"Ni lava ni presta la batea",
			"Hijo de gato, caza ratón",
			"Más vale pájaro en mano que cien volando",
		},
	}

	usuarios := []Usuario{u1, u2, u3}

	fmt.Println("1. Usuarios sin ordenar:", usuarios)

	// Forma de ordenar usando el método Slice de la libreria sort

	/*
		sort.Slice(usuarios, func(i, j int) bool {
			return usuarios[i].Edad > usuarios[j].Edad
		})

		fmt.Println("2. Usuarios ordenados por Edad:", usuarios)

		sort.Slice(usuarios, func(i, j int) bool {
			return usuarios[i].Apellido > usuarios[j].Apellido
		})
		fmt.Println("3. Usuarios ordenados por Apellido:", usuarios)
	*/

	// forma usando la interfaz de Interface(), para esto asociamos la structura del tipo de dato Usuario..

	//sort.Sort(Ordenar_Usuario_Edad(usuarios))
	//fmt.Println("2. Por Edad-->", usuarios)
	//sort.Sort(Ordenar_Usuario_Nombre(usuarios))
	//fmt.Println("3. Por Nombre-->", usuarios)
	for _, v := range usuarios {
		sort.Strings(v.Dichos)
	}

	fmt.Println("4. Por Dichos-->", usuarios)
}
