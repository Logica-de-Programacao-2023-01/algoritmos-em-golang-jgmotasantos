package main

import "fmt"

func main() {
	var pesoKg float64
	fmt.Print("Qual o seu peso em Kg?")
	fmt.Scan(&pesoKg)

	var altura float64
	fmt.Print("Qual a sua altura em metros?")
	fmt.Scan(&altura)

	quadrado := altura * altura

	IMC := pesoKg * 1 / quadrado
	fmt.Println("O seu IMC e", IMC)
}
