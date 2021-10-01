package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1

	//nao tem while no go

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 { //incrementando de 2 em 2
		fmt.Println(i)
	}

	//misturando

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	//laco infinito

	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second * 5)
	}
}
