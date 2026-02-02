# Control Flow

## a) Overview

### What this topic is
How to make decisions (if/else) and repeat code (loops) in Go. Go has simple, clear control flow statements.

### Why it exists in Go
Control flow lets your program make decisions and repeat actions. Go keeps it simple - only `if`, `for`, and `switch` (no `while`, no `do-while`).

---

## b) Syntax

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
    // x is available here
}
```

### For Loops
```go
// Traditional for loop
for i := 0; i < 10; i++ {
    // code
}

// While-style loop
for condition {
    // code
}

// Infinite loop
for {
    // code
    break  // Exit loop
}

// Range loop (for slices, maps, etc.)
for index, value := range slice {
    // code
}
```

### Switch
```go
switch value {
case option1:
    // code
case option2:
    // code
default:
    // code
}

// Switch with no value (like if/else chain)
switch {
case condition1:
    // code
case condition2:
    // code
}
```

---

## c) Explanation

### Step-by-Step Control Flow

**1. If statements**
- `if` checks a condition (must be `bool`)
- `else` runs if condition is false
- `else if` checks another condition
- Curly braces `{ }` are required (even for one line)

**2. For loops**
- Go only has `for` (no `while`, no `do-while`)
- Three parts: initialization, condition, increment
- Can omit parts to create different loop types
- `break` exits loop
- `continue` skips to next iteration

**3. Switch statements**
- Like `if/else` chain but cleaner
- Can switch on value or condition
- `default` is optional
- No `break` needed (unlike C/Java) - Go breaks automatically

### Characteristics

- **Simple**: Only `if`, `for`, `switch` - that's it!
- **No parentheses**: Conditions don't need `()`
- **Curly braces required**: Even for one-line blocks
- **Break/continue**: Control loop flow
- **Range**: Special loop for collections

---

## d) Python Comparison

### Python Control Flow
```python
# If/else
if x > 0:
    print("positive")
elif x < 0:
    print("negative")
else:
    print("zero")

# For loop
for i in range(10):
    print(i)

# While loop
while condition:
    print("looping")

# Switch (Python 3.10+)
match value:
    case 1:
        print("one")
    case 2:
        print("two")
```

### Go Control Flow
```go
// If/else
if x > 0 {
    fmt.Println("positive")
} else if x < 0 {
    fmt.Println("negative")
} else {
    fmt.Println("zero")
}

// For loop (traditional)
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// While-style (using for)
for condition {
    fmt.Println("looping")
}

// Switch
switch value {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
}
```

**Key Differences:**

| Feature | Python | Go |
|---------|--------|-----|
| **If parentheses** | Not needed | Not needed |
| **Indentation** | Required | Not used (uses `{ }`) |
| **For loop** | `for i in range(n)` | `for i := 0; i < n; i++` |
| **While loop** | `while condition` | `for condition` |
| **Switch** | `match` (Python 3.10+) | `switch` (always available) |
| **Break in switch** | Not needed | Not needed (auto-break) |

**Thinking Difference:**
- Python: Indentation-based blocks
- Go: Curly brace blocks
- Python: `for` and `while` are separate
- Go: `for` does everything (simpler!)

---

## e) Visual Flow / Mental Model

### If Statement Flow

```
if x > 0 {
    // Execute this
}
```

**Flow:**
```
1. Check: x > 0?
   ↓
2. If true → Execute block
   ↓
3. If false → Skip block
   ↓
4. Continue to next statement
```

### For Loop Flow

```
for i := 0; i < 10; i++ {
    // code
}
```

**Flow:**
```
1. Initialize: i = 0
   ↓
2. Check: i < 10? Yes
   ↓
3. Execute block
   ↓
4. Increment: i++
   ↓
5. Check: i < 10? Yes
   ↓
6. Execute block
   ↓
... (repeat until i >= 10)
   ↓
7. Exit loop
```

### Range Loop Flow

```
for index, value := range items {
    // code
}
```

**Flow:**
```
1. Get first item: index=0, value=items[0]
   ↓
2. Execute block
   ↓
3. Get next item: index=1, value=items[1]
   ↓
4. Execute block
   ↓
... (repeat for all items)
   ↓
5. Exit loop
```

---

## f) Demo Example

### Complete Example with All Control Flow

```go
package main

import "fmt"

func main() {
    // 1. Basic if
    x := 10
    if x > 5 {
        fmt.Println("x is greater than 5")
    }
    
    // 2. If/else
    if x%2 == 0 {
        fmt.Println("x is even")
    } else {
        fmt.Println("x is odd")
    }
    
    // 3. If/else if/else
    score := 85
    if score >= 90 {
        fmt.Println("Grade: A")
    } else if score >= 80 {
        fmt.Println("Grade: B")
    } else if score >= 70 {
        fmt.Println("Grade: C")
    } else {
        fmt.Println("Grade: F")
    }
    
    // 4. If with initialization
    if y := 20; y > 10 {
        fmt.Printf("y (%d) is greater than 10\n", y)
        // y is only available in this if block
    }
    
    // 5. Traditional for loop
    fmt.Println("Counting 1 to 5:")
    for i := 1; i <= 5; i++ {
        fmt.Printf("%d ", i)
    }
    fmt.Println()
    
    // 6. While-style loop (using for)
    fmt.Println("Countdown:")
    count := 5
    for count > 0 {
        fmt.Printf("%d ", count)
        count--
    }
    fmt.Println("Blast off!")
    
    // 7. Infinite loop with break
    fmt.Println("Finding first even number > 10:")
    for {
        num := 11
        if num%2 == 0 {
            fmt.Println("Found:", num)
            break
        }
        num++
        if num > 20 {
            break
        }
    }
    
    // 8. Continue statement
    fmt.Println("Odd numbers 1-10:")
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            continue  // Skip even numbers
        }
        fmt.Printf("%d ", i)
    }
    fmt.Println()
    
    // 9. Range loop (slices)
    fruits := []string{"apple", "banana", "cherry"}
    fmt.Println("Fruits:")
    for index, fruit := range fruits {
        fmt.Printf("  %d: %s\n", index, fruit)
    }
    
    // 10. Range loop (ignore index)
    fmt.Println("Fruits (values only):")
    for _, fruit := range fruits {
        fmt.Printf("  %s\n", fruit)
    }
    
    // 11. Switch statement
    day := "Monday"
    switch day {
    case "Monday":
        fmt.Println("Start of work week")
    case "Friday":
        fmt.Println("TGIF!")
    case "Saturday", "Sunday":
        fmt.Println("Weekend!")
    default:
        fmt.Println("Regular day")
    }
    
    // 12. Switch with no value (like if/else)
    hour := 14
    switch {
    case hour < 12:
        fmt.Println("Good morning!")
    case hour < 18:
        fmt.Println("Good afternoon!")
    default:
        fmt.Println("Good evening!")
    }
}
```

**Line-by-line explanation:**

1. **Basic if**: Simple condition check
2. **If/else**: Two paths
3. **If/else if/else**: Multiple conditions
4. **If with init**: Declare variable in if (scoped to if block)
5. **Traditional for**: Three-part loop
6. **While-style**: `for condition` (Go's way of doing while)
7. **Infinite loop**: `for { }` with `break`
8. **Continue**: Skip to next iteration
9. **Range with index**: Get both index and value
10. **Range without index**: Use `_` to ignore index
11. **Switch on value**: Match against specific value
12. **Switch on condition**: Like if/else chain

**Output:**
```
x is greater than 5
x is even
Grade: B
y (20) is greater than 10
Counting 1 to 5:
1 2 3 4 5 
Countdown:
5 4 3 2 1 Blast off!
Finding first even number > 10:
Found: 12
Odd numbers 1-10:
1 3 5 7 9 
Fruits:
  0: apple
  1: banana
  2: cherry
Fruits (values only):
  apple
  banana
  cherry
Start of work week
Good afternoon!
```

---

## g) Use Cases

### When to use different control structures

**If/else:**
- Simple conditions
- Two-way decisions
- Multiple conditions with else if

**For loop (traditional):**
- Known number of iterations
- Counting loops
- Index-based access

**For loop (while-style):**
- Unknown number of iterations
- Condition-based loops
- Event-driven loops

**Range loop:**
- Iterating over slices
- Iterating over maps
- Iterating over strings
- When you need index/value

**Switch:**
- Multiple discrete values
- Cleaner than long if/else chains
- Pattern matching

---

## h) Do's and Don'ts / Best Practices

### ✅ Do's

1. **Use if with initialization for scoped variables**
   ```go
   if err := doSomething(); err != nil {
       // err only available here
   }
   ```

2. **Use range for slices/maps**
   ```go
   for _, item := range items {  // ✅
       // process item
   }
   ```

3. **Use switch for multiple values**
   ```go
   switch status {
   case "active", "pending":  // Multiple values
       // code
   }
   ```

4. **Use break/continue appropriately**
   ```go
   for {
       if condition {
           break  // Exit loop
       }
       if skipCondition {
           continue  // Skip to next iteration
       }
   }
   ```

5. **Keep loops simple**
   - One responsibility
   - Easy to understand
   - Avoid nested loops when possible

### ❌ Don'ts

1. **Don't use parentheses in conditions**
   ```go
   if (x > 0) {  // ❌ Unnecessary
   if x > 0 {    // ✅ Correct
   ```

2. **Don't forget curly braces**
   ```go
   if x > 0
       fmt.Println("positive")  // ❌ ERROR!
   
   if x > 0 {  // ✅ Correct
       fmt.Println("positive")
   }
   ```

3. **Don't use while (doesn't exist)**
   ```go
   while condition {  // ❌ Doesn't exist
   for condition {    // ✅ Correct
   ```

4. **Don't create infinite loops accidentally**
   ```go
   for i := 0; i < 10; {  // ❌ Missing increment!
       // This loops forever
   }
   ```

5. **Don't overuse nested loops**
   - Hard to read
   - Consider extracting to function

---

## i) Solved Practice Examples

### Example 1: Number Classifier

**Task:** Classify a number as positive, negative, or zero, and as even or odd.

**Solution:**
```go
package main

import "fmt"

func classifyNumber(n int) {
    // Classify sign
    if n > 0 {
        fmt.Printf("%d is positive", n)
    } else if n < 0 {
        fmt.Printf("%d is negative", n)
    } else {
        fmt.Printf("%d is zero", n)
    }
    
    // Classify parity
    if n%2 == 0 {
        fmt.Println(" and even")
    } else {
        fmt.Println(" and odd")
    }
}

func main() {
    classifyNumber(10)
    classifyNumber(-5)
    classifyNumber(0)
}
```

### Example 2: Sum with Loop

**Task:** Sum numbers from 1 to n using a for loop.

**Solution:**
```go
package main

import "fmt"

func sumToN(n int) int {
    total := 0
    for i := 1; i <= n; i++ {
        total += i
    }
    return total
}

func main() {
    fmt.Println("Sum 1 to 10:", sumToN(10))
    fmt.Println("Sum 1 to 100:", sumToN(100))
}
```

### Example 3: FizzBuzz

**Task:** Print numbers 1-20, but:
- "Fizz" for multiples of 3
- "Buzz" for multiples of 5
- "FizzBuzz" for multiples of both

**Solution:**
```go
package main

import "fmt"

func main() {
    for i := 1; i <= 20; i++ {
        if i%3 == 0 && i%5 == 0 {
            fmt.Println("FizzBuzz")
        } else if i%3 == 0 {
            fmt.Println("Fizz")
        } else if i%5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
    }
}
```

---

## j) Quiz / Practice for Me

### Quiz Questions

1. **What loop type does Go use for "while" loops?**
   - [ ] `while`
   - [ ] `for` (with condition only)
   - [ ] `loop`
   - [ ] `repeat`

2. **Do you need `break` in Go switch statements?**
   - [ ] Yes, always
   - [ ] No, Go breaks automatically
   - [ ] Only for default case
   - [ ] Only in some cases

3. **What does `continue` do?**
   - [ ] Exits the loop
   - [ ] Skips to next iteration
   - [ ] Restarts the loop
   - [ ] Does nothing

4. **Can you declare a variable in an if statement?**
   - [ ] No
   - [ ] Yes, with `:=`
   - [ ] Only with `var`
   - [ ] Only for some types

### Practice Tasks

**Task 1: Grade Calculator**
- Write a function that takes a score (0-100)
- Use if/else to return grade: A (90+), B (80+), C (70+), D (60+), F (<60)

**Task 2: Factorial**
- Write a function `factorial(n int) int`
- Use a for loop to calculate n!
- Example: factorial(5) = 5 * 4 * 3 * 2 * 1 = 120

**Task 3: Find Maximum**
- Write a function that takes a slice of integers
- Use range loop to find the maximum value
- Return the maximum

### Answers

**Quiz Answers:**
1. `for` (with condition only)
2. No, Go breaks automatically
3. Skips to next iteration
4. Yes, with `:=`

**Practice Solutions:**

**Task 1 Solution:**
```go
package main

import "fmt"

func getGrade(score int) string {
    if score >= 90 {
        return "A"
    } else if score >= 80 {
        return "B"
    } else if score >= 70 {
        return "C"
    } else if score >= 60 {
        return "D"
    } else {
        return "F"
    }
}

func main() {
    fmt.Println("Score 95:", getGrade(95))
    fmt.Println("Score 85:", getGrade(85))
    fmt.Println("Score 75:", getGrade(75))
    fmt.Println("Score 65:", getGrade(65))
    fmt.Println("Score 55:", getGrade(55))
}
```

**Task 2 Solution:**
```go
package main

import "fmt"

func factorial(n int) int {
    result := 1
    for i := 1; i <= n; i++ {
        result *= i
    }
    return result
}

func main() {
    fmt.Println("Factorial of 5:", factorial(5))
    fmt.Println("Factorial of 7:", factorial(7))
}
```

**Task 3 Solution:**
```go
package main

import "fmt"

func findMax(numbers []int) int {
    if len(numbers) == 0 {
        return 0
    }
    
    max := numbers[0]
    for _, num := range numbers {
        if num > max {
            max = num
        }
    }
    return max
}

func main() {
    nums := []int{3, 7, 2, 9, 1, 5}
    fmt.Println("Maximum:", findMax(nums))
}
```

---

## Key Takeaways

1. **Go is simple** - Only `if`, `for`, `switch` (no while, no do-while)
2. **Curly braces required** - Even for one-line blocks
3. **For does everything** - Traditional, while-style, infinite, range
4. **Switch is powerful** - Can switch on value or condition
5. **Range is essential** - For iterating collections
6. **Break/continue** - Control loop flow

---

## Must Remember Forever

- `if condition { }` - No parentheses needed
- `for i := 0; i < n; i++ { }` - Traditional loop
- `for condition { }` - While-style loop
- `for { }` - Infinite loop (use break)
- `for i, v := range slice { }` - Range loop
- `switch value { case x: }` - Switch statement
- Go breaks automatically in switch (no break needed)

---

---

## k) Additional Practice Problems (Build Strong Logic)

### 10 More Solved Examples

### Solved Problem 1: Number Pattern Generator

**Task:** Print patterns like:
```
1
12
123
1234
```

**Solution:**
```go
package main

import "fmt"

func printPattern(n int) {
    for i := 1; i <= n; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print(j)
        }
        fmt.Println()
    }
}

func main() {
    printPattern(5)
}
```

### Solved Problem 2: Diamond Pattern

**Task:** Print a diamond pattern of asterisks.

**Solution:**
```go
package main

import "fmt"

func printDiamond(n int) {
    // Upper half
    for i := 1; i <= n; i++ {
        // Spaces
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }
        // Stars
        for j := 1; j <= 2*i-1; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
    
    // Lower half
    for i := n - 1; i >= 1; i-- {
        // Spaces
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }
        // Stars
        for j := 1; j <= 2*i-1; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}

func main() {
    printDiamond(5)
}
```

### Solved Problem 3: Multiplication Table

**Task:** Generate multiplication table for a number.

**Solution:**
```go
package main

import "fmt"

func multiplicationTable(n int) {
    fmt.Printf("Multiplication table for %d:\n", n)
    for i := 1; i <= 10; i++ {
        fmt.Printf("%d × %d = %d\n", n, i, n*i)
    }
}

func main() {
    multiplicationTable(7)
}
```

### Solved Problem 4: Guess the Number Game Logic

**Task:** Implement logic for a number guessing game.

**Solution:**
```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func guessNumber() {
    rand.Seed(time.Now().UnixNano())
    target := rand.Intn(100) + 1
    attempts := 0
    maxAttempts := 7
    
    fmt.Println("Guess a number between 1 and 100!")
    
    for attempts < maxAttempts {
        attempts++
        fmt.Printf("Attempt %d/%d: Enter your guess: ", attempts, maxAttempts)
        
        var guess int
        fmt.Scan(&guess)
        
        if guess == target {
            fmt.Printf("Congratulations! You guessed it in %d attempts!\n", attempts)
            return
        } else if guess < target {
            fmt.Println("Too low! Try again.")
        } else {
            fmt.Println("Too high! Try again.")
        }
    }
    
    fmt.Printf("Game over! The number was %d\n", target)
}

func main() {
    guessNumber()
}
```

### Solved Problem 5: Prime Numbers in Range

**Task:** Find all prime numbers between two numbers.

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

func primesInRange(start, end int) []int {
    var primes []int
    for i := start; i <= end; i++ {
        if isPrime(i) {
            primes = append(primes, i)
        }
    }
    return primes
}

func main() {
    primes := primesInRange(10, 50)
    fmt.Println("Prime numbers between 10 and 50:", primes)
}
```

### Solved Problem 6: Pascal's Triangle

**Task:** Generate first n rows of Pascal's triangle.

**Solution:**
```go
package main

import "fmt"

func pascalTriangle(n int) {
    for i := 0; i < n; i++ {
        // Print spaces
        for j := 0; j < n-i-1; j++ {
            fmt.Print(" ")
        }
        
        // Calculate and print values
        value := 1
        for j := 0; j <= i; j++ {
            fmt.Printf("%d ", value)
            value = value * (i - j) / (j + 1)
        }
        fmt.Println()
    }
}

func main() {
    pascalTriangle(6)
}
```

### Solved Problem 7: Number to Binary Converter

**Task:** Convert decimal number to binary using loops.

**Solution:**
```go
package main

import "fmt"

func decimalToBinary(n int) string {
    if n == 0 {
        return "0"
    }
    
    var binary string
    for n > 0 {
        remainder := n % 2
        binary = fmt.Sprintf("%d%s", remainder, binary)
        n = n / 2
    }
    return binary
}

func main() {
    numbers := []int{10, 25, 37, 100}
    for _, n := range numbers {
        fmt.Printf("%d in binary: %s\n", n, decimalToBinary(n))
    }
}
```

### Solved Problem 8: Perfect Square Checker

**Task:** Check if a number is a perfect square and find its square root.

**Solution:**
```go
package main

import (
    "fmt"
    "math"
)

func isPerfectSquare(n int) (bool, int) {
    if n < 0 {
        return false, 0
    }
    
    sqrt := int(math.Sqrt(float64(n)))
    if sqrt*sqrt == n {
        return true, sqrt
    }
    return false, 0
}

func main() {
    numbers := []int{16, 25, 30, 49, 100}
    for _, n := range numbers {
        isSquare, root := isPerfectSquare(n)
        if isSquare {
            fmt.Printf("%d is a perfect square (√%d = %d)\n", n, n, root)
        } else {
            fmt.Printf("%d is not a perfect square\n", n)
        }
    }
}
```

### Solved Problem 9: Sum of Digits

**Task:** Calculate sum of digits of a number until single digit.

**Solution:**
```go
package main

import "fmt"

func sumOfDigits(n int) int {
    sum := 0
    for n > 0 {
        sum += n % 10
        n = n / 10
    }
    return sum
}

func digitalRoot(n int) int {
    for n >= 10 {
        n = sumOfDigits(n)
    }
    return n
}

func main() {
    numbers := []int{123, 456, 789, 999}
    for _, n := range numbers {
        root := digitalRoot(n)
        fmt.Printf("Digital root of %d: %d\n", n, root)
    }
}
```

### Solved Problem 10: Calendar Month Display

**Task:** Display a calendar month (simplified - just days in grid).

**Solution:**
```go
package main

import "fmt"

func printCalendar(month, year int, daysInMonth, startDay int) {
    monthNames := []string{"", "January", "February", "March", "April", "May", "June",
        "July", "August", "September", "October", "November", "December"}
    
    fmt.Printf("\n   %s %d\n", monthNames[month], year)
    fmt.Println("Su Mo Tu We Th Fr Sa")
    
    // Print leading spaces
    for i := 0; i < startDay; i++ {
        fmt.Print("   ")
    }
    
    // Print days
    for day := 1; day <= daysInMonth; day++ {
        fmt.Printf("%2d ", day)
        if (day+startDay)%7 == 0 {
            fmt.Println()
        }
    }
    fmt.Println()
}

func main() {
    // Example: January 2024 (31 days, starts on Monday = 1)
    printCalendar(1, 2024, 31, 1)
}
```

---

### 10 More Practice Problems (Solve These!)

**Problem 1: Pyramid of Numbers**
- Print:
```
    1
   123
  12345
 1234567
123456789
```

**Problem 2: Right Triangle Pattern**
- Print:
```
*
**
***
****
*****
```

**Problem 3: Hollow Square**
- Print a hollow square of asterisks (only border)

**Problem 4: Number Spiral**
- Print numbers in spiral pattern:
```
1  2  3  4
12 13 14 5
11 16 15 6
10 9  8  7
```

**Problem 5: Prime Factorization Display**
- Display prime factors in format: 60 = 2² × 3 × 5

**Problem 6: Number Pyramid (Reverse)**
- Print:
```
12345
 1234
  123
   12
    1
```

**Problem 7: Checkerboard Pattern**
- Print checkerboard pattern using X and O

**Problem 8: Fibonacci Triangle**
- Print first n rows of Fibonacci triangle

**Problem 9: Number Spiral (Alternate)**
- Create spiral pattern with numbers 1 to n²

**Problem 10: Pattern Matching**
- Given pattern string like "1-2-3-4", generate sequence
- Handle patterns: ascending, descending, alternating

---

## Next Steps

Now you understand control flow. Next topics to explore:
- Pointers (understanding memory)
- Structs (custom types)
- Methods (functions on types)
- Interfaces (polymorphism)

**Continue learning! → Check [README.md](./README.md) for next topics**
