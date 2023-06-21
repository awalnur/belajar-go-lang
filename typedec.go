package main

import "fmt"

func main() {
	type NoKtp string

	var ktp NoKtp = "11209129812"

	fmt.Println(ktp)
}
