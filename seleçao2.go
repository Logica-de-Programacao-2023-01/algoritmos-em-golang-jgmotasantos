package main

import "fmt"

func main() {
	var x int
	fmt.Print("Qual o seu primeiro numero")
	fmt.Scan(&x)
	var y int
	fmt.Print("Qual o seu segundo numero")
	fmt.Scan(&y)
	var z int
	fmt.Print("Qual o seu erceiro numero")
	fmt.Scan(&z)
	if x < y && x < z {
		fmt.Println("Seu primeiro valor é o menor")
		if y < x && y < z {
			fmt.Println("Seu segundo valor é o menor")
			if z < x && z < y {
				fmt.Println("Seu terceiro valor é o menor")
			}

		}
	}
}






