package main

import "fmt"

func main() {
	//Faça um algoritmo que imprima os múltiplos de 3 de 0 a 30.
	for i := 0; i < 31; i++ {
		if i%3 != 0 {
			continue

		}

		fmt.Println(i)

	}

	fmt.Println("Os multiplos de 3 ate 30 sao: ")

}
