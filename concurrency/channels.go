package main

import "fmt"

func main() {
	// canal sin buffer
	//c := make(chan int) // crear un canal
	//c <- 1              // enviar valor al canal, se bloquea esperando que alguien reciba ese mensaje
	//fmt.Println(<-c)    // nunca llega a este punto

	c := make(chan int, 3) // canal con buffer
	c <- 1
	c <- 2
	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
