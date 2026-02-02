# Interfaces

## a) Overview

### What this topic is
Interfaces define behavior - they specify what methods a type must have. If a type has all the methods an interface requires, it automatically implements that interface (no explicit declaration needed).

### Why it exists in Go
Interfaces provide polymorphism - you can write code that works with any type that has certain methods. This is Go's way of achieving flexibility without inheritance.

### 🎯 Layman's Explanation (Simple Terms)

**Think of interfaces like a job description or a contract:**

**Real-world analogy - Job Description:**
- A job posting says: "Must be able to: Drive, Cook, Clean"
- It doesn't care WHO you are (chef, driver, cleaner)
- It only cares that you CAN DO these things
- If you can do all three, you qualify for the job!

**Another analogy - USB Port:**
- A USB port has a **contract**: "If you can plug in and transfer data, you work"
- It doesn't care if you're a mouse, keyboard, or flash drive
- As long as you follow the USB "interface" (can plug in and transfer data), you work
- This is **duck typing**: "If it looks like a duck and quacks like a duck, it's a duck!"

**Simple example:**
```
Interface requirement: "Must be able to make sound"

Dog has Bark() method → ✅ Implements interface
Cat has Meow() method → ✅ Implements interface  
Car has Honk() method → ✅ Implements interface

All can be used wherever "something that makes sound" is needed!
```

**Interface = A promise/contract:**
- Like a **recipe** that says "anything that can be mixed and baked works"
- Like a **power outlet** that works with any device that has the right plug
- Like a **job requirement** that says "anyone who can do X, Y, Z qualifies"

**Why use interfaces?**
1. **Flexibility**: Write code that works with many different types (like a universal charger)
2. **Simplicity**: Don't need to know the exact type, just that it can do certain things
3. **Testing**: Easy to create "fake" versions for testing (like a mock object)

**Key concept - "Duck Typing":**
- "If it walks like a duck and quacks like a duck, it's a duck"
- In Go: "If it has the methods the interface needs, it implements the interface"
- No need to say "I implement this interface" - Go figures it out automatically!

---

## b) Syntax

### Basic Interface
```go
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
```

### Empty Interface
```go
// interface{} can hold any type
var i interface{} = 42
var i interface{} = "hello"
var i interface{} = []int{1, 2, 3}

// Type assertion
value, ok := i.(int)
```

---

## c) Explanation

### Step-by-Step Interface Understanding

**1. Define interface**
```go
type Shape interface {
    Area() float64
}
```
- Interface specifies method signature
- No implementation (just signature)
- Any type with `Area() float64` implements Shape

**2. Implement interface (implicitly)**
```go
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}
```
- Circle has `Area() float64` method
- Circle automatically implements Shape
- No explicit "implements" keyword!

**3. Use interface**
```go
var s Shape = Circle{Radius: 5}
area := s.Area()
```
- Variable of interface type
- Can hold any type that implements interface
- Call methods through interface

### Characteristics

- **Implicit implementation**: No "implements" keyword
- **Duck typing**: "If it walks like a duck, it's a duck"
- **Composable**: Interfaces can embed other interfaces
- **Empty interface**: `interface{}` accepts any type
- **Type assertion**: Get concrete type from interface

---

## d) Python Comparison

### Python (Duck Typing)
```python
# No explicit interface
class Circle:
    def area(self):
        return 3.14159 * self.radius ** 2

class Rectangle:
    def area(self):
        return self.width * self.height

# Function accepts any object with area() method
def print_area(shape):
    print(shape.area())  # Duck typing - if it has area(), it works

print_area(Circle(5))
print_area(Rectangle(10, 5))
```

### Go (Interfaces)
```go
// Explicit interface
type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

// Function accepts interface
func printArea(s Shape) {
    fmt.Println(s.Area())
}

printArea(Circle{Radius: 5})
```

**Key Differences:**

| Feature | Python | Go |
|---------|--------|-----|
| **Interface definition** | No (duck typing) | Yes (explicit) |
| **Implementation** | Implicit | Implicit (but interface is explicit) |
| **Type checking** | Runtime | Compile time |
| **Polymorphism** | Duck typing | Interface-based |
| **Multiple interfaces** | N/A | Yes (type can implement many) |

**Thinking Difference:**
- Python: "If it has the method, it works" (runtime check)
- Go: "If it has the methods, it implements interface" (compile-time check)
- Python: Flexible but errors at runtime
- Go: Strict but catches errors early

---

## e) Visual Flow / Mental Model

### Interface Implementation

```
1. Define interface:
   type Shape interface {
       Area() float64
   }

2. Create type with method:
   type Circle struct { Radius float64 }
   func (c Circle) Area() float64 { ... }

3. Circle automatically implements Shape!
   (No explicit declaration needed)

4. Use interface:
   var s Shape = Circle{Radius: 5}
   area := s.Area()
```

### Interface Variable

```
var s Shape = Circle{Radius: 5}

Memory:
  [s: Shape interface]
    ├── Type: Circle
    ├── Value: {Radius: 5}
    └── Methods: Area(), Perimeter()

When calling s.Area():
  1. Go looks at s's type (Circle)
  2. Calls Circle's Area() method
  3. Returns result
```

### Type Assertion

```
var i interface{} = 42

value, ok := i.(int)
  ↓
1. Check if i holds int
   ↓
2. If yes: value = 42, ok = true
   ↓
3. If no: value = zero value, ok = false
```

---

## f) Demo Example

### Complete Example

```go
package main

import (
    "fmt"
    "math"
)

// 1. Define interface
type Shape interface {
    Area() float64
    Perimeter() float64
}

// 2. Implement with Circle
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// 3. Implement with Rectangle
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

// 4. Function that accepts interface
func printShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n",
        s.Area(), s.Perimeter())
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
        r.Width, r.Height)
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
    rect := Rectangle{Width: 10, Height: 5}
    
    // Use as interface
    var s1 Shape = circle
    var s2 Shape = rect
    
    fmt.Println("Shape 1:")
    printShapeInfo(s1)
    
    fmt.Println("Shape 2:")
    printShapeInfo(s2)
    
    // Stringer interface (built-in)
    fmt.Println("String representation:")
    fmt.Println(circle)
    fmt.Println(rect)
    
    // Empty interface
    fmt.Println("\nEmpty interface:")
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
```

**Line-by-line explanation:**

1. **Interface definition**: Specifies required methods
2. **Implicit implementation**: Circle has required methods, so it implements Shape
3. **Multiple implementations**: Rectangle also implements Shape
4. **Interface as parameter**: Function accepts any Shape
5. **Built-in interfaces**: Stringer is a common interface
6. **Empty interface**: `interface{}` accepts any type
7. **Type assertion**: Get concrete type from interface

**Output:**
```
Shape 1:
Area: 78.54, Perimeter: 31.42
Shape 2:
Area: 50.00, Perimeter: 30.00
String representation:
Circle(radius=5.00)
Rectangle(width=10.00, height=5.00)

Empty interface:
Value: 42, Type: int
Value: hello, Type: string
Value: {5}, Type: main.Circle

Type assertion:
Integer: 42
String: hello
Circle with radius: 5.00
Unknown type: float64
It's an int: 42
```

---

## g) Use Cases

### When to use interfaces

**1. Polymorphism**
```go
type Writer interface {
    Write([]byte) (int, error)
}

// Works with any Writer (file, network, buffer, etc.)
func writeData(w Writer, data []byte) {
    w.Write(data)
}
```

**2. Testing**
```go
type Database interface {
    GetUser(id int) (*User, error)
}

// Can use real database or mock in tests
```

**3. Plugin architecture**
```go
type Plugin interface {
    Execute() error
}

// Different plugins implement same interface
```

**4. Common behavior**
```go
type Stringer interface {
    String() string
}

// Any type with String() method
```

---

## h) Do's and Don'ts / Best Practices

### ✅ Do's

1. **Keep interfaces small**
   ```go
   type Reader interface {
       Read([]byte) (int, error)  // ✅ One method
   }
   ```

2. **Accept interfaces, return structs**
   ```go
   func process(r Reader) { }  // ✅ Accept interface
   func create() *MyStruct { }  // ✅ Return concrete type
   ```

3. **Use interface{} sparingly**
   ```go
   // ✅ Prefer specific interface
   func process(s Stringer) { }
   
   // ❌ Avoid if possible
   func process(i interface{}) { }
   ```

4. **Check type assertions**
   ```go
   if value, ok := i.(int); ok {
       // use value
   }
   ```

### ❌ Don'ts

1. **Don't create interfaces before you need them**
   ```go
   // ❌ Premature
   type MyInterface interface { ... }
   
   // ✅ Create when you have multiple implementations
   ```

2. **Don't make interfaces too large**
   ```go
   // ❌ Too many methods
   type Everything interface {
       Method1()
       Method2()
       // ... 20 methods
   }
   ```

3. **Don't ignore type assertion errors**
   ```go
   value := i.(int)  // ❌ Can panic!
   value, ok := i.(int)  // ✅ Check ok
   ```

---

## i) Solved Practice Examples

### Example 1: Animal Interface

**Task:** Create Animal interface with Speak() method. Implement with Dog and Cat.

**Solution:**
```go
package main

import "fmt"

type Animal interface {
    Speak() string
}

type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return fmt.Sprintf("%s says: Woof!", d.Name)
}

type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return fmt.Sprintf("%s says: Meow!", c.Name)
}

func makeAnimalSpeak(a Animal) {
    fmt.Println(a.Speak())
}

func main() {
    dog := Dog{"Buddy"}
    cat := Cat{"Whiskers"}
    
    makeAnimalSpeak(dog)
    makeAnimalSpeak(cat)
}
```

### Example 2: Calculator Interface

**Task:** Create Calculator interface. Implement with BasicCalculator.

**Solution:**
```go
package main

import "fmt"

type Calculator interface {
    Add(a, b float64) float64
    Subtract(a, b float64) float64
    Multiply(a, b float64) float64
    Divide(a, b float64) (float64, error)
}

type BasicCalculator struct{}

func (c BasicCalculator) Add(a, b float64) float64 {
    return a + b
}

func (c BasicCalculator) Subtract(a, b float64) float64 {
    return a - b
}

func (c BasicCalculator) Multiply(a, b float64) float64 {
    return a * b
}

func (c BasicCalculator) Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

func main() {
    var calc Calculator = BasicCalculator{}
    
    fmt.Println("Add:", calc.Add(10, 5))
    fmt.Println("Subtract:", calc.Subtract(10, 5))
    fmt.Println("Multiply:", calc.Multiply(10, 5))
    
    result, err := calc.Divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Divide:", result)
    }
}
```

---

## j) Quiz / Practice for Me

### Quiz Questions

1. **How do you implement an interface in Go?**
   - [ ] Use "implements" keyword
   - [ ] Just have the required methods
   - [ ] Inherit from interface
   - [ ] Register with interface

2. **What is an empty interface?**
   - [ ] Interface with no methods
   - [ ] interface{}
   - [ ] Accepts any type
   - [ ] All of the above

3. **When is type assertion checked?**
   - [ ] Compile time
   - [ ] Runtime
   - [ ] Never
   - [ ] Both

### Practice Tasks

**Task 1: Drawable Interface**
- Create Drawable interface with Draw() method
- Implement with Circle and Square
- Function that draws any Drawable

**Task 2: Storage Interface**
- Create Storage interface with Save() and Load() methods
- Implement with MemoryStorage (uses map)
- Test save and load operations

### Answers

**Quiz Answers:**
1. Just have the required methods
2. All of the above
3. Runtime

**Practice Solutions:**

**Task 1 Solution:**
```go
package main

import "fmt"

type Drawable interface {
    Draw()
}

type Circle struct {
    Radius float64
}

func (c Circle) Draw() {
    fmt.Printf("Drawing circle with radius %.2f\n", c.Radius)
}

type Square struct {
    Side float64
}

func (s Square) Draw() {
    fmt.Printf("Drawing square with side %.2f\n", s.Side)
}

func drawShape(d Drawable) {
    d.Draw()
}

func main() {
    circle := Circle{Radius: 5}
    square := Square{Side: 10}
    
    drawShape(circle)
    drawShape(square)
}
```

**Task 2 Solution:**
```go
package main

import (
    "errors"
    "fmt"
)

type Storage interface {
    Save(key string, value string) error
    Load(key string) (string, error)
}

type MemoryStorage struct {
    data map[string]string
}

func NewMemoryStorage() *MemoryStorage {
    return &MemoryStorage{
        data: make(map[string]string),
    }
}

func (m *MemoryStorage) Save(key string, value string) error {
    m.data[key] = value
    return nil
}

func (m *MemoryStorage) Load(key string) (string, error) {
    value, exists := m.data[key]
    if !exists {
        return "", errors.New("key not found")
    }
    return value, nil
}

func main() {
    var storage Storage = NewMemoryStorage()
    
    storage.Save("name", "Alice")
    storage.Save("age", "30")
    
    name, _ := storage.Load("name")
    age, _ := storage.Load("age")
    
    fmt.Printf("Name: %s, Age: %s\n", name, age)
}
```

---

## Key Takeaways

1. **Interfaces define behavior** - What methods a type must have
2. **Implicit implementation** - No "implements" keyword needed
3. **Duck typing** - "If it has the methods, it implements"
4. **Empty interface** - `interface{}` accepts any type
5. **Type assertion** - Get concrete type from interface
6. **Polymorphism** - Write code that works with multiple types

---

## Must Remember Forever

- `type InterfaceName interface { Method() }` - Define interface
- Any type with required methods implements interface (implicit)
- `interface{}` - Empty interface, accepts any type
- `value, ok := i.(Type)` - Type assertion with check
- Accept interfaces, return structs (common pattern)
- Keep interfaces small (1-3 methods usually)

---

---

## k) Additional Practice Problems (Build Strong Logic)

### 10 More Solved Examples

### Solved Problem 1: Payment Processor Interface

**Task:** Create payment processor interface with multiple implementations.

**Solution:**
```go
package main

import "fmt"

type PaymentProcessor interface {
    ProcessPayment(amount float64) error
    Refund(transactionID string) error
}

type CreditCardProcessor struct {
    CardNumber string
}

func (c *CreditCardProcessor) ProcessPayment(amount float64) error {
    fmt.Printf("Processing $%.2f with credit card %s\n", amount, c.CardNumber)
    return nil
}

func (c *CreditCardProcessor) Refund(transactionID string) error {
    fmt.Printf("Refunding transaction %s\n", transactionID)
    return nil
}

type PayPalProcessor struct {
    Email string
}

func (p *PayPalProcessor) ProcessPayment(amount float64) error {
    fmt.Printf("Processing $%.2f via PayPal (%s)\n", amount, p.Email)
    return nil
}

func (p *PayPalProcessor) Refund(transactionID string) error {
    fmt.Printf("PayPal refund for %s\n", transactionID)
    return nil
}

func processPayment(processor PaymentProcessor, amount float64) {
    processor.ProcessPayment(amount)
}

func main() {
    cc := &CreditCardProcessor{CardNumber: "1234-5678"}
    pp := &PayPalProcessor{Email: "user@example.com"}
    
    processPayment(cc, 100.0)
    processPayment(pp, 50.0)
}
```

### Solved Problem 2: Logger Interface

**Task:** Create logger interface with multiple implementations.

**Solution:**
```go
package main

import "fmt"

type Logger interface {
    Log(level string, message string)
    Error(message string)
    Info(message string)
}

type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(level, message string) {
    fmt.Printf("[%s] %s\n", level, message)
}

func (c *ConsoleLogger) Error(message string) {
    c.Log("ERROR", message)
}

func (c *ConsoleLogger) Info(message string) {
    c.Log("INFO", message)
}

type FileLogger struct {
    filename string
}

func (f *FileLogger) Log(level, message string) {
    fmt.Printf("Writing to %s: [%s] %s\n", f.filename, level, message)
}

func (f *FileLogger) Error(message string) {
    f.Log("ERROR", message)
}

func (f *FileLogger) Info(message string) {
    f.Log("INFO", message)
}

func logMessage(logger Logger, message string) {
    logger.Info(message)
}

func main() {
    consoleLogger := &ConsoleLogger{}
    fileLogger := &FileLogger{filename: "app.log"}
    
    logMessage(consoleLogger, "Application started")
    logMessage(fileLogger, "Application started")
}
```

### Solved Problem 3: Shape Calculator Interface

**Task:** Create shape interface with area and perimeter calculations.

**Solution:**
```go
package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
    Perimeter() float64
}

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

type Triangle struct {
    A, B, C float64
}

func (t Triangle) Area() float64 {
    s := t.Perimeter() / 2
    return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
    return t.A + t.B + t.C
}

func printShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    triangle := Triangle{A: 3, B: 4, C: 5}
    
    printShapeInfo(rect)
    printShapeInfo(triangle)
}
```

### Solved Problem 4: Storage Interface

**Task:** Create storage interface with in-memory and file implementations.

**Solution:**
```go
package main

import (
    "errors"
    "fmt"
)

type Storage interface {
    Save(key string, value string) error
    Load(key string) (string, error)
    Delete(key string) error
}

type MemoryStorage struct {
    data map[string]string
}

func NewMemoryStorage() *MemoryStorage {
    return &MemoryStorage{data: make(map[string]string)}
}

func (m *MemoryStorage) Save(key, value string) error {
    m.data[key] = value
    return nil
}

func (m *MemoryStorage) Load(key string) (string, error) {
    if value, exists := m.data[key]; exists {
        return value, nil
    }
    return "", errors.New("key not found")
}

func (m *MemoryStorage) Delete(key string) error {
    delete(m.data, key)
    return nil
}

type FileStorage struct {
    filename string
}

func NewFileStorage(filename string) *FileStorage {
    return &FileStorage{filename: filename}
}

func (f *FileStorage) Save(key, value string) error {
    fmt.Printf("Saving to %s: %s=%s\n", f.filename, key, value)
    return nil
}

func (f *FileStorage) Load(key string) (string, error) {
    fmt.Printf("Loading from %s: key=%s\n", f.filename, key)
    return "value", nil
}

func (f *FileStorage) Delete(key string) error {
    fmt.Printf("Deleting from %s: key=%s\n", f.filename, key)
    return nil
}

func useStorage(storage Storage) {
    storage.Save("name", "Alice")
    value, _ := storage.Load("name")
    fmt.Printf("Loaded: %s\n", value)
}

func main() {
    memStorage := NewMemoryStorage()
    fileStorage := NewFileStorage("data.txt")
    
    useStorage(memStorage)
    useStorage(fileStorage)
}
```

### Solved Problem 5: Sortable Interface

**Task:** Create interface for sortable types.

**Solution:**
```go
package main

import (
    "fmt"
    "sort"
)

type Sortable interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}

type IntSlice []int

func (is IntSlice) Len() int {
    return len(is)
}

func (is IntSlice) Less(i, j int) bool {
    return is[i] < is[j]
}

func (is IntSlice) Swap(i, j int) {
    is[i], is[j] = is[j], is[i]
}

type StringSlice []string

func (ss StringSlice) Len() int {
    return len(ss)
}

func (ss StringSlice) Less(i, j int) bool {
    return ss[i] < ss[j]
}

func (ss StringSlice) Swap(i, j int) {
    ss[i], ss[j] = ss[j], ss[i]
}

func sortData(s Sortable) {
    sort.Sort(s)
}

func main() {
    ints := IntSlice{3, 1, 4, 1, 5, 9, 2, 6}
    strings := StringSlice{"banana", "apple", "cherry"}
    
    sortData(ints)
    sortData(strings)
    
    fmt.Printf("Sorted ints: %v\n", ints)
    fmt.Printf("Sorted strings: %v\n", strings)
}
```

### Solved Problem 6: Reader Interface Implementation

**Task:** Implement io.Reader interface.

**Solution:**
```go
package main

import (
    "fmt"
    "io"
)

type StringReader struct {
    s string
    i int
}

func NewStringReader(s string) *StringReader {
    return &StringReader{s: s, i: 0}
}

func (sr *StringReader) Read(p []byte) (n int, err error) {
    if sr.i >= len(sr.s) {
        return 0, io.EOF
    }
    n = copy(p, sr.s[sr.i:])
    sr.i += n
    return n, nil
}

func readAll(r io.Reader) ([]byte, error) {
    result := make([]byte, 0)
    buf := make([]byte, 4)
    
    for {
        n, err := r.Read(buf)
        if n > 0 {
            result = append(result, buf[:n]...)
        }
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, err
        }
    }
    return result, nil
}

func main() {
    reader := NewStringReader("Hello, World!")
    data, _ := readAll(reader)
    fmt.Printf("Read: %s\n", string(data))
}
```

### Solved Problem 7: Comparable Interface

**Task:** Create interface for comparable types.

**Solution:**
```go
package main

import "fmt"

type Comparable interface {
    Compare(other Comparable) int // -1: less, 0: equal, 1: greater
}

type Person struct {
    Age int
}

func (p Person) Compare(other Comparable) int {
    otherPerson := other.(Person)
    if p.Age < otherPerson.Age {
        return -1
    } else if p.Age > otherPerson.Age {
        return 1
    }
    return 0
}

func findMax(items []Comparable) Comparable {
    if len(items) == 0 {
        return nil
    }
    max := items[0]
    for _, item := range items[1:] {
        if item.Compare(max) > 0 {
            max = item
        }
    }
    return max
}

func main() {
    people := []Comparable{
        Person{Age: 25},
        Person{Age: 30},
        Person{Age: 20},
    }
    max := findMax(people)
    fmt.Printf("Oldest person: %+v\n", max)
}
```

### Solved Problem 8: Cloner Interface

**Task:** Create interface for clonable objects.

**Solution:**
```go
package main

import "fmt"

type Cloner interface {
    Clone() Cloner
}

type Document struct {
    Title   string
    Content string
}

func (d *Document) Clone() Cloner {
    return &Document{
        Title:   d.Title,
        Content: d.Content,
    }
}

func cloneMultiple(original Cloner, count int) []Cloner {
    clones := make([]Cloner, count)
    for i := 0; i < count; i++ {
        clones[i] = original.Clone()
    }
    return clones
}

func main() {
    doc := &Document{Title: "Original", Content: "Content"}
    clones := cloneMultiple(doc, 3)
    
    for i, clone := range clones {
        fmt.Printf("Clone %d: %+v\n", i+1, clone)
    }
}
```

### Solved Problem 9: Validator Interface

**Task:** Create validator interface for different validation types.

**Solution:**
```go
package main

import "fmt"

type Validator interface {
    Validate(value interface{}) error
}

type EmailValidator struct{}

func (e *EmailValidator) Validate(value interface{}) error {
    email := value.(string)
    if len(email) == 0 || !contains(email, "@") {
        return fmt.Errorf("invalid email: %s", email)
    }
    return nil
}

type AgeValidator struct {
    Min, Max int
}

func (a *AgeValidator) Validate(value interface{}) error {
    age := value.(int)
    if age < a.Min || age > a.Max {
        return fmt.Errorf("age %d out of range [%d, %d]", age, a.Min, a.Max)
    }
    return nil
}

func contains(s, substr string) bool {
    for i := 0; i <= len(s)-len(substr); i++ {
        if s[i:i+len(substr)] == substr {
            return true
        }
    }
    return false
}

func validate(validator Validator, value interface{}) error {
    return validator.Validate(value)
}

func main() {
    emailValidator := &EmailValidator{}
    ageValidator := &AgeValidator{Min: 18, Max: 100}
    
    validate(emailValidator, "user@example.com")
    validate(ageValidator, 25)
}
```

### Solved Problem 10: Factory Pattern with Interface

**Task:** Use interface with factory pattern.

**Solution:**
```go
package main

import "fmt"

type Animal interface {
    Speak() string
    Move() string
}

type Dog struct{}

func (d *Dog) Speak() string {
    return "Woof!"
}

func (d *Dog) Move() string {
    return "Running"
}

type Cat struct{}

func (c *Cat) Speak() string {
    return "Meow!"
}

func (c *Cat) Move() string {
    return "Walking"
}

func CreateAnimal(animalType string) Animal {
    switch animalType {
    case "dog":
        return &Dog{}
    case "cat":
        return &Cat{}
    default:
        return nil
    }
}

func main() {
    dog := CreateAnimal("dog")
    cat := CreateAnimal("cat")
    
    fmt.Printf("Dog: %s, %s\n", dog.Speak(), dog.Move())
    fmt.Printf("Cat: %s, %s\n", cat.Speak(), cat.Move())
}
```

---

### 10 More Practice Problems (Solve These!)

**Problem 1: Database Interface**
- Create Database interface: Connect, Query, Execute, Close
- Implement: MySQLDatabase, PostgreSQLDatabase

**Problem 2: Renderer Interface**
- Create Renderer interface: Render(data interface{}) string
- Implement: JSONRenderer, XMLRenderer, HTMLRenderer

**Problem 3: Cache Interface**
- Create Cache interface: Get, Set, Delete, Clear
- Implement: MemoryCache, FileCache

**Problem 4: Notifier Interface**
- Create Notifier interface: Send(message string) error
- Implement: EmailNotifier, SMSNotifier, PushNotifier

**Problem 5: Parser Interface**
- Create Parser interface: Parse(data []byte) (interface{}, error)
- Implement: JSONParser, XMLParser, CSVParser

**Problem 6: Encoder Interface**
- Create Encoder interface: Encode(data interface{}) ([]byte, error)
- Implement: Base64Encoder, JSONEncoder

**Problem 7: Filter Interface**
- Create Filter interface: Filter(items []interface{}) []interface{}
- Implement: EvenFilter, PositiveFilter, StringFilter

**Problem 8: Transformer Interface**
- Create Transformer interface: Transform(input interface{}) interface{}
- Implement: UppercaseTransformer, ReverseTransformer

**Problem 9: Repository Interface**
- Create Repository interface: Create, Read, Update, Delete
- Implement: UserRepository, ProductRepository

**Problem 10: Strategy Pattern**
- Create Strategy interface: Execute()
- Implement multiple strategies and choose at runtime

---

## Next Steps

Now you understand interfaces. Next:
- Error handling (Go's unique approach)
- How errors work in Go
- Best practices for error handling

**Ready? → [10_ERROR_HANDLING.md](./10_ERROR_HANDLING.md)**
