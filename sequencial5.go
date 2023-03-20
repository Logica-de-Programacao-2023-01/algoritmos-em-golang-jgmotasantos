package main

import "fmt"

func main() {
	var idade int
	fmt.Print("Qual a sua idade?")
	fmt.Scan(&idade)

	resultado := idade * 365
	fmt.Println("Voce tem esta quantidade de dias de idade:", resultado)

}
