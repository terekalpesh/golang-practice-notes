# JSON Handling

## a) Overview

### What this topic is
Working with JSON data - encoding Go structs to JSON and decoding JSON to Go structs.

### Why it exists in Go
JSON is the standard format for APIs and data exchange. Go has excellent built-in JSON support.

---

## b) Syntax

### Encoding (Go to JSON)
```go
import "encoding/json"

data, err := json.Marshal(structValue)
jsonString := string(data)

// Pretty print
data, err := json.MarshalIndent(structValue, "", "  ")
```

### Decoding (JSON to Go)
```go
var result MyStruct
err := json.Unmarshal(jsonData, &result)
```

### Struct Tags
```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    City string `json:"city,omitempty"`
}
```

---

## c) Explanation

### Step-by-Step JSON Operations

**1. Define struct with tags**
```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```
- Tags control JSON field names
- `omitempty` excludes zero values

**2. Marshal (Go → JSON)**
```go
data, err := json.Marshal(person)
```
- Converts struct to JSON bytes
- Returns `[]byte`

**3. Unmarshal (JSON → Go)**
```go
err := json.Unmarshal(data, &person)
```
- Converts JSON to struct
- Must pass pointer

---

## d) Python Comparison

### Python JSON
```python
import json

# Encode
data = json.dumps({"name": "Alice", "age": 30})

# Decode
person = json.loads(data)
```

### Go JSON
```go
import "encoding/json"

// Encode
data, _ := json.Marshal(Person{Name: "Alice", Age: 30})

// Decode
var person Person
json.Unmarshal(data, &person)
```

**Key Differences:**
- Python: Dicts automatically become JSON
- Go: Need structs with tags
- Python: `dumps`/`loads` with strings
- Go: `Marshal`/`Unmarshal` with bytes

---

## e) Visual Flow / Mental Model

### JSON Encoding Flow

```
Go Struct
  ↓
json.Marshal()
  ↓
JSON bytes
  ↓
string(data) → JSON string
```

### JSON Decoding Flow

```
JSON string/bytes
  ↓
json.Unmarshal()
  ↓
Go Struct (via pointer)
```

---

## f) Demo Example

### Complete Example

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    City string `json:"city,omitempty"`
}

func main() {
    // 1. Marshal (Go → JSON)
    person := Person{
        Name: "Alice",
        Age:  30,
    }
    
    data, err := json.Marshal(person)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("JSON:", string(data))
    
    // 2. Pretty print
    pretty, _ := json.MarshalIndent(person, "", "  ")
    fmt.Println("Pretty JSON:")
    fmt.Println(string(pretty))
    
    // 3. Unmarshal (JSON → Go)
    jsonStr := `{"name":"Bob","age":25,"city":"NYC"}`
    var person2 Person
    err = json.Unmarshal([]byte(jsonStr), &person2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Person: %+v\n", person2)
    
    // 4. Array of objects
    people := []Person{
        {Name: "Alice", Age: 30},
        {Name: "Bob", Age: 25},
    }
    data, _ = json.Marshal(people)
    fmt.Println("Array JSON:", string(data))
}
```

---

## g) Use Cases

- API requests/responses
- Configuration files
- Data serialization
- Inter-service communication

---

## h) Do's and Don'ts

### ✅ Do's

1. **Use struct tags for field names**
2. **Use omitempty for optional fields**
3. **Always check errors**

### ❌ Don'ts

1. **Don't forget pointer for Unmarshal**
2. **Don't ignore errors**

---

## i) Solved Practice Examples

### Example 1: User API Response

**Task:** Create a struct for API response and marshal/unmarshal JSON.

**Solution:**
```go
package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Active   bool   `json:"active"`
}

type APIResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Data    *User  `json:"data,omitempty"`
}

func main() {
    // Create user
    user := User{
        ID:       1,
        Username: "alice",
        Email:    "alice@example.com",
        Active:   true,
    }
    
    // Marshal to JSON
    response := APIResponse{
        Status:  "success",
        Message: "User retrieved",
        Data:    &user,
    }
    
    jsonData, err := json.MarshalIndent(response, "", "  ")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("JSON Response:")
    fmt.Println(string(jsonData))
    
    // Unmarshal from JSON
    jsonStr := `{"status":"success","message":"User created","data":{"id":2,"username":"bob","email":"bob@example.com","active":true}}`
    var newResponse APIResponse
    err = json.Unmarshal([]byte(jsonStr), &newResponse)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    fmt.Printf("\nParsed Response:\n")
    fmt.Printf("Status: %s\n", newResponse.Status)
    fmt.Printf("Message: %s\n", newResponse.Message)
    if newResponse.Data != nil {
        fmt.Printf("User: %+v\n", *newResponse.Data)
    }
}
```

### Example 2: Nested JSON Structures

**Task:** Work with nested JSON structures (user with address).

**Solution:**
```go
package main

import (
    "encoding/json"
    "fmt"
)

type Address struct {
    Street  string `json:"street"`
    City    string `json:"city"`
    State   string `json:"state"`
    ZipCode string `json:"zip_code"`
}

type Person struct {
    Name    string   `json:"name"`
    Age     int      `json:"age"`
    Address Address  `json:"address"`
    Phones  []string `json:"phones,omitempty"`
}

func main() {
    person := Person{
        Name: "John Doe",
        Age:  30,
        Address: Address{
            Street:  "123 Main St",
            City:    "New York",
            State:   "NY",
            ZipCode: "10001",
        },
        Phones: []string{"555-0100", "555-0101"},
    }
    
    jsonData, _ := json.MarshalIndent(person, "", "  ")
    fmt.Println("Person JSON:")
    fmt.Println(string(jsonData))
    
    // Unmarshal nested JSON
    jsonStr := `{
        "name": "Jane Smith",
        "age": 25,
        "address": {
            "street": "456 Oak Ave",
            "city": "Los Angeles",
            "state": "CA",
            "zip_code": "90001"
        },
        "phones": ["555-0200"]
    }`
    
    var newPerson Person
    json.Unmarshal([]byte(jsonStr), &newPerson)
    fmt.Printf("\nParsed Person: %+v\n", newPerson)
}
```

### Example 3: JSON Array Handling

**Task:** Marshal and unmarshal arrays of objects.

**Solution:**
```go
package main

import (
    "encoding/json"
    "fmt"
)

type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

func main() {
    // Marshal array
    products := []Product{
        {ID: 1, Name: "Laptop", Price: 999.99},
        {ID: 2, Name: "Mouse", Price: 29.99},
        {ID: 3, Name: "Keyboard", Price: 79.99},
    }
    
    jsonData, _ := json.MarshalIndent(products, "", "  ")
    fmt.Println("Products JSON:")
    fmt.Println(string(jsonData))
    
    // Unmarshal array
    jsonStr := `[
        {"id": 4, "name": "Monitor", "price": 299.99},
        {"id": 5, "name": "Webcam", "price": 49.99}
    ]`
    
    var newProducts []Product
    json.Unmarshal([]byte(jsonStr), &newProducts)
    
    fmt.Println("\nParsed Products:")
    for _, p := range newProducts {
        fmt.Printf("  %s: $%.2f\n", p.Name, p.Price)
    }
}
```

### Example 4: Custom JSON Marshaling

**Task:** Implement custom JSON marshaling for a type.

**Solution:**
```go
package main

import (
    "encoding/json"
    "fmt"
    "strings"
)

type Email string

func (e Email) MarshalJSON() ([]byte, error) {
    // Mask email for security
    parts := strings.Split(string(e), "@")
    if len(parts) != 2 {
        return json.Marshal(string(e))
    }
    masked := parts[0][:1] + "***@" + parts[1]
    return json.Marshal(masked)
}

type User struct {
    Name  string `json:"name"`
    Email Email  `json:"email"`
}

func main() {
    user := User{
        Name:  "Alice",
        Email: "alice@example.com",
    }
    
    jsonData, _ := json.MarshalIndent(user, "", "  ")
    fmt.Println("User with masked email:")
    fmt.Println(string(jsonData))
}
```

### Example 5: JSON with Optional Fields

**Task:** Handle JSON with optional/nullable fields.

**Solution:**
```go
package main

import (
    "encoding/json"
    "fmt"
)

type Config struct {
    Host     string  `json:"host"`
    Port     int     `json:"port"`
    Timeout  *int    `json:"timeout,omitempty"`  // Pointer for optional
    Debug    bool    `json:"debug,omitempty"`
    LogLevel *string `json:"log_level,omitempty"`
}

func main() {
    // Config with all fields
    timeout := 30
    logLevel := "info"
    config1 := Config{
        Host:     "localhost",
        Port:     8080,
        Timeout:  &timeout,
        Debug:    true,
        LogLevel: &logLevel,
    }
    
    json1, _ := json.MarshalIndent(config1, "", "  ")
    fmt.Println("Full config:")
    fmt.Println(string(json1))
    
    // Config with minimal fields
    config2 := Config{
        Host: "example.com",
        Port: 443,
    }
    
    json2, _ := json.MarshalIndent(config2, "", "  ")
    fmt.Println("\nMinimal config:")
    fmt.Println(string(json2))
    
    // Unmarshal with missing fields
    jsonStr := `{"host":"test.com","port":3000,"timeout":60}`
    var config3 Config
    json.Unmarshal([]byte(jsonStr), &config3)
    fmt.Printf("\nParsed config: %+v\n", config3)
    if config3.Timeout != nil {
        fmt.Printf("Timeout: %d\n", *config3.Timeout)
    }
}
```

### Example 6: JSON Streaming (for Large Data)

**Task:** Use json.Decoder for streaming large JSON files.

**Solution:**
```go
package main

import (
    "encoding/json"
    "fmt"
    "strings"
)

type Item struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func main() {
    // Simulate large JSON stream
    jsonStream := `{"id":1,"name":"Item 1"}{"id":2,"name":"Item 2"}{"id":3,"name":"Item 3"}`
    reader := strings.NewReader(jsonStream)
    decoder := json.NewDecoder(reader)
    
    for {
        var item Item
        if err := decoder.Decode(&item); err != nil {
            break // End of stream
        }
        fmt.Printf("Decoded: %+v\n", item)
    }
}
```

---

## j) Quiz / Practice for Me

### Quiz Questions

1. **What does `omitempty` do in JSON tags?**
   - [ ] Makes field required
   - [ ] Omits field if value is zero/empty
   - [ ] Makes field optional
   - [ ] Encrypts the field

2. **What must you pass to `json.Unmarshal`?**
   - [ ] Value
   - [ ] Pointer
   - [ ] Either
   - [ ] Reference

3. **What does `json.MarshalIndent` do?**
   - [ ] Encrypts JSON
   - [ ] Formats JSON with indentation
   - [ ] Validates JSON
   - [ ] Compresses JSON

4. **Can you use slices in JSON structs?**
   - [ ] No, never
   - [ ] Yes, always
   - [ ] Only strings
   - [ ] Only numbers

### Practice Tasks

**Task 1: Book Library JSON**
- Create `Book` struct with: Title, Author, ISBN, Price
- Create `Library` struct containing slice of Books
- Marshal library to JSON
- Unmarshal JSON back to Library

**Task 2: API Error Response**
- Create `ErrorResponse` struct with: Code, Message, Details (map)
- Handle both success and error JSON responses
- Marshal error response with details

**Task 3: Configuration Parser**
- Read JSON config file
- Parse into Config struct
- Handle missing optional fields gracefully
- Print configuration

**Task 4: JSON Validator**
- Create function that validates JSON structure
- Check required fields
- Return validation errors

**Task 5: JSON Transformer**
- Read JSON with one structure
- Transform to different structure
- Write transformed JSON

### Answers

**Quiz Answers:**
1. Omits field if value is zero/empty
2. Pointer
3. Formats JSON with indentation
4. Yes, always

**Practice Solutions:**

**Task 1 Solution:**
```go
package main

import (
    "encoding/json"
    "fmt"
)

type Book struct {
    Title  string  `json:"title"`
    Author string  `json:"author"`
    ISBN   string  `json:"isbn"`
    Price  float64 `json:"price"`
}

type Library struct {
    Name  string `json:"name"`
    Books []Book `json:"books"`
}

func main() {
    library := Library{
        Name: "My Library",
        Books: []Book{
            {Title: "Go Programming", Author: "Author A", ISBN: "123-456", Price: 29.99},
            {Title: "Python Basics", Author: "Author B", ISBN: "789-012", Price: 24.99},
        },
    }
    
    jsonData, _ := json.MarshalIndent(library, "", "  ")
    fmt.Println("Library JSON:")
    fmt.Println(string(jsonData))
    
    // Unmarshal
    jsonStr := `{"name":"Public Library","books":[{"title":"Rust Guide","author":"Author C","isbn":"345-678","price":34.99}]}`
    var newLibrary Library
    json.Unmarshal([]byte(jsonStr), &newLibrary)
    fmt.Printf("\nParsed Library: %+v\n", newLibrary)
}
```

**Task 2 Solution:**
```go
package main

import (
    "encoding/json"
    "fmt"
)

type ErrorResponse struct {
    Code    int               `json:"code"`
    Message string            `json:"message"`
    Details map[string]string `json:"details,omitempty"`
}

type SuccessResponse struct {
    Status string      `json:"status"`
    Data   interface{} `json:"data"`
}

func main() {
    // Error response
    errorResp := ErrorResponse{
        Code:    400,
        Message: "Validation failed",
        Details: map[string]string{
            "email": "Invalid email format",
            "age":   "Age must be positive",
        },
    }
    
    jsonData, _ := json.MarshalIndent(errorResp, "", "  ")
    fmt.Println("Error Response:")
    fmt.Println(string(jsonData))
    
    // Success response
    successResp := SuccessResponse{
        Status: "success",
        Data:   map[string]string{"user_id": "123"},
    }
    
    jsonData2, _ := json.MarshalIndent(successResp, "", "  ")
    fmt.Println("\nSuccess Response:")
    fmt.Println(string(jsonData2))
}
```

**Task 3 Solution:**
```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Config struct {
    DatabaseURL *string `json:"database_url,omitempty"`
    Port        int     `json:"port"`
    Debug       bool    `json:"debug"`
    LogLevel    string  `json:"log_level,omitempty"`
}

func loadConfig(filename string) (*Config, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    
    var config Config
    if err := json.Unmarshal(data, &config); err != nil {
        return nil, err
    }
    
    return &config, nil
}

func main() {
    config, err := loadConfig("config.json")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    fmt.Printf("Config: %+v\n", config)
    if config.DatabaseURL != nil {
        fmt.Printf("Database URL: %s\n", *config.DatabaseURL)
    }
}
```

**Task 4 Solution:**
```go
package main

import (
    "encoding/json"
    "fmt"
    "strings"
)

type ValidationError struct {
    Field   string
    Message string
}

func validateUserJSON(jsonStr string) ([]ValidationError, error) {
    var user map[string]interface{}
    if err := json.Unmarshal([]byte(jsonStr), &user); err != nil {
        return nil, err
    }
    
    var errors []ValidationError
    required := []string{"name", "email", "age"}
    
    for _, field := range required {
        if _, exists := user[field]; !exists {
            errors = append(errors, ValidationError{
                Field:   field,
                Message: fmt.Sprintf("%s is required", field),
            })
        }
    }
    
    return errors, nil
}

func main() {
    jsonStr := `{"name":"Alice","email":"alice@example.com"}`
    errors, _ := validateUserJSON(jsonStr)
    
    if len(errors) > 0 {
        fmt.Println("Validation errors:")
        for _, err := range errors {
            fmt.Printf("  %s: %s\n", err.Field, err.Message)
        }
    } else {
        fmt.Println("JSON is valid")
    }
}
```

**Task 5 Solution:**
```go
package main

import (
    "encoding/json"
    "fmt"
)

type OldFormat struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Age       int    `json:"age"`
}

type NewFormat struct {
    FullName string `json:"full_name"`
    Age      int    `json:"age"`
}

func transform(old OldFormat) NewFormat {
    return NewFormat{
        FullName: old.FirstName + " " + old.LastName,
        Age:      old.Age,
    }
}

func main() {
    oldJSON := `{"first_name":"John","last_name":"Doe","age":30}`
    var old OldFormat
    json.Unmarshal([]byte(oldJSON), &old)
    
    new := transform(old)
    newJSON, _ := json.MarshalIndent(new, "", "  ")
    
    fmt.Println("Transformed JSON:")
    fmt.Println(string(newJSON))
}
```

---

## Key Takeaways

1. **Marshal = Go → JSON** - `json.Marshal()`
2. **Unmarshal = JSON → Go** - `json.Unmarshal()`
3. **Struct tags control JSON** - `json:"field_name"`
4. **Always pass pointer to Unmarshal** - `&struct`

---

## Next Steps

**Ready? → [17_TESTING.md](./17_TESTING.md)**
