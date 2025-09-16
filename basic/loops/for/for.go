package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Iteration:", i)
	}

	for i := range 5 {
		fmt.Println("Range iteration:", i)
	}

	x := 0
	for x < 5 {
		fmt.Println("While-like loop iteration:", x)
		x++
	}
}
