package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 9, 1, 2, 0} // short hand declaration to create array

	//CHECK CASE
	if len(a) > 15 {
		fmt.Println("Array Length is exceeding than 15")
		return
	}

	frequency_count := make(map[int]int)

	//POPULATING MAP
	for i := 0; i < len(a); i++ {
		value := frequency_count[a[i]]

		frequency_count[a[i]] = value + 1
		


	}

	//ITERATNG OVER MAP
	fmt.Println("Integer Frequency Counts:", frequency_count)
	fmt.Println("Integer Frequency:")

	for i := 0; i < len(a); i++ {

		fmt.Print(" ", frequency_count[a[i]])

	}

}
