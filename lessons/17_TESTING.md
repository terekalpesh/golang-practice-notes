# Testing

## a) Overview

### What this topic is
Writing and running tests in Go. Go has excellent built-in testing support.

### Why it exists in Go
Testing is essential for reliable code. Go makes testing simple and integrated into the language.

---

## b) Syntax

### Basic Test
```go
// file_test.go
package mypackage

import "testing"

func TestFunction(t *testing.T) {
    result := Function(5)
    if result != 10 {
        t.Errorf("Expected 10, got %d", result)
    }
}
```

### Table-Driven Tests
```go
func TestFunction(t *testing.T) {
    tests := []struct {
        input    int
        expected int
    }{
        {5, 10},
        {10, 20},
        {0, 0},
    }
    
    for _, tt := range tests {
        result := Function(tt.input)
        if result != tt.expected {
            t.Errorf("Function(%d) = %d, want %d", tt.input, result, tt.expected)
        }
    }
}
```

### Running Tests
```bash
go test
go test -v
go test ./...
go test -cover
```

---

## c) Explanation

### Step-by-Step Testing

**1. Create test file**
```go
// myfile_test.go
package mypackage
```
- Test files end with `_test.go`
- Same package (or `_test` package)

**2. Write test function**
```go
func TestFunction(t *testing.T) {
    // test code
}
```
- Must start with `Test`
- Takes `*testing.T` parameter

**3. Run tests**
```bash
go test
```
- Finds all `*_test.go` files
- Runs all `Test*` functions

### Characteristics

#### Test Structure Characteristics
- **Test files**: Files ending with `_test.go` are test files
- **Test functions**: Functions starting with `Test` are test functions
- **Test package**: Can be same package or `_test` package
- **Test naming**: `TestFunctionName` convention
- **Table-driven tests**: Idiomatic Go pattern for multiple test cases

#### Test Execution Characteristics
- **Automatic discovery**: Go automatically finds and runs tests
- **Parallel execution**: Tests can run in parallel (with `t.Parallel()`)
- **Test isolation**: Each test runs independently
- **Test ordering**: Tests run in unpredictable order (don't depend on order)
- **Fast execution**: Tests compile and run quickly

#### Test Data Characteristics
- **Test fixtures**: Can set up test data in test functions
- **Test helpers**: Helper functions for common test setup
- **Test tables**: Slice of test cases for table-driven tests
- **Test cleanup**: Use `defer` for cleanup in tests
- **Test data isolation**: Each test should be independent

#### Testing Framework Characteristics
- **Built-in**: Testing built into Go (no external framework needed)
- **Simple API**: Minimal testing API (`testing.T`)
- **No assertions**: Manual if statements (no assert library)
- **Benchmarking**: Built-in benchmark support
- **Coverage**: Built-in code coverage tools
- **Examples**: Can write example code that doubles as tests

---

## d) Python Comparison

### Python Testing
```python
import unittest

class TestFunction(unittest.TestCase):
    def test_add(self):
        self.assertEqual(add(2, 3), 5)

if __name__ == '__main__':
    unittest.main()
```

### Go Testing
```go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Expected 5, got %d", result)
    }
}
```

**Key Differences:**
- Python: Class-based (unittest)
- Go: Function-based (simpler)
- Python: Assert methods
- Go: Manual if statements

---

## e) Visual Flow / Mental Model

### Test Execution Flow

```
go test
  ↓
1. Find all *_test.go files
  ↓
2. Compile test files
  ↓
3. Run all Test* functions
  ↓
4. Report results
```

---

## f) Demo Example

### Complete Example

```go
// math.go
package main

func Add(a, b int) int {
    return a + b
}

func Multiply(a, b int) int {
    return a * b
}

// math_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d, want 5", result)
    }
}

func TestMultiply(t *testing.T) {
    tests := []struct {
        a, b, expected int
    }{
        {2, 3, 6},
        {5, 4, 20},
        {0, 10, 0},
    }
    
    for _, tt := range tests {
        result := Multiply(tt.a, tt.b)
        if result != tt.expected {
            t.Errorf("Multiply(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
        }
    }
}
```

---

## g) Use Cases

- Unit testing
- Integration testing
- Benchmarking
- Example code

---

## h) Do's and Don'ts

### ✅ Do's

1. **Use table-driven tests**
2. **Test edge cases**
3. **Use descriptive test names**

### ❌ Don'ts

1. **Don't skip error cases**
2. **Don't write tests that always pass**

---

## i) Solved Practice Examples

### Example 1: Basic Function Testing

**Task:** Write tests for a calculator function.

**Solution:**
```go
// calculator.go
package main

func Add(a, b int) int {
    return a + b
}

func Multiply(a, b int) int {
    return a * b
}

// calculator_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d, want 5", result)
    }
}

func TestMultiply(t *testing.T) {
    result := Multiply(4, 5)
    if result != 20 {
        t.Errorf("Multiply(4, 5) = %d, want 20", result)
    }
}
```

### Example 2: Table-Driven Tests

**Task:** Write comprehensive table-driven tests for a function.

**Solution:**
```go
// math.go
package main

func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// math_test.go
package main

import (
    "errors"
    "testing"
)

func TestDivide(t *testing.T) {
    tests := []struct {
        name      string
        a, b      float64
        want      float64
        wantError bool
        errorMsg  string
    }{
        {
            name:      "normal division",
            a:          10,
            b:          2,
            want:       5,
            wantError:  false,
        },
        {
            name:      "division by zero",
            a:          10,
            b:          0,
            want:       0,
            wantError:  true,
            errorMsg:   "division by zero",
        },
        {
            name:      "negative numbers",
            a:          -10,
            b:          2,
            want:       -5,
            wantError:  false,
        },
        {
            name:      "decimal result",
            a:          7,
            b:          2,
            want:       3.5,
            wantError:  false,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Divide(tt.a, tt.b)
            
            if tt.wantError {
                if err == nil {
                    t.Errorf("Divide() expected error, got nil")
                } else if err.Error() != tt.errorMsg {
                    t.Errorf("Divide() error = %v, want %s", err, tt.errorMsg)
                }
            } else {
                if err != nil {
                    t.Errorf("Divide() unexpected error: %v", err)
                }
                if got != tt.want {
                    t.Errorf("Divide() = %v, want %v", got, tt.want)
                }
            }
        })
    }
}
```

### Example 3: Testing with Setup and Teardown

**Task:** Write tests that need setup and cleanup.

**Solution:**
```go
// cache.go
package main

type Cache struct {
    data map[string]string
}

func NewCache() *Cache {
    return &Cache{
        data: make(map[string]string),
    }
}

func (c *Cache) Set(key, value string) {
    c.data[key] = value
}

func (c *Cache) Get(key string) (string, bool) {
    value, exists := c.data[key]
    return value, exists
}

func (c *Cache) Clear() {
    c.data = make(map[string]string)
}

// cache_test.go
package main

import "testing"

func setupCache() *Cache {
    cache := NewCache()
    cache.Set("key1", "value1")
    cache.Set("key2", "value2")
    return cache
}

func TestCacheGet(t *testing.T) {
    cache := setupCache()
    defer cache.Clear() // Cleanup
    
    tests := []struct {
        key      string
        want     string
        wantOk   bool
    }{
        {"key1", "value1", true},
        {"key2", "value2", true},
        {"key3", "", false},
    }
    
    for _, tt := range tests {
        got, ok := cache.Get(tt.key)
        if ok != tt.wantOk {
            t.Errorf("Get(%s) ok = %v, want %v", tt.key, ok, tt.wantOk)
        }
        if got != tt.want {
            t.Errorf("Get(%s) = %s, want %s", tt.key, got, tt.want)
        }
    }
}
```

### Example 4: Testing Error Cases

**Task:** Write tests that verify error handling.

**Solution:**
```go
// validator.go
package main

import "errors"

var ErrEmptyString = errors.New("string cannot be empty")
var ErrTooLong = errors.New("string too long")

func ValidateString(s string, maxLen int) error {
    if s == "" {
        return ErrEmptyString
    }
    if len(s) > maxLen {
        return ErrTooLong
    }
    return nil
}

// validator_test.go
package main

import "testing"

func TestValidateString(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        maxLen  int
        wantErr error
    }{
        {
            name:    "valid string",
            input:   "hello",
            maxLen:  10,
            wantErr: nil,
        },
        {
            name:    "empty string",
            input:   "",
            maxLen:  10,
            wantErr: ErrEmptyString,
        },
        {
            name:    "too long",
            input:   "this is too long",
            maxLen:  5,
            wantErr: ErrTooLong,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := ValidateString(tt.input, tt.maxLen)
            if err != tt.wantErr {
                t.Errorf("ValidateString() error = %v, want %v", err, tt.wantErr)
            }
        })
    }
}
```

### Example 5: Benchmark Testing

**Task:** Write benchmark tests to measure performance.

**Solution:**
```go
// string_utils.go
package main

func ReverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

// string_utils_test.go
package main

import "testing"

func BenchmarkReverseString(b *testing.B) {
    input := "Hello, World! This is a test string."
    for i := 0; i < b.N; i++ {
        ReverseString(input)
    }
}

func BenchmarkReverseStringLong(b *testing.B) {
    input := "A" + string(make([]byte, 1000))
    for i := 0; i < b.N; i++ {
        ReverseString(input)
    }
}
```

### Example 6: Testing HTTP Handlers

**Task:** Write tests for HTTP handler functions.

**Solution:**
```go
// handler.go
package main

import (
    "encoding/json"
    "net/http"
)

type Response struct {
    Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    response := Response{Message: "Hello, World!"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// handler_test.go
package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHelloHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/hello", nil)
    if err != nil {
        t.Fatal(err)
    }
    
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(HelloHandler)
    handler.ServeHTTP(rr, req)
    
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status: got %v want %v", status, http.StatusOK)
    }
    
    var response Response
    json.Unmarshal(rr.Body.Bytes(), &response)
    
    expected := "Hello, World!"
    if response.Message != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", response.Message, expected)
    }
}
```

### Example 7: Test Helpers

**Task:** Create reusable test helper functions.

**Solution:**
```go
// user.go
package main

type User struct {
    ID   int
    Name string
}

func NewUser(id int, name string) *User {
    return &User{ID: id, Name: name}
}

// user_test.go
package main

import (
    "reflect"
    "testing"
)

// Test helper function
func assertUser(t *testing.T, got, want *User) {
    t.Helper() // Marks this as a helper function
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %+v, want %+v", got, want)
    }
}

func TestNewUser(t *testing.T) {
    got := NewUser(1, "Alice")
    want := &User{ID: 1, Name: "Alice"}
    assertUser(t, got, want)
}
```

---

## j) Quiz / Practice for Me

### Quiz Questions

1. **What must test function names start with?**
   - [ ] `test`
   - [ ] `Test`
   - [ ] `_test`
   - [ ] `check`

2. **What parameter do test functions take?**
   - [ ] `*testing.T`
   - [ ] `testing.T`
   - [ ] `*test.T`
   - [ ] No parameter

3. **What is the command to run tests?**
   - [ ] `go test`
   - [ ] `go run test`
   - [ ] `go check`
   - [ ] `go verify`

4. **What does `t.Helper()` do?**
   - [ ] Marks function as test helper
   - [ ] Runs helper function
   - [ ] Skips test
   - [ ] Fails test

### Practice Tasks

**Task 1: Test String Functions**
- Create functions: `UpperCase(s string) string`, `LowerCase(s string) string`
- Write table-driven tests for both
- Test edge cases (empty string, special characters)

**Task 2: Test Stack Data Structure**
- Create Stack with Push, Pop, Peek methods
- Write tests for:
  - Pushing elements
  - Popping elements
  - Peeking without removing
  - Empty stack behavior

**Task 3: Test Calculator with Errors**
- Create calculator with Add, Subtract, Divide
- Write tests for:
  - Normal operations
  - Division by zero
  - Negative numbers
  - Large numbers

**Task 4: Benchmark Comparison**
- Create two functions that do the same thing differently
- Write benchmarks for both
- Compare performance

**Task 5: Integration Test**
- Create a function that reads a file and processes it
- Write test that:
  - Creates temporary file
  - Tests function
  - Cleans up file

### Answers

**Quiz Answers:**
1. `Test`
2. `*testing.T`
3. `go test`
4. Marks function as test helper

**Practice Solutions:**

**Task 1 Solution:**
```go
// string_ops.go
package main

import "strings"

func UpperCase(s string) string {
    return strings.ToUpper(s)
}

func LowerCase(s string) string {
    return strings.ToLower(s)
}

// string_ops_test.go
package main

import "testing"

func TestUpperCase(t *testing.T) {
    tests := []struct {
        name  string
        input string
        want  string
    }{
        {"lowercase", "hello", "HELLO"},
        {"mixed case", "HeLLo", "HELLO"},
        {"empty", "", ""},
        {"numbers", "hello123", "HELLO123"},
        {"special chars", "hello!", "HELLO!"},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := UpperCase(tt.input)
            if got != tt.want {
                t.Errorf("UpperCase(%q) = %q, want %q", tt.input, got, tt.want)
            }
        })
    }
}

func TestLowerCase(t *testing.T) {
    tests := []struct {
        name  string
        input string
        want  string
    }{
        {"uppercase", "HELLO", "hello"},
        {"mixed case", "HeLLo", "hello"},
        {"empty", "", ""},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := LowerCase(tt.input)
            if got != tt.want {
                t.Errorf("LowerCase(%q) = %q, want %q", tt.input, got, tt.want)
            }
        })
    }
}
```

**Task 2 Solution:**
```go
// stack.go
package main

import "errors"

var ErrEmptyStack = errors.New("stack is empty")

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
    if len(s.items) == 0 {
        return 0, ErrEmptyStack
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, nil
}

func (s *Stack) Peek() (int, error) {
    if len(s.items) == 0 {
        return 0, ErrEmptyStack
    }
    return s.items[len(s.items)-1], nil
}

// stack_test.go
package main

import "testing"

func TestStack(t *testing.T) {
    s := NewStack()
    
    // Test Push
    s.Push(1)
    s.Push(2)
    s.Push(3)
    
    // Test Peek
    peek, err := s.Peek()
    if err != nil {
        t.Errorf("Peek() error = %v, want nil", err)
    }
    if peek != 3 {
        t.Errorf("Peek() = %d, want 3", peek)
    }
    
    // Test Pop
    pop, err := s.Pop()
    if err != nil {
        t.Errorf("Pop() error = %v, want nil", err)
    }
    if pop != 3 {
        t.Errorf("Pop() = %d, want 3", pop)
    }
    
    // Test empty stack
    s.Pop()
    s.Pop()
    _, err = s.Pop()
    if err != ErrEmptyStack {
        t.Errorf("Pop() error = %v, want ErrEmptyStack", err)
    }
}
```

**Task 3 Solution:**
```go
// calculator.go
package main

import "errors"

var ErrDivideByZero = errors.New("division by zero")

type Calculator struct{}

func (c Calculator) Add(a, b int) int {
    return a + b
}

func (c Calculator) Subtract(a, b int) int {
    return a - b
}

func (c Calculator) Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, ErrDivideByZero
    }
    return a / b, nil
}

// calculator_test.go
package main

import "testing"

func TestCalculator(t *testing.T) {
    calc := Calculator{}
    
    // Test Add
    if got := calc.Add(5, 3); got != 8 {
        t.Errorf("Add(5, 3) = %d, want 8", got)
    }
    
    // Test Subtract
    if got := calc.Subtract(10, 4); got != 6 {
        t.Errorf("Subtract(10, 4) = %d, want 6", got)
    }
    
    // Test Divide - normal
    if got, err := calc.Divide(10, 2); got != 5 || err != nil {
        t.Errorf("Divide(10, 2) = %d, %v, want 5, nil", got, err)
    }
    
    // Test Divide - by zero
    if _, err := calc.Divide(10, 0); err != ErrDivideByZero {
        t.Errorf("Divide(10, 0) error = %v, want ErrDivideByZero", err)
    }
    
    // Test Divide - negative
    if got, err := calc.Divide(-10, 2); got != -5 || err != nil {
        t.Errorf("Divide(-10, 2) = %d, %v, want -5, nil", got, err)
    }
}
```

**Task 4 Solution:**
```go
// algorithms.go
package main

func SumLoop(numbers []int) int {
    sum := 0
    for _, n := range numbers {
        sum += n
    }
    return sum
}

func SumRecursive(numbers []int) int {
    if len(numbers) == 0 {
        return 0
    }
    return numbers[0] + SumRecursive(numbers[1:])
}

// algorithms_test.go
package main

import "testing"

func BenchmarkSumLoop(b *testing.B) {
    numbers := make([]int, 1000)
    for i := range numbers {
        numbers[i] = i
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        SumLoop(numbers)
    }
}

func BenchmarkSumRecursive(b *testing.B) {
    numbers := make([]int, 100)
    for i := range numbers {
        numbers[i] = i
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        SumRecursive(numbers)
    }
}
```

**Task 5 Solution:**
```go
// processor.go
package main

import (
    "bufio"
    "os"
    "strings"
)

func ProcessFile(filename string) (int, error) {
    file, err := os.Open(filename)
    if err != nil {
        return 0, err
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    count := 0
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line != "" {
            count++
        }
    }
    return count, scanner.Err()
}

// processor_test.go
package main

import (
    "os"
    "testing"
)

func TestProcessFile(t *testing.T) {
    // Create temporary file
    tmpfile, err := os.CreateTemp("", "test-*.txt")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove(tmpfile.Name()) // Cleanup
    
    // Write test data
    content := "line 1\nline 2\n\nline 3\n"
    if _, err := tmpfile.WriteString(content); err != nil {
        t.Fatal(err)
    }
    tmpfile.Close()
    
    // Test
    count, err := ProcessFile(tmpfile.Name())
    if err != nil {
        t.Errorf("ProcessFile() error = %v", err)
    }
    if count != 3 {
        t.Errorf("ProcessFile() = %d, want 3", count)
    }
}
```

---

## Key Takeaways

1. **Test files end with `_test.go`**
2. **Test functions start with `Test`**
3. **Use `t.Error` or `t.Fatal` for failures**
4. **Table-driven tests are idiomatic**

---

## Next Steps

**Ready? → [18_BEST_PRACTICES.md](./18_BEST_PRACTICES.md)**
