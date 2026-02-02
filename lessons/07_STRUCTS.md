# Structs

## a) Overview

### What this topic is
Structs are custom types that group together related data. They're like classes in Python (but simpler - no inheritance).

### Why it exists in Go
Structs let you create your own data types, organize related data, and build complex data structures. They're the foundation for object-oriented programming in Go.

### 🎯 Layman's Explanation (Simple Terms)

**Think of structs like a form, a card, or a template:**

**Real-world analogy - A Contact Card:**
- Instead of having separate pieces of paper: one with name, one with phone, one with email
- You create a **Contact Card** that has spaces for: Name, Phone, Email, Address
- All related information in one organized place!
- Like a **form** you fill out - the form defines what information goes where

**Another analogy - A Recipe Card:**
- A recipe has: Name, Ingredients, Instructions, Cooking Time
- Instead of separate notes, you have a **Recipe Card** with all sections
- Like a **template** - you define what information a recipe should have
- Then you can create many recipe cards, each with different values

**Simple example:**
```
Without struct (messy - like loose papers):
name1 = "Alice"
age1 = 30
email1 = "alice@email.com"
name2 = "Bob"  
age2 = 25
email2 = "bob@email.com"
// Hard to keep track! Which name goes with which age and email?

With struct (organized - like a card):
person1 = Person{Name: "Alice", Age: 30, Email: "alice@email.com"}
person2 = Person{Name: "Bob", Age: 25, Email: "bob@email.com"}
// Everything grouped together - clear and organized!
// Like having a contact card for each person
```

**Struct = A custom container/template:**
- Like a **shopping list template** with sections: Groceries, Toiletries, Electronics
- Like a **student report card** with: Name, Grades, Comments
- Like a **form** with multiple fields to fill in
- Like a **blueprint** that defines what information something should have

**Why use structs?**
1. **Organization**: Keep related data together (like keeping all your contact info on one card instead of scattered papers)
2. **Reusability**: Create many instances (like printing many contact cards from the same template)
3. **Clarity**: Clear what data belongs together (like a form shows exactly what information is needed)
4. **Type Safety**: Go knows what fields exist (like a form won't let you put phone number in the name field)

---

## b) Syntax

### Basic Struct
```go
type Person struct {
    Name string
    Age  int
}

// Create instance
p := Person{"John", 30}
p := Person{Name: "John", Age: 30}
p := Person{}  // Zero values

// Access fields
p.Name = "Alice"
age := p.Age
```

### Struct with Methods
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
```

---

## c) Explanation

### Step-by-Step Struct Creation

**1. Define struct type**
```go
type Person struct {
    Name string
    Age  int
}
```
- `type Person struct` = define new type called Person
- Fields inside `{ }`
- Each field has name and type

**2. Create instance**
```go
p := Person{"John", 30}  // Positional
p := Person{Name: "John", Age: 30}  // Named (recommended)
```
- Can initialize with values
- Named fields are clearer and safer

**3. Access fields**
```go
p.Name = "Alice"
age := p.Age
```
- Use dot notation
- Can read and write

**4. Zero values**
```go
p := Person{}  // Name="", Age=0
```
- Uninitialized fields get zero values
- String = "", int = 0, etc.

### Characteristics

#### Struct Data Structure Characteristics
- **Value type**: Structs are copied when passed to functions (unless using pointer)
- **Memory layout**: Fields stored in order (contiguous in memory)
- **Zero values**: All fields get zero values if not initialized
- **Type safety**: Strong typing - each field has specific type
- **Size**: Sum of all field sizes (plus padding for alignment)
- **Comparable**: Structs are comparable if all fields are comparable
- **Copy behavior**: Assignment creates copy (not reference)

#### Design Characteristics
- **No inheritance**: Go doesn't have classes/inheritance
- **Composition over inheritance**: Build complex types by embedding structs
- **Methods**: Can attach functions to structs (value or pointer receiver)
- **Exported fields**: Capital letter = exported (public), lowercase = unexported (private)
- **Field tags**: Can add metadata tags (e.g., JSON tags)
- **Embedded structs**: Can embed other structs (like inheritance, but composition)

#### Data Collection Characteristics
- **Heterogeneous data**: Can store different types of data together
- **Named fields**: Access by name (not index like arrays)
- **Type definition**: Defines structure once, create many instances
- **Memory efficient**: Only stores actual data (no overhead like classes)
- **Immutable structure**: Field names/types fixed after definition
- **Flexible values**: Field values can change, structure cannot
- **Grouping**: Groups related data into logical units

---

## d) Python Comparison

### Python Class
```python
class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age
    
    def greet(self):
        return f"Hello, I'm {self.name}"
    
    def have_birthday(self):
        self.age += 1

p = Person("John", 30)
p.greet()
p.have_birthday()
```

### Go Struct
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

p := Person{"John", 30}
p.Greet()
p.HaveBirthday()
```

**Key Differences:**

| Feature | Python | Go |
|---------|--------|-----|
| **Definition** | `class Person:` | `type Person struct { }` |
| **Initialization** | `Person(name, age)` | `Person{Name: name, Age: age}` |
| **Methods** | `def method(self):` | `func (p Person) method() { }` |
| **Inheritance** | Yes | No (use composition) |
| **Self/this** | `self` | Receiver name (any name) |
| **Private fields** | `_name` (convention) | `name` (lowercase) |
| **Public fields** | `name` | `Name` (capitalize) |

**Thinking Difference:**
- Python: Classes with inheritance
- Go: Structs with composition
- Python: Everything is an object
- Go: Structs are just data, methods are separate

---

## e) Visual Flow / Mental Model

### Struct in Memory

```
type Person struct {
    Name string
    Age  int
}

p := Person{"Alice", 30}
```

**Memory layout:**
```
Address     Field    Value
0x1000      Name     "Alice"
0x1010      Age      30
```

### Method Call

```
p := Person{"Alice", 30}
greeting := p.Greet()
```

**Flow:**
```
1. p.Greet() called
   ↓
2. Go finds Greet method for Person
   ↓
3. Passes p as receiver
   ↓
4. Method executes with p's data
   ↓
5. Returns result
```

### Pointer Receiver

```
p := Person{"Alice", 30}
p.HaveBirthday()  // Method with pointer receiver
```

**Flow:**
```
1. p.HaveBirthday() called
   ↓
2. Go automatically takes address: &p
   ↓
3. Method receives *Person
   ↓
4. Modifies p.Age directly
   ↓
5. Original p is modified
```

---

## f) Demo Example

### Complete Example

```go
package main

import "fmt"

// 1. Basic struct
type Person struct {
    Name string
    Age  int
}

// 2. Method with value receiver (can't modify)
func (p Person) Greet() string {
    return fmt.Sprintf("Hello, I'm %s, age %d", p.Name, p.Age)
}

// 3. Method with pointer receiver (can modify)
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
    Person           // Embedded (has Name and Age)
    Address          // Embedded (has Street, City, Zip)
    EmployeeID int
    Salary    float64
}

func main() {
    // Create Person
    p1 := Person{"Alice", 25}
    fmt.Println(p1.Greet())
    fmt.Printf("Is adult: %t\n", p1.IsAdult())
    
    // Modify with pointer receiver
    p1.HaveBirthday()
    
    // Create with named fields
    p2 := Person{
        Name: "Bob",
        Age:  30,
    }
    fmt.Println(p2.Greet())
    
    // Zero value
    var p3 Person
    fmt.Printf("Zero value: %+v\n", p3)
    
    // Embedded structs
    emp := Employee{
        Person: Person{
            Name: "Charlie",
            Age:  35,
        },
        Address: Address{
            Street: "123 Main St",
            City:   "New York",
            Zip:    "10001",
        },
        EmployeeID: 12345,
        Salary:     75000,
    }
    
    // Access embedded fields directly
    fmt.Printf("Employee: %s, City: %s, Salary: $%.2f\n",
        emp.Name,      // From Person
        emp.City,      // From Address
        emp.Salary)    // From Employee
    
    // Can also access through embedded type name
    fmt.Printf("Full address: %s, %s %s\n",
        emp.Address.Street,
        emp.Address.City,
        emp.Address.Zip)
}
```

**Line-by-line explanation:**

1. **Basic struct**: Define type with fields
2. **Value receiver**: `(p Person)` - gets copy, can't modify original
3. **Pointer receiver**: `(p *Person)` - can modify original
4. **Return value**: Methods can return values
5. **Embedded structs**: Composition (like inheritance, but better)
6. **Field access**: Direct access to embedded fields
7. **Named initialization**: Clear and safe

**Output:**
```
Hello, I'm Alice, age 25
Is adult: true
Alice is now 26 years old!
Hello, I'm Bob, age 30
Zero value: {Name: Age:0}
Employee: Charlie, City: New York, Salary: $75000.00
Full address: 123 Main St, New York 10001
```

---

## g) Use Cases

### When to use structs

**1. Group related data**
```go
type Point struct {
    X, Y float64
}
```

**2. Represent real-world entities**
```go
type User struct {
    ID       int
    Username string
    Email    string
    Created  time.Time
}
```

**3. Configuration**
```go
type Config struct {
    Host string
    Port int
    Debug bool
}
```

**4. API responses**
```go
type Response struct {
    Status  int
    Message string
    Data    interface{}
}
```

### Value vs Pointer Receiver

**Value receiver** (use when):
- Method doesn't modify struct
- Struct is small
- Want to work with copy

**Pointer receiver** (use when):
- Method modifies struct
- Struct is large (avoid copying)
- Consistency (if one method uses pointer, all should)

---

## h) Do's and Don'ts / Best Practices

### ✅ Do's

1. **Use named fields in initialization**
   ```go
   p := Person{Name: "John", Age: 30}  // ✅ Clear
   ```

2. **Use pointer receivers for modification**
   ```go
   func (p *Person) SetAge(age int) { p.Age = age }  // ✅
   ```

3. **Use value receivers for read-only**
   ```go
   func (p Person) GetAge() int { return p.Age }  // ✅
   ```

4. **Be consistent with receivers**
   - If one method uses pointer, all should (usually)

5. **Use embedded structs for composition**
   ```go
   type Employee struct {
       Person  // ✅ Composition
       // fields
   }
   ```

### ❌ Don'ts

1. **Don't mix value and pointer receivers unnecessarily**
   ```go
   // ❌ Inconsistent
   func (p Person) Method1() { }
   func (p *Person) Method2() { }
   ```

2. **Don't use positional initialization for many fields**
   ```go
   p := Person{"John", 30, "email", ...}  // ❌ Hard to read
   ```

3. **Don't export fields unnecessarily**
   ```go
   type Person struct {
       name string  // ✅ Unexported (private)
       Name string  // Only if needed outside package
   }
   ```

4. **Don't create deep inheritance hierarchies**
   - Go doesn't have inheritance
   - Use composition instead

---

## i) Solved Practice Examples

### Example 1: Rectangle Struct

**Task:** Create a Rectangle struct with methods to calculate area and perimeter.

**Solution:**
```go
package main

import "fmt"

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    fmt.Printf("Area: %.2f\n", rect.Area())
    fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())
}
```

### Example 2: Bank Account

**Task:** Create a BankAccount struct with deposit and withdraw methods.

**Solution:**
```go
package main

import (
    "errors"
    "fmt"
)

type BankAccount struct {
    Balance float64
    Owner   string
}

func (ba *BankAccount) Deposit(amount float64) {
    ba.Balance += amount
}

func (ba *BankAccount) Withdraw(amount float64) error {
    if amount > ba.Balance {
        return errors.New("insufficient funds")
    }
    ba.Balance -= amount
    return nil
}

func main() {
    account := BankAccount{
        Balance: 1000,
        Owner:   "Alice",
    }
    
    account.Deposit(500)
    fmt.Printf("Balance after deposit: $%.2f\n", account.Balance)
    
    err := account.Withdraw(200)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Balance after withdrawal: $%.2f\n", account.Balance)
    }
}
```

---

## j) Quiz / Practice for Me

### Quiz Questions

1. **What is the zero value of a struct?**
   - [ ] nil
   - [ ] All fields get their zero values
   - [ ] Error
   - [ ] Undefined

2. **When should you use a pointer receiver?**
   - [ ] Always
   - [ ] When method modifies struct
   - [ ] Never
   - [ ] Only for large structs

3. **How do you access embedded struct fields?**
   - [ ] Through the embedded type name only
   - [ ] Directly or through embedded type name
   - [ ] Only directly
   - [ ] Can't access

### Practice Tasks

**Task 1: Circle Struct**
- Create Circle struct with Radius field
- Add methods: Area() and Circumference()
- Test it

**Task 2: Student Struct**
- Create Student with Name, Age, Grades (slice of ints)
- Add method AverageGrade() that calculates average
- Add method IsPassing() that returns true if average >= 60

### Answers

**Quiz Answers:**
1. All fields get their zero values
2. When method modifies struct
3. Directly or through embedded type name

**Practice Solutions:**

**Task 1 Solution:**
```go
package main

import (
    "fmt"
    "math"
)

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Circumference() float64 {
    return 2 * math.Pi * c.Radius
}

func main() {
    c := Circle{Radius: 5}
    fmt.Printf("Area: %.2f\n", c.Area())
    fmt.Printf("Circumference: %.2f\n", c.Circumference())
}
```

**Task 2 Solution:**
```go
package main

import "fmt"

type Student struct {
    Name   string
    Age    int
    Grades []int
}

func (s Student) AverageGrade() float64 {
    if len(s.Grades) == 0 {
        return 0
    }
    sum := 0
    for _, grade := range s.Grades {
        sum += grade
    }
    return float64(sum) / float64(len(s.Grades))
}

func (s Student) IsPassing() bool {
    return s.AverageGrade() >= 60
}

func main() {
    student := Student{
        Name:   "Alice",
        Age:    20,
        Grades: []int{85, 90, 78, 92, 88},
    }
    
    fmt.Printf("Average: %.2f\n", student.AverageGrade())
    fmt.Printf("Passing: %t\n", student.IsPassing())
}
```

---

## Key Takeaways

1. **Structs group related data** - Like classes, but simpler
2. **Methods attach to structs** - Value or pointer receiver
3. **Composition over inheritance** - Embed structs for reuse
4. **Zero values** - All fields get default values
5. **Exported fields** - Capital letter = public
6. **Pointer receivers** - Use when modifying struct

---

## Must Remember Forever

- `type Name struct { }` - Define struct
- `p := Person{Name: "John"}` - Create instance
- `p.Field` - Access field
- `func (p Person) method()` - Value receiver
- `func (p *Person) method()` - Pointer receiver
- Embedded structs allow composition
- No inheritance - use composition instead

---

---

## k) Additional Practice Problems (Build Strong Logic)

### 10 More Solved Examples

### Solved Problem 1: Student Management System

**Task:** Create a student management system with structs.

**Solution:**
```go
package main

import "fmt"

type Student struct {
    ID      int
    Name    string
    Grades  []float64
    Average float64
}

func NewStudent(id int, name string) *Student {
    return &Student{
        ID:     id,
        Name:   name,
        Grades: []float64{},
    }
}

func (s *Student) AddGrade(grade float64) {
    s.Grades = append(s.Grades, grade)
    s.calculateAverage()
}

func (s *Student) calculateAverage() {
    if len(s.Grades) == 0 {
        s.Average = 0
        return
    }
    sum := 0.0
    for _, grade := range s.Grades {
        sum += grade
    }
    s.Average = sum / float64(len(s.Grades))
}

func (s *Student) GetStatus() string {
    if s.Average >= 90 {
        return "Excellent"
    } else if s.Average >= 80 {
        return "Good"
    } else if s.Average >= 70 {
        return "Average"
    }
    return "Needs Improvement"
}

func main() {
    student := NewStudent(1, "Alice")
    student.AddGrade(85)
    student.AddGrade(90)
    student.AddGrade(88)
    
    fmt.Printf("Student: %s\n", student.Name)
    fmt.Printf("Average: %.2f\n", student.Average)
    fmt.Printf("Status: %s\n", student.GetStatus())
}
```

### Solved Problem 2: Library Book System

**Task:** Create a library system with books and borrowing.

**Solution:**
```go
package main

import (
    "fmt"
    "time"
)

type Book struct {
    ISBN        string
    Title       string
    Author      string
    IsAvailable bool
    BorrowedBy  *string
    DueDate     *time.Time
}

type Library struct {
    Books []Book
}

func NewLibrary() *Library {
    return &Library{Books: []Book{}}
}

func (l *Library) AddBook(book Book) {
    l.Books = append(l.Books, book)
}

func (l *Library) BorrowBook(isbn string, borrower string) error {
    for i := range l.Books {
        if l.Books[i].ISBN == isbn {
            if !l.Books[i].IsAvailable {
                return fmt.Errorf("book %s is already borrowed", isbn)
            }
            l.Books[i].IsAvailable = false
            l.Books[i].BorrowedBy = &borrower
            dueDate := time.Now().Add(14 * 24 * time.Hour)
            l.Books[i].DueDate = &dueDate
            return nil
        }
    }
    return fmt.Errorf("book %s not found", isbn)
}

func (l *Library) ReturnBook(isbn string) {
    for i := range l.Books {
        if l.Books[i].ISBN == isbn {
            l.Books[i].IsAvailable = true
            l.Books[i].BorrowedBy = nil
            l.Books[i].DueDate = nil
        }
    }
}

func main() {
    lib := NewLibrary()
    lib.AddBook(Book{ISBN: "123", Title: "Go Guide", Author: "Author A", IsAvailable: true})
    lib.AddBook(Book{ISBN: "456", Title: "Python Basics", Author: "Author B", IsAvailable: true})
    
    lib.BorrowBook("123", "Alice")
    fmt.Println("Book borrowed successfully")
}
```

### Solved Problem 3: Bank Account with Transactions

**Task:** Create bank account with transaction history.

**Solution:**
```go
package main

import (
    "errors"
    "fmt"
    "time"
)

type Transaction struct {
    Type      string
    Amount    float64
    Timestamp time.Time
}

type BankAccount struct {
    AccountNumber string
    Balance       float64
    Transactions  []Transaction
}

func NewBankAccount(accountNumber string) *BankAccount {
    return &BankAccount{
        AccountNumber: accountNumber,
        Balance:       0,
        Transactions:  []Transaction{},
    }
}

func (ba *BankAccount) Deposit(amount float64) {
    ba.Balance += amount
    ba.Transactions = append(ba.Transactions, Transaction{
        Type:      "Deposit",
        Amount:    amount,
        Timestamp: time.Now(),
    })
}

func (ba *BankAccount) Withdraw(amount float64) error {
    if amount > ba.Balance {
        return errors.New("insufficient funds")
    }
    ba.Balance -= amount
    ba.Transactions = append(ba.Transactions, Transaction{
        Type:      "Withdrawal",
        Amount:    amount,
        Timestamp: time.Now(),
    })
    return nil
}

func (ba *BankAccount) GetStatement() {
    fmt.Printf("Account: %s\n", ba.AccountNumber)
    fmt.Printf("Balance: $%.2f\n\n", ba.Balance)
    fmt.Println("Transactions:")
    for _, txn := range ba.Transactions {
        fmt.Printf("  %s: $%.2f at %s\n", txn.Type, txn.Amount, txn.Timestamp.Format("2006-01-02 15:04:05"))
    }
}

func main() {
    account := NewBankAccount("ACC001")
    account.Deposit(1000)
    account.Deposit(500)
    account.Withdraw(200)
    account.GetStatement()
}
```

### Solved Problem 4: Shopping Cart System

**Task:** Create shopping cart with items and calculations.

**Solution:**
```go
package main

import "fmt"

type Item struct {
    ID       int
    Name     string
    Price    float64
    Quantity int
}

type ShoppingCart struct {
    Items []Item
}

func NewShoppingCart() *ShoppingCart {
    return &ShoppingCart{Items: []Item{}}
}

func (sc *ShoppingCart) AddItem(item Item) {
    // Check if item already exists
    for i := range sc.Items {
        if sc.Items[i].ID == item.ID {
            sc.Items[i].Quantity += item.Quantity
            return
        }
    }
    sc.Items = append(sc.Items, item)
}

func (sc *ShoppingCart) RemoveItem(itemID int) {
    for i := range sc.Items {
        if sc.Items[i].ID == itemID {
            sc.Items = append(sc.Items[:i], sc.Items[i+1:]...)
            return
        }
    }
}

func (sc *ShoppingCart) GetTotal() float64 {
    total := 0.0
    for _, item := range sc.Items {
        total += item.Price * float64(item.Quantity)
    }
    return total
}

func (sc *ShoppingCart) GetItemCount() int {
    count := 0
    for _, item := range sc.Items {
        count += item.Quantity
    }
    return count
}

func main() {
    cart := NewShoppingCart()
    cart.AddItem(Item{ID: 1, Name: "Laptop", Price: 999.99, Quantity: 1})
    cart.AddItem(Item{ID: 2, Name: "Mouse", Price: 29.99, Quantity: 2})
    
    fmt.Printf("Items in cart: %d\n", cart.GetItemCount())
    fmt.Printf("Total: $%.2f\n", cart.GetTotal())
}
```

### Solved Problem 5: Employee Management

**Task:** Create employee management system with hierarchy.

**Solution:**
```go
package main

import "fmt"

type Employee struct {
    ID       int
    Name     string
    Position string
    Salary   float64
    Manager  *Employee
    Team     []*Employee
}

func NewEmployee(id int, name, position string, salary float64) *Employee {
    return &Employee{
        ID:       id,
        Name:     name,
        Position: position,
        Salary:   salary,
        Team:     []*Employee{},
    }
}

func (e *Employee) SetManager(manager *Employee) {
    e.Manager = manager
    manager.Team = append(manager.Team, e)
}

func (e *Employee) GetTeamSize() int {
    return len(e.Team)
}

func (e *Employee) GetTotalTeamSalary() float64 {
    total := e.Salary
    for _, member := range e.Team {
        total += member.GetTotalTeamSalary()
    }
    return total
}

func main() {
    ceo := NewEmployee(1, "John", "CEO", 200000)
    manager := NewEmployee(2, "Alice", "Manager", 100000)
    employee1 := NewEmployee(3, "Bob", "Developer", 80000)
    employee2 := NewEmployee(4, "Charlie", "Developer", 75000)
    
    employee1.SetManager(manager)
    employee2.SetManager(manager)
    manager.SetManager(ceo)
    
    fmt.Printf("CEO team size: %d\n", ceo.GetTeamSize())
    fmt.Printf("Total team salary: $%.2f\n", ceo.GetTotalTeamSalary())
}
```

### Solved Problem 6: Product Catalog

**Task:** Create product catalog with categories.

**Solution:**
```go
package main

import "fmt"

type Product struct {
    ID          int
    Name        string
    Price       float64
    Category    string
    InStock     bool
    StockCount  int
}

type Catalog struct {
    Products []Product
}

func NewCatalog() *Catalog {
    return &Catalog{Products: []Product{}}
}

func (c *Catalog) AddProduct(product Product) {
    c.Products = append(c.Products, product)
}

func (c *Catalog) GetProductsByCategory(category string) []Product {
    var result []Product
    for _, product := range c.Products {
        if product.Category == category {
            result = append(result, product)
        }
    }
    return result
}

func (c *Catalog) GetAvailableProducts() []Product {
    var result []Product
    for _, product := range c.Products {
        if product.InStock && product.StockCount > 0 {
            result = append(result, product)
        }
    }
    return result
}

func (c *Catalog) UpdateStock(productID int, quantity int) {
    for i := range c.Products {
        if c.Products[i].ID == productID {
            c.Products[i].StockCount += quantity
            c.Products[i].InStock = c.Products[i].StockCount > 0
        }
    }
}

func main() {
    catalog := NewCatalog()
    catalog.AddProduct(Product{ID: 1, Name: "Laptop", Price: 999, Category: "Electronics", InStock: true, StockCount: 5})
    catalog.AddProduct(Product{ID: 2, Name: "Book", Price: 20, Category: "Books", InStock: true, StockCount: 10})
    
    electronics := catalog.GetProductsByCategory("Electronics")
    fmt.Printf("Electronics: %d products\n", len(electronics))
}
```

### Solved Problem 7: Time Tracking System

**Task:** Create time tracking system for tasks.

**Solution:**
```go
package main

import (
    "fmt"
    "time"
)

type TimeEntry struct {
    TaskName    string
    StartTime   time.Time
    EndTime     *time.Time
    Duration    time.Duration
}

type TimeTracker struct {
    Entries []TimeEntry
    Current *TimeEntry
}

func NewTimeTracker() *TimeTracker {
    return &TimeTracker{Entries: []TimeEntry{}}
}

func (tt *TimeTracker) StartTask(taskName string) {
    if tt.Current != nil {
        tt.StopTask()
    }
    tt.Current = &TimeEntry{
        TaskName:  taskName,
        StartTime: time.Now(),
    }
}

func (tt *TimeTracker) StopTask() {
    if tt.Current != nil {
        now := time.Now()
        tt.Current.EndTime = &now
        tt.Current.Duration = now.Sub(tt.Current.StartTime)
        tt.Entries = append(tt.Entries, *tt.Current)
        tt.Current = nil
    }
}

func (tt *TimeTracker) GetTotalTime() time.Duration {
    total := time.Duration(0)
    for _, entry := range tt.Entries {
        total += entry.Duration
    }
    return total
}

func main() {
    tracker := NewTimeTracker()
    tracker.StartTask("Coding")
    time.Sleep(100 * time.Millisecond) // Simulate work
    tracker.StopTask()
    
    fmt.Printf("Total time tracked: %v\n", tracker.GetTotalTime())
}
```

### Solved Problem 8: Contact Management

**Task:** Create contact management system.

**Solution:**
```go
package main

import "fmt"

type Contact struct {
    Name    string
    Email   string
    Phone   string
    Address string
}

type ContactBook struct {
    Contacts []Contact
}

func NewContactBook() *ContactBook {
    return &ContactBook{Contacts: []Contact{}}
}

func (cb *ContactBook) AddContact(contact Contact) {
    cb.Contacts = append(cb.Contacts, contact)
}

func (cb *ContactBook) FindContact(name string) *Contact {
    for i := range cb.Contacts {
        if cb.Contacts[i].Name == name {
            return &cb.Contacts[i]
        }
    }
    return nil
}

func (cb *ContactBook) UpdateContact(name string, updated Contact) bool {
    for i := range cb.Contacts {
        if cb.Contacts[i].Name == name {
            cb.Contacts[i] = updated
            return true
        }
    }
    return false
}

func (cb *ContactBook) DeleteContact(name string) bool {
    for i := range cb.Contacts {
        if cb.Contacts[i].Name == name {
            cb.Contacts = append(cb.Contacts[:i], cb.Contacts[i+1:]...)
            return true
        }
    }
    return false
}

func main() {
    book := NewContactBook()
    book.AddContact(Contact{Name: "Alice", Email: "alice@example.com", Phone: "123-456-7890"})
    
    contact := book.FindContact("Alice")
    if contact != nil {
        fmt.Printf("Found: %s - %s\n", contact.Name, contact.Email)
    }
}
```

### Solved Problem 9: Game Character System

**Task:** Create game character with stats and leveling.

**Solution:**
```go
package main

import "fmt"

type Character struct {
    Name      string
    Level     int
    Health    int
    MaxHealth int
    Attack    int
    Defense   int
    Experience int
}

func NewCharacter(name string) *Character {
    return &Character{
        Name:      name,
        Level:     1,
        Health:    100,
        MaxHealth: 100,
        Attack:    10,
        Defense:   5,
        Experience: 0,
    }
}

func (c *Character) TakeDamage(damage int) {
    actualDamage := damage - c.Defense
    if actualDamage < 0 {
        actualDamage = 0
    }
    c.Health -= actualDamage
    if c.Health < 0 {
        c.Health = 0
    }
}

func (c *Character) Heal(amount int) {
    c.Health += amount
    if c.Health > c.MaxHealth {
        c.Health = c.MaxHealth
    }
}

func (c *Character) GainExperience(exp int) {
    c.Experience += exp
    for c.Experience >= c.Level*100 {
        c.LevelUp()
    }
}

func (c *Character) LevelUp() {
    c.Level++
    c.MaxHealth += 20
    c.Health = c.MaxHealth
    c.Attack += 5
    c.Defense += 2
    c.Experience -= (c.Level - 1) * 100
}

func (c *Character) IsAlive() bool {
    return c.Health > 0
}

func main() {
    hero := NewCharacter("Hero")
    hero.GainExperience(150)
    fmt.Printf("Level: %d, Health: %d/%d\n", hero.Level, hero.Health, hero.MaxHealth)
}
```

### Solved Problem 10: Inventory Management

**Task:** Create inventory system with stock tracking.

**Solution:**
```go
package main

import "fmt"

type InventoryItem struct {
    ID          int
    Name        string
    Quantity    int
    MinQuantity int
    Price       float64
}

type Inventory struct {
    Items []InventoryItem
}

func NewInventory() *Inventory {
    return &Inventory{Items: []InventoryItem{}}
}

func (inv *Inventory) AddItem(item InventoryItem) {
    inv.Items = append(inv.Items, item)
}

func (inv *Inventory) UpdateQuantity(itemID int, quantity int) bool {
    for i := range inv.Items {
        if inv.Items[i].ID == itemID {
            inv.Items[i].Quantity += quantity
            return true
        }
    }
    return false
}

func (inv *Inventory) GetLowStockItems() []InventoryItem {
    var result []InventoryItem
    for _, item := range inv.Items {
        if item.Quantity <= item.MinQuantity {
            result = append(result, item)
        }
    }
    return result
}

func (inv *Inventory) GetTotalValue() float64 {
    total := 0.0
    for _, item := range inv.Items {
        total += float64(item.Quantity) * item.Price
    }
    return total
}

func main() {
    inv := NewInventory()
    inv.AddItem(InventoryItem{ID: 1, Name: "Widget", Quantity: 50, MinQuantity: 10, Price: 5.99})
    inv.AddItem(InventoryItem{ID: 2, Name: "Gadget", Quantity: 5, MinQuantity: 10, Price: 12.99})
    
    lowStock := inv.GetLowStockItems()
    fmt.Printf("Low stock items: %d\n", len(lowStock))
    fmt.Printf("Total inventory value: $%.2f\n", inv.GetTotalValue())
}
```

---

### 10 More Practice Problems (Solve These!)

**Problem 1: Movie Ticket System**
- Create Movie struct with: Title, ShowTime, SeatsAvailable, Price
- Create Booking struct with: Movie, CustomerName, Seats, TotalPrice
- Implement booking and cancellation

**Problem 2: Restaurant Menu System**
- Create MenuItem struct: Name, Category, Price, Ingredients
- Create Order struct: Items, Total, Status
- Implement order management

**Problem 3: Vehicle Rental System**
- Create Vehicle struct: ID, Type, Brand, DailyRate, IsAvailable
- Create Rental struct: Vehicle, Customer, StartDate, EndDate, TotalCost
- Calculate rental cost based on days

**Problem 4: Course Management**
- Create Course struct: Code, Name, Credits, Students
- Create Student struct: ID, Name, EnrolledCourses
- Implement enrollment and grade tracking

**Problem 5: Hotel Reservation System**
- Create Room struct: Number, Type, Price, IsAvailable
- Create Reservation struct: Room, Guest, CheckIn, CheckOut
- Handle room availability

**Problem 6: Task Management System**
- Create Task struct: ID, Title, Description, Priority, Status, DueDate
- Create Project struct: Name, Tasks, Progress
- Track task completion

**Problem 7: Social Media Post System**
- Create Post struct: ID, Author, Content, Likes, Comments, Timestamp
- Create Comment struct: Author, Content, Timestamp
- Implement like and comment functionality

**Problem 8: Library Management Enhanced**
- Extend library system with: Member management, Fines calculation, Book reservations

**Problem 9: E-commerce Order System**
- Create Order with: Items, ShippingAddress, PaymentMethod, Status
- Calculate totals, taxes, shipping costs

**Problem 10: Fitness Tracker**
- Create Workout struct: Type, Duration, Calories, Date
- Create Goal struct: Type, Target, Current, Deadline
- Track progress toward goals

---

## Next Steps

Now you understand structs. Next topics:
- Methods (more details)
- Interfaces (polymorphism)
- Error handling (Go's way)
- Concurrency (goroutines and channels)

**Continue learning! → Check [README.md](../README.md) for next topics**
