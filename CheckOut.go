package main

import "fmt"

func main() {
	//declaring and initializing a variable in go
	var letters = "It is lovely here!"
	//declaring a variable in go
	var letter string
	letter = letters
	fmt.Println(letter)
	//another way is to let the compiler infer the type of the variable
	name := "Laptop"
	fmt.Println(name)
	//another way is using the const keyword
	//const play = "she is out there playing"
	//play = letter //throws an error
	//ps: const value cannot be changed like reassigned
	//fmt.Println(play)
	//const values are declared immediately
	var (
		a = 5
		b = 6
		c = 8
	)
	fmt.Println(a + b + c)

	//take user input from the console
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	var output = input + 2
	fmt.Println("output is", output)
}
