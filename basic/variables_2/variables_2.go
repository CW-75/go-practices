package main

import "fmt"

func main() {
	var name string
	fmt.Println("var name state", name)
	name = "John"
	fmt.Println("var name after assign", name)
	age := 4 // Using short variable declaration syntax
	fmt.Println("short variable declaration of age variable: ", age)

	// Uncommenting the following lines will cause a compilation error
	// because 'age' is already declared with a short variable declaration syntax.
	// age = "nombre"
	// fmt.Println("age variable after assign", age)

	// Uncommenting the following lines will cause a compilation error, because
	// the variable 'age' is not initialized before use.
	// In Go, variables must be initialized before they can be used.
	// If you want to declare a variable without initializing it, you can use the
	// short variable declaration syntax with ':=' or declare it with 'var' and assign later.
	//
	// var age int
	// var age;
	// fmt.Println("var age state", age)
	// age = 30
	// fmt.Println("var age after assign", age)
}
