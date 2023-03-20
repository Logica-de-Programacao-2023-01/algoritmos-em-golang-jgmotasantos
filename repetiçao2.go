package main

import "fmt"

func main() {
	for i := 0; i < 21; i++ {
		if i%2 != 0 {
			continue

		}

		fmt.Println(i)
	}

	fmt.Println("Estes sao seu numeros pares de 0 a 20")

}
