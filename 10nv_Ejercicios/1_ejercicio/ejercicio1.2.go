package main

// como usamos bufer no necesitamos la go rutine.
import (
	"fmt"
)

func main() {
	c := make(chan int, 2)

	c <- 42
	c <- 50

	fmt.Println(<-c)
	fmt.Println(<-c)

}
