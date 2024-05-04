//paquete CONTEXT, para manejar request

package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

// funcion para generar contexto

func gen(ctx context.Context) <-chan int {
	worker := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case worker <- n:
				n++
			}

		}
	}()
	return worker
}

/*

//otro ejercicio de FAN OUT, pero esta vez con definimos el número de goutines para operar.
// un poco diferente al visto en el curso.
// en este ejercicio cada gorutine ejecuta una funcion anónima, la cuál captura el valor actual de valor en el momento que se crea,
// x lo tanto cada GR tiene su propia copia del valor2
//esta forma es mas segura, para el manejo de los datos, pero usa mas memoria al crear funciones anónimas..

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Add_Chs(ch1)
	//leerCanal(ch1)
	fmt.Println("1.#Gorutines:", runtime.NumGoroutine())
	go fanOutIn(ch1, ch2)
	fmt.Println("2.#Gorutines:", runtime.NumGoroutine())

	for value := range ch2 {
		fmt.Println("leyendo canal2:", value)
		fmt.Println("3.#Gorutines:", runtime.NumGoroutine())
	}
	//close(ch2)
	fmt.Println("4.#Gorutines:", runtime.NumGoroutine())
	fmt.Println("END")
}

func Add_Chs(c chan int) {
	for i := 0; i <= 20; i += 5 {
		c <- 1
	}
	close(c)
}

func leerCanal(c chan int) {
	for v := range c {
		fmt.Println("leyendo canal", v)
	}

}

// entradas 2 canales
func fanOutIn(x1, x2 chan int) {
	var wg sync.WaitGroup
	num_gr := 2
	wg.Add(num_gr)

	for valor := 0; valor < num_gr; valor++ {
		go func() {
			for i := range x1 {
				//cada gorutine tiene su propio valor leído del canal x1 ch1
				func(valor2 int) {
					x2 <- tarea_Xtiempo(valor2)
				}(i)

			}
			wg.Done()
		}()
	}

	wg.Wait()
	close(x2)

}

func tarea_Xtiempo(v int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return (v + rand.Intn(1000))
}


*/
/*
//otro ejercicio de FAN OUT, pero esta vez con definimos el número de goutines para operar.
// un poco diferente al visto en el curso.
// en este ejercicio cada gorutine tiene acceso a la variable valor, puede causar problemas..
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Add_Chs(ch1)
	//leerCanal(ch1)
	fmt.Println("1.#Gorutines:", runtime.NumGoroutine())
	go fanOutIn(ch1, ch2)
	fmt.Println("2.#Gorutines:", runtime.NumGoroutine())

	for value := range ch2 {
		fmt.Println("leyendo canal2:", value)
		fmt.Println("3.#Gorutines:", runtime.NumGoroutine())
	}
	//close(ch2)
	fmt.Println("4.#Gorutines:", runtime.NumGoroutine())
	fmt.Println("END")
}

func Add_Chs(c chan int) {
	for i := 0; i <= 20; i += 5 {
		c <- 1
	}
	close(c)
}

func leerCanal(c chan int) {
	for v := range c {
		fmt.Println("leyendo canal", v)
	}

}

// entradas 2 canales
func fanOutIn(x1, x2 chan int) {
	var wg sync.WaitGroup
	num_gr := 2
	wg.Add(num_gr)
	//worker := make(chan int)
	for valor := 0; valor < num_gr; valor++ {
		go func() {
			for i := range x1 {
				x2 <- tarea_Xtiempo(i)
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(x2)
	}()

}

func tarea_Xtiempo(v int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return (v + rand.Intn(1000))
}

/*
// FAN OUT
// En este ejemplo, ejecutamos go routines según el valor almacenado en canal x. El resultado de procesar la info del canal X, se guarda en otro canal
//procesamiento de tablas??
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Add_Chs(ch1)
	//leerCanal(ch1)
	fmt.Println("1.#Gorutines:", runtime.NumGoroutine())
	go fanOutIn(ch1, ch2)
	fmt.Println("2.#Gorutines:", runtime.NumGoroutine())

	for value := range ch2 {
		fmt.Println("leyendo canal2:", value)
		fmt.Println("3.#Gorutines:", runtime.NumGoroutine())
	}
	fmt.Println("4.#Gorutines:", runtime.NumGoroutine())
	fmt.Println("END")
}

func Add_Chs(c chan int) {
	for i := 0; i <= 15; i += 5 {
		c <- 1
	}
	close(c)
}

func leerCanal(c chan int) {
	for v := range c {
		fmt.Println("leyendo canal", v)
	}

}

// entradas 2 canales
func fanOutIn(x1, x2 chan int) {
	var wg sync.WaitGroup
	for valor := range x1 {
		wg.Add(1)
		go func(in int) {
			x2 <- tarea_Xtiempo(valor)
			wg.Done()
		}(valor)
	}
	wg.Wait()
	close(x2)
}

func tarea_Xtiempo(v int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return (v + rand.Intn(1000))
}

*/

/*
// Este codigo lo que hace es unificar la salida de multiples canales a uno sólo fan in !!,
//el canal fan in estara leyendo durante x iteraciones. los otros canales siempre están enviando data.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := fanIn(boring("Joe"), boring("ANa"))
	for i := 0; i < 100; i++ {
		fmt.Println(<-canal)
	}
	fmt.Println("Me voy")
}

func boring(mensaje string) <-chan string {
	canal := make(chan string)
	go func() {
		for i := 0; ; i++ {
			canal <- fmt.Sprintf("%s %d", mensaje, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

		}
	}()
	return canal
}

func fanIn(entrada1, entrada2 <-chan string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- <-entrada1
		}
	}()
	go func() {
		for {
			canal <- <-entrada2
		}
	}()
	return canal

}

/*
// llevar de multicanales a un sólo canal
package main

import (
	"fmt"
	"sync"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	fanin := make(chan int)

	go enviar(par, impar)
	go recibir(par, impar, fanin)
	for valor := range fanin {
		fmt.Printf("Valor: %v\n", valor)
	}
	fmt.Println("END")
}

// función para enviar al canal valores pares e impares
func enviar(p, i chan<- int) {
	for y := 0; y < 100; y++ {
		if y%2 == 0 {
			p <- y
		} else {
			i <- y
		}

	}
	close(p)
	close(i)

}

// vamos a recibir del canal
func recibir(p, i <-chan int, f chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	//range bloquea el canal si este no se cierra
	go func() {
		for x := range p {
			f <- x
		}
		wg.Done()
	}()

	go func() {
		for x := range i {
			f <- x
		}
		wg.Done()
	}()

	wg.Wait()
	close(f)
}


*/

/*
// con el idioma ok, verificamos que un canal tiene valor para ser leído
package main

import "fmt"

func main() {
	canal := make(chan int)

	go func() {
		canal <- 100
		canal <- 101
		close(canal)

	}()

	valor, ok := <-canal
	fmt.Println(valor, ok)

	valor, ok = <-canal
	fmt.Println(valor, ok)

	valor, ok = <-canal
	fmt.Println(valor, ok)
}

*/

/*
//Idioma coma OK

package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan int)

	go enviar(par, impar, salir)
	recibir(par, impar, salir)
	fmt.Println("END")
}

func enviar(p, i, s chan<- int) {
	for x := 0; x < 100; x++ {
		//verficar que es par

		if x%2 == 0 {
			p <- x
		} else {
			i <- x
		}
	}
	close(s)

}

// de esta manera se define que vamos p recibir desde el canal
func recibir(p, i, s <-chan int) {
	for {
		select {
		case valor := <-p:
			fmt.Println("canal par:", valor)
		case valor := <-i:
			fmt.Println("canal impar:", valor)
		case x, ok := <-s:
			if !ok {
				fmt.Println("DESDE OK", x, ok)
				return
			} else {
				fmt.Println("DESDE OK", x)
			}

		}
	}

}


*/
/*

Como usar canales para obtener numerps pares e impares
package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan bool)

	go enviar(par, impar, salir)
	recibir(par, impar, salir)
	fmt.Println("END")
}

func enviar(p, i chan<- int, s chan<- bool) {
	for x := 0; x < 100; x++ {
		//verficar que es par

		if x%2 == 0 {
			p <- x
		} else {
			i <- x
		}
	}
	close(p)
	close(i)
	s <- false

}

// de esta manera se define que vamos a recibir desde loa canales
func recibir(p, i <-chan int, s <-chan bool) {
	for {
		select {
		case valor := <-p:
			fmt.Println("canal par:", valor)
		case valor := <-i:
			fmt.Println("canal impar:", valor)
		case x, ok := <-s:
			fmt.Println("canal impar:", valor)
			return
		}
	}

}

*/

/*

package main

import "fmt"

func main() {
    par := make(chan int)
    impar := make(chan int)
    salir := make(chan int)

    go enviar(par, impar, salir)
    recibir(par, impar, salir)
    fmt.Println("END")
}

func enviar(p, i, s chan<- int) {
    for x := 0; x < 100; x++ {
        if x%2 == 0 { // Corregir la condición para verificar si x es par
            p <- x
        } else {
            i <- x
        }
    }
    close(p)
    close(i)
    s <- 0
}

func recibir(p, i, s <-chan int) {
    for {
        select {
        case valor := <-p:
            fmt.Println("canal par:", valor)
        case valor := <-i:
            fmt.Println("canal impar:", valor)
        case valor := <-s:
            fmt.Println("canal salir:", valor)
            return
        }
    }
}

*/

/*

leer un canal usando un range
package main

import "fmt"

func main() {
	s := make(chan int)
	//enviar hacia el canal
	go func() {
		for i := 0; i <= 10; i++ {
			s <- i
		}
		close(s)
	}()
	//recibir desde el canal
	for v := range s {
		fmt.Println(v)
	}
	fmt.Println("END")

}

*/

// USO de CANALES
/*
package main

import "fmt"

func main() {
	s := make(chan int)
	go enviar(s)
	recibir(s)

	fmt.Println("END")
}


//Enviar al canal
func enviar(s chan<- int) {
	s <- 42
}

//recibir o pedir del canal

func recibir(s <-chan int) {
	fmt.Println(<-s)
}
*/

/*
//CANALES DIRECCIONALES

package main

import "fmt"

func main() {
	//canal con capacidad o bufer.
	canal_send := make(chan<- int)    // solo para enviar
	canal_receive := make(<-chan int) // solo para recibir

	go func() {
		canal <- 42
	}()
	//acá está pidiendo que envie datos desde el canal hacia
	//desde el print, y el canal es solo para recibir
	//fmt.Println(<-canal_receptor)  ERROR
	fmt.Println(<-canal_emisor)
}

*/

/*

package main

import "fmt"

func main() {
	//canal con capacidad o bufer.
	canal := make(chan int)
	canal_recibir := make(<-chan int)
	canal_enviar := make(chan<- int)
	fmt.Println("<----->")
	//vamos p tratar de mandar un dato más al canal
	fmt.Printf("C1: %T \n", canal)
	fmt.Printf("C2: %T \n", canal_recibir)
	fmt.Printf("C3: %T \n", canal_enviar)
	//fmt.Println(<-canal2)
}

*/

/*
package main

import "fmt"

func main() {
	//canal con capacidad o bufer.
	canal2 := make(chan int, 3)
	canal2 <- 45
	//vamos p tratar de mandar un dato más al canal
	canal2 <- 50
	fmt.Println(<-canal2)
	//fmt.Println(<-canal2)
}
*/

/*
package main

import (
	"fmt"
	"runtime"
)

func main() {

	//canal sin bufer
	//en este código tenemos 2 go rutines el main y la go func, se manda al canal desde la go func y se recibe desde el main
	fmt.Println("1. Número de Gorutinas: ", runtime.NumGoroutine())
	canal1 := make(chan int)
	go func() {
		canal1 <- 42
	}()
	fmt.Println("2. Número de Gorutinas: ", runtime.NumGoroutine())
	fmt.Println("Lectura de canal1:", <-canal1)
	fmt.Println("3. Número de Gorutinas: ", runtime.NumGoroutine())
}

*/
