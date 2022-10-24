package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go PrintA()
	go PrintB()

	wg.Wait()
}

func PrintA() {
	for i := 0; i < 5; i++ {
		fmt.Println("Função A ", i)
	}
	wg.Done()
}

func PrintB() {
	for i := 0; i < 5; i++ {
		fmt.Println("Função B ", i)
	}
	wg.Done()
}

/*
- Alem da goroutine principal, crie duas outras goroutines.
- Cada goroutine adicional devem fazer um print separado.
- Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.
*/
