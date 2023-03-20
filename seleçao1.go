package main

import "fmt"

func main() {
	var x int
	fmt.Print("Qual o seu primeiro valor")
	fmt.Scan(&x)
	var y int
	fmt.Print("Qual o seu segundo valor")
	fmt.Scan(&y)
	if x > y {
		fmt.Println("Seu primeiro valor é maior que o segundo")
		if y > x {
			fmt.Println("Seu segundo valor é maior que o primeiro")
		}

	}

}
