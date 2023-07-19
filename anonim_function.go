package main

import (
	"fmt"
	"strings"
)

type Blocked func(string) bool

func RegUser(name string, block Blocked) {
	if block(name) {

		uppername := strings.Replace(name, "s", "STR", 9)
		fmt.Println(uppername, "Blocked")

	} else {

		fmt.Println(name, "Not Blocked")
	}
}

func main() {
	RegUser("str", func(name string) bool {
		uppername := strings.Replace(name, "S", "STR", 1)
		return uppername == "str"
	})
}
