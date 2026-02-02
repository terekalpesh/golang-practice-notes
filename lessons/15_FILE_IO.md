# File I/O

## a) Overview

### What this topic is
Reading from and writing to files in Go. Working with the file system.

### Why it exists in Go
File I/O is essential for most applications - reading configs, writing logs, processing data files.

---

## b) Syntax

### Reading Files
```go
import "io/ioutil"
data, err := ioutil.ReadFile("file.txt")

import "os"
file, err := os.Open("file.txt")
defer file.Close()

import "bufio"
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
}
```

### Writing Files
```go
import "io/ioutil"
err := ioutil.WriteFile("file.txt", data, 0644)

import "os"
file, err := os.Create("file.txt")
defer file.Close()
file.WriteString("content")
```

---

## c) Explanation

### Step-by-Step File Operations

**1. Read entire file**
```go
data, err := ioutil.ReadFile("file.txt")
```
- Reads all content at once
- Returns `[]byte`
- Simple but uses memory for large files

**2. Read line by line**
```go
file, _ := os.Open("file.txt")
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
}
```
- Memory efficient
- Processes one line at a time

**3. Write file**
```go
err := ioutil.WriteFile("file.txt", data, 0644)
```
- Writes all data at once
- `0644` = file permissions

### Characteristics

#### File Operation Characteristics
- **Error handling**: All file operations return errors (must check)
- **Resource management**: Files must be closed (use `defer`)
- **Byte-oriented**: Files read/write as `[]byte`
- **Text vs Binary**: No distinction in Go (treat as bytes)
- **File modes**: Read, write, append, create modes
- **Permissions**: Unix-style file permissions (octal notation)

#### Reading Characteristics
- **Read entire file**: Simple but memory-intensive for large files
- **Streaming read**: Line-by-line or chunk-by-chunk (memory efficient)
- **Scanner**: Convenient for line-by-line reading
- **Buffer**: Can use buffered I/O for performance
- **Error handling**: EOF (end of file) is normal, not error

#### Writing Characteristics
- **Write entire file**: Simple but overwrites existing file
- **Append mode**: Add to end of file without overwriting
- **Buffered writing**: Use `bufio.Writer` for performance
- **Atomic writes**: Write to temp file then rename (for safety)
- **File creation**: Creates file if doesn't exist (with appropriate mode)

#### Data Characteristics
- **Byte slices**: Files work with `[]byte` (convert strings as needed)
- **Text encoding**: Handle encoding explicitly (UTF-8, etc.)
- **Line endings**: Handle different line endings (Unix, Windows, Mac)
- **File size**: Can check file size before reading
- **File metadata**: Can access file info (size, mod time, permissions)

---

## d) Python Comparison

### Python File I/O
```python
# Read entire file
with open("file.txt", "r") as f:
    content = f.read()

# Read line by line
with open("file.txt", "r") as f:
    for line in f:
        print(line)

# Write file
with open("file.txt", "w") as f:
    f.write("content")
```

### Go File I/O
```go
// Read entire file
data, err := ioutil.ReadFile("file.txt")

// Read line by line
file, _ := os.Open("file.txt")
defer file.Close()
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
}

// Write file
err := ioutil.WriteFile("file.txt", data, 0644)
```

**Key Differences:**
- Python: `with` statement for auto-close
- Go: `defer file.Close()` for cleanup
- Python: String mode ("r", "w")
- Go: Use different functions (ReadFile, WriteFile)

---

## e) Visual Flow / Mental Model

### File Read Flow

```
1. Open file
   ↓
2. Read data (all at once or line by line)
   ↓
3. Process data
   ↓
4. Close file (defer ensures this)
```

---

## f) Demo Example

### Complete Example

```go
package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    // 1. Read entire file
    data, err := ioutil.ReadFile("example.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("File content:", string(data))
    
    // 2. Write file
    content := []byte("Hello, Go!\nThis is a test file.")
    err = ioutil.WriteFile("output.txt", content, 0644)
    if err != nil {
        fmt.Println("Error:", err)
    }
    
    // 3. Read line by line
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    lineNum := 1
    for scanner.Scan() {
        fmt.Printf("Line %d: %s\n", lineNum, scanner.Text())
        lineNum++
    }
    
    // 4. Write line by line
    outFile, err := os.Create("output2.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer outFile.Close()
    
    writer := bufio.NewWriter(outFile)
    writer.WriteString("Line 1\n")
    writer.WriteString("Line 2\n")
    writer.Flush()
}
```

---

## g) Use Cases

- Reading configuration files
- Processing log files
- Writing reports
- Data import/export

---

## h) Do's and Don'ts

### ✅ Do's

1. **Always check errors**
2. **Use defer to close files**
3. **Use bufio for line-by-line reading**

### ❌ Don'ts

1. **Don't forget to close files**
2. **Don't ignore errors**

---

## i) Solved Practice Examples

### Example 1: Read and Count Lines

**Task:** Read a file and count the number of lines.

**Solution:**
```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func countLines(filename string) (int, error) {
    file, err := os.Open(filename)
    if err != nil {
        return 0, err
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    count := 0
    for scanner.Scan() {
        count++
    }
    
    if err := scanner.Err(); err != nil {
        return 0, err
    }
    
    return count, nil
}

func main() {
    count, err := countLines("data.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("File has %d lines\n", count)
}
```

### Example 2: Copy File

**Task:** Copy contents from one file to another.

**Solution:**
```go
package main

import (
    "fmt"
    "io"
    "os"
)

func copyFile(src, dst string) error {
    sourceFile, err := os.Open(src)
    if err != nil {
        return fmt.Errorf("opening source: %w", err)
    }
    defer sourceFile.Close()
    
    destFile, err := os.Create(dst)
    if err != nil {
        return fmt.Errorf("creating destination: %w", err)
    }
    defer destFile.Close()
    
    _, err = io.Copy(destFile, sourceFile)
    if err != nil {
        return fmt.Errorf("copying: %w", err)
    }
    
    return nil
}

func main() {
    err := copyFile("source.txt", "destination.txt")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("File copied successfully")
    }
}
```

### Example 3: Write Log File

**Task:** Append log messages to a log file with timestamps.

**Solution:**
```go
package main

import (
    "fmt"
    "os"
    "time"
)

func writeLog(filename, message string) error {
    file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()
    
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    logEntry := fmt.Sprintf("[%s] %s\n", timestamp, message)
    
    _, err = file.WriteString(logEntry)
    return err
}

func main() {
    messages := []string{
        "Application started",
        "Processing request",
        "Request completed",
    }
    
    for _, msg := range messages {
        if err := writeLog("app.log", msg); err != nil {
            fmt.Printf("Error writing log: %v\n", err)
        }
    }
    fmt.Println("Logs written successfully")
}
```

### Example 4: Read CSV-like File

**Task:** Read a file with comma-separated values and parse each line.

**Solution:**
```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func readCSV(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    lineNum := 1
    
    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Split(line, ",")
        
        fmt.Printf("Line %d: ", lineNum)
        for i, field := range fields {
            fmt.Printf("Field%d=%s ", i+1, strings.TrimSpace(field))
        }
        fmt.Println()
        lineNum++
    }
    
    return scanner.Err()
}

func main() {
    err := readCSV("data.csv")
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

### Example 5: File Exists Check

**Task:** Check if a file exists before reading it.

**Solution:**
```go
package main

import (
    "fmt"
    "os"
)

func fileExists(filename string) bool {
    _, err := os.Stat(filename)
    return !os.IsNotExist(err)
}

func readFileIfExists(filename string) (string, error) {
    if !fileExists(filename) {
        return "", fmt.Errorf("file %s does not exist", filename)
    }
    
    data, err := os.ReadFile(filename)
    if err != nil {
        return "", err
    }
    
    return string(data), nil
}

func main() {
    content, err := readFileIfExists("config.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("File content:")
    fmt.Println(content)
}
```

### Example 6: Read Configuration File

**Task:** Read a configuration file with key-value pairs (one per line).

**Solution:**
```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func readConfig(filename string) (map[string]string, error) {
    config := make(map[string]string)
    
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" || strings.HasPrefix(line, "#") {
            continue // Skip empty lines and comments
        }
        
        parts := strings.SplitN(line, "=", 2)
        if len(parts) == 2 {
            key := strings.TrimSpace(parts[0])
            value := strings.TrimSpace(parts[1])
            config[key] = value
        }
    }
    
    return config, scanner.Err()
}

func main() {
    config, err := readConfig("config.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    fmt.Println("Configuration:")
    for key, value := range config {
        fmt.Printf("  %s = %s\n", key, value)
    }
}
```

---

## j) Quiz / Practice for Me

### Quiz Questions

1. **What package is used for reading entire files?**
   - [ ] `os`
   - [ ] `io/ioutil` or `os`
   - [ ] `bufio`
   - [ ] `file`

2. **What does `defer file.Close()` do?**
   - [ ] Closes file immediately
   - [ ] Closes file when function returns
   - [ ] Doesn't close file
   - [ ] Closes file on error only

3. **What is the difference between `os.Open` and `os.Create`?**
   - [ ] No difference
   - [ ] Open reads, Create writes
   - [ ] Open appends, Create overwrites
   - [ ] Open for reading, Create for writing (creates if not exists)

4. **What permission mode is `0644`?**
   - [ ] Read-only
   - [ ] Write-only
   - [ ] Read-write for owner, read for others
   - [ ] Execute only

### Practice Tasks

**Task 1: Word Counter**
- Read a text file
- Count total words
- Count unique words
- Print statistics

**Task 2: File Merger**
- Read multiple files
- Merge their contents into one file
- Add file separators between contents

**Task 3: Log Parser**
- Read a log file
- Count lines by log level (INFO, ERROR, WARNING)
- Print summary statistics

**Task 4: File Backup**
- Create a backup function that:
  - Reads a file
  - Creates a backup with `.bak` extension
  - Verifies backup was created

**Task 5: Line Numberer**
- Read a file
- Add line numbers to each line
- Write to new file with numbered lines

### Answers

**Quiz Answers:**
1. `io/ioutil` or `os` (os.ReadFile in newer Go)
2. Closes file when function returns
3. Open for reading, Create for writing (creates if not exists)
4. Read-write for owner, read for others

**Practice Solutions:**

**Task 1 Solution:**
```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func countWords(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    wordCount := 0
    uniqueWords := make(map[string]int)
    
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)
    
    for scanner.Scan() {
        word := strings.ToLower(strings.Trim(scanner.Text(), ".,!?;:\""))
        if word != "" {
            wordCount++
            uniqueWords[word]++
        }
    }
    
    fmt.Printf("Total words: %d\n", wordCount)
    fmt.Printf("Unique words: %d\n", len(uniqueWords))
    
    return scanner.Err()
}

func main() {
    err := countWords("text.txt")
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

**Task 2 Solution:**
```go
package main

import (
    "fmt"
    "io"
    "os"
)

func mergeFiles(filenames []string, output string) error {
    outFile, err := os.Create(output)
    if err != nil {
        return err
    }
    defer outFile.Close()
    
    for i, filename := range filenames {
        file, err := os.Open(filename)
        if err != nil {
            return fmt.Errorf("opening %s: %w", filename, err)
        }
        
        fmt.Fprintf(outFile, "=== File %d: %s ===\n", i+1, filename)
        _, err = io.Copy(outFile, file)
        file.Close()
        
        if err != nil {
            return fmt.Errorf("copying %s: %w", filename, err)
        }
        fmt.Fprintf(outFile, "\n")
    }
    
    return nil
}

func main() {
    files := []string{"file1.txt", "file2.txt", "file3.txt"}
    err := mergeFiles(files, "merged.txt")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Files merged successfully")
    }
}
```

**Task 3 Solution:**
```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func parseLogs(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    stats := map[string]int{
        "INFO":    0,
        "ERROR":   0,
        "WARNING": 0,
    }
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.ToUpper(scanner.Text())
        if strings.Contains(line, "[INFO]") {
            stats["INFO"]++
        } else if strings.Contains(line, "[ERROR]") {
            stats["ERROR"]++
        } else if strings.Contains(line, "[WARNING]") {
            stats["WARNING"]++
        }
    }
    
    fmt.Println("Log Statistics:")
    for level, count := range stats {
        fmt.Printf("  %s: %d\n", level, count)
    }
    
    return scanner.Err()
}

func main() {
    err := parseLogs("app.log")
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

**Task 4 Solution:**
```go
package main

import (
    "fmt"
    "io"
    "os"
)

func backupFile(filename string) error {
    source, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer source.Close()
    
    backupName := filename + ".bak"
    dest, err := os.Create(backupName)
    if err != nil {
        return err
    }
    defer dest.Close()
    
    _, err = io.Copy(dest, source)
    if err != nil {
        return err
    }
    
    // Verify backup
    if _, err := os.Stat(backupName); os.IsNotExist(err) {
        return fmt.Errorf("backup file was not created")
    }
    
    fmt.Printf("Backup created: %s\n", backupName)
    return nil
}

func main() {
    err := backupFile("important.txt")
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

**Task 5 Solution:**
```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func numberLines(input, output string) error {
    inFile, err := os.Open(input)
    if err != nil {
        return err
    }
    defer inFile.Close()
    
    outFile, err := os.Create(output)
    if err != nil {
        return err
    }
    defer outFile.Close()
    
    scanner := bufio.NewScanner(inFile)
    lineNum := 1
    
    for scanner.Scan() {
        fmt.Fprintf(outFile, "%d: %s\n", lineNum, scanner.Text())
        lineNum++
    }
    
    return scanner.Err()
}

func main() {
    err := numberLines("input.txt", "numbered.txt")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Lines numbered successfully")
    }
}
```

---

## Key Takeaways

1. **ioutil for simple operations** - ReadFile, WriteFile
2. **bufio for line-by-line** - Scanner
3. **Always defer Close()** - Ensures cleanup
4. **Check errors** - File operations can fail

---

## Next Steps

**Ready? → [16_JSON_HANDLING.md](./16_JSON_HANDLING.md)**
