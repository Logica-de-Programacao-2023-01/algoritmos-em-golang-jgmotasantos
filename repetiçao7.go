package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")

		} else if i%5 == 0 {
			println("Buzz")

		} else if i%3 == 0 {
			println("Fizz")

		} else {
			fmt.Println(i)
		}

	}

}
