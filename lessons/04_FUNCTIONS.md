# Functions

## a) Overview

### What this topic is
Functions in Go - how to create reusable blocks of code, pass parameters, and return values.

### Why it exists in Go
Functions are the building blocks of Go programs. Go's function syntax is simple but powerful, with support for multiple return values (especially for errors).

---

## b) Syntax

### Basic Function
```go
func functionName() {
    // code
}

func functionName(param1 type1, param2 type2) {
    // code
}

func functionName() returnType {
    return value
}

func functionName(param type) (returnType1, returnType2) {
    return value1, value2
}
```

### Function Variants
```go
// Named return values
func divide(a, b float64) (result float64, err error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    result = a / b
    return  // Naked return (uses named values)
}

// Variadic functions (variable arguments)
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

// Functions as values
var myFunc func(int, int) int
myFunc = add
```

---

## c) Explanation

### Step-by-Step Function Creation

**1. Basic function (no parameters, no return)**
```go
func greet() {
    fmt.Println("Hello!")
}
```
- `func` = function keyword
- `greet` = function name
- `()` = parameters (empty)
- `{ }` = function body

**2. Function with parameters**
```go
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
```
- `name string` = parameter name and type
- Type comes after name (unlike some languages)

**3. Function with return value**
```go
func add(a int, b int) int {
    return a + b
}
```
- `int` after `)` = return type
- `return` = send value back

**4. Multiple return values (Go's special feature!)**
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```
- `(float64, error)` = two return types
- Return multiple values: `return value1, value2`
- Common pattern: `(result, error)`

**5. Named return values**
```go
func calculate(x, y int) (sum int, product int) {
    sum = x + y
    product = x * y
    return  // Naked return
}
```
- Names in return type: `(sum int, product int)`
- Variables `sum` and `product` are created automatically
- `return` without values = returns named variables

### Function Characteristics

- **First-class citizens**: Functions are values (can assign to variables)
- **Multiple returns**: Go's signature feature
- **No default parameters**: Must pass all parameters
- **No function overloading**: One function name = one function
- **Can return functions**: Functions can return other functions

---

## d) Python Comparison

### Python Functions
```python
def greet(name="World"):
    print(f"Hello, {name}!")

def add(a, b):
    return a + b

def divide(a, b):
    if b == 0:
        raise ValueError("division by zero")
    return a / b
```

### Go Functions
```go
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

func add(a int, b int) int {
    return a + b
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

**Key Differences:**

| Feature | Python | Go |
|---------|--------|-----|
| **Default parameters** | Yes | No |
| **Multiple returns** | Yes (tuples) | Yes (built-in) |
| **Error handling** | Exceptions | Return error value |
| **Type hints** | Optional | Required |
| **Function overloading** | No (but can simulate) | No |
| **Variadic args** | `*args` | `...type` |

**Thinking Difference:**
- Python: Use exceptions for errors
- Go: Return error as value (explicit)
- Python: Default parameters for flexibility
- Go: Simplicity (no defaults, but can use structs)

---

## e) Visual Flow / Mental Model

### Function Execution Flow

```
Call: result := add(5, 3)
      ↓
1. Go to add function
      ↓
2. Create local variables: a=5, b=3
      ↓
3. Execute function body: a + b = 8
      ↓
4. Return value: 8
      ↓
5. Assign to result: result = 8
      ↓
6. Continue with next statement
```

### Multiple Return Values

```
Call: result, err := divide(10, 2)
      ↓
1. Go to divide function
      ↓
2. Check: b == 0? No
      ↓
3. Calculate: 10 / 2 = 5.0
      ↓
4. Return: (5.0, nil)
      ↓
5. Assign: result = 5.0, err = nil
```

### Named Return Values

```
func calculate(x, y int) (sum int, product int) {
    // sum and product are already declared!
    sum = x + y        // Assign to named return
    product = x * y    // Assign to named return
    return             // Returns (sum, product)
}
```

**Memory:**
- Named returns are variables in the function
- You can modify them
- `return` without values uses them automatically

---

## f) Demo Example

### Complete Example with All Function Types

```go
package main

import (
    "errors"
    "fmt"
)

// 1. Basic function
func greet() {
    fmt.Println("Hello, World!")
}

// 2. Function with parameter
func greetPerson(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// 3. Function with return value
func add(a int, b int) int {
    return a + b
}

// 4. Multiple parameters (same type can be shortened)
func multiply(a, b int) int {
    return a * b
}

// 5. Multiple return values (Go's special feature!)
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// 6. Named return values
func calculate(x, y int) (sum int, product int) {
    sum = x + y
    product = x * y
    return  // Naked return - uses named values
}

// 7. Variadic function (variable arguments)
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

// 8. Function as value
var subtract = func(a, b int) int {
    return a - b
}

func main() {
    // Call basic function
    greet()
    
    // Call with parameter
    greetPerson("Alice")
    
    // Call with return value
    result := add(5, 3)
    fmt.Printf("5 + 3 = %d\n", result)
    
    // Call with multiple returns (must handle both!)
    quotient, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("10 / 2 = %.2f\n", quotient)
    }
    
    // Error case
    _, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    }
    
    // Named returns
    s, p := calculate(4, 5)
    fmt.Printf("Sum: %d, Product: %d\n", s, p)
    
    // Variadic function
    total := sum(1, 2, 3, 4, 5)
    fmt.Printf("Sum of 1,2,3,4,5 = %d\n", total)
    
    // Function as value
    diff := subtract(10, 3)
    fmt.Printf("10 - 3 = %d\n", diff)
}
```

**Line-by-line explanation:**

1. **Basic function**: No parameters, no return
2. **Parameter**: `name string` - type after name
3. **Return value**: `int` after parameters
4. **Short parameter syntax**: `a, b int` instead of `a int, b int`
5. **Multiple returns**: `(float64, error)` - Go's error handling pattern
6. **Named returns**: Variables created automatically, can use naked return
7. **Variadic**: `...int` accepts any number of ints
8. **Function value**: Assign function to variable

**Output:**
```
Hello, World!
Hello, Alice!
5 + 3 = 8
10 / 2 = 5.00
Error: division by zero
Sum: 9, Product: 20
Sum of 1,2,3,4,5 = 15
10 - 3 = 7
```

---

## g) Use Cases

### When to use different function patterns

**Basic function:**
```go
func initialize() {
    // Setup code
}
```
- One-time setup
- No input/output needed

**Single return:**
```go
func calculate(x int) int {
    return x * 2
}
```
- Simple calculations
- No error possibility

**Multiple returns (result, error):**
```go
func readFile(name string) ([]byte, error) {
    // Can fail, so return error
}
```
- Operations that can fail
- Go's standard pattern
- **Always check the error!**

**Named returns:**
```go
func parse(data string) (name string, age int, err error) {
    // Multiple related returns
}
```
- Multiple related values
- Makes code more readable
- Use sparingly (can be confusing)

**Variadic functions:**
```go
func log(level string, messages ...string) {
    // Accept any number of strings
}
```
- When number of arguments varies
- Like Python's `*args`

---

## h) Do's and Don'ts / Best Practices

### ✅ Do's

1. **Always handle errors**
   ```go
   result, err := divide(10, 2)
   if err != nil {
       // Handle error
   }
   ```

2. **Use meaningful function names**
   ```go
   func calculateTotal()  // ✅ Clear
   func calc()           // ❌ Unclear
   ```

3. **Keep functions small and focused**
   - One responsibility
   - Easy to test
   - Easy to understand

4. **Use multiple returns for errors**
   ```go
   func doSomething() (result Type, err error) {
       // Standard Go pattern
   }
   ```

5. **Document complex functions**
   ```go
   // calculateTotal computes the sum of all items
   // Returns the total and any calculation error
   func calculateTotal(items []Item) (float64, error) {
   ```

### ❌ Don'ts

1. **Don't ignore errors**
   ```go
   result, _ := divide(10, 0)  // ❌ Bad!
   result, err := divide(10, 0)
   if err != nil {  // ✅ Good!
       // handle
   }
   ```

2. **Don't overuse named returns**
   ```go
   // Simple case - don't need named returns
   func add(a, b int) int {  // ✅
       return a + b
   }
   ```

3. **Don't make functions too long**
   - If > 50 lines, consider breaking up
   - Each function should do one thing

4. **Don't use global variables**
   - Pass parameters instead
   - Makes functions testable

5. **Don't return nil for errors when successful**
   ```go
   return result, nil  // ✅ Correct
   return result, err  // Only if err is nil
   ```

---

## i) Solved Practice Examples

### Example 1: Basic Calculator Functions

**Task:** Create functions for add, subtract, multiply, divide with error handling.

**Solution:**
```go
package main

import (
    "errors"
    "fmt"
)

func add(a, b float64) float64 {
    return a + b
}

func subtract(a, b float64) float64 {
    return a - b
}

func multiply(a, b float64) float64 {
    return a * b
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    x, y := 10.0, 3.0
    
    fmt.Printf("%.2f + %.2f = %.2f\n", x, y, add(x, y))
    fmt.Printf("%.2f - %.2f = %.2f\n", x, y, subtract(x, y))
    fmt.Printf("%.2f * %.2f = %.2f\n", x, y, multiply(x, y))
    
    result, err := divide(x, y)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("%.2f / %.2f = %.2f\n", x, y, result)
    }
}
```

### Example 2: Variadic Function

**Task:** Create a function that finds the maximum of any number of integers.

**Solution:**
```go
package main

import "fmt"

func max(numbers ...int) int {
    if len(numbers) == 0 {
        return 0
    }
    
    maxVal := numbers[0]
    for _, n := range numbers {
        if n > maxVal {
            maxVal = n
        }
    }
    return maxVal
}

func main() {
    fmt.Println("Max of 1, 5, 3:", max(1, 5, 3))
    fmt.Println("Max of 10, 20, 30, 5:", max(10, 20, 30, 5))
    fmt.Println("Max of single value:", max(42))
}
```

### Example 3: Function Returning Function

**Task:** Create a function that returns a greeting function.

**Solution:**
```go
package main

import "fmt"

func makeGreeter(greeting string) func(string) {
    return func(name string) {
        fmt.Printf("%s, %s!\n", greeting, name)
    }
}

func main() {
    helloGreeter := makeGreeter("Hello")
    hiGreeter := makeGreeter("Hi")
    
    helloGreeter("Alice")  // Output: Hello, Alice!
    hiGreeter("Bob")       // Output: Hi, Bob!
}
```

---

## j) Quiz / Practice for Me

### Quiz Questions

1. **How do you return multiple values in Go?**
   - [ ] `return (a, b)`
   - [ ] `return a, b`
   - [ ] `return [a, b]`
   - [ ] `return {a, b}`

2. **What is the standard error handling pattern in Go?**
   - [ ] Exceptions
   - [ ] Return error as last value
   - [ ] Global error variable
   - [ ] Error callbacks

3. **Can Go functions have default parameters?**
   - [ ] Yes
   - [ ] No
   - [ ] Only for some types
   - [ ] Only in structs

4. **What does `...int` mean in a function parameter?**
   - [ ] Optional parameter
   - [ ] Variadic (any number of ints)
   - [ ] Array parameter
   - [ ] Pointer parameter

### Practice Tasks

**Task 1: Temperature Converter**
- Create function `celsiusToFahrenheit(c float64) float64`
- Formula: `F = C * 9/5 + 32`
- Create function `fahrenheitToCelsius(f float64) float64`
- Test both functions

**Task 2: Safe Division with Multiple Returns**
- Create function `safeDivide(a, b float64) (float64, error)`
- Return error if b is 0
- Return result and nil error if successful
- Test both success and error cases

**Task 3: Variadic Average**
- Create function `average(numbers ...float64) float64`
- Calculate average of any number of floats
- Handle empty input (return 0)

### Answers

**Quiz Answers:**
1. `return a, b`
2. Return error as last value
3. No
4. Variadic (any number of ints)

**Practice Solutions:**

**Task 1 Solution:**
```go
package main

import "fmt"

func celsiusToFahrenheit(c float64) float64 {
    return c*9/5 + 32
}

func fahrenheitToCelsius(f float64) float64 {
    return (f - 32) * 5 / 9
}

func main() {
    c := 25.0
    f := celsiusToFahrenheit(c)
    fmt.Printf("%.1f°C = %.1f°F\n", c, f)
    
    f2 := 77.0
    c2 := fahrenheitToCelsius(f2)
    fmt.Printf("%.1f°F = %.1f°C\n", f2, c2)
}
```

**Task 2 Solution:**
```go
package main

import (
    "errors"
    "fmt"
)

func safeDivide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    // Success case
    result, err := safeDivide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("10 / 2 = %.2f\n", result)
    }
    
    // Error case
    result, err = safeDivide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Result: %.2f\n", result)
    }
}
```

**Task 3 Solution:**
```go
package main

import "fmt"

func average(numbers ...float64) float64 {
    if len(numbers) == 0 {
        return 0
    }
    
    sum := 0.0
    for _, n := range numbers {
        sum += n
    }
    return sum / float64(len(numbers))
}

func main() {
    fmt.Printf("Average of 1,2,3,4,5: %.2f\n", average(1, 2, 3, 4, 5))
    fmt.Printf("Average of 10,20,30: %.2f\n", average(10, 20, 30))
    fmt.Printf("Average of empty: %.2f\n", average())
}
```

---

## Key Takeaways

1. **Functions are first-class** - Can assign to variables, pass as parameters
2. **Multiple returns** - Go's signature feature, especially for errors
3. **Error handling** - Return error as value, always check it
4. **No default parameters** - Must pass all parameters (use structs if needed)
5. **Variadic functions** - `...type` for variable arguments
6. **Named returns** - Can make code clearer, use sparingly

---

## Must Remember Forever

- Function syntax: `func name(params) returnType { }`
- Multiple returns: `func name() (type1, type2) { return val1, val2 }`
- Error pattern: `(result, error)` - always check error!
- Variadic: `func name(args ...type) { }`
- No default parameters - Go values simplicity
- Functions are values - can assign, pass, return

---

---

## k) Additional Practice Problems (Build Strong Logic)

### 10 More Solved Examples

### Solved Problem 1: Prime Number Checker

**Task:** Create a function that checks if a number is prime.

**Solution:**
```go
package main

import (
    "fmt"
    "math"
)

func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    if n == 2 {
        return true
    }
    if n%2 == 0 {
        return false
    }
    
    sqrt := int(math.Sqrt(float64(n)))
    for i := 3; i <= sqrt; i += 2 {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func main() {
    numbers := []int{2, 3, 4, 17, 25, 29, 100}
    for _, n := range numbers {
        fmt.Printf("%d is prime: %t\n", n, isPrime(n))
    }
}
```

### Solved Problem 2: Fibonacci Sequence Generator

**Task:** Create a function that generates the first n Fibonacci numbers.

**Solution:**
```go
package main

import "fmt"

func fibonacci(n int) []int {
    if n <= 0 {
        return []int{}
    }
    if n == 1 {
        return []int{0}
    }
    if n == 2 {
        return []int{0, 1}
    }
    
    fib := []int{0, 1}
    for i := 2; i < n; i++ {
        fib = append(fib, fib[i-1]+fib[i-2])
    }
    return fib
}

func main() {
    result := fibonacci(10)
    fmt.Println("First 10 Fibonacci numbers:", result)
}
```

### Solved Problem 3: String Reverser with Multiple Returns

**Task:** Create a function that reverses a string and returns both the reversed string and its length.

**Solution:**
```go
package main

import "fmt"

func reverseString(s string) (string, int) {
    runes := []rune(s)
    length := len(runes)
    
    for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    
    return string(runes), length
}

func main() {
    original := "Hello, World!"
    reversed, length := reverseString(original)
    fmt.Printf("Original: %s\n", original)
    fmt.Printf("Reversed: %s\n", reversed)
    fmt.Printf("Length: %d\n", length)
}
```

### Solved Problem 4: Greatest Common Divisor (GCD)

**Task:** Implement Euclidean algorithm to find GCD of two numbers.

**Solution:**
```go
package main

import "fmt"

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func main() {
    fmt.Printf("GCD of 48 and 18: %d\n", gcd(48, 18))
    fmt.Printf("GCD of 100 and 25: %d\n", gcd(100, 25))
    fmt.Printf("GCD of 17 and 13: %d\n", gcd(17, 13))
}
```

### Solved Problem 5: Factorial with Error Handling

**Task:** Create a factorial function that handles negative numbers and overflow.

**Solution:**
```go
package main

import (
    "errors"
    "fmt"
)

func factorial(n int) (int, error) {
    if n < 0 {
        return 0, errors.New("factorial is not defined for negative numbers")
    }
    if n > 20 {
        return 0, errors.New("number too large, would cause overflow")
    }
    
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result, nil
}

func main() {
    tests := []int{-5, 0, 5, 10, 25}
    for _, n := range tests {
        result, err := factorial(n)
        if err != nil {
            fmt.Printf("factorial(%d): Error - %v\n", n, err)
        } else {
            fmt.Printf("factorial(%d) = %d\n", n, result)
        }
    }
}
```

### Solved Problem 6: Password Validator

**Task:** Create a function that validates password strength with multiple criteria.

**Solution:**
```go
package main

import (
    "errors"
    "fmt"
    "strings"
    "unicode"
)

func validatePassword(password string) error {
    if len(password) < 8 {
        return errors.New("password must be at least 8 characters")
    }
    
    hasUpper := false
    hasLower := false
    hasDigit := false
    hasSpecial := false
    
    for _, char := range password {
        switch {
        case unicode.IsUpper(char):
            hasUpper = true
        case unicode.IsLower(char):
            hasLower = true
        case unicode.IsDigit(char):
            hasDigit = true
        case strings.ContainsRune("!@#$%^&*", char):
            hasSpecial = true
        }
    }
    
    var missing []string
    if !hasUpper {
        missing = append(missing, "uppercase letter")
    }
    if !hasLower {
        missing = append(missing, "lowercase letter")
    }
    if !hasDigit {
        missing = append(missing, "digit")
    }
    if !hasSpecial {
        missing = append(missing, "special character")
    }
    
    if len(missing) > 0 {
        return fmt.Errorf("password missing: %v", missing)
    }
    
    return nil
}

func main() {
    passwords := []string{
        "weak",
        "Weak123",
        "Strong@123",
        "NoSpecial123",
    }
    
    for _, pwd := range passwords {
        err := validatePassword(pwd)
        if err != nil {
            fmt.Printf("'%s': %v\n", pwd, err)
        } else {
            fmt.Printf("'%s': Valid password\n", pwd)
        }
    }
}
```

### Solved Problem 7: Number to Words Converter

**Task:** Convert a number (0-99) to its word representation.

**Solution:**
```go
package main

import "fmt"

func numberToWords(n int) (string, error) {
    if n < 0 || n > 99 {
        return "", fmt.Errorf("number must be between 0 and 99")
    }
    
    ones := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    teens := []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
    tens := []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
    
    if n == 0 {
        return "zero", nil
    }
    
    if n < 10 {
        return ones[n], nil
    }
    
    if n < 20 {
        return teens[n-10], nil
    }
    
    ten := n / 10
    one := n % 10
    
    if one == 0 {
        return tens[ten], nil
    }
    
    return tens[ten] + "-" + ones[one], nil
}

func main() {
    numbers := []int{0, 5, 12, 25, 33, 99, 100}
    for _, n := range numbers {
        word, err := numberToWords(n)
        if err != nil {
            fmt.Printf("%d: Error - %v\n", n, err)
        } else {
            fmt.Printf("%d: %s\n", n, word)
        }
    }
}
```

### Solved Problem 8: Binary Search

**Task:** Implement binary search on a sorted slice.

**Solution:**
```go
package main

import "fmt"

func binarySearch(arr []int, target int) (int, bool) {
    left, right := 0, len(arr)-1
    
    for left <= right {
        mid := (left + right) / 2
        
        if arr[mid] == target {
            return mid, true
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return -1, false
}

func main() {
    arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
    
    targets := []int{7, 10, 1, 19, 20}
    for _, target := range targets {
        index, found := binarySearch(arr, target)
        if found {
            fmt.Printf("Found %d at index %d\n", target, index)
        } else {
            fmt.Printf("%d not found\n", target)
        }
    }
}
```

### Solved Problem 9: Palindrome Checker

**Task:** Check if a string is a palindrome (ignoring case and spaces).

**Solution:**
```go
package main

import (
    "fmt"
    "strings"
)

func isPalindrome(s string) bool {
    // Remove spaces and convert to lowercase
    cleaned := strings.ReplaceAll(strings.ToLower(s), " ", "")
    
    runes := []rune(cleaned)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        if runes[i] != runes[j] {
            return false
        }
    }
    return true
}

func main() {
    tests := []string{
        "racecar",
        "A man a plan a canal Panama",
        "hello",
        "Madam",
        "not a palindrome",
    }
    
    for _, test := range tests {
        fmt.Printf("'%s' is palindrome: %t\n", test, isPalindrome(test))
    }
}
```

### Solved Problem 10: Matrix Operations

**Task:** Create functions to add and multiply matrices.

**Solution:**
```go
package main

import (
    "errors"
    "fmt"
)

func addMatrices(a, b [][]int) ([][]int, error) {
    if len(a) != len(b) || len(a[0]) != len(b[0]) {
        return nil, errors.New("matrices must have same dimensions")
    }
    
    result := make([][]int, len(a))
    for i := range result {
        result[i] = make([]int, len(a[0]))
        for j := range result[i] {
            result[i][j] = a[i][j] + b[i][j]
        }
    }
    return result, nil
}

func multiplyMatrices(a, b [][]int) ([][]int, error) {
    if len(a[0]) != len(b) {
        return nil, errors.New("invalid dimensions for multiplication")
    }
    
    result := make([][]int, len(a))
    for i := range result {
        result[i] = make([]int, len(b[0]))
        for j := range result[i] {
            for k := 0; k < len(b); k++ {
                result[i][j] += a[i][k] * b[k][j]
            }
        }
    }
    return result, nil
}

func printMatrix(m [][]int) {
    for _, row := range m {
        fmt.Println(row)
    }
}

func main() {
    a := [][]int{{1, 2}, {3, 4}}
    b := [][]int{{5, 6}, {7, 8}}
    
    sum, _ := addMatrices(a, b)
    fmt.Println("Sum:")
    printMatrix(sum)
    
    product, _ := multiplyMatrices(a, b)
    fmt.Println("\nProduct:")
    printMatrix(product)
}
```

---

### 10 More Practice Problems (Solve These!)

**Problem 1: Perfect Number Checker**
- A perfect number equals the sum of its proper divisors
- Example: 6 = 1 + 2 + 3
- Create function `isPerfect(n int) bool`
- Test with: 6, 28, 496, 10

**Problem 2: Armstrong Number**
- An Armstrong number equals the sum of its digits each raised to the power of the number of digits
- Example: 153 = 1³ + 5³ + 3³
- Create function `isArmstrong(n int) bool`

**Problem 3: Roman Numeral Converter**
- Convert integer (1-3999) to Roman numeral
- Create function `intToRoman(n int) (string, error)`
- Test with: 4, 9, 58, 1994

**Problem 4: Anagram Checker**
- Check if two strings are anagrams (same letters, different order)
- Create function `areAnagrams(s1, s2 string) bool`
- Ignore case and spaces

**Problem 5: Collatz Sequence**
- For any number n: if even, divide by 2; if odd, multiply by 3 and add 1
- Repeat until you reach 1
- Create function `collatzSequence(n int) []int` that returns the sequence
- Test with: 6, 11, 27

**Problem 6: Prime Factorization**
- Find all prime factors of a number
- Create function `primeFactors(n int) []int`
- Example: 60 = 2 × 2 × 3 × 5
- Return: [2, 2, 3, 5]

**Problem 7: String Compression**
- Compress string: "aaabbbccc" → "a3b3c3"
- Create function `compress(s string) string`
- If compressed is longer, return original

**Problem 8: Calculator with Operations**
- Create function that takes two numbers and operation (+, -, *, /)
- Return result and error
- Handle division by zero

**Problem 9: Find Missing Number**
- Given array of n-1 numbers from 1 to n, find the missing number
- Create function `findMissing(arr []int) int`
- Example: [1, 2, 4, 5] → missing 3

**Problem 10: Power Function**
- Implement power function: x^n
- Handle negative exponents
- Create function `power(x float64, n int) float64`
- Use efficient algorithm (don't just multiply n times)

---

## Next Steps

Now you understand functions. Next:
- Control flow
- if/else, for loops, switch statements
- How to make decisions and repeat code

**Ready? → [05_CONTROL_FLOW.md](./05_CONTROL_FLOW.md)**
