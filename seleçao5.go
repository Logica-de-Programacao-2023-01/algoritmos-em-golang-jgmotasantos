package main

import (
	"fmt"
)

func main() {
	var numero int
	fmt.Print("Qual o valor do seu numero inteiro?")
	fmt.Scan(&numero)
	if numero%3 == 0 && numero%5 == 0 {
		fmt.Println("Seu numero e divisivel por 3 e por 5 ao mesmo tempo")
	} else {
		fmt.Println("O seu numero nao e divisivel por 3 e por 5 ao mesmo tempo")
	}

}
