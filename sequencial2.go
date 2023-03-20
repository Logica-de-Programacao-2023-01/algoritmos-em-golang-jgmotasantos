package main

import "fmt"

func main() {
	var x int
	fmt.Print("Qual o valor do seu numero")
	fmt.Scan(&x)
	fmt.Print("O dobro do seu numero e")
	fmt.Println(x * 2)
	fmt.Print("O triplo do seu numero e")
	fmt.Println(x * 3)
	fmt.Print("O quadruplo do seu numero e")
	fmt.Println(x * 4)

}
