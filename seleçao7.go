package main

import "fmt"

func main() {
	var salario float64
	fmt.Print("Qual o seu salario?")
	fmt.Scan(&salario)

	if salario >= 1000 {
		fmt.Println("Este e o seu salario com aumento de 5%", salario+salario*0.05)
	} else {
		fmt.Println("Este e o seu salario com aumento de 10%", salario+salario*0.1)

	}

}
