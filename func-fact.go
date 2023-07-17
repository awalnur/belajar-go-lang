package main

import "fmt"

func faktorial(number int) int {
	result := 1

	for i := number; i > 0; i-- {
		result *= i
	}
	return result
}
func faktorialRec(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * faktorialRec(value-1)
	}
}
func main() {
	fmt.Println(faktorial(5))

	fmt.Println(faktorialRec(5))
}
