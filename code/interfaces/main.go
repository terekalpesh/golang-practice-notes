package main

// INTERFACES
/*
-	Interfaces define behavior - they specify what methods a type must have.
-	If a type has all the methods an interface requires, it automatically implements that interface (no explicit declaration needed).
-	Interfaces provide polymorphism.
*/


// BASIC INTERFACE

/*
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Any type with Area() and Perimeter() methods implements Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

// Usage
var s Shape = Circle{Radius: 5}
area := s.Area()
*/


// EMPTY INTERFACE

/*
// interface{} can hold any type
var i interface{} = 42
var i interface{} = "hello"
var i interface{} = []int{1, 2, 3}

// Type assertion
value, ok := i.(int)
*/


import (
	"fmt"
	"math"
)

//	1. Define interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 	2. Implement with Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

//	3. Implement with Rectangle
type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 4. Function that accepts interface
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n",
	s.Area(), s.Perimeter(),)
}

// 5. Interface with one method
type Stringer interface {
	String() string
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(radius=%.2f)", c.Radius)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(width=%.2f, height=%.2f)",
	r.Width, r.Width)
}

// 6. Empty interface
func printAnything(i interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", i, i)
}

// 7. Type assertion
func processValue(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case Circle:
		fmt.Printf("Circle with radius: %.2f\n", v.Radius)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func main() {
	// Create shapes
	circle := Circle{Radius: 5}
	rect := Rectangle{Width: 10, Height: 5,}

	// Use as interface
	var s1 Shape = circle
	var s2 Shape = rect

	fmt.Println("Shape 1: ")
	printShapeInfo(s1)

	fmt.Println("Shape 2: ")
	printAnything(s2)

	// Stringer interface (built-in)
	fmt.Println("String representation:")
	fmt.Println(circle)
	fmt.Println(rect)

	// Empty interface
	fmt.Println("\nEmmpty interface:")
	printAnything(42)
	printAnything("hello")
	printAnything(circle)

	// Type assertion
	fmt.Println("\nType assertion:")
	processValue(42)
	processValue("hello")
	processValue(circle)
	processValue(3.14)

	// Type assertion with ok check
	var i interface{} = 42
	if value, ok := i.(int); ok {
		fmt.Printf("It's an int: %d\n", value)
	}
}
