package main

import (
	"fmt"
)

type User struct {
	name, email string
	age         int
}

func (uss User) hello() {
	fmt.Println("hello ", uss.name)
}
func main() {
	var john User

	john.name = "John"
	john.email = "joh@mail"
	john.age = 15

	data_user := User{"johns", "j@fl", 12}
	data_user.hello()
	john.hello()

}
