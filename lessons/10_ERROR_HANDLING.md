# Error Handling

## a) Overview

### What this topic is
Error handling in Go - how Go treats errors as values that must be explicitly checked, rather than using exceptions like Python.

### Why it exists in Go
Go's philosophy: "Errors are values." This makes error handling explicit and visible, preventing hidden failures. You must check errors - the compiler doesn't let you ignore them easily.

### 🎯 Layman's Explanation (Simple Terms)

**Think of error handling like checking if something worked before using it:**

**Real-world analogy - Making a Phone Call:**
- You call someone: `result, err := callFriend()`
- Two things can happen:
  - **Success**: Friend answers, you get a conversation (result)
  - **Failure**: Phone is busy, you get an error message (err)
- **You MUST check**: Did the call succeed or fail?
- If it failed, you can't have a conversation - you need to handle the error!

**Another analogy - Opening a Door:**
- You try to open a door: `door, err := openDoor(key)`
- Two outcomes:
  - **Success**: Door opens, you get access (door)
  - **Failure**: Wrong key, door stays locked (err)
- **You MUST check**: Did the door open?
- If it failed, you can't walk through - you need to handle it (get the right key, call a locksmith, etc.)

**Simple example:**
```
Python (exceptions - hidden):
try:
    file = open("file.txt")
    # If error happens, jumps to except block
    # Error is "thrown" - can be hidden
except:
    print("Error!")

Go (explicit - visible):
file, err := openFile("file.txt")
if err != nil {  // MUST check!
    // Error is right here - can't ignore it
    return err
}
// Only use file if no error
```

**Error = A return value, not an exception:**
- Like a **receipt** that says "Success" or "Error: Item not found"
- Like a **status report** - "Mission accomplished" or "Mission failed: reason"
- Like a **yes/no answer** with an explanation if "no"

**Why Go's way?**
1. **Explicit**: You see errors immediately - can't hide them
2. **Visible**: Errors are part of the function signature - you know a function can fail
3. **Safe**: Forces you to handle errors - program won't crash unexpectedly
4. **Clear**: Error handling is right there in the code - easy to see what happens on failure

**Key concept:**
- Functions return `(result, error)` - two values
- `nil` error = success (like "no error" = "everything worked")
- Non-nil error = failure (like "error message" = "something went wrong")
- **You MUST check** the error before using the result!

---

## b) Syntax

### Basic Error Handling
```go
result, err := doSomething()
if err != nil {
    // handle error
    return err
}
// use result
```

### Creating Errors
```go
import "errors"
err := errors.New("something went wrong")

import "fmt"
err := fmt.Errorf("error: %v", value)
```

### Custom Error Type
```go
type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
```

---

## c) Explanation

### Step-by-Step Error Handling

**1. Functions return errors**
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```
- Error is last return value
- `nil` means no error
- Non-nil means error occurred

**2. Check errors explicitly**
```go
result, err := divide(10, 2)
if err != nil {
    // handle error
    return err
}
// use result
```
- Must check `err != nil`
- Handle error or return it
- Can't ignore errors easily

**3. Error propagation**
```go
func process() error {
    result, err := divide(10, 2)
    if err != nil {
        return err  // Pass error up
    }
    // use result
    return nil
}
```
- Return error to caller
- Let caller decide how to handle

### Characteristics

- **Explicit**: Errors are visible, not hidden
- **No exceptions**: No try/catch, no panic (usually)
- **Error is value**: Can pass around, store, compare
- **Multiple returns**: `(result, error)` pattern
- **Must check**: Can't easily ignore errors

---

## d) Python Comparison

### Python Exception Handling
```python
def divide(a, b):
    if b == 0:
        raise ValueError("division by zero")
    return a / b

try:
    result = divide(10, 2)
    print(result)
except ValueError as e:
    print(f"Error: {e}")
```

### Go Error Handling
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 2)
if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}
fmt.Println(result)
```

**Key Differences:**

| Feature | Python | Go |
|---------|--------|-----|
| **Error mechanism** | Exceptions (raise/try/except) | Return error value |
| **Visibility** | Can be hidden (uncaught) | Must be explicit |
| **Control flow** | Jumps to except block | Normal flow (if/else) |
| **Error type** | Exception object | error interface |
| **Default behavior** | Crashes if uncaught | Must check or ignore explicitly |

**Thinking Difference:**
- Python: "I'll raise an exception and handle it somewhere"
- Go: "I'll return an error and you must check it"
- Python: Errors can be hidden
- Go: Errors are explicit and visible

---

## e) Visual Flow / Mental Model

### Error Flow in Go

```
1. Function call: result, err := doSomething()
      ↓
2. Function executes
      ↓
3. If error occurs:
      return nil, errors.New("error message")
      ↓
4. err is not nil
      ↓
5. Check: if err != nil
      ↓
6. Handle error (log, return, etc.)
      ↓
7. If no error:
      return value, nil
      ↓
8. err is nil
      ↓
9. Use result
```

### Error Propagation

```
main()
  ↓ calls
function1()
  ↓ calls
function2()
  ↓ returns error
function1() receives error
  ↓ returns error
main() receives error
  ↓ handles error
```

**Each level can:**
- Handle the error
- Add context and return
- Return as-is

---

## f) Demo Example

### Complete Example

```go
package main

import (
    "errors"
    "fmt"
    "strconv"
)

// 1. Basic error return
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// 2. Error with context
func parseAge(ageStr string) (int, error) {
    age, err := strconv.Atoi(ageStr)
    if err != nil {
        return 0, fmt.Errorf("failed to parse age '%s': %w", ageStr, err)
    }
    if age < 0 || age > 150 {
        return 0, fmt.Errorf("invalid age: %d (must be 0-150)", age)
    }
    return age, nil
}

// 3. Custom error type
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error in %s: %s", e.Field, e.Message)
}

func validateUser(name string, age int) error {
    if name == "" {
        return &ValidationError{
            Field:   "name",
            Message: "name cannot be empty",
        }
    }
    if age < 0 {
        return &ValidationError{
            Field:   "age",
            Message: "age cannot be negative",
        }
    }
    return nil
}

// 4. Error wrapping (Go 1.13+)
func processUser(name string, ageStr string) error {
    age, err := parseAge(ageStr)
    if err != nil {
        return fmt.Errorf("processing user: %w", err)  // Wrap error
    }
    
    err = validateUser(name, age)
    if err != nil {
        return fmt.Errorf("processing user: %w", err)
    }
    
    fmt.Printf("User %s, age %d is valid\n", name, age)
    return nil
}

// 5. Checking error types
func handleError(err error) {
    var validationErr *ValidationError
    if errors.As(err, &validationErr) {
        fmt.Printf("Validation error in field '%s': %s\n",
            validationErr.Field, validationErr.Message)
        return
    }
    
    fmt.Printf("Other error: %v\n", err)
}

// 6. Ignoring errors (rare, use carefully)
func riskyOperation() (int, error) {
    return 42, errors.New("something went wrong")
}

func main() {
    // Basic error handling
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Result: %.2f\n", result)
    }
    
    // Error case
    _, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    }
    
    // Error with context
    age, err := parseAge("25")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Age: %d\n", age)
    }
    
    // Invalid age
    _, err = parseAge("200")
    if err != nil {
        fmt.Println("Error:", err)
    }
    
    // Custom error
    err = validateUser("", 30)
    if err != nil {
        fmt.Println("Error:", err)
        handleError(err)
    }
    
    // Error wrapping
    err = processUser("Alice", "30")
    if err != nil {
        fmt.Printf("Processing error: %v\n", err)
    }
    
    // Ignoring error (not recommended, but possible)
    value, _ := riskyOperation()
    fmt.Printf("Value (ignored error): %d\n", value)
}
```

**Line-by-line explanation:**

1. **Basic error**: Return `(value, error)`
2. **Error with context**: Use `fmt.Errorf` to add context
3. **Custom error**: Create type that implements `error` interface
4. **Error wrapping**: Wrap errors to add context (`%w` verb)
5. **Error checking**: Use `errors.As` to check error type
6. **Ignoring errors**: Use `_` (not recommended)

**Output:**
```
Result: 5.00
Error: division by zero
Age: 25
Error: invalid age: 200 (must be 0-150)
Error: validation error in name: name cannot be empty
Validation error in field 'name': name cannot be empty
User Alice, age 30 is valid
Value (ignored error): 42
```

---

## g) Use Cases

### When to return errors

**1. Operations that can fail**
```go
func readFile(filename string) ([]byte, error) {
    // Can fail - return error
}
```

**2. Validation**
```go
func validateEmail(email string) error {
    if !strings.Contains(email, "@") {
        return errors.New("invalid email")
    }
    return nil
}
```

**3. External dependencies**
```go
func connectDB(host string) (*DB, error) {
    // Network can fail - return error
}
```

### Error handling patterns

**1. Return early**
```go
if err != nil {
    return err
}
```

**2. Add context**
```go
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}
```

**3. Log and continue**
```go
if err != nil {
    log.Printf("Warning: %v", err)
    // continue with default value
}
```

---

## h) Do's and Don'ts / Best Practices

### ✅ Do's

1. **Always check errors**
   ```go
   result, err := doSomething()
   if err != nil {
       return err
   }
   ```

2. **Add context to errors**
   ```go
   if err != nil {
       return fmt.Errorf("failed to process: %w", err)
   }
   ```

3. **Return errors, don't log them**
   ```go
   // ✅ Let caller decide
   return err
   
   // ❌ Don't log in library code
   log.Println(err)
   return err
   ```

4. **Use custom errors for specific cases**
   ```go
   type NotFoundError struct {
       Resource string
   }
   ```

5. **Wrap errors with context**
   ```go
   return fmt.Errorf("user service: %w", err)
   ```

### ❌ Don'ts

1. **Don't ignore errors**
   ```go
   result, _ := doSomething()  // ❌ Bad!
   ```

2. **Don't panic for normal errors**
   ```go
   if err != nil {
       panic(err)  // ❌ Only for unrecoverable
   }
   ```

3. **Don't return error and value when error occurs**
   ```go
   if err != nil {
       return result, err  // ❌ result might be invalid
       return zeroValue, err  // ✅ Correct
   }
   ```

4. **Don't create errors in tight loops**
   ```go
   // ❌ Creates many error objects
   for i := 0; i < 1000000; i++ {
       if condition {
           return errors.New("error")
       }
   }
   ```

---

## i) Solved Practice Examples

### Example 1: File Reader with Error Handling

**Task:** Create function that reads file with proper error handling.

**Solution:**
```go
package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func readFile(filename string) (string, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return "", fmt.Errorf("failed to read file '%s': %w", filename, err)
    }
    return string(data), nil
}

func main() {
    content, err := readFile("example.txt")
    if err != nil {
        if os.IsNotExist(err) {
            fmt.Println("File does not exist")
        } else {
            fmt.Printf("Error: %v\n", err)
        }
        return
    }
    fmt.Println("File content:", content)
}
```

### Example 2: Calculator with Error Handling

**Task:** Create calculator functions with comprehensive error handling.

**Solution:**
```go
package main

import (
    "errors"
    "fmt"
)

type Calculator struct{}

func (c Calculator) Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func (c Calculator) Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, errors.New("cannot take square root of negative number")
    }
    // Simple approximation
    return x * 0.5, nil  // Simplified
}

func main() {
    calc := Calculator{}
    
    result, err := calc.Divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("10 / 2 = %.2f\n", result)
    }
    
    result, err = calc.Divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    }
    
    sqrt, err := calc.Sqrt(16)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("sqrt(16) = %.2f\n", sqrt)
    }
}
```

---

## j) Quiz / Practice for Me

### Quiz Questions

1. **What is the standard error return pattern in Go?**
   - [ ] `(error, result)`
   - [ ] `(result, error)`
   - [ ] `error` only
   - [ ] `(result)` with panic

2. **What does `nil` mean for an error?**
   - [ ] Error occurred
   - [ ] No error
   - [ ] Unknown
   - [ ] Fatal error

3. **Should you panic for normal errors?**
   - [ ] Yes, always
   - [ ] No, only for unrecoverable
   - [ ] Sometimes
   - [ ] Never

### Practice Tasks

**Task 1: Age Validator**
- Create function `validateAge(age int) error`
- Return error if age < 0 or age > 150
- Test with valid and invalid ages

**Task 2: Email Validator**
- Create function `validateEmail(email string) error`
- Check if email contains "@"
- Check if email is not empty
- Return appropriate errors

### Answers

**Quiz Answers:**
1. `(result, error)` - error is last
2. No error
3. No, only for unrecoverable

**Practice Solutions:**

**Task 1 Solution:**
```go
package main

import (
    "errors"
    "fmt"
)

func validateAge(age int) error {
    if age < 0 {
        return errors.New("age cannot be negative")
    }
    if age > 150 {
        return errors.New("age cannot be greater than 150")
    }
    return nil
}

func main() {
    ages := []int{-5, 0, 25, 100, 200}
    for _, age := range ages {
        err := validateAge(age)
        if err != nil {
            fmt.Printf("Age %d: Error - %v\n", age, err)
        } else {
            fmt.Printf("Age %d: Valid\n", age)
        }
    }
}
```

**Task 2 Solution:**
```go
package main

import (
    "errors"
    "fmt"
    "strings"
)

func validateEmail(email string) error {
    if email == "" {
        return errors.New("email cannot be empty")
    }
    if !strings.Contains(email, "@") {
        return errors.New("email must contain @")
    }
    return nil
}

func main() {
    emails := []string{"", "invalid", "user@example.com", "test@test"}
    for _, email := range emails {
        err := validateEmail(email)
        if err != nil {
            fmt.Printf("'%s': Error - %v\n", email, err)
        } else {
            fmt.Printf("'%s': Valid\n", email)
        }
    }
}
```

---

## Key Takeaways

1. **Errors are values** - Return `(result, error)`
2. **Must check errors** - `if err != nil`
3. **No exceptions** - Errors are explicit
4. **Add context** - Wrap errors with `fmt.Errorf`
5. **Custom errors** - Create types implementing `error` interface
6. **Error is last** - Standard pattern: `(result, error)`

---

## Must Remember Forever

- `(result, error)` - Standard return pattern
- `nil` error = no error
- Always check: `if err != nil { }`
- Add context: `fmt.Errorf("context: %w", err)`
- Don't panic for normal errors
- Return errors, let caller decide how to handle

---

---

## k) Additional Practice Problems (Build Strong Logic)

### 10 More Solved Examples

### Solved Problem 1: File Operations with Error Handling

**Task:** Implement file operations with comprehensive error handling.

**Solution:**
```go
package main

import (
    "fmt"
    "os"
)

func readFileWithErrorHandling(filename string) (string, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        if os.IsNotExist(err) {
            return "", fmt.Errorf("file '%s' does not exist", filename)
        }
        if os.IsPermission(err) {
            return "", fmt.Errorf("permission denied reading '%s'", filename)
        }
        return "", fmt.Errorf("error reading file '%s': %w", filename, err)
    }
    return string(data), nil
}

func main() {
    content, err := readFileWithErrorHandling("nonexistent.txt")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Println(content)
    }
}
```

### Solved Problem 2: Validation with Multiple Errors

**Task:** Collect multiple validation errors and return them all.

**Solution:**
```go
package main

import (
    "errors"
    "fmt"
    "strings"
)

type ValidationErrors struct {
    Errors []error
}

func (ve *ValidationErrors) Add(err error) {
    ve.Errors = append(ve.Errors, err)
}

func (ve *ValidationErrors) Error() string {
    var messages []string
    for _, err := range ve.Errors {
        messages = append(messages, err.Error())
    }
    return strings.Join(messages, "; ")
}

func (ve *ValidationErrors) HasErrors() bool {
    return len(ve.Errors) > 0
}

func validateUser(name, email string, age int) error {
    var ve ValidationErrors
    
    if name == "" {
        ve.Add(errors.New("name is required"))
    }
    if len(name) < 3 {
        ve.Add(errors.New("name must be at least 3 characters"))
    }
    if !strings.Contains(email, "@") {
        ve.Add(errors.New("email must contain @"))
    }
    if age < 0 || age > 150 {
        ve.Add(errors.New("age must be between 0 and 150"))
    }
    
    if ve.HasErrors() {
        return &ve
    }
    return nil
}

func main() {
    err := validateUser("", "invalid-email", 200)
    if err != nil {
        fmt.Printf("Validation errors: %v\n", err)
    }
}
```

### Solved Problem 3: Retry with Error Handling

**Task:** Implement retry logic with exponential backoff.

**Solution:**
```go
package main

import (
    "errors"
    "fmt"
    "time"
)

var ErrTemporary = errors.New("temporary error")

func operation() error {
    // Simulate operation that might fail
    return ErrTemporary
}

func retryWithBackoff(fn func() error, maxRetries int) error {
    var err error
    for i := 0; i < maxRetries; i++ {
        err = fn()
        if err == nil {
            return nil
        }
        if !errors.Is(err, ErrTemporary) {
            return err // Don't retry non-temporary errors
        }
        backoff := time.Duration(1<<uint(i)) * time.Second
        fmt.Printf("Retry %d after %v\n", i+1, backoff)
        time.Sleep(backoff)
    }
    return fmt.Errorf("failed after %d retries: %w", maxRetries, err)
}

func main() {
    err := retryWithBackoff(operation, 3)
    if err != nil {
        fmt.Printf("Final error: %v\n", err)
    }
}
```

### Solved Problem 4: Error Wrapping Chain

**Task:** Wrap errors through multiple layers with context.

**Solution:**
```go
package main

import (
    "errors"
    "fmt"
)

var ErrDatabase = errors.New("database error")
var ErrNotFound = errors.New("not found")

func databaseQuery(id int) error {
    if id < 0 {
        return fmt.Errorf("query failed: %w", ErrNotFound)
    }
    return fmt.Errorf("query failed: %w", ErrDatabase)
}

func getUser(id int) error {
    err := databaseQuery(id)
    if err != nil {
        return fmt.Errorf("getUser(%d): %w", id, err)
    }
    return nil
}

func processUser(id int) error {
    err := getUser(id)
    if err != nil {
        return fmt.Errorf("processUser: %w", err)
    }
    return nil
}

func main() {
    err := processUser(-1)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        if errors.Is(err, ErrNotFound) {
            fmt.Println("Handling not found error")
        }
    }
}
```

### Solved Problem 5: Error Recovery

**Task:** Implement panic recovery with error conversion.

**Solution:**
```go
package main

import (
    "errors"
    "fmt"
)

func riskyOperation() {
    panic("something went wrong")
}

func safeOperation() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered from panic: %v", r)
        }
    }()
    
    riskyOperation()
    return nil
}

func main() {
    err := safeOperation()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}
```

### Solved Problem 6: Custom Error Types

**Task:** Create custom error types for different error categories.

**Solution:**
```go
package main

import "fmt"

type ErrorCode int

const (
    ErrCodeNotFound ErrorCode = iota
    ErrCodeInvalidInput
    ErrCodePermissionDenied
    ErrCodeInternal
)

type AppError struct {
    Code    ErrorCode
    Message string
    Details map[string]interface{}
}

func (e *AppError) Error() string {
    return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func NewNotFoundError(resource string) *AppError {
    return &AppError{
        Code:    ErrCodeNotFound,
        Message: fmt.Sprintf("%s not found", resource),
        Details: map[string]interface{}{"resource": resource},
    }
}

func NewInvalidInputError(field string, reason string) *AppError {
    return &AppError{
        Code:    ErrCodeInvalidInput,
        Message: fmt.Sprintf("invalid %s: %s", field, reason),
        Details: map[string]interface{}{"field": field, "reason": reason},
    }
}

func handleError(err error) {
    if appErr, ok := err.(*AppError); ok {
        switch appErr.Code {
        case ErrCodeNotFound:
            fmt.Printf("Not found: %s\n", appErr.Message)
        case ErrCodeInvalidInput:
            fmt.Printf("Invalid input: %s\n", appErr.Message)
        default:
            fmt.Printf("Error: %s\n", appErr.Message)
        }
    }
}

func main() {
    err1 := NewNotFoundError("user")
    err2 := NewInvalidInputError("email", "missing @")
    
    handleError(err1)
    handleError(err2)
}
```

### Solved Problem 7: Error Aggregation

**Task:** Collect and handle multiple errors from batch operations.

**Solution:**
```go
package main

import (
    "errors"
    "fmt"
)

type ErrorList struct {
    errors []error
}

func (el *ErrorList) Add(err error) {
    if err != nil {
        el.errors = append(el.errors, err)
    }
}

func (el *ErrorList) Error() string {
    if len(el.errors) == 0 {
        return ""
    }
    msg := fmt.Sprintf("%d errors: ", len(el.errors))
    for i, err := range el.errors {
        if i > 0 {
            msg += "; "
        }
        msg += err.Error()
    }
    return msg
}

func (el *ErrorList) HasErrors() bool {
    return len(el.errors) > 0
}

func processItems(items []string) error {
    var errorList ErrorList
    
    for i, item := range items {
        if item == "" {
            errorList.Add(fmt.Errorf("item %d is empty", i))
        }
        if len(item) < 3 {
            errorList.Add(fmt.Errorf("item %d too short: %s", i, item))
        }
    }
    
    if errorList.HasErrors() {
        return &errorList
    }
    return nil
}

func main() {
    items := []string{"abc", "", "xy", "def"}
    err := processItems(items)
    if err != nil {
        fmt.Printf("Errors: %v\n", err)
    }
}
```

### Solved Problem 8: Error with Stack Trace

**Task:** Create error type that includes stack trace information.

**Solution:**
```go
package main

import (
    "fmt"
    "runtime"
)

type StackError struct {
    Err     error
    File    string
    Line    int
    Func    string
}

func (se *StackError) Error() string {
    return fmt.Sprintf("%s at %s:%d in %s", se.Err, se.File, se.Line, se.Func)
}

func newStackError(err error) *StackError {
    pc, file, line, _ := runtime.Caller(1)
    fn := runtime.FuncForPC(pc)
    return &StackError{
        Err:  err,
        File: file,
        Line: line,
        Func: fn.Name(),
    }
}

func riskyFunction() error {
    return newStackError(fmt.Errorf("operation failed"))
}

func main() {
    err := riskyFunction()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}
```

### Solved Problem 9: Error Context Builder

**Task:** Build error with context information.

**Solution:**
```go
package main

import (
    "errors"
    "fmt"
)

type ErrorBuilder struct {
    base    error
    context []string
}

func NewErrorBuilder(base error) *ErrorBuilder {
    return &ErrorBuilder{base: base, context: []string{}}
}

func (eb *ErrorBuilder) WithContext(key, value string) *ErrorBuilder {
    eb.context = append(eb.context, fmt.Sprintf("%s=%s", key, value))
    return eb
}

func (eb *ErrorBuilder) Build() error {
    if len(eb.context) == 0 {
        return eb.base
    }
    contextStr := ""
    for i, ctx := range eb.context {
        if i > 0 {
            contextStr += ", "
        }
        contextStr += ctx
    }
    return fmt.Errorf("%w [context: %s]", eb.base, contextStr)
}

func main() {
    baseErr := errors.New("operation failed")
    err := NewErrorBuilder(baseErr).
        WithContext("user", "alice").
        WithContext("action", "delete").
        WithContext("resource", "file.txt").
        Build()
    
    fmt.Printf("Error: %v\n", err)
}
```

### Solved Problem 10: Error Handler Chain

**Task:** Implement error handling chain with different handlers.

**Solution:**
```go
package main

import (
    "errors"
    "fmt"
    "log"
)

type ErrorHandler func(error) error

func logError(err error) error {
    log.Printf("Error logged: %v", err)
    return err
}

func retryError(err error) error {
    if errors.Is(err, errors.New("temporary")) {
        fmt.Println("Retrying...")
        return nil
    }
    return err
}

func notifyError(err error) error {
    fmt.Println("Sending notification about error")
    return err
}

func handleError(err error, handlers ...ErrorHandler) error {
    for _, handler := range handlers {
        err = handler(err)
        if err == nil {
            return nil
        }
    }
    return err
}

func main() {
    err := errors.New("temporary error")
    finalErr := handleError(err, logError, retryError, notifyError)
    if finalErr != nil {
        fmt.Printf("Final error: %v\n", finalErr)
    }
}
```

---

### 10 More Practice Problems (Solve These!)

**Problem 1: Database Transaction Error Handling**
- Implement transaction with rollback on error
- Handle connection errors, query errors, commit errors

**Problem 2: API Client with Error Handling**
- Create HTTP client with retry logic
- Handle network errors, timeout errors, status code errors

**Problem 3: Configuration Loader**
- Load config with validation
- Return all validation errors at once
- Handle missing file, invalid format, invalid values

**Problem 4: Batch Processor**
- Process items in batch
- Continue on individual errors
- Collect all errors and report at end

**Problem 5: Error Recovery System**
- Implement circuit breaker pattern
- Handle errors and prevent cascading failures

**Problem 6: Validation Framework**
- Create validation framework with custom validators
- Return detailed validation errors

**Problem 7: Error Metrics Collector**
- Track error rates by type
- Implement error reporting system

**Problem 8: Graceful Degradation**
- Handle errors gracefully
- Provide fallback functionality

**Problem 9: Error Logging System**
- Log errors with different levels
- Include context and stack traces

**Problem 10: Error Transformation**
- Transform errors between layers
- Map internal errors to user-friendly messages

---

## Next Steps

Now you understand error handling. Next:
- Arrays and Slices
- Working with collections
- Dynamic arrays in Go

**Ready? → [11_ARRAYS_AND_SLICES.md](./11_ARRAYS_AND_SLICES.md)**
