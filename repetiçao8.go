package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&num)
	fmt.Printf("Os divisores de %d são: ", num)
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Printf("d", i)
		}
	}
}
