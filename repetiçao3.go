package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			continue

		}

		fmt.Println(i)
	}

	fmt.Println("Estes sao seu numeros impares de 0 a 19")

}
