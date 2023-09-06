package main

import "fmt"

func main() {
	//number := 1
	//for number <= 10 {
	//	fmt.Print(number, " ")
	//	number += 1
	//}
	//for number := 0; number < 10; number++ {
	//	fmt.Print(number)
	//}
	//for number := 0; number < 10; number++ {
	//	if number%2 == 0 {
	//		fmt.Println(number, "even")
	//	} else {
	//		fmt.Println(number, "odd")
	//	}
	//}
	//creating a switch statement
	//var number = 0
	//for ; number <= 10; number++ {
	//	switch number {
	//	case 0:
	//		fmt.Println(number, "zero")
	//	case 1:
	//		fmt.Println(number, "One")
	//	case 2:
	//		fmt.Println(number, "Two")
	//	case 3:
	//		fmt.Println(number, "Three")
	//	default:
	//		fmt.Println(number, "unknown")
	//
	//	}
	//}
	//i := 10
	//if i > 10 {
	//	fmt.Println("big")
	//} else {
	//	fmt.Println("small")
	//}
	//numbers divisible by 3
	//for index := 1; index <= 100; index++ {
	//	if index%3 == 0 {
	//		fmt.Print(index, " ")
	//	}
	//}
	//printing fizzBuzz
	for index := 0 + 1; index <= 100; index++ {
		if index%3 == 0 && index%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if index%5 == 0 {
			fmt.Println("Buzz")
		} else if index%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(index)
		}
	}
}
