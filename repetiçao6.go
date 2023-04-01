package main

import "fmt"

func main() {
	var valor int
	fmt.Println("Digite um valor para ser o multiplicador: ")
	fmt.Scan(&valor)
	for i := 1; i <= 10; i++ {
		fmt.Println(i * valor)
	}
	fmt.Println("Esta e sua sequencia dos numeros de 1 a 10 multiplicados pelo seu valor")
}
