# Methods

## a) Overview

### What this topic is
Methods are functions that belong to a specific type. They let you attach behavior to your data types (like structs).

### Why it exists in Go
Methods provide a way to organize code by associating functions with the data they operate on. They're Go's way of doing object-oriented programming without classes.

### 🎯 Layman's Explanation (Simple Terms)

**Think of methods like actions that belong to an object:**

**Real-world analogy - A Car:**
- A **Car** (struct) has properties: Color, Brand, Speed
- A **Car** can do actions: Start(), Drive(), Stop()
- These actions are **methods** - they belong to the car
- You don't say "Start the car" to a person - you say it to the car itself!

**Another analogy - A Bank Account:**
- A **Bank Account** (struct) has: Account Number, Balance
- A **Bank Account** can do: Deposit(), Withdraw(), CheckBalance()
- These are **methods** - they operate on the account
- Like buttons on an ATM - each button does something to YOUR account

**Simple example:**
```
Without methods (separate functions):
account = BankAccount{Balance: 100}
deposit(account, 50)  // Function takes account as parameter
withdraw(account, 20)  // Function takes account as parameter
// Functions are separate from the data

With methods (attached to data):
account = BankAccount{Balance: 100}
account.Deposit(50)  // Method belongs to account
account.Withdraw(20)  // Method belongs to account
// Methods are part of the account - more natural!
```

**Method = An action that belongs to something:**
- Like a **remote control** for a TV - the buttons (methods) control that specific TV
- Like **functions on a calculator** - Add, Subtract belong to the calculator
- Like **verbs** for a noun - "The car drives" - "drives" is a method of "car"

**Value receiver vs Pointer receiver:**
- **Value receiver** (copy): Like taking a photo of a document - you can look at it, but can't change the original
- **Pointer receiver** (original): Like having the actual document - you can modify it

**Why use methods?**
1. **Organization**: Keep actions with the data they work on (like keeping car controls with the car)
2. **Clarity**: Clear what can be done with something (like a TV remote shows what the TV can do)
3. **Natural**: More intuitive - "account.Deposit()" reads like English
4. **Encapsulation**: Data and behavior together (like a phone has both data and functions)

---

## b) Syntax

### Basic Method
```go
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
```

### Method on Any Type
```go
type MyInt int

func (m MyInt) Double() int {
    return int(m * 2)
}

value := MyInt(5)
result := value.Double()  // 10
```

---

## c) Explanation

### Step-by-Step Method Creation

**1. Value receiver**
```go
func (p Person) Greet() string {
    return fmt.Sprintf("Hello, I'm %s", p.Name)
}
```
- `(p Person)` = receiver (like `self` in Python)
- `p` is a copy of the struct
- Cannot modify original struct
- Use when method doesn't need to modify

**2. Pointer receiver**
```go
func (p *Person) HaveBirthday() {
    p.Age++
}
```
- `(p *Person)` = pointer receiver
- `p` points to original struct
- Can modify original struct
- Use when method needs to modify

**3. Method call**
```go
p := Person{"Alice", 30}
p.Greet()  // Go automatically handles value/pointer
```
- Go automatically converts between value and pointer
- `p.Greet()` works whether receiver is value or pointer
- Go handles `&p` or `*p` automatically

### Characteristics

- **Attached to types**: Methods belong to a type
- **Two receiver types**: Value or pointer
- **Automatic conversion**: Go handles value/pointer conversion
- **Any type**: Can add methods to any type (not just structs)
- **No `this` or `self`**: Receiver name is your choice

---

## d) Python Comparison

### Python Methods
```python
class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age
    
    def greet(self):
        return f"Hello, I'm {self.name}"
    
    def have_birthday(self):
        self.age += 1

p = Person("Alice", 30)
p.greet()
p.have_birthday()
```

### Go Methods
```go
type Person struct {
    Name string
    Age  int
}

func (p Person) Greet() string {
    return fmt.Sprintf("Hello, I'm %s", p.Name)
}

func (p *Person) HaveBirthday() {
    p.Age++
}

p := Person{"Alice", 30}
p.Greet()
p.HaveBirthday()
```

**Key Differences:**

| Feature | Python | Go |
|---------|--------|-----|
| **Receiver name** | `self` (convention) | Any name (often type initial) |
| **Explicit receiver** | `self` always first | `(p Person)` before function name |
| **Value vs pointer** | Always reference | Choose value or pointer |
| **Method definition** | Inside class | Outside struct |
| **Automatic conversion** | N/A | Go converts value/pointer |

**Thinking Difference:**
- Python: Methods defined inside class
- Go: Methods defined outside struct (but attached to type)
- Python: Always works with reference
- Go: Choose value (copy) or pointer (reference)

---

## e) Visual Flow / Mental Model

### Method Call with Value Receiver

```
p := Person{"Alice", 30}
greeting := p.Greet()
```

**Flow:**
```
1. p.Greet() called
   ↓
2. Go creates copy of p
   ↓
3. Passes copy to Greet method
   ↓
4. Method operates on copy
   ↓
5. Returns result
   ↓
6. Original p unchanged
```

### Method Call with Pointer Receiver

```
p := Person{"Alice", 30}
p.HaveBirthday()
```

**Flow:**
```
1. p.HaveBirthday() called
   ↓
2. Go automatically takes address: &p
   ↓
3. Passes pointer to HaveBirthday method
   ↓
4. Method operates on original p
   ↓
5. Modifies p.Age directly
   ↓
6. Original p is modified
```

### Automatic Conversion

```
p := Person{"Alice", 30}

// If method has value receiver
func (p Person) Greet() { }
p.Greet()        // ✅ Works (value)
(&p).Greet()     // ✅ Works (Go converts to value)

// If method has pointer receiver
func (p *Person) HaveBirthday() { }
p.HaveBirthday()     // ✅ Works (Go takes address)
(*p).HaveBirthday()  // ✅ Works (pointer)
```

**Go is smart**: It automatically converts between value and pointer!

---

## f) Demo Example

### Complete Example

```go
package main

import "fmt"

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

// Method with multiple receivers (not possible - one method per type)
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
    rect.Scale(2)  // Go automatically takes address
    fmt.Printf("After scaling: Width=%.2f, Height=%.2f\n", rect.Width, rect.Height)
    
    rect.SetDimensions(20, 10)
    fmt.Printf("New dimensions: Width=%.2f, Height=%.2f\n", rect.Width, rect.Height)
    
    // Method on non-struct type
    f := MyFloat(-5.5)
    fmt.Printf("Absolute value: %.2f\n", f.Abs())
    
    // Bank account example
    account := BankAccount{
        Balance: 1000,
        Owner:   "Alice",
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
```

**Line-by-line explanation:**

1. **Value receiver**: Gets copy, can't modify original
2. **Pointer receiver**: Gets reference, can modify original
3. **Methods on any type**: Can add methods to int, string, etc. (via type alias)
4. **Multiple methods**: One type can have many methods
5. **Automatic conversion**: Go handles value/pointer automatically
6. **Consistent receivers**: Usually all methods use same receiver type

**Output:**
```
Area: 50.00
Perimeter: 30.00
After scaling: Width=20.00, Height=10.00
New dimensions: Width=20.00, Height=10.00
Absolute value: 5.50
Balance after deposit: $1500.00
Balance after withdrawal: $1300.00
```

---

## g) Use Cases

### When to use value receiver

**1. Read-only operations**
```go
func (p Person) GetName() string {
    return p.Name
}
```

**2. Small structs**
```go
type Point struct {
    X, Y float64
}

func (p Point) Distance() float64 {
    return math.Sqrt(p.X*p.X + p.Y*p.Y)
}
```

**3. Immutable operations**
```go
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

### When to use pointer receiver

**1. Modify struct**
```go
func (p *Person) SetAge(age int) {
    p.Age = age
}
```

**2. Large structs (avoid copying)**
```go
type LargeStruct struct {
    // many fields
}

func (ls *LargeStruct) Process() {
    // modify ls
}
```

**3. Consistency**
- If one method uses pointer, all should (usually)
- Makes API consistent

---

## h) Do's and Don'ts / Best Practices

### ✅ Do's

1. **Be consistent with receivers**
   ```go
   // ✅ All methods use pointer
   func (p *Person) Method1() { }
   func (p *Person) Method2() { }
   ```

2. **Use pointer receiver for modification**
   ```go
   func (p *Person) SetName(name string) {  // ✅
       p.Name = name
   }
   ```

3. **Use value receiver for read-only**
   ```go
   func (p Person) GetName() string {  // ✅
       return p.Name
   }
   ```

4. **Choose receiver name wisely**
   ```go
   func (p Person) Greet() { }      // ✅ p for Person
   func (r Rectangle) Area() { }    // ✅ r for Rectangle
   ```

5. **Use pointer for large structs**
   ```go
   func (ls *LargeStruct) Process() { }  // ✅ Avoid copying
   ```

### ❌ Don'ts

1. **Don't mix receiver types unnecessarily**
   ```go
   // ❌ Inconsistent
   func (p Person) Method1() { }
   func (p *Person) Method2() { }
   ```

2. **Don't use pointer receiver for tiny structs**
   ```go
   type Point struct { X, Y float64 }
   func (p Point) Distance() float64 { }  // ✅ Value is fine
   ```

3. **Don't use value receiver when you need to modify**
   ```go
   func (p Person) SetAge(age int) {  // ❌ Won't modify!
       p.Age = age  // Only modifies copy
   }
   ```

4. **Don't create methods on built-in types directly**
   ```go
   // ❌ Can't do this
   func (i int) Double() int { }
   
   // ✅ Create type alias first
   type MyInt int
   func (m MyInt) Double() int { }
   ```

---

## i) Solved Practice Examples

### Example 1: Counter with Methods

**Task:** Create a Counter struct with Increment, Decrement, and GetValue methods.

**Solution:**
```go
package main

import "fmt"

type Counter struct {
    value int
}

func (c *Counter) Increment() {
    c.value++
}

func (c *Counter) Decrement() {
    c.value--
}

func (c Counter) GetValue() int {
    return c.value
}

func main() {
    counter := Counter{}
    counter.Increment()
    counter.Increment()
    counter.Increment()
    fmt.Println("Value:", counter.GetValue())  // 3
    
    counter.Decrement()
    fmt.Println("Value:", counter.GetValue())  // 2
}
```

### Example 2: Temperature Converter

**Task:** Create a Temperature type with methods to convert between Celsius and Fahrenheit.

**Solution:**
```go
package main

import "fmt"

type Temperature float64

func (t Temperature) ToFahrenheit() float64 {
    return float64(t)*9/5 + 32
}

func (t Temperature) ToCelsius() float64 {
    return float64(t)
}

func FahrenheitToCelsius(f float64) Temperature {
    return Temperature((f - 32) * 5 / 9)
}

func main() {
    temp := Temperature(25)  // 25°C
    fmt.Printf("%.1f°C = %.1f°F\n", temp, temp.ToFahrenheit())
    
    f := 77.0
    c := FahrenheitToCelsius(f)
    fmt.Printf("%.1f°F = %.1f°C\n", f, c)
}
```

---

## j) Quiz / Practice for Me

### Quiz Questions

1. **What is a receiver in Go?**
   - [ ] A return value
   - [ ] The parameter that attaches method to type
   - [ ] A function parameter
   - [ ] A type definition

2. **When should you use a pointer receiver?**
   - [ ] Always
   - [ ] When method modifies struct
   - [ ] Never
   - [ ] Only for small structs

3. **Can you add methods to built-in types like int?**
   - [ ] Yes, directly
   - [ ] No, never
   - [ ] Yes, with type alias
   - [ ] Only for some types

### Practice Tasks

**Task 1: Stack Methods**
- Create Stack struct with slice field
- Add methods: Push(item int), Pop() (int, error), IsEmpty() bool
- Test all methods

**Task 2: String Methods**
- Create type MyString string
- Add method Reverse() string that returns reversed string
- Add method Uppercase() string
- Test both methods

### Answers

**Quiz Answers:**
1. The parameter that attaches method to type
2. When method modifies struct
3. Yes, with type alias

**Practice Solutions:**

**Task 1 Solution:**
```go
package main

import (
    "errors"
    "fmt"
)

type Stack struct {
    items []int
}

func (s *Stack) Push(item int) {
    s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, error) {
    if s.IsEmpty() {
        return 0, errors.New("stack is empty")
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, nil
}

func (s Stack) IsEmpty() bool {
    return len(s.items) == 0
}

func main() {
    stack := Stack{}
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)
    
    for !stack.IsEmpty() {
        item, _ := stack.Pop()
        fmt.Println("Popped:", item)
    }
}
```

**Task 2 Solution:**
```go
package main

import (
    "fmt"
    "strings"
)

type MyString string

func (s MyString) Reverse() string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func (s MyString) Uppercase() string {
    return strings.ToUpper(string(s))
}

func main() {
    str := MyString("Hello, World!")
    fmt.Println("Original:", str)
    fmt.Println("Reversed:", str.Reverse())
    fmt.Println("Uppercase:", str.Uppercase())
}
```

---

## Key Takeaways

1. **Methods attach to types** - Functions that belong to a type
2. **Two receiver types** - Value (copy) or pointer (reference)
3. **Automatic conversion** - Go handles value/pointer conversion
4. **Any type can have methods** - Not just structs (use type alias)
5. **Be consistent** - Usually all methods use same receiver type
6. **Pointer for modification** - Use pointer receiver when modifying

---

## Must Remember Forever

- `func (p Person) method()` - Value receiver
- `func (p *Person) method()` - Pointer receiver
- Go automatically converts between value and pointer
- Use pointer receiver to modify struct
- Use value receiver for read-only operations
- Methods are defined outside struct (unlike Python classes)

---

---

## k) Additional Practice Problems (Build Strong Logic)

### 10 More Solved Examples

### Solved Problem 1: Stack with Methods

**Task:** Implement stack data structure with methods.

**Solution:**
```go
package main

import (
    "errors"
    "fmt"
)

type Stack struct {
    items []int
}

func NewStack() *Stack {
    return &Stack{items: []int{}}
}

func (s *Stack) Push(item int) {
    s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, error) {
    if s.IsEmpty() {
        return 0, errors.New("stack is empty")
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, nil
}

func (s *Stack) Peek() (int, error) {
    if s.IsEmpty() {
        return 0, errors.New("stack is empty")
    }
    return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
    return len(s.items) == 0
}

func (s *Stack) Size() int {
    return len(s.items)
}

func main() {
    stack := NewStack()
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)
    
    fmt.Printf("Stack size: %d\n", stack.Size())
    peek, _ := stack.Peek()
    fmt.Printf("Top element: %d\n", peek)
    
    for !stack.IsEmpty() {
        item, _ := stack.Pop()
        fmt.Printf("Popped: %d\n", item)
    }
}
```

### Solved Problem 2: Queue with Methods

**Task:** Implement queue data structure with methods.

**Solution:**
```go
package main

import (
    "errors"
    "fmt"
)

type Queue struct {
    items []int
}

func NewQueue() *Queue {
    return &Queue{items: []int{}}
}

func (q *Queue) Enqueue(item int) {
    q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (int, error) {
    if q.IsEmpty() {
        return 0, errors.New("queue is empty")
    }
    item := q.items[0]
    q.items = q.items[1:]
    return item, nil
}

func (q *Queue) Front() (int, error) {
    if q.IsEmpty() {
        return 0, errors.New("queue is empty")
    }
    return q.items[0], nil
}

func (q *Queue) IsEmpty() bool {
    return len(q.items) == 0
}

func main() {
    queue := NewQueue()
    queue.Enqueue(1)
    queue.Enqueue(2)
    queue.Enqueue(3)
    
    for !queue.IsEmpty() {
        item, _ := queue.Dequeue()
        fmt.Printf("Dequeued: %d\n", item)
    }
}
```

### Solved Problem 3: Calculator with Methods

**Task:** Create calculator with method chaining.

**Solution:**
```go
package main

import "fmt"

type Calculator struct {
    result float64
}

func NewCalculator() *Calculator {
    return &Calculator{result: 0}
}

func (c *Calculator) Add(value float64) *Calculator {
    c.result += value
    return c
}

func (c *Calculator) Subtract(value float64) *Calculator {
    c.result -= value
    return c
}

func (c *Calculator) Multiply(value float64) *Calculator {
    c.result *= value
    return c
}

func (c *Calculator) Divide(value float64) *Calculator {
    if value != 0 {
        c.result /= value
    }
    return c
}

func (c *Calculator) GetResult() float64 {
    return c.result
}

func (c *Calculator) Reset() {
    c.result = 0
}

func main() {
    calc := NewCalculator()
    result := calc.Add(10).Multiply(2).Subtract(5).GetResult()
    fmt.Printf("Result: %.2f\n", result)
}
```

### Solved Problem 4: String Builder with Methods

**Task:** Create string builder with useful methods.

**Solution:**
```go
package main

import "fmt"

type StringBuilder struct {
    parts []string
}

func NewStringBuilder() *StringBuilder {
    return &StringBuilder{parts: []string{}}
}

func (sb *StringBuilder) Append(s string) *StringBuilder {
    sb.parts = append(sb.parts, s)
    return sb
}

func (sb *StringBuilder) AppendLine(s string) *StringBuilder {
    sb.parts = append(sb.parts, s+"\n")
    return sb
}

func (sb *StringBuilder) Clear() {
    sb.parts = []string{}
}

func (sb *StringBuilder) String() string {
    result := ""
    for _, part := range sb.parts {
        result += part
    }
    return result
}

func main() {
    sb := NewStringBuilder()
    sb.Append("Hello").Append(" ").Append("World")
    fmt.Println(sb.String())
    
    sb.Clear()
    sb.AppendLine("Line 1").AppendLine("Line 2")
    fmt.Println(sb.String())
}
```

### Solved Problem 5: Counter with Methods

**Task:** Create counter with increment, decrement, and reset.

**Solution:**
```go
package main

import "fmt"

type Counter struct {
    value int
    max   int
    min   int
}

func NewCounter(min, max int) *Counter {
    return &Counter{value: 0, min: min, max: max}
}

func (c *Counter) Increment() bool {
    if c.value < c.max {
        c.value++
        return true
    }
    return false
}

func (c *Counter) Decrement() bool {
    if c.value > c.min {
        c.value--
        return true
    }
    return false
}

func (c *Counter) Reset() {
    c.value = 0
}

func (c *Counter) Value() int {
    return c.value
}

func (c *Counter) IsAtMax() bool {
    return c.value >= c.max
}

func (c *Counter) IsAtMin() bool {
    return c.value <= c.min
}

func main() {
    counter := NewCounter(0, 10)
    for i := 0; i < 12; i++ {
        if !counter.Increment() {
            fmt.Println("Reached maximum!")
            break
        }
    }
    fmt.Printf("Counter value: %d\n", counter.Value())
}
```

### Solved Problem 6: Date with Methods

**Task:** Create Date struct with useful date methods.

**Solution:**
```go
package main

import "fmt"

type Date struct {
    Year  int
    Month int
    Day   int
}

func NewDate(year, month, day int) *Date {
    return &Date{Year: year, Month: month, Day: day}
}

func (d *Date) IsLeapYear() bool {
    return (d.Year%4 == 0 && d.Year%100 != 0) || (d.Year%400 == 0)
}

func (d *Date) DaysInMonth() int {
    days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
    if d.Month == 2 && d.IsLeapYear() {
        return 29
    }
    return days[d.Month-1]
}

func (d *Date) AddDays(days int) {
    d.Day += days
    for d.Day > d.DaysInMonth() {
        d.Day -= d.DaysInMonth()
        d.Month++
        if d.Month > 12 {
            d.Month = 1
            d.Year++
        }
    }
}

func (d *Date) String() string {
    return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}

func main() {
    date := NewDate(2024, 2, 28)
    fmt.Printf("Date: %s\n", date)
    fmt.Printf("Is leap year: %t\n", date.IsLeapYear())
    date.AddDays(1)
    fmt.Printf("After adding 1 day: %s\n", date)
}
```

### Solved Problem 7: Fraction Calculator

**Task:** Create Fraction type with arithmetic methods.

**Solution:**
```go
package main

import "fmt"

type Fraction struct {
    Numerator   int
    Denominator int
}

func NewFraction(num, den int) *Fraction {
    if den == 0 {
        return nil
    }
    return &Fraction{Numerator: num, Denominator: den}
}

func (f *Fraction) Add(other *Fraction) *Fraction {
    num := f.Numerator*other.Denominator + other.Numerator*f.Denominator
    den := f.Denominator * other.Denominator
    return NewFraction(num, den)
}

func (f *Fraction) Multiply(other *Fraction) *Fraction {
    return NewFraction(f.Numerator*other.Numerator, f.Denominator*other.Denominator)
}

func (f *Fraction) String() string {
    return fmt.Sprintf("%d/%d", f.Numerator, f.Denominator)
}

func (f *Fraction) ToFloat() float64 {
    return float64(f.Numerator) / float64(f.Denominator)
}

func main() {
    f1 := NewFraction(1, 2)
    f2 := NewFraction(1, 3)
    
    sum := f1.Add(f2)
    fmt.Printf("%s + %s = %s (%.3f)\n", f1, f2, sum, sum.ToFloat())
    
    product := f1.Multiply(f2)
    fmt.Printf("%s × %s = %s (%.3f)\n", f1, f2, product, product.ToFloat())
}
```

### Solved Problem 8: Bank Account with Methods

**Task:** Create bank account with transaction methods.

**Solution:**
```go
package main

import (
    "errors"
    "fmt"
)

type BankAccount struct {
    accountNumber string
    balance       float64
    transactions  int
}

func NewBankAccount(accountNumber string) *BankAccount {
    return &BankAccount{
        accountNumber: accountNumber,
        balance:       0,
        transactions:  0,
    }
}

func (ba *BankAccount) Deposit(amount float64) error {
    if amount <= 0 {
        return errors.New("deposit amount must be positive")
    }
    ba.balance += amount
    ba.transactions++
    return nil
}

func (ba *BankAccount) Withdraw(amount float64) error {
    if amount <= 0 {
        return errors.New("withdrawal amount must be positive")
    }
    if amount > ba.balance {
        return errors.New("insufficient funds")
    }
    ba.balance -= amount
    ba.transactions++
    return nil
}

func (ba *BankAccount) GetBalance() float64 {
    return ba.balance
}

func (ba *BankAccount) GetTransactionCount() int {
    return ba.transactions
}

func (ba *BankAccount) Transfer(to *BankAccount, amount float64) error {
    if err := ba.Withdraw(amount); err != nil {
        return err
    }
    return to.Deposit(amount)
}

func main() {
    acc1 := NewBankAccount("ACC001")
    acc2 := NewBankAccount("ACC002")
    
    acc1.Deposit(1000)
    acc1.Transfer(acc2, 300)
    
    fmt.Printf("Account 1 balance: $%.2f\n", acc1.GetBalance())
    fmt.Printf("Account 2 balance: $%.2f\n", acc2.GetBalance())
}
```

### Solved Problem 9: Temperature with Methods

**Task:** Create Temperature type with conversion methods.

**Solution:**
```go
package main

import "fmt"

type Temperature struct {
    celsius float64
}

func NewTemperatureCelsius(c float64) *Temperature {
    return &Temperature{celsius: c}
}

func NewTemperatureFahrenheit(f float64) *Temperature {
    return &Temperature{celsius: (f - 32) * 5 / 9}
}

func (t *Temperature) Celsius() float64 {
    return t.celsius
}

func (t *Temperature) Fahrenheit() float64 {
    return t.celsius*9/5 + 32
}

func (t *Temperature) Kelvin() float64 {
    return t.celsius + 273.15
}

func (t *Temperature) Add(other *Temperature) *Temperature {
    return NewTemperatureCelsius(t.celsius + other.celsius)
}

func (t *Temperature) String() string {
    return fmt.Sprintf("%.2f°C (%.2f°F)", t.celsius, t.Fahrenheit())
}

func main() {
    temp1 := NewTemperatureCelsius(25)
    temp2 := NewTemperatureFahrenheit(77)
    
    fmt.Printf("Temp1: %s\n", temp1)
    fmt.Printf("Temp2: %s\n", temp2)
    fmt.Printf("Sum: %s\n", temp1.Add(temp2))
}
```

### Solved Problem 10: Complex Number Operations

**Task:** Create Complex number type with arithmetic methods.

**Solution:**
```go
package main

import "fmt"

type Complex struct {
    Real      float64
    Imaginary float64
}

func NewComplex(real, imaginary float64) *Complex {
    return &Complex{Real: real, Imaginary: imaginary}
}

func (c *Complex) Add(other *Complex) *Complex {
    return NewComplex(c.Real+other.Real, c.Imaginary+other.Imaginary)
}

func (c *Complex) Multiply(other *Complex) *Complex {
    real := c.Real*other.Real - c.Imaginary*other.Imaginary
    imag := c.Real*other.Imaginary + c.Imaginary*other.Real
    return NewComplex(real, imag)
}

func (c *Complex) Magnitude() float64 {
    return c.Real*c.Real + c.Imaginary*c.Imaginary
}

func (c *Complex) String() string {
    if c.Imaginary >= 0 {
        return fmt.Sprintf("%.2f + %.2fi", c.Real, c.Imaginary)
    }
    return fmt.Sprintf("%.2f - %.2fi", c.Real, -c.Imaginary)
}

func main() {
    c1 := NewComplex(3, 4)
    c2 := NewComplex(1, 2)
    
    sum := c1.Add(c2)
    product := c1.Multiply(c2)
    
    fmt.Printf("c1 = %s\n", c1)
    fmt.Printf("c2 = %s\n", c2)
    fmt.Printf("c1 + c2 = %s\n", sum)
    fmt.Printf("c1 × c2 = %s\n", product)
}
```

---

### 10 More Practice Problems (Solve These!)

**Problem 1: Linked List with Methods**
- Create Node and LinkedList structs
- Methods: Insert, Delete, Search, Display, Reverse

**Problem 2: Binary Tree with Methods**
- Create TreeNode and BinaryTree
- Methods: Insert, Search, InOrder, PreOrder, PostOrder

**Problem 3: Set Data Structure**
- Create Set with methods: Add, Remove, Contains, Union, Intersection

**Problem 4: Matrix Operations**
- Create Matrix struct with methods: Add, Multiply, Transpose, Determinant

**Problem 5: Polynomial Calculator**
- Create Polynomial with methods: Add, Multiply, Evaluate, Derivative

**Problem 6: Time Duration Calculator**
- Create Duration with methods: Add, Subtract, Multiply, Format

**Problem 7: Vector Operations**
- Create Vector with methods: Add, DotProduct, CrossProduct, Magnitude

**Problem 8: Priority Queue**
- Create PriorityQueue with methods: Enqueue, Dequeue, Peek

**Problem 9: Graph Data Structure**
- Create Graph with methods: AddVertex, AddEdge, BFS, DFS

**Problem 10: Cache with Methods**
- Create Cache with methods: Get, Set, Delete, Clear, Size

---

## Next Steps

Now you understand methods. Next:
- Interfaces (polymorphism in Go)
- How interfaces work with methods
- Duck typing in Go

**Ready? → [09_INTERFACES.md](./09_INTERFACES.md)**
