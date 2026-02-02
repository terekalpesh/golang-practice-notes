# Basic Syntax

## a) Overview

### What this topic is
The fundamental rules and structure of Go code - how to write valid Go programs, organize code, and use basic language features.

### Why it exists in Go
Go has strict syntax rules that make code consistent and readable. Understanding these basics is essential before learning advanced features.

---

## b) Syntax

### Basic Program Structure
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

### Key Syntax Elements
- **Package declaration**: `package name`
- **Import statements**: `import "package"`
- **Functions**: `func name() { }`
- **Statements end with**: No semicolon needed (but allowed)
- **Comments**: `// single line` or `/* multi-line */`

---

## c) Explanation

### Step-by-Step Breakdown

**1. Package Declaration**
```go
package main
```
- Every Go file starts with `package`
- `package main` = executable program
- Other packages = libraries (reusable code)
- Package name matches directory name (usually)

**2. Import Statements**
```go
import "fmt"
import "os"
// Or multiple:
import (
    "fmt"
    "os"
)
```
- Import packages you need
- Standard library packages (like `fmt`, `os`) come with Go
- Third-party packages from internet
- Unused imports = compile error (Go is strict!)

**3. Functions**
```go
func main() {
    // code here
}
```
- `func` = function keyword
- `main` = special function (program starts here)
- `()` = parameters (empty here)
- `{ }` = function body

**4. Statements and Semicolons**
- Go automatically inserts semicolons
- You don't need to type them
- But they're allowed if you want

**5. Comments**
```go
// This is a single-line comment

/* This is a
   multi-line comment */
```

### Characteristics

#### Language Syntax Characteristics
- **Case-sensitive**: `Name` and `name` are different identifiers
- **No semicolons needed**: Go adds them automatically (but allowed)
- **Curly braces required**: `{ }` for blocks (even one-line blocks)
- **Strict formatting**: Go has opinions about style (enforced by `go fmt`)
- **Unused code = error**: Can't have unused variables/imports (compiler error)
- **Explicit syntax**: Clear and unambiguous syntax

#### Package Characteristics
- **Package declaration**: Every file starts with `package` declaration
- **Main package**: `package main` creates executable program
- **Package scope**: All files in same directory = same package
- **Package naming**: Package name usually matches directory name

#### Import Characteristics
- **Explicit imports**: Must import what you use
- **Unused imports**: Compile error (not warning)
- **Import grouping**: Standard library, third-party, local
- **Import aliases**: Can alias imports to avoid conflicts
- **Blank imports**: `_` import for side effects only

#### Code Organization Characteristics
- **File-based**: Code organized in files
- **Package-based**: Code organized in packages (directories)
- **Entry point**: `main()` function is program entry point
- **Compilation unit**: Package is compilation unit
- **Visibility**: Exported (capital) vs unexported (lowercase)

---

## d) Python Comparison

### Python Program
```python
# Python - simple script
print("Hello, Python!")

# Or with main guard
if __name__ == "__main__":
    print("Hello, Python!")
```

### Go Program
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

**Key Differences:**

| Feature | Python | Go |
|---------|--------|-----|
| **Entry point** | Top-level code or `if __name__ == "__main__"` | `func main()` |
| **Imports** | `import os` or `from os import path` | `import "os"` |
| **Package system** | Modules (files) | Packages (directories) |
| **Unused imports** | Warning | Compile error |
| **Semicolons** | Never used | Automatic |
| **Formatting** | Flexible (PEP 8 is style guide) | Enforced (gofmt) |

**Thinking Difference:**
- Python: Write code, run it
- Go: Write code, compile it, then run it
- Python: Flexible structure
- Go: Strict structure (but simpler)

---

## e) Visual Flow / Mental Model

### How a Go Program Executes

```
1. Program starts
   ↓
2. Go to main() function
   ↓
3. Execute statements in order
   ↓
4. When main() ends, program exits
```

### File Structure

```
main.go
│
├── package main          ← Declare this is an executable
│
├── import "fmt"          ← Bring in fmt package
│
└── func main() {         ← Program entry point
        fmt.Println(...)  ← Execute this
    }
```

### Memory Model (Simplified)

```
When program runs:
│
├── Code section          (Your compiled code)
├── Stack                 (Function calls, local variables)
└── Heap                  (Dynamic memory)
```

**Execution Flow:**
1. OS loads executable
2. Go runtime initializes
3. Calls `main()` function
4. Executes your code
5. Program exits

---

## f) Demo Example

### Complete Example with Explanation

```go
package main  // Line 1: This file is part of "main" package
             // "main" package = executable program

import (     // Line 3: Import block (multiple packages)
    "fmt"    // fmt = formatting package (like Python's print)
    "time"   // time = time-related functions
)

func main() {  // Line 8: main() is where program starts
               // Similar to Python's if __name__ == "__main__"
    
    // Print current time
    now := time.Now()                    // Get current time
    fmt.Println("Current time:", now)   // Print it
    
    // Print formatted message
    name := "Go Learner"
    fmt.Printf("Hello, %s!\n", name)   // Formatted print (like Python's f-string)
    
    // Print multiple values
    fmt.Println("Learning", "Go", "is", "fun!")
}
```

**Line-by-line explanation:**

1. **`package main`**: Declares this as an executable (not a library)
2. **`import (...)`**: Brings in packages we need
3. **`func main()`**: Entry point - program starts here
4. **`now := time.Now()`**: Get current time (short variable declaration)
5. **`fmt.Println(...)`**: Print with newline
6. **`fmt.Printf(...)`**: Print with formatting (like Python's f-strings)

**Run it:**
```bash
go run main.go
```

**Output:**
```
Current time: 2024-01-15 10:30:45.123456789 +0000 UTC
Hello, Go Learner!
Learning Go is fun!
```

---

## g) Use Cases

### When to use different import styles

**Single import:**
```go
import "fmt"
```
- When you only need one package
- Simple and clear

**Multiple imports:**
```go
import (
    "fmt"
    "os"
    "strings"
)
```
- When you need several packages
- More readable
- Recommended style

**Import with alias:**
```go
import f "fmt"
import "database/sql"
import _ "github.com/lib/pq"  // Blank import (side effects only)
```
- When package names conflict
- When you only want side effects (like database drivers)

### Package naming

**Executable (main package):**
```go
package main  // Creates runnable program
```

**Library (other packages):**
```go
package calculator  // Reusable code
package utils       // Utility functions
```
- Package name = directory name
- Lowercase, one word
- Descriptive

---

## h) Do's and Don'ts / Best Practices

### ✅ Do's

1. **Use `go fmt`**
   ```bash
   go fmt ./...
   ```
   - Formats your code automatically
   - Everyone's code looks the same
   - Run before committing

2. **Group related imports**
   ```go
   import (
       // Standard library
       "fmt"
       "os"
       
       // Third-party
       "github.com/gin-gonic/gin"
   )
   ```

3. **Use meaningful package names**
   - `package calculator` ✅
   - `package calc` ✅ (short is okay)
   - `package mypackage` ❌ (too generic)

4. **Keep main() simple**
   - Main should orchestrate, not do everything
   - Move logic to other functions

5. **Use comments for "why", not "what"**
   ```go
   // Calculate tax (rate changed in 2024)
   tax := price * 0.08
   ```

### ❌ Don'ts

1. **Don't leave unused imports**
   ```go
   import "os"  // Error if not used!
   ```
   - Go compiler will error
   - Remove unused imports

2. **Don't use semicolons unnecessarily**
   ```go
   x := 5;  // Semicolon not needed
   y := 10; // But allowed
   ```

3. **Don't mix package styles**
   ```go
   package main  // ✅ Consistent
   package Main  // ❌ Wrong (exported names start with capital)
   ```

4. **Don't ignore compiler errors**
   - Fix them immediately
   - Go compiler is helpful

5. **Don't write overly long functions**
   - Keep functions focused
   - One responsibility per function

---

## i) Solved Practice Examples

### Example 1: Basic Program Structure

**Task:** Write a program that prints your name and age.

**Solution:**
```go
package main

import "fmt"

func main() {
    name := "Alice"
    age := 30
    
    fmt.Printf("My name is %s and I am %d years old.\n", name, age)
}
```

**Explanation:**
- `package main`: Executable program
- `import "fmt"`: Need fmt for printing
- `func main()`: Entry point
- `:=`: Short variable declaration (we'll learn this next)
- `fmt.Printf`: Formatted printing (like Python's f-strings)

### Example 2: Multiple Imports

**Task:** Write a program that prints current time and a random number.

**Solution:**
```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Seed random number generator
    rand.Seed(time.Now().UnixNano())
    
    // Print current time
    fmt.Println("Current time:", time.Now())
    
    // Print random number
    randomNum := rand.Intn(100)
    fmt.Printf("Random number: %d\n", randomNum)
}
```

**Explanation:**
- Multiple imports in block
- `rand.Seed()`: Initialize random generator
- `time.Now()`: Get current time
- `rand.Intn(100)`: Random number 0-99

### Example 3: Package Organization

**Task:** Create a simple calculator with separate function.

**Solution:**
```go
package main

import "fmt"

// Function defined outside main
func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(5, 3)
    fmt.Printf("5 + 3 = %d\n", result)
}
```

**Explanation:**
- Functions can be defined before or after main
- `add()` function takes two integers, returns one
- `int` is the type (we'll learn types next)

---

## j) Quiz / Practice for Me

### Quiz Questions

1. **What package name creates an executable program?**
   - [ ] `package program`
   - [ ] `package main`
   - [ ] `package executable`
   - [ ] `package run`

2. **What happens if you import a package but don't use it?**
   - [ ] Warning
   - [ ] Compile error
   - [ ] Nothing
   - [ ] Runtime error

3. **Where does a Go program start executing?**
   - [ ] Top of file
   - [ ] `func init()`
   - [ ] `func main()`
   - [ ] `func start()`

4. **What does `go fmt` do?**
   - [ ] Runs your program
   - [ ] Formats your code
   - [ ] Installs packages
   - [ ] Checks for errors

### Practice Tasks

**Task 1: Hello World with Details**
- Create a program that prints:
  - Your name
  - Today's date
  - A welcome message
- Use `fmt.Printf` for formatted output

**Task 2: Import Practice**
- Import three different packages
- Use at least one function from each
- Make sure all imports are used

**Task 3: Function Organization**
- Create a function called `greet()` that takes a name
- Call it from `main()`
- Print the greeting

### Answers

**Quiz Answers:**
1. `package main`
2. Compile error
3. `func main()`
4. Formats your code

**Practice Solutions:**

**Task 1 Solution:**
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    name := "Go Learner"
    today := time.Now().Format("2006-01-02")
    
    fmt.Printf("Hello, my name is %s.\n", name)
    fmt.Printf("Today is %s.\n", today)
    fmt.Println("Welcome to Go programming!")
}
```

**Task 2 Solution:**
```go
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    // Use fmt
    fmt.Println("Using fmt package")
    
    // Use os
    hostname, _ := os.Hostname()
    fmt.Printf("Hostname: %s\n", hostname)
    
    // Use strings
    text := "hello world"
    upper := strings.ToUpper(text)
    fmt.Println(upper)
}
```

**Task 3 Solution:**
```go
package main

import "fmt"

func greet(name string) {
    fmt.Printf("Hello, %s! Welcome to Go!\n", name)
}

func main() {
    greet("Alice")
    greet("Bob")
}
```

---

## Key Takeaways

1. **Every Go file starts with `package`** - `package main` for executables
2. **Imports are explicit** - Must import what you use, unused = error
3. **`main()` is entry point** - Program starts here
4. **Go is strict** - Unused code = compile error (this is good!)
5. **Use `go fmt`** - Keeps code consistent
6. **Simple structure** - Package → Imports → Functions

---

## Must Remember Forever

- `package main` = executable program
- `func main()` = where program starts
- Unused imports/variables = compile error
- Go automatically formats code (use `go fmt`)
- Curly braces `{ }` are required for blocks

---

## Next Steps

Now you understand Go's basic structure. Next:
- Variables and types
- How to declare and use variables
- Go's type system

**Ready? → [03_VARIABLES_AND_TYPES.md](./03_VARIABLES_AND_TYPES.md)**
