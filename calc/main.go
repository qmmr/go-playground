package main

import "fmt"

func add(n1, n2 int) int {
	return n1 + n2
}

func main() {
	var method string
	var firstNumber int
	var secondNumber int
	var result string

	fmt.Println("What type of calculation would you like to perform?")
	fmt.Scanf("%s", &method)
	fmt.Println("What is the first number?")
	fmt.Scanf("%d", &firstNumber)
	fmt.Println("What is the second number?")
	fmt.Scanf("%d", &secondNumber)

	switch method {
	case "add":
		result = fmt.Sprintf("The result of %d + %d is %d", firstNumber, secondNumber, add(firstNumber, secondNumber))
	case "subtract":
		result = fmt.Sprintf("The result of %d - %d is %d", firstNumber, secondNumber, firstNumber-secondNumber)
	case "multiply":
		result = fmt.Sprintf("The result of %d * %d is %d", firstNumber, secondNumber, firstNumber*secondNumber)
	case "divide":
		result = fmt.Sprintf("The result of %d / %d is %d", firstNumber, secondNumber, firstNumber/secondNumber)
	}

	fmt.Println(result)
}
