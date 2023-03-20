package main

import "fmt"

func main() {
	var X int
	fmt.Print("Qual o seu primeiro numero?")
	fmt.Scan(&X)
	var Y int
	fmt.Print("Qual o seu segundo numero?")
	fmt.Scan(&Y)
	var Z int
	fmt.Print("Qual seu terceiro numero")
	fmt.Scan(&Z)

	resultado := X + Y + Z
	fmt.Println(resultado)

}
