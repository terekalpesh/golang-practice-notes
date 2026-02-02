# Installation and Setup

## a) Overview

### What this topic is
Setting up Go on your computer so you can write and run Go programs.

### Why it exists in Go
Unlike Python (which is usually pre-installed), Go needs to be installed separately. The setup is simple but important to get right.

---

## b) Syntax

### Installation Commands

**Linux/Mac:**
```bash
# Download and install
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz

# Add to PATH
export PATH=$PATH:/usr/local/go/bin
```

**Or use package manager:**
```bash
# Ubuntu/Debian
sudo apt update
sudo apt install golang-go

# Mac (with Homebrew)
brew install go
```

**Verify installation:**
```bash
go version
```

---

## c) Explanation

### Step-by-Step Setup

1. **Download Go**
   - Go to https://go.dev/dl/
   - Download the version for your operating system
   - Or use package manager (easier)

2. **Install Go**
   - Extract files to `/usr/local/go` (Linux/Mac)
   - Or run installer (Windows)

3. **Set PATH**
   - Add Go's `bin` directory to your PATH
   - This lets you run `go` command from anywhere

4. **Verify**
   - Run `go version`
   - Should show something like: `go version go1.21.5 linux/amd64`

5. **Set GOPATH (optional, Go 1.11+)**
   - Modern Go uses modules, so GOPATH is less important
   - But it's good to know where Go stores things

### Go Environment Variables

**GOROOT:**
- Where Go is installed
- Usually `/usr/local/go`
- Don't change this

**GOPATH:**
- Where your Go code lives
- Default: `~/go` (home directory + go)
- Contains: `src/`, `bin/`, `pkg/`

**GOBIN:**
- Where compiled binaries go
- Usually `$GOPATH/bin`

**Check your environment:**
```bash
go env
```

---

## d) Python Comparison

### Python Setup
```bash
# Python is often pre-installed
python3 --version

# Or install via package manager
sudo apt install python3

# Virtual environments (common practice)
python3 -m venv myenv
source myenv/bin/activate
```

### Go Setup
```bash
# Go must be installed
go version

# No virtual environments needed
# Go modules handle dependencies
go mod init myproject
```

**Key Differences:**
- **Python**: Often pre-installed, uses virtual environments
- **Go**: Must install, uses modules (built-in)
- **Python**: Multiple versions can coexist
- **Go**: One version at a time (but easy to switch)

---

## e) Visual Flow / Mental Model

### Go Workspace Structure

```
Your Computer
│
├── /usr/local/go          (GOROOT - Go installation)
│   ├── bin/               (Go tools)
│   ├── src/               (Go source code)
│   └── ...
│
└── ~/go                   (GOPATH - Your workspace, optional)
    ├── bin/               (Your compiled programs)
    ├── pkg/               (Package cache)
    └── src/               (Your source code - old way)
        └── github.com/
            └── username/
                └── project/
```

### Modern Go (Modules - Recommended)

```
Your Project Directory
│
├── go.mod                 (Dependencies file)
├── go.sum                 (Checksums for security)
├── main.go                (Your code)
└── other files...
```

**Flow:**
1. Create project directory
2. Run `go mod init projectname`
3. Write code
4. Run `go run main.go` or `go build`

---

## f) Demo Example

### Complete Setup and First Program

**Step 1: Verify Go is installed**
```bash
go version
```

**Step 2: Create project directory**
```bash
mkdir my-first-go-program
cd my-first-go-program
```

**Step 3: Initialize Go module**
```bash
go mod init my-first-go-program
```

This creates `go.mod` file:
```go
module my-first-go-program

go 1.21
```

**Step 4: Create main.go**
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

**Step 5: Run the program**
```bash
# Option 1: Run directly
go run main.go

# Option 2: Build then run
go build main.go
./main
```

**Line-by-line explanation:**
```go
package main  // This file is part of the "main" package
             // "main" package creates an executable program

import "fmt"  // Import the "fmt" package
             // "fmt" is for formatting and printing (like Python's print)

func main() { // The main function - where program starts
             // Similar to Python's if __name__ == "__main__"
    fmt.Println("Hello, Go!")  // Print "Hello, Go!" to screen
}
```

---

## g) Use Cases

### When to use different commands

**`go run`**
- Quick testing
- Development
- Run without creating executable

**`go build`**
- Create executable file
- Production deployment
- Share compiled program

**`go install`**
- Install program to GOBIN
- Make command available system-wide
- Like `pip install` but for Go programs

**`go mod init`**
- Start new project
- Enable dependency management
- Required for modern Go projects

**`go get`**
- Download and add dependency
- Updates go.mod
- Like `pip install package`

---

## h) Do's and Don'ts / Best Practices

### ✅ Do's

1. **Use Go modules** (Go 1.11+)
   - Don't use GOPATH for new projects
   - Modules are the modern way

2. **Keep Go updated**
   ```bash
   # Check for updates
   go version
   # Download latest from go.dev/dl
   ```

3. **Use `go fmt`**
   - Automatically formats your code
   - Run: `go fmt ./...`
   - Or configure editor to format on save

4. **Use `go vet`**
   - Checks for common mistakes
   - Run: `go vet ./...`

5. **Set up your editor**
   - VS Code: Install Go extension
   - Vim: Use vim-go plugin
   - GoLand: Full IDE for Go

### ❌ Don'ts

1. **Don't ignore compiler errors**
   - Go compiler is strict for good reason
   - Fix errors, don't work around them

2. **Don't use old GOPATH style**
   - Use modules instead
   - GOPATH is legacy

3. **Don't mix Go versions**
   - One project = one Go version
   - Specify in go.mod

4. **Don't skip `go mod tidy`**
   - Cleans up dependencies
   - Removes unused, adds missing

---

## i) Solved Practice Examples

### Example 1: Check Your Setup

**Task:** Verify Go is properly installed and configured.

**Solution:**
```bash
# Check Go version
go version

# Check Go environment
go env

# Check if go is in PATH
which go

# Expected output locations:
# Linux/Mac: /usr/local/go/bin/go
# Or: /usr/bin/go (if installed via package manager)
```

### Example 2: Create Your First Module

**Task:** Create a Go module called "calculator" and write a simple program.

**Solution:**
```bash
# Create directory
mkdir calculator
cd calculator

# Initialize module
go mod init calculator

# Create main.go
cat > main.go << 'EOF'
package main

import "fmt"

func main() {
    result := 10 + 5
    fmt.Printf("10 + 5 = %d\n", result)
}
EOF

# Run it
go run main.go
```

**Output:**
```
10 + 5 = 15
```

### Example 3: Install a Go Tool

**Task:** Install a popular Go tool (like `golint`) and use it.

**Solution:**
```bash
# Install golint
go install golang.org/x/lint/golint@latest

# Check if installed
which golint

# Use it on your code
golint ./...
```

---

## j) Quiz / Practice for Me

### Quiz Questions

1. **What command checks your Go version?**
   - [ ] `go --version`
   - [ ] `go version`
   - [ ] `go check`
   - [ ] `go info`

2. **What does `go mod init` do?**
   - [ ] Installs Go
   - [ ] Creates a new Go module
   - [ ] Runs your program
   - [ ] Formats your code

3. **What's the difference between `go run` and `go build`?**
   - [ ] No difference
   - [ ] `go run` compiles and runs, `go build` only compiles
   - [ ] `go build` is faster
   - [ ] `go run` creates an executable

4. **Where does Go store installed tools?**
   - [ ] GOROOT
   - [ ] GOPATH/bin or GOBIN
   - [ ] /usr/bin
   - [ ] Nowhere, tools are global

### Practice Tasks

**Task 1: Setup Verification**
- Install Go (if not already installed)
- Verify installation with `go version`
- Check your Go environment with `go env`
- Note your GOROOT and GOPATH values

**Task 2: Create Hello World**
- Create a directory called `hello-world`
- Initialize a Go module
- Write a program that prints "Hello, World!"
- Run it with `go run`
- Build it with `go build`
- Run the compiled executable

**Task 3: Explore Go Tools**
- Run `go help` to see all commands
- Try `go doc fmt.Println` to see documentation
- Try `go fmt` on your hello-world program
- Try `go vet` on your hello-world program

### Answers

**Quiz Answers:**
1. `go version`
2. Creates a new Go module
3. `go run` compiles and runs, `go build` only compiles
4. GOPATH/bin or GOBIN

**Practice Task Solutions:**

**Task 2 Solution:**
```bash
mkdir hello-world
cd hello-world
go mod init hello-world

# Create main.go
cat > main.go << 'EOF'
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
EOF

go run main.go
go build main.go
./main
```

---

## Key Takeaways

1. **Go must be installed** - Unlike Python, it's not usually pre-installed
2. **Use Go modules** - Modern way to manage projects (Go 1.11+)
3. **Simple setup** - Just install and start coding
4. **Go tools are helpful** - `go fmt`, `go vet`, `go doc` are your friends
5. **One executable** - `go build` creates a single file you can share

---

## Next Steps

Now that Go is installed, let's learn:
- Basic syntax
- Package structure
- Import statements
- Your first real Go program

**Ready? Let's code! → [02_BASIC_SYNTAX.md](./02_BASIC_SYNTAX.md)**
