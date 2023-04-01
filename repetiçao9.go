package main

import "fmt"

func main() {
	var soma, count, numeros int
	fmt.Println("Digite uma sequencia de numeros inteiros")
	fmt.Println("Caso queira finalizar a sequencia digite 0")
	for {
		fmt.Scan(&numeros)
		if numeros == 0 {
			break

		}
		count++
		soma += numeros

	}
	if count > 0 {
		media := soma / count
		fmt.Println("A media dos numeros digitados eh:", media)

	} else {
		println("Um numero invalido foi digitado!")
	}

}
