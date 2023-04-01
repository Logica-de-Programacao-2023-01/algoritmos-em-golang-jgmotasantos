package main

import "fmt"

func main() {
	var numeros, maior int
	fmt.Println("Digite uma sequencia de numeros inteiros")
	fmt.Println("Caso queira finalizar a sequencia digite 0")
	for {
		fmt.Scan(&numeros)
		if numeros == 0 {
			break

		}
		if numeros > maior {
			maior = numeros
		}
		if maior > 0 {
			fmt.Println("O maior numero digitado foi: ", maior)
		}

	}

}
