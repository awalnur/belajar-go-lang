package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "namse",
	}

	per := make(map[string]string)

	per["test"] = "test"

	fmt.Println(per["test"])

	fmt.Println(person["name"])
}
