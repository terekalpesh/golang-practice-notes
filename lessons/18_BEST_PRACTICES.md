# Best Practices

## a) Overview

### What this topic is
Production-ready Go code practices - how to write maintainable, efficient, and idiomatic Go code.

### Why it exists in Go
Following best practices makes code readable, maintainable, and performant. Go has strong opinions about code style.

---

## b) Key Practices

### Code Organization
```go
// ✅ Good: Clear package structure
package user

type User struct { }
func NewUser() *User { }
func (u *User) Validate() error { }
```

### Error Handling
```go
// ✅ Good: Always check errors
result, err := doSomething()
if err != nil {
    return err
}

// ❌ Bad: Ignoring errors
result, _ := doSomething()
```

### Naming Conventions
```go
// ✅ Good: Clear, descriptive names
func CalculateTotal(items []Item) float64 { }

// ❌ Bad: Unclear names
func calc(x []y) float64 { }
```

---

## c) Explanation

### Best Practices Categories

**1. Code Style**
- Use `go fmt` (automatic formatting)
- Follow naming conventions
- Keep functions small

**2. Error Handling**
- Always check errors
- Add context to errors
- Return errors, don't log them

**3. Performance**
- Pre-allocate slices when size known
- Use pointers for large structs
- Avoid unnecessary allocations

**4. Concurrency**
- Use channels for communication
- Don't share memory
- Use WaitGroup for synchronization

---

## d) Python Comparison

### Python Best Practices
```python
# PEP 8 style guide
def calculate_total(items):
    return sum(item.price for item in items)
```

### Go Best Practices
```go
// go fmt style (automatic)
func CalculateTotal(items []Item) float64 {
    total := 0.0
    for _, item := range items {
        total += item.Price
    }
    return total
}
```

**Key Differences:**
- Python: PEP 8 (style guide)
- Go: `go fmt` (enforced style)
- Python: Flexible naming
- Go: Strict conventions (exported = capital)

---

## e) Visual Flow / Mental Model

### Code Review Checklist

```
1. Does it compile? (go build)
   ↓
2. Does it pass tests? (go test)
   ↓
3. Is it formatted? (go fmt)
   ↓
4. Are errors handled?
   ↓
5. Is it readable?
   ↓
6. Is it efficient?
```

---

## f) Demo Example

### Good vs Bad Examples

```go
// ✅ Good: Clear, handles errors
func ReadConfig(filename string) (*Config, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("reading config: %w", err)
    }
    
    var config Config
    if err := json.Unmarshal(data, &config); err != nil {
        return nil, fmt.Errorf("parsing config: %w", err)
    }
    
    return &config, nil
}

// ❌ Bad: Unclear, ignores errors
func read(f string) *Config {
    d, _ := ioutil.ReadFile(f)
    var c Config
    json.Unmarshal(d, &c)
    return &c
}
```

---

## g) Use Cases

- Production code
- Library development
- Team collaboration
- Code reviews

---

## h) Do's and Don'ts

### ✅ Do's

1. **Use go fmt**
2. **Handle all errors**
3. **Write tests**
4. **Use meaningful names**
5. **Keep functions small**

### ❌ Don'ts

1. **Don't ignore errors**
2. **Don't use global variables**
3. **Don't write overly complex code**
4. **Don't skip tests**

---

## i) Solved Practice Examples

### Example 1: Good Error Handling

**Task:** Compare good vs bad error handling patterns.

**Solution:**

**❌ Bad:**
```go
func readConfig(filename string) *Config {
    data, _ := ioutil.ReadFile(filename)  // Ignoring error!
    var config Config
    json.Unmarshal(data, &config)  // Ignoring error!
    return &config
}
```

**✅ Good:**
```go
func readConfig(filename string) (*Config, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("reading config file: %w", err)
    }
    
    var config Config
    if err := json.Unmarshal(data, &config); err != nil {
        return nil, fmt.Errorf("parsing config: %w", err)
    }
    
    return &config, nil
}
```

### Example 2: Proper Function Design

**Task:** Show good function design principles.

**Solution:**

**❌ Bad:**
```go
func process(data []int, flag bool, count int, name string) {
    // Too many parameters, unclear purpose
    // Does too many things
    // No return value for errors
}
```

**✅ Good:**
```go
type ProcessOptions struct {
    Count int
    Name  string
}

func ProcessData(data []int, options ProcessOptions) error {
    // Clear purpose: process data
    // Options struct for multiple parameters
    // Returns error for failure cases
    if len(data) == 0 {
        return errors.New("data cannot be empty")
    }
    // Process...
    return nil
}
```

### Example 3: Efficient Slice Usage

**Task:** Show efficient slice operations.

**Solution:**

**❌ Bad:**
```go
func processItems(items []Item) []Result {
    results := []Result{}  // Grows multiple times
    for _, item := range items {
        result := process(item)
        results = append(results, result)
    }
    return results
}
```

**✅ Good:**
```go
func processItems(items []Item) []Result {
    results := make([]Result, 0, len(items))  // Pre-allocate capacity
    for _, item := range items {
        result := process(item)
        results = append(results, result)
    }
    return results
}
```

### Example 4: Proper Interface Design

**Task:** Show good interface design.

**Solution:**

**❌ Bad:**
```go
type Everything interface {
    Method1()
    Method2()
    Method3()
    Method4()
    Method5()
    // Too many methods!
}
```

**✅ Good:**
```go
// Small, focused interfaces
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type ReadWriter interface {
    Reader
    Writer
}
```

### Example 5: Context Usage

**Task:** Show proper context usage for cancellation.

**Solution:**

**❌ Bad:**
```go
func longRunningTask() {
    // No way to cancel
    for {
        // work...
    }
}
```

**✅ Good:**
```go
func longRunningTask(ctx context.Context) error {
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()  // Respect cancellation
        default:
            // work...
        }
    }
}
```

### Example 6: Resource Cleanup

**Task:** Show proper resource management.

**Solution:**

**❌ Bad:**
```go
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    // What if error occurs? File not closed!
    process(file)
    file.Close()
    return nil
}
```

**✅ Good:**
```go
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()  // Always closes, even on error
    
    return process(file)
}
```

### Example 7: Naming Conventions

**Task:** Show good naming practices.

**Solution:**

**❌ Bad:**
```go
func calc(x []y) z {
    // Unclear names
}

var d = 10  // What is d?
var tmp = []string{}  // Temporary what?
```

**✅ Good:**
```go
func CalculateTotal(items []Item) float64 {
    // Clear, descriptive names
}

var userCount = 10  // Clear purpose
var userNames = []string{}  // Clear what it contains
```

### Example 8: Testing Best Practices

**Task:** Show good test structure.

**Solution:**

**❌ Bad:**
```go
func TestEverything(t *testing.T) {
    // One huge test
    // Tests multiple things
    // Hard to debug
}
```

**✅ Good:**
```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        want     int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
        {"zero", 0, 5, 5},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("Add() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### Example 9: Concurrency Best Practices

**Task:** Show safe concurrency patterns.

**Solution:**

**❌ Bad:**
```go
var counter int  // Shared variable

func increment() {
    counter++  // Race condition!
}
```

**✅ Good:**
```go
// Option 1: Use channels
func increment(ch chan<- int) {
    ch <- 1
}

// Option 2: Use mutex if needed
type SafeCounter struct {
    mu sync.Mutex
    value int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}
```

### Example 10: Package Organization

**Task:** Show good package structure.

**Solution:**

**❌ Bad:**
```
project/
├── file1.go
├── file2.go
├── file3.go
└── file4.go
// Everything in one package
```

**✅ Good:**
```
project/
├── go.mod
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handlers/
│   ├── models/
│   └── services/
├── pkg/
│   └── utils/
└── README.md
```

---

## j) Quiz / Practice for Me

### Quiz Questions

1. **What is the Go way to handle errors?**
   - [ ] Exceptions
   - [ ] Return error as value
   - [ ] Global error variable
   - [ ] Error callbacks

2. **What should you use for resource cleanup?**
   - [ ] try/finally
   - [ ] defer
   - [ ] Manual cleanup
   - [ ] Garbage collector only

3. **How should you design interfaces?**
   - [ ] Large with many methods
   - [ ] Small and focused
   - [ ] One interface for everything
   - [ ] No interfaces

4. **What is the preferred way to test in Go?**
   - [ ] One test per function
   - [ ] Table-driven tests
   - [ ] No tests needed
   - [ ] Only integration tests

### Practice Tasks

**Task 1: Refactor Bad Code**
- Given bad code with multiple issues:
  - Ignored errors
  - Poor naming
  - No error handling
  - Resource leaks
- Refactor to follow best practices

**Task 2: Design API Handler**
- Create HTTP handler that:
  - Handles errors properly
  - Uses context for cancellation
  - Returns proper status codes
  - Has good error messages

**Task 3: Create Reusable Package**
- Design a package with:
  - Clear exports
  - Good documentation
  - Proper error handling
  - Example usage

**Task 4: Optimize Performance**
- Given inefficient code:
  - Multiple allocations
  - Unnecessary copies
  - Inefficient loops
- Optimize following Go best practices

**Task 5: Write Production-Ready Code**
- Create a complete feature with:
  - Proper error handling
  - Tests
  - Documentation
  - Logging
  - Configuration

### Answers

**Quiz Answers:**
1. Return error as value
2. defer
3. Small and focused
4. Table-driven tests

**Practice Solutions:**

**Task 1 Solution:**

**Before (Bad):**
```go
func process(f string) *Result {
    d, _ := ioutil.ReadFile(f)
    var r Result
    json.Unmarshal(d, &r)
    return &r
}
```

**After (Good):**
```go
func ProcessFile(filename string) (*Result, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("reading file %s: %w", filename, err)
    }
    
    var result Result
    if err := json.Unmarshal(data, &result); err != nil {
        return nil, fmt.Errorf("parsing result: %w", err)
    }
    
    return &result, nil
}
```

**Task 2 Solution:**
```go
package handlers

import (
    "context"
    "encoding/json"
    "net/http"
    "time"
)

type UserHandler struct {
    service UserService
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()
    
    userID := r.URL.Query().Get("id")
    if userID == "" {
        http.Error(w, "user ID required", http.StatusBadRequest)
        return
    }
    
    user, err := h.service.GetUser(ctx, userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
```

**Task 3 Solution:**
```go
// Package: validator
// Package validator provides email and string validation functions.

package validator

import (
    "errors"
    "strings"
)

var ErrInvalidEmail = errors.New("invalid email format")

// ValidateEmail checks if an email address is valid.
// It returns an error if the email is invalid.
func ValidateEmail(email string) error {
    email = strings.TrimSpace(email)
    if email == "" {
        return ErrInvalidEmail
    }
    
    if !strings.Contains(email, "@") {
        return ErrInvalidEmail
    }
    
    parts := strings.Split(email, "@")
    if len(parts) != 2 {
        return ErrInvalidEmail
    }
    
    if parts[0] == "" || parts[1] == "" {
        return ErrInvalidEmail
    }
    
    return nil
}

// Example usage:
//   if err := validator.ValidateEmail("user@example.com"); err != nil {
//       log.Fatal(err)
//   }
```

**Task 4 Solution:**

**Before (Inefficient):**
```go
func process(items []Item) []string {
    result := []string{}
    for i := 0; i < len(items); i++ {
        item := items[i]  // Unnecessary copy
        result = append(result, item.Name)
    }
    return result
}
```

**After (Efficient):**
```go
func process(items []Item) []string {
    result := make([]string, 0, len(items))  // Pre-allocate
    for _, item := range items {  // Range is more efficient
        result = append(result, item.Name)
    }
    return result
}
```

**Task 5 Solution:**
```go
package main

import (
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "os"
    "time"
)

type Config struct {
    Host    string        `json:"host"`
    Port    int           `json:"port"`
    Timeout time.Duration `json:"timeout"`
}

func loadConfig(filename string) (*Config, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("loading config: %w", err)
    }
    
    var config Config
    if err := json.Unmarshal(data, &config); err != nil {
        return nil, fmt.Errorf("parsing config: %w", err)
    }
    
    return &config, nil
}

type Service struct {
    config *Config
    logger *log.Logger
}

func NewService(config *Config) *Service {
    return &Service{
        config: config,
        logger: log.New(os.Stdout, "[SERVICE] ", log.LstdFlags),
    }
}

func (s *Service) Process(ctx context.Context, data string) error {
    s.logger.Printf("Processing data: %s", data)
    
    select {
    case <-ctx.Done():
        return ctx.Err()
    case <-time.After(s.config.Timeout):
        // Simulate work
        s.logger.Printf("Processing complete")
        return nil
    }
}

func main() {
    config, err := loadConfig("config.json")
    if err != nil {
        log.Fatal(err)
    }
    
    service := NewService(config)
    
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    if err := service.Process(ctx, "test data"); err != nil {
        if errors.Is(err, context.DeadlineExceeded) {
            log.Println("Processing timed out")
        } else {
            log.Fatal(err)
        }
    }
}
```

---

## Key Takeaways

1. **Use go fmt** - Automatic formatting
2. **Handle errors** - Always check
3. **Write tests** - Essential for reliability
4. **Follow conventions** - Capital = exported
5. **Keep it simple** - Go values simplicity

---

## Complete Learning Path Summary

You've now learned:
- ✅ Go fundamentals
- ✅ Variables and types
- ✅ Functions and control flow
- ✅ Pointers and structs
- ✅ Methods and interfaces
- ✅ Error handling
- ✅ Collections (slices, maps)
- ✅ Concurrency
- ✅ Packages and modules
- ✅ File I/O and JSON
- ✅ Testing
- ✅ Best practices

**Congratulations! You're ready to build production Go applications! 🎉**

---

## Next Steps

1. **Build projects** - Apply what you learned
2. **Read Go code** - Study standard library
3. **Join Go community** - Learn from others
4. **Keep practicing** - Mastery comes with practice

**Use the [Cheat Sheet](../reference/99_CHEAT_SHEET.md) for quick reference!**
