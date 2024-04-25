package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "Clave1234"
	password_Login := "Clave123455"
	bs, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	err_compare := bcrypt.CompareHashAndPassword(bs, []byte(password_Login))

	// para evitar la parte del else, se puede hacer un return despues de la condición if se cumpla, así no imprimir doble
	/*
		if err_compare != nil {
			fmt.Println("Claves errada")
		} else {
			fmt.Println("Claves iguales")
		}
	*/

	if err_compare != nil {
		fmt.Println("Claves errada")
		return
	}
	fmt.Println("Claves iguales")

}

/*

//Funcionalidad muy util para poder organizar información de algun tipo, en este caso la data de tipo Persona

package main

import (
	"fmt"
	"sort"
)

type Persona struct {
	Nombre string
	Edad   int
}

//función para imprimir un objeto de un tipo de dato, como persona, cada que se llame a esa persona, se va imprimir según el siguiente formato

func (p Persona) String() string {
	return fmt.Sprintf("\nNombre: %s- Edad: %d", p.Nombre, p.Edad)
}

type Por_edad []Persona

func (e Por_edad) Len() int           { return len(e) }
func (e Por_edad) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e Por_edad) Less(i, j int) bool { return e[i].Edad < e[j].Edad }

type Por_Nombre []Persona

func (e Por_Nombre) Less(i, j int) bool { return e[i].Nombre < e[j].Nombre }
func (e Por_Nombre) Len() int           { return len(e) }
func (e Por_Nombre) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

//func(e Por_edad) Max() int {return len(e)}

func main() {
	p1 := Persona{"Otoniel", 50}
	p2 := Persona{"Bertha", 10}
	p3 := Persona{"Fernando", 36}
	p4 := Persona{"Natalia", 18}

	grupo_personas := []Persona{p1, p2, p3, p4}

	fmt.Println("1. Normal-->", grupo_personas)


	// esta es la otra forma de usar el paquete de sort. usando el método SLice, que no requiere usar de la interfaz Interface
	/*
		fmt.Println("1. Normal-->", grupo_personas)
		sort.Slice(grupo_personas, func(i int, j int) bool {
			return grupo_personas[i].Nombre > grupo_personas[j].Nombre
		})
		fmt.Println("2. Por nombre-->", grupo_personas)
		sort.Slice(grupo_personas, func(i int, j int) bool {
			return grupo_personas[i].Edad < grupo_personas[j].Edad
		})

		data_int, err := fmt.Println("3. Por Edad-->", grupo_personas)

		fmt.Println("Data Int:", data_int, " err: ", err)

	sort.Sort(Por_edad(grupo_personas))
	fmt.Println("3. Por Edad-->", grupo_personas)
	sort.Sort(Por_Nombre(grupo_personas))

	fmt.Println("3. Por Nombre-->", grupo_personas)

}
*/

/* función obtenida de la librerias de fmt,para revisar si un slice de strings está ordenado alfabeticamente
package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func main() {
	names := []string{"Alice", "Bob", "Ana", "VERA"}
	isSortedInsensitive := slices.IsSortedFunc(names, func(a, b string) int {
		return cmp.Compare(strings.ToLower(a), strings.ToLower(b))
	})
	fmt.Println(isSortedInsensitive)
	fmt.Println(slices.IsSorted(names))
}


*/
/*
package main

import (
	"fmt"
	"sort"
)

func main() {
	x_int := []int{1, 5, 9, 1, 2, 3, 8, 90, 10}
	x_string := []string{"Beta", "Alfa", "Omega", "Epsilon", "Theta"}

	fmt.Println(x_int)
	fmt.Println(x_string)
	sort.Ints(x_int)
	sort.Strings(x_string)

	fmt.Println(x_int)
	fmt.Println(x_string)

}

*/

/*
// Interfaz writer
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Hola mundo fmt print normal")
	fmt.Fprintln(os.Stdout, "Hola mundo fmt.Fprint")
	io.WriteString(os.Stdout, "Hola mundo io.WriteString")

}
*/

//DECODER y ENCODER SE usan para trabajar con JSON en la web directamente
// una interfaz es una colección de métodos

/*UnMarshal

package main

import (
	"encoding/json"
	"fmt"
)

type superh struct {
	Id       int    `json:"Id"`
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
}

func main() {

	//el json se debe poner entre comillas raras ``
	salida := `[{"Id":1,"Nombre":"Acuaman","Apellido":"Pantano"},{"Id":2,"Nombre":"Bati","Apellido":"chica"},{"Id":3,"Nombre":"Green","Apellido":"Lantern"}]`
	//el string de JSON se convierte a un slice de bytes

	bs := []byte(salida)
	//fmt.Printf("%T tipo JSON: %v\n", salida, salida)
	//fmt.Printf("\n %T slice de bytes: %v\n", bs, bs)

	var super_heros []superh
	// realiza la conversión de bytes a slice de algun tipo
	err := json.Unmarshal(bs, &super_heros)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Data obtenido del JSON: ", super_heros)

	for _, v := range super_heros {
		fmt.Println("ID:", v.Id, "Nombre: ", v.Nombre, "Apellido:", v.Apellido)
	}
}
*/

/*
 Marshal JSON
import (
	"encoding/json"
	"fmt"
)

// sección de manejo de JSON

//documentación:= https://pkg.go.dev/encoding/json

type superh struct {
	Id       int
	Nombre   string
	Apellido string
}

func main() {

	sh1 := superh{
		Id:       001,
		Nombre:   "Acuaman",
		Apellido: "Pantano",
	}
	sh2 := superh{
		Id:       002,
		Nombre:   "Bati",
		Apellido: "chica",
	}
	sh3 := superh{
		Id:       003,
		Nombre:   "Green",
		Apellido: "Lantern",
	}

	peloton1 := []superh{sh1, sh2, sh3}

	fmt.Println(peloton1)

	//Función para convert a JSON, desde un slice de datos de tipo superhero.
	bs, err := json.Marshal(peloton1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
*/
