package main

import "fmt"

func withTwoReturnValues(num int) (int, int) {
	return num * 2, num * 3
}

// Or named return values with clean return
func withTwoNamedReturnValues(num int) (num1 int, num2 int) {
	num1 = num + 2
	num2 = num + 3
	return
}

func main() {
	fmt.Println("Hello from first.go")
	double, triple := withTwoReturnValues(2)
	plusTwo, plusThree := withTwoNamedReturnValues(2)
	fmt.Println(double, triple)
	fmt.Println(plusTwo, plusThree)
}
