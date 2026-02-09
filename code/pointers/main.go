package main

// POINTERS

/*

Key concepts:

& (Address Operator) = "Give me the address" (like "What's your home address?")
* (Dereference Operator) = "Go to that address and get the value" (like "Go to 123 Main St and get the house")
*p = 100 (Modify through pointer) = "Go to that address and change the value" (like "Go to that house and paint it red")

*/

// SYNTAX

// Basic Pointer Operations
// var x int = 42
// var p *int = &x	// p is a pointer to x
// value := *p		// Dereference: get value p points to
// *p = 100		// Modify value through pointer

//Pointer in Functions
// func modifuValue(x *int) {
// 	*x = 100
// }

// func main() {
// 	value := 42
// 	modifuValue(&value)		// Pass address
// 	// Value is now 100
// }


import "fmt"

// Function that modifies value (need pointer)
func increment(x *int) {
	*x++	// Increment value at address x
}

// Function that doesn't modify (no pointer needed)
func addOne(x int) int {
	return x + 1
}

// Function returning pointer
func createInt(value int) *int {
	return &value	// Return address of local variable (safe in go)
}

func main() {
	// 1. Basic pointer operations
	x := 42
	fmt.Printf("x = %d\n", x)

	p := &x
	fmt.Printf("p points to address: %p\n", p)
	fmt.Printf("Value at p: %d\n", *p)

	// 2. Modify through pointer
	*p = 100
	fmt.Printf("After *p = 100, x = %d\n", x)

	// 3. Pointer in function
	value := 100
	fmt.Printf("Before increment: %d\n", value)
	increment(&value)
	fmt.Printf("After increament: %d\n", value)

	// 4. Compare with value copy
	value2 := 10
	fmt.Printf("Before addOne: %d\n", value2)
	result := addOne(value2)
	fmt.Printf("After addOne: %d (original unchanged)\n", value2)
	fmt.Printf("Result: %d\n", result)

	// 5. Nil pointer
	var nilPtr *int
	fmt.Printf("nilPtr is nil: %v\n", nilPtr == nil)
	// *nilPtr would cause panic! Always check:
	if nilPtr != nil {
        fmt.Println(*nilPtr)
    }

	// 6. New function (allocates memory)
	ptr := new(int)
	*ptr = 42
	fmt.Printf("Value at new pointer: %d\n", *ptr)

	// 7. Pointer to pointer
	x2 := 42
	p2 := &x2
	pp := &p2	//	Pointer to pointer
	fmt.Printf("x2 = %d\n", x2)
	fmt.Printf("p2 = %d\n", *p2)
	fmt.Printf("pp = %d\n", **pp)
}