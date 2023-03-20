package main

import "fmt"

func main() {
	var valorProduto float64
	fmt.Print("Qual o valor do produto?")
	fmt.Scan(&valorProduto)

	desconto := valorProduto * 0.1
	resultado := valorProduto - desconto
	fmt.Println("O valo com desconto é de", resultado)

}
