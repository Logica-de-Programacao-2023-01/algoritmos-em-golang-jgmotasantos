package main

import "fmt"

func main() {
	var x float64
	fmt.Print("Qual o valor do seu primeiro numero real?")
	fmt.Scan(&x)

	var y float64
	fmt.Print("Qual o valor do seu segundo numero real?")
	fmt.Scan(&y)

	var z float64
	fmt.Print("Qual o valor do seu terceiro numero real?")
	fmt.Scan(&z)

	mediaPeso := x*2 + y*3 + z*5
	Resultado := mediaPeso / 3
	fmt.Println("A sua media ponderada com pesos 2, 3 e 5 respectivamente e:", Resultado)

}
