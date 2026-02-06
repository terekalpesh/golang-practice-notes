package main

// SYNTAX

// func functionName() {
// 	// code
// }

// func functionName(param1 type1, param2 type2) {
// 	// code
// }

// func functionName() returnType {
// 	return value
// }

// func functionName(params types) (returnType1, returnType2) {
// 	return  value1, value2
// }

// Named return values
// func divide(a, b float64) (result float64, err error) {
// 	if b == 0 {
// 		return 0, error.New("division by zero")
// 	}
// 	result a / b
// 	return
// }

// Variadic function (variable arguments)
// func sum(numbers ...int) int {
// 	total := 0
// 	for _, n := range numbers {
// 		total += n
// 	}
// 	return total
// }

// Function as values
// var myFunc func(int, int) int
// myFunc = add

import (
	"errors"
	"fmt"
)

// 1. Basic function
func greet() {
	fmt.Println("Hello, from planet Earth.")
}

// 2. Function with parameter
func greetPerson(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// 3. Function with return values
func add(a, b int) int {
	return a + b
}

// 4. Multiple parameters (same type can be shortened)
// func multiply(a, b int) int {
// 	return a * b
// }

// 5. Multiple return values (Go's special feature!)
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// 6. Named return values
func calculate(x, y int) (sum int, product int) {
	sum = x + y
	product = x * y
	return	// Naked return - uses named values
}

// 7. Variadic function
func sum(numbers ...int) (total int) {
	for _, n := range numbers {
		total += n
	}
	return total
}

// 8. Function as value
var subtract = func(a, b int) int {
	return a - b
}

func main() {
	// Call basic function
	greet()

	// Call with parameter
	greetPerson("Alita")

	// Call with return values
	var result = add(10, 11)
	fmt.Printf("10 + 11 = %d\n", result)

	// Call with multiple returns (must handle both)
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("quotient: ", quotient)
	}

	// Error case
	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Named returns
	s, p := calculate(4, 5)
	fmt.Printf("Sum: %d, Product: %d\n", s, p)

	// Variadic function
	// fmt.Printf("Sum: %d\n", sum(1,2,3,4,5))
	total := sum(1,2,3,4,5)
	fmt.Printf("Sum: %d\n", total)

	// Function as value
	// fmt.Printf("Subtraction: %d", subtract(10, 3))
	diff := subtract(10, 3)
	fmt.Printf("Diff: %d\n", diff)
}
