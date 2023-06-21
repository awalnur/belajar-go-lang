package main

import "fmt"

func main() {
	var names [10]string
	var sb = [...]string{}
	names[0] = "adb"
	names[1] = "adb2"
	names[2] = "adb3"
	names[3] = "adb4"
	names[4] = "adb4"

	fmt.Println(names[2])

	fmt.Println(len(names))
	fmt.Println(len(sb))
}
