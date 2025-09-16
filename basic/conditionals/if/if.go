package main

import "fmt"

func main() {
	var number int

	fmt.Println("Type a number:")
	fmt.Scan(&number)
	if number > 10 {
		fmt.Println("Number is too high")
	} else if number < 5 {
		fmt.Println("Number is too low")
	} else {
		fmt.Println("Number is acceptable")
	}
}
