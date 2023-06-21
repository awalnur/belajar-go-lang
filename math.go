package main

import "fmt"

func main() {
	var angka int32 = 2
	var angka2 int32 = 10
	var tst bool = angka < angka2
	var stt bool = angka < angka2
	var res int32

	res = angka + angka2 /// /bagi * kali -kurang
	fmt.Println(res)

	res = res % 2
	fmt.Println(res)

	println(tst && stt) // hasil boolean true false

}
