package main

import "fmt"

// provide action to structs (like class-methods in python)
// Value receiver → copy → no change
// Pointer receiver → points to original → changes persist
// Receiver -> variable that represents struct

// SYNTAX
/*
// BASIC METHOD
type Person struct {
	Name string
	Age  int
}

// Value receiver
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s", p.Name)
}

// Pointer receiver
func (p *Person) HaveBirthday() {
	p.Age++
}

// Usage
p := Person{"Alice", 30}
greeting := p.Greet()
p.HaveBirthday()
*/

// METHOD ON ANY TYPE
/*
type MyInt int

func (m MyInt) Double() int {
	return int(m * 2)
}

value := MyInt(5)
result := value.Double()	// 10
*/

type Rectangle struct {
	Width  float64
	Height float64
}

// Value receiver - doesn't modify
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Pointer receiver - can modify
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func (r *Rectangle) SetDimensions(width, height float64) {
	r.Width = width
	r.Height = height
}

// Method on non-struct type
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Method with multiiple receivers (not possible - one method per type)
// But you can have multiple methods on same type

type BankAccount struct {
	Balance float64
	Owner   string
}

func (ba *BankAccount) Deposit(amount float64) {
	ba.Balance += amount
}

func (ba *BankAccount) Withdraw(amount float64) error {
	if amount > ba.Balance {
		return fmt.Errorf("insufficient funds")
	}
	ba.Balance -= amount
	return nil
}

func (ba BankAccount) GetBalance() float64 {
	return ba.Balance
}

func main() {
	// Value receiver methods
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())

	// Pointer receiver methods
	rect.Scale(2)	//	Go automatically takes address
	fmt.Printf("After scaling: Width=%.2f, Height=%.2f\n", rect.Width, rect.Height)

	// Method on non-struct type
	f := MyFloat(-5.5)
	fmt.Printf("Absolute value: %.2f\n", f.Abs())

	// Bank account example
	account := BankAccount{
		Balance: 1000,
		Owner: "Alice",
	}

	account.Deposit(500)
	fmt.Printf("Balance after deposit: $%.2f\n", account.GetBalance())

	err := account.Withdraw(200)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Balance after withdrawal: $%.2f\n", account.GetBalance())
	}
}