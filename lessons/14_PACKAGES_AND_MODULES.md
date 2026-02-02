# Packages and Modules

## a) Overview

### What this topic is
How to organize Go code into packages and manage dependencies with modules. This is essential for building real applications.

### Why it exists in Go
Packages organize code into logical units. Modules manage dependencies and versions. This keeps code maintainable and reusable.

---

## b) Syntax

### Package Declaration
```go
package mypackage

// Exported (capital letter)
func PublicFunction() { }

// Unexported (lowercase)
func privateFunction() { }
```

### Module Setup
```bash
go mod init myproject
go get package@version
go mod tidy
```

### Importing
```go
import "fmt"                    // Standard library
import "myproject/mypackage"    // Local package
import "github.com/user/repo"   // External package
```

---

## c) Explanation

### Step-by-Step Module Creation

**1. Initialize module**
```bash
go mod init myproject
```
- Creates `go.mod` file
- Defines module name
- Enables dependency management

**2. Create package**
```go
// mypackage/mypackage.go
package mypackage

func PublicFunc() { }
```
- Package name = directory name (usually)
- Files in same directory = same package

**3. Import and use**
```go
import "myproject/mypackage"
mypackage.PublicFunc()
```

### Characteristics

- **Package = directory**: One package per directory
- **Exports**: Capital letter = exported, lowercase = private
- **Modules**: Dependency management (Go 1.11+)
- **go.mod**: Defines module and dependencies
- **go.sum**: Checksums for security

---

## d) Python Comparison

### Python Modules
```python
# mymodule.py
def public_function():
    pass

def _private_function():
    pass

# Import
import mymodule
from mymodule import public_function
```

### Go Packages
```go
// mypackage/mypackage.go
package mypackage

func PublicFunction() { }  // Exported
func privateFunction() { }  // Unexported

// Import
import "myproject/mypackage"
mypackage.PublicFunction()
```

**Key Differences:**

| Feature | Python | Go |
|---------|--------|-----|
| **Organization** | File = module | Directory = package |
| **Exports** | `_name` (convention) | Capital letter (enforced) |
| **Import** | `import module` | `import "path/to/package"` |
| **Dependencies** | requirements.txt | go.mod |

---

## e) Visual Flow / Mental Model

### Package Structure

```
myproject/
├── go.mod
├── go.sum
├── main.go
├── utils/
│   └── helper.go
└── models/
    └── user.go
```

### Import Resolution

```
import "myproject/utils"
  ↓
1. Look in go.mod for module name
  ↓
2. Resolve to myproject/utils/
  ↓
3. Load all .go files in utils/
  ↓
4. Make exported functions available
```

---

## f) Demo Example

### Complete Example

```go
// go.mod
module myproject

go 1.21

// main.go
package main

import (
    "fmt"
    "myproject/utils"
    "myproject/models"
)

func main() {
    result := utils.Add(5, 3)
    fmt.Println("Result:", result)
    
    user := models.NewUser("Alice", 30)
    fmt.Println(user)
}

// utils/helper.go
package utils

func Add(a, b int) int {
    return a + b
}

func subtract(a, b int) int {  // Private
    return a - b
}

// models/user.go
package models

type User struct {
    Name string
    Age  int
}

func NewUser(name string, age int) *User {
    return &User{Name: name, Age: age}
}

func (u *User) String() string {
    return fmt.Sprintf("User{Name: %s, Age: %d}", u.Name, u.Age)
}
```

---

## g) Use Cases

### When to create packages

**1. Logical grouping**
```go
// utils/ - utility functions
// models/ - data models
// handlers/ - HTTP handlers
```

**2. Reusability**
```go
// shared/ - shared code
```

**3. Separation of concerns**
```go
// database/ - database code
// api/ - API code
```

---

## h) Do's and Don'ts

### ✅ Do's

1. **Use meaningful package names**
2. **Keep packages focused**
3. **Export only what's needed**
4. **Use go mod tidy regularly**

### ❌ Don'ts

1. **Don't create too many small packages**
2. **Don't export everything**
3. **Don't use circular dependencies**

---

## i) Solved Practice Examples

### Example 1: Basic Package Structure

**Task:** Create a calculator package with basic math operations.

**Solution:**

```go
// go.mod
module calculator

go 1.21

// calculator/calculator.go
package calculator

// Add adds two numbers
func Add(a, b float64) float64 {
    return a + b
}

// Subtract subtracts b from a
func Subtract(a, b float64) float64 {
    return a - b
}

// Multiply multiplies two numbers
func Multiply(a, b float64) float64 {
    return a * b
}

// Divide divides a by b
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// main.go
package main

import (
    "calculator"
    "fmt"
)

func main() {
    sum := calculator.Add(10, 5)
    fmt.Println("10 + 5 =", sum)
    
    diff := calculator.Subtract(10, 5)
    fmt.Println("10 - 5 =", diff)
    
    product := calculator.Multiply(10, 5)
    fmt.Println("10 * 5 =", product)
    
    quotient, err := calculator.Divide(10, 5)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("10 / 5 =", quotient)
    }
}
```

### Example 2: Multi-Package Project

**Task:** Create a project with multiple packages: models, utils, and handlers.

**Solution:**

```go
// go.mod
module myapp

go 1.21

// models/user.go
package models

import "fmt"

type User struct {
    ID    int
    Name  string
    Email string
}

func NewUser(id int, name, email string) *User {
    return &User{
        ID:    id,
        Name:  name,
        Email: email,
    }
}

func (u *User) String() string {
    return fmt.Sprintf("User{ID: %d, Name: %s, Email: %s}", u.ID, u.Name, u.Email)
}

// utils/validator.go
package utils

import "strings"

func ValidateEmail(email string) bool {
    return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func ValidateName(name string) bool {
    return len(name) > 0 && len(name) <= 100
}

// handlers/user_handler.go
package handlers

import (
    "myapp/models"
    "myapp/utils"
    "fmt"
)

func CreateUser(id int, name, email string) (*models.User, error) {
    if !utils.ValidateName(name) {
        return nil, fmt.Errorf("invalid name")
    }
    if !utils.ValidateEmail(email) {
        return nil, fmt.Errorf("invalid email")
    }
    return models.NewUser(id, name, email), nil
}

// main.go
package main

import (
    "fmt"
    "myapp/handlers"
)

func main() {
    user, err := handlers.CreateUser(1, "Alice", "alice@example.com")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Created:", user)
}
```

### Example 3: Using External Packages

**Task:** Create a project that uses an external package (like a UUID generator).

**Solution:**

```bash
# Initialize module
go mod init uuid-example

# Get external package
go get github.com/google/uuid
```

```go
// go.mod
module uuid-example

go 1.21

require github.com/google/uuid v1.3.0

// main.go
package main

import (
    "fmt"
    "github.com/google/uuid"
)

func main() {
    // Generate new UUID
    id := uuid.New()
    fmt.Println("Generated UUID:", id)
    
    // Parse existing UUID
    parsedID, err := uuid.Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Parsed UUID:", parsedID)
}
```

### Example 4: Package with Internal Package

**Task:** Create a package with internal subpackage that can't be imported from outside.

**Solution:**

```go
// mylib/mylib.go
package mylib

import "mylib/internal/helper"

func PublicFunction() string {
    return helper.InternalHelper() + " - public"
}

// mylib/internal/helper.go
package helper

func InternalHelper() string {
    return "internal helper"
}

// main.go (in different module)
package main

import (
    "fmt"
    "mylib"
    // "mylib/internal/helper"  // ❌ ERROR! Can't import internal
)

func main() {
    result := mylib.PublicFunction()
    fmt.Println(result)
}
```

---

## j) Quiz / Practice for Me

### Quiz Questions

1. **What is the difference between a package and a module?**
   - [ ] No difference
   - [ ] Package = code organization, Module = dependency management
   - [ ] Module = code organization, Package = dependency management
   - [ ] They're the same thing

2. **How do you export a function from a package?**
   - [ ] Use `export` keyword
   - [ ] Start name with capital letter
   - [ ] Use `public` keyword
   - [ ] Add `@export` annotation

3. **What does `go mod tidy` do?**
   - [ ] Installs all packages
   - [ ] Removes unused dependencies, adds missing ones
   - [ ] Updates Go version
   - [ ] Formats code

4. **Can you import an `internal` package from outside the module?**
   - [ ] Yes, always
   - [ ] No, never
   - [ ] Only if it's exported
   - [ ] Only from parent directory

### Practice Tasks

**Task 1: Create Math Package**
- Create a `math` package with functions: `Power(base, exp float64) float64` and `Sqrt(x float64) float64`
- Create a main program that uses this package
- Test all functions

**Task 2: Create String Utilities Package**
- Create a `strutil` package with:
  - `Reverse(s string) string` - reverses a string
  - `ToTitle(s string) string` - converts to title case
  - `Contains(s, substr string) bool` - checks if string contains substring
- Export all functions
- Create main program to test them

**Task 3: Multi-Module Project**
- Create two modules: `lib` and `app`
- `lib` module exports a `Logger` struct with `Log(message string)` method
- `app` module imports and uses `lib` module
- Set up proper go.mod files

**Task 4: Package with Constants**
- Create a `config` package that exports:
  - `MaxUsers = 100`
  - `DefaultPort = 8080`
  - `AppName = "MyApp"`
- Import and use these constants in main

### Answers

**Quiz Answers:**
1. Package = code organization, Module = dependency management
2. Start name with capital letter
3. Removes unused dependencies, adds missing ones
4. No, never

**Practice Solutions:**

**Task 1 Solution:**
```go
// math/math.go
package math

import "math"

func Power(base, exp float64) float64 {
    return math.Pow(base, exp)
}

func Sqrt(x float64) float64 {
    return math.Sqrt(x)
}

// main.go
package main

import (
    "fmt"
    "myproject/math"
)

func main() {
    result := math.Power(2, 3)
    fmt.Printf("2^3 = %.2f\n", result)
    
    sqrt := math.Sqrt(16)
    fmt.Printf("sqrt(16) = %.2f\n", sqrt)
}
```

**Task 2 Solution:**
```go
// strutil/strutil.go
package strutil

import (
    "strings"
    "unicode"
)

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func ToTitle(s string) string {
    return strings.Title(strings.ToLower(s))
}

func Contains(s, substr string) bool {
    return strings.Contains(s, substr)
}

// main.go
package main

import (
    "fmt"
    "myproject/strutil"
)

func main() {
    fmt.Println("Reverse:", strutil.Reverse("Hello"))
    fmt.Println("ToTitle:", strutil.ToTitle("hello world"))
    fmt.Println("Contains:", strutil.Contains("Hello World", "World"))
}
```

**Task 3 Solution:**
```go
// lib/go.mod
module lib

go 1.21

// lib/logger.go
package lib

import "fmt"

type Logger struct {
    prefix string
}

func NewLogger(prefix string) *Logger {
    return &Logger{prefix: prefix}
}

func (l *Logger) Log(message string) {
    fmt.Printf("[%s] %s\n", l.prefix, message)
}

// app/go.mod
module app

go 1.21

require lib v0.0.0
replace lib => ../lib

// app/main.go
package main

import "lib"

func main() {
    logger := lib.NewLogger("APP")
    logger.Log("Application started")
    logger.Log("Processing data")
}
```

**Task 4 Solution:**
```go
// config/config.go
package config

const (
    MaxUsers   = 100
    DefaultPort = 8080
    AppName    = "MyApp"
)

// main.go
package main

import (
    "fmt"
    "myproject/config"
)

func main() {
    fmt.Printf("App: %s\n", config.AppName)
    fmt.Printf("Max Users: %d\n", config.MaxUsers)
    fmt.Printf("Default Port: %d\n", config.DefaultPort)
}
```

---

## Key Takeaways

1. **Package = directory** - One package per directory
2. **Exports are case-sensitive** - Capital = exported
3. **Modules manage dependencies** - go.mod file
4. **Import by path** - Relative to module root

---

## Next Steps

**Ready? → [15_FILE_IO.md](./15_FILE_IO.md)**
