package main

import "fmt"

func main() {
	var salario float64
	fmt.Print("Qual é o seu salario")
	fmt.Scan(&salario)

	aumento := salario * 0.15
	total := salario + aumento
	fmt.Println("Seu salario com aumento é", total)

}
