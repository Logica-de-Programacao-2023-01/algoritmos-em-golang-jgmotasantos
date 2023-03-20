package main

import "fmt"

func main() {
	var numeroDias int
	fmt.Print("Quantos dias o funcionario trabalhou?")
	fmt.Scan(&numeroDias)
	var valorDiaria float64
	fmt.Print("Qual o valo da diaria do funcionario")
	fmt.Scan(&valorDiaria)

	salario := float64(numeroDias) * valorDiaria

}
