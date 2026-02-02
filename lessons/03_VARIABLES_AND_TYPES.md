# Variables and Types

## a) Overview

### What this topic is
How to create and use variables in Go, and understanding Go's type system (static typing).

### Why it exists in Go
Go is statically typed - every variable has a type that's known at compile time. This catches errors early and makes code more reliable.

---

## b) Syntax

### Variable Declaration
```go
// Explicit type
var name string = "John"

// Type inference
var age = 25

// Short declaration (most common)
name := "John"
age := 25

// Multiple variables
var x, y int = 1, 2
a, b := 10, 20

// Constants
const pi = 3.14159
const greeting = "Hello"
```

### Basic Types
```go
// Numbers
var i int = 42
var f float64 = 3.14
var b byte = 255

// Strings
var s string = "Hello"

// Booleans
var isTrue bool = true

// Zero values (default values)
var x int        // 0
var y string     // ""
var z bool       // false
```

---

## c) Explanation

### Step-by-Step Variable Declaration

**1. Using `var` (explicit)**
```go
var name string
name = "John"
```
- Declare variable with type
- Assign value later
- Can declare without initial value (gets "zero value")

**2. Using `var` with initialization**
```go
var name string = "John"
```
- Declare and assign in one line
- Type is explicit

**3. Type inference**
```go
var name = "John"
```
- Go figures out the type
- `"John"` is a string, so `name` is `string`

**4. Short declaration (most common)**
```go
name := "John"
```
- `:=` means "declare and assign"
- Type is inferred
- Only works inside functions
- Most idiomatic Go style

### Zero Values

Every type has a "zero value" (default value):

| Type | Zero Value |
|------|------------|
| `int`, `float64`, etc. | `0` |
| `string` | `""` (empty string) |
| `bool` | `false` |
| Pointers, slices, maps, channels | `nil` |

**Example:**
```go
var x int        // x = 0
var s string     // s = ""
var b bool       // b = false
```

### Type System

**Go's types:**
- **Basic types**: `int`, `string`, `bool`, `float64`, etc.
- **Composite types**: arrays, slices, maps, structs (we'll learn later)
- **Reference types**: pointers, slices, maps, channels, functions
- **Interface types**: interfaces (we'll learn later)

**Type safety:**
- Can't mix types without conversion
- Compiler catches type errors
- No implicit conversions (unlike Python)

### Variable Characteristics

#### Declaration Characteristics
- **Static typing**: Type is known at compile time (not runtime)
- **Type inference**: Go can infer type from value (but type is fixed)
- **Immutable type**: Once declared, type cannot change
- **Scope**: Variables have scope (package-level or function-level)
- **Lifetime**: Variables exist until out of scope

#### Memory Characteristics
- **Stack vs Heap**: Local variables on stack, large data on heap
- **Zero values**: Uninitialized variables get default values
- **Memory safety**: Go manages memory (garbage collection)
- **No null**: Go uses zero values, not null/None

#### Type Characteristics
- **Strong typing**: Types are strictly enforced
- **No type coercion**: Must explicitly convert types
- **Type checking**: Happens at compile time (catches errors early)
- **Type aliases**: Can create type aliases for clarity

#### Data Characteristics
- **Value semantics**: Variables hold values (not references, except for reference types)
- **Copy behavior**: Assignment copies value (except reference types)
- **Mutability**: Can modify variable value (but not type)
- **Initialization**: Can declare without initializing (gets zero value)

---

## d) Python Comparison

### Python Variables
```python
# Python - dynamic typing
name = "John"        # Type: str (inferred)
age = 25             # Type: int (inferred)
price = 19.99        # Type: float (inferred)

# Can change type
x = 10               # x is int
x = "hello"          # Now x is str (allowed!)
```

### Go Variables
```go
// Go - static typing
var name string = "John"  // Type: string (explicit)
age := 25                 // Type: int (inferred, but fixed)
price := 19.99           // Type: float64 (inferred, but fixed)

// Cannot change type
x := 10                   // x is int
x = "hello"               // ERROR! Can't assign string to int
```

**Key Differences:**

| Feature | Python | Go |
|---------|--------|-----|
| **Type checking** | Runtime | Compile time |
| **Type changes** | Allowed | Not allowed |
| **Type declaration** | Optional | Required (explicit or inferred) |
| **Type errors** | Found when running | Found before running |
| **Performance** | Slower (type checking at runtime) | Faster (types known at compile time) |

**Thinking Difference:**
- Python: "I'll figure out the type when I run the code"
- Go: "I need to know the type before I run the code"
- Python: Flexible but can have runtime errors
- Go: Strict but catches errors early

---

## e) Visual Flow / Mental Model

### Variable Declaration in Memory

**Python (runtime):**
```
Code: name = "John"
      ↓
Runtime checks: Is "John" a string? Yes.
      ↓
Memory: [name] → "John" (type: str, stored in memory)
      ↓
Later: name = 25
      ↓
Runtime checks: Is 25 an int? Yes.
      ↓
Memory: [name] → 25 (type: int, changed!)
```

**Go (compile time):**
```
Code: name := "John"
      ↓
Compiler: "John" is string, so name is string (forever)
      ↓
Memory: [name: string] → "John"
      ↓
Later: name = 25  ← COMPILE ERROR!
      ↓
Compiler: "Can't assign int to string variable"
      ↓
Program doesn't even compile!
```

### Type Inference Process

```go
x := 42
```

**What Go does:**
1. See `42` - it's an integer literal
2. Default integer type is `int`
3. Assign `x` type `int`
4. `x` is now `int` forever (can't change)

```go
y := 3.14
```

**What Go does:**
1. See `3.14` - it's a float literal
2. Default float type is `float64`
3. Assign `y` type `float64`
4. `y` is now `float64` forever

---

## f) Demo Example

### Complete Example with All Variable Types

```go
package main

import "fmt"

func main() {
    // 1. Basic types with short declaration
    name := "Alice"           // string
    age := 30                 // int
    height := 5.6             // float64
    isStudent := false        // bool
    
    fmt.Printf("Name: %s, Age: %d, Height: %.1f, Student: %t\n",
        name, age, height, isStudent)
    
    // 2. Explicit type declaration
    var count int
    count = 10
    fmt.Println("Count:", count)
    
    // 3. Multiple variables
    x, y := 10, 20
    fmt.Printf("x = %d, y = %d\n", x, y)
    
    // 4. Type conversion (explicit)
    var a int = 42
    var b float64 = float64(a)  // Convert int to float64
    fmt.Printf("a (int): %d, b (float64): %.1f\n", a, b)
    
    // 5. Constants
    const pi = 3.14159
    const greeting = "Hello, Go!"
    fmt.Println("Pi:", pi)
    fmt.Println(greeting)
    
    // 6. Zero values
    var zeroInt int
    var zeroString string
    var zeroBool bool
    fmt.Printf("Zero int: %d, Zero string: '%s', Zero bool: %t\n",
        zeroInt, zeroString, zeroBool)
    
    // 7. Type checking (compile-time)
    // This would cause error:
    // name = 123  // ERROR: cannot use 123 (type int) as string
}
```

**Line-by-line explanation:**

1. **Short declarations**: `:=` declares and assigns, type inferred
2. **Explicit declaration**: `var count int` then assign later
3. **Multiple assignment**: `x, y := 10, 20` declares both
4. **Type conversion**: `float64(a)` converts int to float64 (explicit!)
5. **Constants**: `const` for values that don't change
6. **Zero values**: Uninitialized variables get default values
7. **Type safety**: Can't assign wrong type (commented out error)

**Output:**
```
Name: Alice, Age: 30, Height: 5.6, Student: false
Count: 10
x = 10, y = 20
a (int): 42, b (float64): 42.0
Pi: 3.14159
Hello, Go!
Zero int: 0, Zero string: '', Zero bool: false
```

---

## g) Use Cases

### When to use different declaration styles

**Short declaration (`:=`):**
```go
name := "John"  // Most common, inside functions
```
- Inside functions
- When type is obvious
- Most idiomatic Go

**`var` with type:**
```go
var count int  // When you need zero value
```
- When you need zero value initially
- Package-level variables
- When type is important to show

**`var` with initialization:**
```go
var name string = "John"  // Explicit type
```
- When you want to be explicit
- Less common (short declaration preferred)

**Constants:**
```go
const pi = 3.14159
```
- Values that never change
- Compile-time constants
- No `:=` for constants (use `const`)

### Type conversions

**When you need them:**
```go
var i int = 42
var f float64 = float64(i)  // int to float64
var s string = string(i)    // int to string (careful! converts to Unicode)
```

**Common conversions:**
- `int` ↔ `float64`
- `string` ↔ `[]byte` (byte slice)
- Between number types

---

## h) Do's and Don'ts / Best Practices

### ✅ Do's

1. **Use short declaration (`:=`) inside functions**
   ```go
   name := "John"  // ✅ Preferred
   ```

2. **Use `var` for package-level variables**
   ```go
   var GlobalCount int  // ✅ Package level
   
   func main() {
       local := 10  // ✅ Function level
   }
   ```

3. **Use meaningful variable names**
   ```go
   userName := "John"     // ✅ Clear
   u := "John"           // ❌ Unclear
   ```

4. **Use constants for magic numbers**
   ```go
   const maxUsers = 100  // ✅
   if count > 100 {      // ❌ What is 100?
   ```

5. **Group related declarations**
   ```go
   var (
       name string
       age  int
   )
   ```

### ❌ Don'ts

1. **Don't use `:=` for package-level variables**
   ```go
   // ❌ ERROR: syntax error
   global := 10
   
   // ✅ Correct
   var global = 10
   ```

2. **Don't redeclare with `:=` if variable exists**
   ```go
   x := 10
   x := 20  // ❌ ERROR: no new variables
   x = 20   // ✅ Correct (assignment)
   ```

3. **Don't ignore type conversions**
   ```go
   var i int = 42
   var f float64 = i  // ❌ ERROR: need conversion
   var f float64 = float64(i)  // ✅ Correct
   ```

4. **Don't use `var` when `:=` is clearer**
   ```go
   var name = "John"  // Works, but...
   name := "John"     // ✅ More idiomatic
   ```

5. **Don't use unused variables**
   ```go
   x := 10  // ❌ ERROR if not used
   _ = 10   // ✅ Use blank identifier if needed
   ```

---

## i) Solved Practice Examples

### Example 1: Basic Variable Operations

**Task:** Create variables for a person's information and print them.

**Solution:**
```go
package main

import "fmt"

func main() {
    firstName := "John"
    lastName := "Doe"
    age := 30
    salary := 75000.50
    
    fmt.Printf("Name: %s %s\n", firstName, lastName)
    fmt.Printf("Age: %d\n", age)
    fmt.Printf("Salary: $%.2f\n", salary)
}
```

**Explanation:**
- Short declarations for all variables
- Types inferred automatically
- `fmt.Printf` for formatted output
- `%s` for strings, `%d` for integers, `%.2f` for floats

### Example 2: Type Conversions

**Task:** Convert between different number types.

**Solution:**
```go
package main

import "fmt"

func main() {
    // Start with int
    wholeNumber := 42
    
    // Convert to float64
    decimalNumber := float64(wholeNumber)
    
    // Convert to string (using fmt.Sprintf)
    numberString := fmt.Sprintf("%d", wholeNumber)
    
    fmt.Printf("Int: %d\n", wholeNumber)
    fmt.Printf("Float64: %.2f\n", decimalNumber)
    fmt.Printf("String: %s\n", numberString)
    
    // Convert float to int (truncates)
    pi := 3.14159
    piInt := int(pi)
    fmt.Printf("Pi as int: %d\n", piInt)  // Output: 3 (truncated)
}
```

**Explanation:**
- `float64(intValue)` converts int to float64
- `int(floatValue)` converts float to int (truncates, doesn't round)
- `fmt.Sprintf` converts to string (better than `string()` for numbers)

### Example 3: Constants and Zero Values

**Task:** Use constants and demonstrate zero values.

**Solution:**
```go
package main

import "fmt"

const (
    maxUsers = 100
    appName  = "MyApp"
)

func main() {
    // Constants
    fmt.Printf("App: %s, Max Users: %d\n", appName, maxUsers)
    
    // Zero values
    var userCount int
    var userName string
    var isActive bool
    
    fmt.Printf("User Count (zero): %d\n", userCount)
    fmt.Printf("User Name (zero): '%s'\n", userName)
    fmt.Printf("Is Active (zero): %t\n", isActive)
    
    // Assign values
    userCount = 5
    userName = "Alice"
    isActive = true
    
    fmt.Printf("After assignment:\n")
    fmt.Printf("User Count: %d\n", userCount)
    fmt.Printf("User Name: %s\n", userName)
    fmt.Printf("Is Active: %t\n", isActive)
}
```

**Explanation:**
- Constants declared with `const`
- Can group multiple constants
- Zero values are automatic defaults
- Assign values later as needed

---

## j) Quiz / Practice for Me

### Quiz Questions

1. **What is the zero value for `int`?**
   - [ ] `nil`
   - [ ] `0`
   - [ ] `""`
   - [ ] `undefined`

2. **What does `:=` do?**
   - [ ] Assigns value
   - [ ] Declares and assigns (short declaration)
   - [ ] Compares values
   - [ ] Converts type

3. **Can you change a variable's type after declaration?**
   - [ ] Yes, always
   - [ ] No, never
   - [ ] Only with conversion
   - [ ] Only for numbers

4. **Where can you use `:=`?**
   - [ ] Anywhere
   - [ ] Only inside functions
   - [ ] Only at package level
   - [ ] Only for constants

### Practice Tasks

**Task 1: Variable Declaration Practice**
- Create variables for: name (string), age (int), height (float64), isMarried (bool)
- Use short declaration
- Print all values

**Task 2: Type Conversion**
- Start with `int` value 100
- Convert it to `float64`
- Convert it to `string`
- Print all three versions

**Task 3: Constants and Calculations**
- Create constants for: `pi = 3.14159`, `radius = 5`
- Calculate circle area: `area = pi * radius * radius`
- Print the result

### Answers

**Quiz Answers:**
1. `0`
2. Declares and assigns (short declaration)
3. No, never (type is fixed)
4. Only inside functions

**Practice Solutions:**

**Task 1 Solution:**
```go
package main

import "fmt"

func main() {
    name := "Alice"
    age := 30
    height := 5.6
    isMarried := true
    
    fmt.Printf("Name: %s\n", name)
    fmt.Printf("Age: %d\n", age)
    fmt.Printf("Height: %.1f\n", height)
    fmt.Printf("Married: %t\n", isMarried)
}
```

**Task 2 Solution:**
```go
package main

import "fmt"

func main() {
    original := 100
    
    asInt := original
    asFloat := float64(original)
    asString := fmt.Sprintf("%d", original)
    
    fmt.Printf("Int: %d\n", asInt)
    fmt.Printf("Float64: %.2f\n", asFloat)
    fmt.Printf("String: %s\n", asString)
}
```

**Task 3 Solution:**
```go
package main

import "fmt"

func main() {
    const pi = 3.14159
    const radius = 5
    
    area := pi * radius * radius
    
    fmt.Printf("Circle area (radius=%.0f): %.2f\n", radius, area)
}
```

---

## Key Takeaways

1. **Go is statically typed** - Types are known at compile time
2. **Short declaration (`:=`)** - Most common, inside functions
3. **Zero values** - Every type has a default value
4. **Type conversion is explicit** - Must convert types manually
5. **Constants use `const`** - Values that don't change
6. **Type safety** - Can't mix types without conversion

---

## Must Remember Forever

- `:=` declares and assigns (inside functions only)
- `var` declares (can be package-level)
- Types are fixed - can't change after declaration
- Zero values: `int=0`, `string=""`, `bool=false`
- Type conversion: `float64(intValue)` (explicit!)
- Constants: `const name = "value"`

---

## Next Steps

Now you understand variables and types. Next:
- Functions
- How to organize code into reusable functions
- Parameters and return values

**Ready? → [04_FUNCTIONS.md](./04_FUNCTIONS.md)**
