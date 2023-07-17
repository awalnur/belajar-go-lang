package main

import "fmt"

func main() {

	//for counts := 1; counts <= 10; counts++ {
	//	fmt.Println(counts)
	//}

	per := make(map[string]string)

	per["test"] = "tester"
	per["test2"] = "tester2"
	per["test3"] = ""

	for _, value := range per {
		fmt.Println(value)
	}

	//for i := 0; i < 10; i++ {
	//
	//	if i%2 == 0 {
	//		continue
	//	}
	//	fmt.Println(i)
	//}

}
