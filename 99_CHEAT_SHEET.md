# Complete Golang Cheat Sheet
## Quick Reference for Python Developers

---

## 📋 Table of Contents

1. [Basic Syntax](#basic-syntax)
2. [Variables & Types](#variables--types)
3. [Functions](#functions)
4. [Control Flow](#control-flow)
5. [Collections](#collections)
6. [Structs & Methods](#structs--methods)
7. [Interfaces](#interfaces)
8. [Error Handling](#error-handling)
9. [Concurrency](#concurrency)
10. [Packages & Modules](#packages--modules)
11. [Common Patterns](#common-patterns)
12. [Python vs Go Quick Reference](#python-vs-go-quick-reference)

---

## Basic Syntax

### Program Structure
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

### Comments
```go
// Single line comment

/* Multi-line
   comment */
```

### Package Declaration
```go
package main        // Executable program
package mypackage   // Library package
```

### Imports
```go
import "fmt"                    // Single import
import "os"
import "strings"

import (                         // Multiple imports
    "fmt"
    "os"
    "strings"
)

import f "fmt"                  // Import with alias
import _ "database/sql"         // Blank import (side effects)
```

---

## Variables & Types

### Variable Declaration
```go
// Explicit type
var name string = "John"

// Type inference
var age = 25

// Short declaration (inside functions only)
name := "John"
age := 25

// Multiple variables
var x, y int = 1, 2
a, b := 10, 20

// Package-level (must use var)
var GlobalCount int
```

### Constants
```go
const pi = 3.14159
const greeting = "Hello"

const (
    maxUsers = 100
    appName  = "MyApp"
)
```

### Basic Types
```go
// Numbers
var i int = 42           // int (32 or 64 bit, depends on system)
var i8 int8 = 127        // 8-bit integer
var i16 int16 = 32767    // 16-bit integer
var i32 int32 = 2147483647
var i64 int64 = 9223372036854775807

var ui uint = 42         // Unsigned integer
var ui8 uint8 = 255      // Also called byte
var ui16 uint16 = 65535
var ui32 uint32          // Also called rune
var ui64 uint64

var f32 float32 = 3.14
var f64 float64 = 3.14159

var c64 complex64 = 1 + 2i
var c128 complex128

// Strings
var s string = "Hello"
var raw string = `Raw string
can span
multiple lines`

// Booleans
var b bool = true

// Byte and Rune
var by byte = 'A'        // Alias for uint8
var r rune = '中'         // Alias for int32 (Unicode)
```

### Zero Values
```go
var i int        // 0
var f float64    // 0.0
var s string     // ""
var b bool       // false
var p *int       // nil
var sl []int     // nil
var m map[string]int  // nil
```

### Type Conversion
```go
var i int = 42
var f float64 = float64(i)      // int to float64
var s string = string(i)        // int to string (Unicode!)
var s2 string = fmt.Sprintf("%d", i)  // Better for numbers

var f2 float64 = 3.14
var i2 int = int(f2)            // float64 to int (truncates)
```

---

## Functions

### Basic Function
```go
func functionName() {
    // code
}

func functionName(param type) {
    // code
}

func functionName() returnType {
    return value
}

func functionName(param type) returnType {
    return value
}
```

### Multiple Return Values
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Usage
result, err := divide(10, 2)
if err != nil {
    // handle error
}
```

### Named Return Values
```go
func calculate(x, y int) (sum int, product int) {
    sum = x + y
    product = x * y
    return  // Naked return
}
```

### Variadic Functions
```go
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

// Usage
total := sum(1, 2, 3, 4, 5)
```

### Function as Value
```go
var add = func(a, b int) int {
    return a + b
}

// Usage
result := add(5, 3)
```

### Function Types
```go
type Operation func(int, int) int

func apply(op Operation, a, b int) int {
    return op(a, b)
}
```

---

## Control Flow

### If/Else
```go
if condition {
    // code
}

if condition {
    // code
} else {
    // code
}

if condition1 {
    // code
} else if condition2 {
    // code
} else {
    // code
}

// If with initialization
if x := getValue(); x > 0 {
    // x available here
}
```

### For Loops
```go
// Traditional
for i := 0; i < 10; i++ {
    // code
}

// While-style
for condition {
    // code
}

// Infinite
for {
    // code
    break
}

// Range (slices)
for index, value := range slice {
    // code
}

// Range (ignore index)
for _, value := range slice {
    // code
}

// Range (maps)
for key, value := range map {
    // code
}

// Range (strings - gives runes)
for i, char := range "Hello" {
    // code
}
```

### Switch
```go
// Switch on value
switch value {
case option1:
    // code
case option2:
    // code
default:
    // code
}

// Multiple cases
switch value {
case 1, 2, 3:
    // code
}

// Switch on condition
switch {
case condition1:
    // code
case condition2:
    // code
}

// Switch with initialization
switch x := getValue(); x {
case 1:
    // code
}
```

### Break & Continue
```go
for {
    if condition {
        break      // Exit loop
    }
    if skipCondition {
        continue   // Skip to next iteration
    }
}

// Labeled break/continue
OuterLoop:
    for i := 0; i < 10; i++ {
        for j := 0; j < 10; j++ {
            if condition {
                break OuterLoop  // Break outer loop
            }
        }
    }
```

---

## Collections

### Arrays
```go
// Fixed size
var arr [5]int                    // [0, 0, 0, 0, 0]
arr := [5]int{1, 2, 3, 4, 5}      // [1, 2, 3, 4, 5]
arr := [...]int{1, 2, 3}          // Size inferred

// Access
arr[0] = 10
value := arr[0]

// Length
len(arr)
```

### Slices (Dynamic Arrays)
```go
// Create
var s []int                       // nil slice
s := []int{1, 2, 3}               // [1, 2, 3]
s := make([]int, 5)               // [0, 0, 0, 0, 0]
s := make([]int, 0, 10)           // Length 0, capacity 10

// Append
s = append(s, 4)                  // Add element
s = append(s, 5, 6, 7)            // Add multiple

// Slice operations
s[1:3]                            // Slice from index 1 to 3 (exclusive)
s[:3]                             // From start to 3
s[2:]                             // From 2 to end
s[:]                              // Copy of entire slice

// Length and capacity
len(s)
cap(s)

// Copy
dest := make([]int, len(s))
copy(dest, s)
```

### Maps
```go
// Create
var m map[string]int              // nil map
m := make(map[string]int)         // Empty map
m := map[string]int{              // With initial values
    "apple": 5,
    "banana": 3,
}

// Access
value := m["apple"]               // Get value
value, exists := m["apple"]       // Check existence

// Set
m["orange"] = 10

// Delete
delete(m, "apple")

// Iterate
for key, value := range m {
    // code
}

// Length
len(m)
```

### Strings
```go
// String operations
s := "Hello"
len(s)                            // Length in bytes
s[0]                              // First byte
s[1:3]                            // Substring

// String conversion
s := string([]byte{'H', 'e', 'l', 'l', 'o'})
b := []byte("Hello")

// String formatting
fmt.Sprintf("Name: %s, Age: %d", name, age)
```

---

## Structs & Methods

### Structs
```go
// Define
type Person struct {
    Name string
    Age  int
}

// Create
p := Person{"John", 30}          // Positional
p := Person{Name: "John", Age: 30}  // Named fields
p := Person{}                     // Zero values

// Access
p.Name = "Alice"
age := p.Age

// Pointer to struct
ptr := &Person{"Bob", 25}
ptr.Name = "Robert"              // Go automatically dereferences

// Anonymous struct
p := struct {
    Name string
    Age  int
}{"John", 30}
```

### Methods
```go
// Value receiver
func (p Person) Greet() string {
    return fmt.Sprintf("Hello, I'm %s", p.Name)
}

// Pointer receiver (can modify)
func (p *Person) HaveBirthday() {
    p.Age++
}

// Usage
p := Person{"Alice", 30}
greeting := p.Greet()
p.HaveBirthday()
```

### Embedded Structs
```go
type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person                      // Embedded
    EmployeeID int
}

// Usage
e := Employee{
    Person: Person{"John", 30},
    EmployeeID: 123,
}
e.Name                         // Access embedded field
```

---

## Interfaces

### Basic Interface
```go
// Define
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Implement (implicit)
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
// Can hold any type
var i interface{} = 42
var i interface{} = "hello"
var i interface{} = []int{1, 2, 3}

// Type assertion
value, ok := i.(int)
if ok {
    // value is int
}

// Type switch
switch v := i.(type) {
case int:
    // v is int
case string:
    // v is string
default:
    // unknown type
}
```

---

## Error Handling

### Error Pattern
```go
// Standard error handling
result, err := doSomething()
if err != nil {
    // handle error
    return err
}

// Create errors
import "errors"
err := errors.New("something went wrong")

// Formatted errors
import "fmt"
err := fmt.Errorf("error: %v", value)

// Custom error type
type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
```

### Common Patterns
```go
// Check and return
if err != nil {
    return nil, err
}

// Check and log
if err != nil {
    log.Printf("Error: %v", err)
    return
}

// Check and panic (rare, only for unrecoverable)
if err != nil {
    panic(err)
}

// Defer for cleanup
func doSomething() error {
    file, err := os.Open("file.txt")
    if err != nil {
        return err
    }
    defer file.Close()  // Always closes
    
    // use file
    return nil
}
```

---

## Concurrency

### Goroutines
```go
// Start goroutine
go functionName()

// Anonymous function
go func() {
    // code
}()

// With parameters
go func(x int) {
    // code
}(10)
```

### Channels
```go
// Create
ch := make(chan int)             // Unbuffered
ch := make(chan int, 10)         // Buffered (capacity 10)

// Send
ch <- value

// Receive
value := <-ch
value, ok := <-ch                // Check if channel closed

// Close
close(ch)

// Range over channel
for value := range ch {
    // code
}

// Select (like switch for channels)
select {
case msg := <-ch1:
    // received from ch1
case ch2 <- value:
    // sent to ch2
case <-time.After(1 * time.Second):
    // timeout
default:
    // non-blocking
}
```

### WaitGroup
```go
import "sync"

var wg sync.WaitGroup

wg.Add(1)                        // Add goroutine
go func() {
    defer wg.Done()              // Mark done
    // work
}()
wg.Wait()                        // Wait for all
```

### Mutex
```go
import "sync"

var mu sync.Mutex
var counter int

mu.Lock()
counter++
mu.Unlock()

// Or use defer
mu.Lock()
defer mu.Unlock()
counter++
```

---

## Packages & Modules

### Module Setup
```bash
go mod init myproject
go mod tidy
go get package@version
go build
go run
```

### Package Structure
```
myproject/
├── go.mod
├── go.sum
├── main.go
└── pkg/
    └── utils/
        └── helper.go
```

### Exports
```go
// Exported (capital letter)
func PublicFunction() { }

// Unexported (lowercase)
func privateFunction() { }

// Same for types, variables, constants
type PublicType struct { }
type privateType struct { }
```

### Importing
```go
import "myproject/pkg/utils"
utils.PublicFunction()
```

---

## Common Patterns

### JSON Handling
```go
import "encoding/json"

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// Marshal (struct to JSON)
p := Person{"John", 30}
data, err := json.Marshal(p)

// Unmarshal (JSON to struct)
var p Person
err := json.Unmarshal(data, &p)
```

### File I/O
```go
import "os"
import "io/ioutil"

// Read entire file
data, err := ioutil.ReadFile("file.txt")

// Write file
err := ioutil.WriteFile("file.txt", data, 0644)

// Open file
file, err := os.Open("file.txt")
defer file.Close()

// Read line by line
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
}
```

### HTTP Server
```go
import "net/http"

http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
})

http.ListenAndServe(":8080", nil)
```

### Testing
```go
import "testing"

func TestFunction(t *testing.T) {
    result := function(5)
    if result != 10 {
        t.Errorf("Expected 10, got %d", result)
    }
}

// Run tests
// go test
// go test -v
// go test ./...
```

---

## Python vs Go Quick Reference

### Variables
| Python | Go |
|--------|-----|
| `name = "John"` | `name := "John"` |
| `age = 25` | `age := 25` |
| `x, y = 1, 2` | `x, y := 1, 2` |

### Functions
| Python | Go |
|--------|-----|
| `def add(a, b): return a + b` | `func add(a, b int) int { return a + b }` |
| `def divide(a, b): return a/b, None` | `func divide(a, b float64) (float64, error)` |

### Lists/Arrays
| Python | Go |
|--------|-----|
| `items = [1, 2, 3]` | `items := []int{1, 2, 3}` |
| `items.append(4)` | `items = append(items, 4)` |
| `len(items)` | `len(items)` |

### Dictionaries
| Python | Go |
|--------|-----|
| `d = {"key": "value"}` | `d := map[string]string{"key": "value"}` |
| `d["key"]` | `d["key"]` |
| `d["new"] = "val"` | `d["new"] = "val"` |

### Loops
| Python | Go |
|--------|-----|
| `for i in range(10):` | `for i := 0; i < 10; i++ { }` |
| `for item in items:` | `for _, item := range items { }` |
| `while condition:` | `for condition { }` |

### Conditionals
| Python | Go |
|--------|-----|
| `if x > 0:` | `if x > 0 { }` |
| `elif x < 0:` | `else if x < 0 { }` |
| `else:` | `else { }` |

### Error Handling
| Python | Go |
|--------|-----|
| `try: ... except: ...` | `result, err := ...; if err != nil { }` |
| `raise ValueError("msg")` | `return nil, errors.New("msg")` |

### Classes
| Python | Go |
|--------|-----|
| `class Person: ...` | `type Person struct { ... }` |
| `def method(self): ...` | `func (p Person) method() { }` |

---

## Common Gotchas

1. **Slices are references** - Modifying a slice affects all references
2. **Range gives copies** - Use pointers if you need to modify
3. **Zero values** - Variables have default values (0, "", false, nil)
4. **No generics (pre-1.18)** - Use interfaces or code generation
5. **No exceptions** - Errors are values, must handle explicitly
6. **No inheritance** - Use composition instead
7. **No default parameters** - Use structs or variadic functions
8. **No function overloading** - One name = one function
9. **Exports are case-sensitive** - Capital = exported, lowercase = private
10. **Unused imports/variables** - Compiler error (not warning)

---

## Quick Commands

```bash
# Build
go build
go build -o myapp

# Run
go run main.go

# Test
go test
go test -v
go test ./...

# Format
go fmt ./...

# Vet (check for errors)
go vet ./...

# Get dependencies
go get package@version
go get -u ./...

# Module management
go mod init project
go mod tidy
go mod download

# Documentation
go doc package
go doc package.Function

# Install tool
go install tool@latest
```

---

## Must Remember Forever

1. **Go is compiled** - Errors caught before running
2. **Go is statically typed** - Types are fixed
3. **Errors are values** - Return `(result, error)`
4. **No exceptions** - Handle errors explicitly
5. **Goroutines are lightweight** - Use for concurrency
6. **Channels for communication** - Share data safely
7. **Composition over inheritance** - No classes
8. **Exports are case-sensitive** - Capital = public
9. **Unused code = error** - Compiler is strict
10. **Simplicity is key** - One clear way to do things

---

**Keep this cheat sheet handy while learning Go! 📚**
