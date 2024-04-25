package main

/*Unmarshal

unmarshal el JSON a un estructura de datos de Go.

*/

import (
	"encoding/json"
	"fmt"
)

type amigos struct {
	Nombre   string
	Apellido string
	Edad     int
	Dichos   []string
}

//func (a amigos) String() string {
//	return fmt.Sprintf("\nNombre: %s- Apellido: %s- Edad: %d- Dichos: %v", a.Nombre, a.Apellido, a.Edad, a.Dichos)
//}

func main() {
	s := `[{"Nombre":"Eduar","Apellido":"Tua","Edad":32,"Dichos":["Cachicamo diciéndole a morrocoy conchudo","La mona, aunque se vista de seda, mona se queda","Poco a poco se anda lejos"]},{"Nombre":"Carlos","Apellido":"Pérez","Edad":27,"Dichos":["Camarón que se duerme se lo lleva la corriente","A ponerse las alpargatas que lo que viene es joropo","No gastes pólvora en zamuro"]},{"Nombre":"Mario","Apellido":"Hmmmm","Edad":54,"Dichos":["Ni lava ni presta la batea","Hijo de gato, caza ratón","Más vale pájaro en mano que cien volando"]}]`
	//fmt.Println(s)

	s_tobytes := []byte(s)

	var amigos []amigos

	err := json.Unmarshal(s_tobytes, &amigos)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("JSON obtenido:", amigos)

	// importante recordar que acá se está recorriendo el objeto de tipo amigos

	for i, amigos := range amigos {
		fmt.Printf("Amigo# %v --> \nNombre: %s- Apellido: %s- Edad: %d ", i+1, amigos.Nombre, amigos.Apellido, amigos.Edad)
		for x, dicho := range amigos.Dichos {
			fmt.Println("Dicho ", x, " ", dicho)
		}

	}

}
