package main

import "fmt"

// STRUCT

/*
-	Structs are custom types that group together related data. They're like classes in Python (but simpler - no inheritance).
-	Structs let you create your own data types, organize related data, and build complex data structures. They're the foundation for object-oriented programming in Go.
-	Structs are copied when passed to functions (unless using pointer)
-	Assignment creates copy (not reference)
-	Can embed other structs (like inheritance, but composition)
*/

// BASIC STRUCT
// Define struct type
// type Person struct {
// 	Name string
// 	Age  int
// }

// Create instance
// p := Person{"Johnson", 30}	// Positional
// p := Person{Name: "John", Age: 30}	// Named (Recommended)
// p := Person{}	//	Zero values (Name="", Age=0)

// // Access fields
// p.Name = "Alice"
// age = p.Age

// STRUCT WITH METHODS
// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) Greet() string {
// 	return fmt.Sprintf("Hello, I'm %s", p.Name)
// }

// func (p *Person) HaveBirthday() {
// 	p.Age++
// }

// 1. Basic struct
type Person struct {
	Name string
	Age  int
}

// 2. Method with value reciever (can't modify)
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s, age %d", p.Name, p.Age)
}

// 3. Method with pointer reciever (can modify)
func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Printf("%s is now %d years old!\n", p.Name, p.Age)
}

// 4. Method that returns info
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// 5. Embedded struct (composition)
type Address struct {
	Street string
	City   string
	Zip    string
}

type Employee struct {
	Person
	Address
	EmployeeId int
	Salary     float64
}

func main() {
	// Create person
	p1 := Person{"Bean", 25}
	fmt.Println(p1.Greet())
	fmt.Printf("Is adult: %t\n", p1.IsAdult())

	// Modify with pointer reciever
	p1.HaveBirthday()

	// Create with named fields
	p2 := Person{
		Name: "Bob",
		Age:  30,
	}
	fmt.Println(p2.Greet())

	// Zero value
	var p3 Person
	fmt.Printf("Zero value: %+v\n", p3) // +v shows which value belongs to which field.

	emp := Employee{
		Person: Person{
			Name: "Robert",
			Age:  35,
		},
		Address: Address{
			Street: "123 KG Road",
			City:   "Mumbai",
			Zip:    "10001",
		},
		EmployeeId: 12345,
		Salary:     75000,
	}

	// Access embedded fields directly
	fmt.Printf(
		"Employee: %s, City: %s, Salary: $%.2f\n",
		emp.Name,   // From Person
		emp.City,   // From Address
		emp.Salary) // From Employee

	// Can also access through embedded type name
	fmt.Printf(
		"Full address: %s, %s %s\n",
		emp.Address.Street,
		emp.Address.City,
		emp.Address.Zip,
	)
}
