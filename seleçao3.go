package main

import "fmt"

func main() {
	var num int
	fmt.Print("Qual o valor do seu numero inteiro?")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println("Seu numero é um numero par")
	} else {
		fmt.Println("Seu numero é impar")
	}

}
