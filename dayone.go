package main

import "fmt"

func FirstFactorial(num int) int {
	var factorial = 1

	for num > 1 {
		factorial = factorial * num
		num--
	}

	return factorial
}

func main() {
	num := 4
	fmt.Println(FirstFactorial(num))
}
