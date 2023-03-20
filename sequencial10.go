package main

import "fmt"

func main() {
	var peso float64
	fmt.Print("Qual é o seu peso?")
	fmt.Scan(&peso)

	resultado := peso * 2.2046
	fmt.Println("Seu peso em lb é: , resultado")

}
