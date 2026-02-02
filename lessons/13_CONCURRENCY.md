# Concurrency: Goroutines and Channels

## a) Overview

### What this topic is
Concurrency in Go - using goroutines (lightweight threads) and channels (communication between goroutines) to run multiple tasks simultaneously.

### Why it exists in Go
Go was designed for concurrency. Goroutines are cheap (can have millions), and channels provide safe communication. This is Go's killer feature!

### 🎯 Layman's Explanation (Simple Terms)

**Think of concurrency like a restaurant with multiple waiters:**

**Real-world analogy - Restaurant:**
- **One waiter (no concurrency)**: Takes order from Table 1, goes to kitchen, waits, comes back, then serves Table 2
  - Slow! Tables wait a long time
- **Multiple waiters (concurrency)**: Waiter 1 serves Table 1, Waiter 2 serves Table 2, Waiter 3 serves Table 3 - all at the same time!
  - Fast! All tables get served simultaneously

**Goroutines = Workers:**
- Like having **multiple employees** working at the same time
- Like **multitasking** - doing multiple things at once
- Very cheap to create (like hiring is free and easy!)

**Channels = Communication:**
- Like a **mailbox** or **message box** between workers
- Worker 1 puts a message in the box
- Worker 2 takes the message from the box
- Safe way to share information without conflicts

**Simple example:**
```
Without concurrency (one at a time):
Task 1: Download file (takes 5 seconds) → Wait
Task 2: Process data (takes 3 seconds) → Wait  
Task 3: Send email (takes 2 seconds) → Wait
Total: 10 seconds ❌

With concurrency (all at once):
Task 1: Download file (5 seconds) ┐
Task 2: Process data (3 seconds)  ├─ All running at same time!
Task 3: Send email (2 seconds)    ┘
Total: 5 seconds (longest task) ✅
```

**Goroutine = A lightweight worker:**
- Like a **helper** that can do a task for you
- Very cheap to create (can have millions!)
- Like having unlimited free workers

**Channel = A safe mailbox:**
- Like a **pipe** or **tube** to send messages
- One end puts messages in, other end takes them out
- Prevents conflicts (like a queue at a store)

**Why use concurrency?**
1. **Speed**: Do multiple things at once (like cooking while the rice is boiling)
2. **Efficiency**: Don't waste time waiting (like doing dishes while laundry runs)
3. **Responsiveness**: Program stays fast even when doing heavy work

**Key concepts:**
- **Goroutine**: A task running in the background (like a worker)
- **Channel**: A safe way to send data between goroutines (like a mailbox)
- **Select**: Choose which channel to listen to (like choosing which phone to answer)

---

## b) Syntax

### Goroutines
```go
// Start goroutine
go functionName()

// Anonymous function
go func() {
    // code
}()

// With parameters
go func(x int) {
    // code
}(10)
```

### Channels
```go
// Create
ch := make(chan int)           // Unbuffered
ch := make(chan int, 10)        // Buffered (capacity 10)

// Send
ch <- value

// Receive
value := <-ch
value, ok := <-ch              // Check if closed

// Close
close(ch)

// Range over channel
for value := range ch {
    // code
}

// Select (like switch for channels)
select {
case msg := <-ch1:
    // received from ch1
case ch2 <- value:
    // sent to ch2
case <-time.After(1 * time.Second):
    // timeout
default:
    // non-blocking
}
```

---

## c) Explanation

### Step-by-Step Concurrency

**1. Goroutine**
```go
go doSomething()
```
- Starts function in new goroutine
- Program continues immediately (doesn't wait)
- Goroutine runs concurrently

**2. Channel creation**
```go
ch := make(chan int)
```
- Creates communication channel
- Unbuffered: sender waits for receiver
- Buffered: can hold values (up to capacity)

**3. Send/receive**
```go
ch <- 42        // Send value
value := <-ch   // Receive value
```
- Sender blocks until receiver ready (unbuffered)
- Receiver blocks until sender ready
- Synchronizes goroutines

**4. Select statement**
```go
select {
case msg := <-ch:
    // handle message
}
```
- Like switch, but for channels
- Can wait on multiple channels
- Non-blocking with `default`

### Characteristics

- **Goroutines**: Lightweight (2KB stack), cheap to create
- **Channels**: Type-safe communication
- **No shared memory**: Communicate by sharing (channels), not sharing memory
- **Built-in**: Part of language, not library
- **Scalable**: Can have millions of goroutines

---

## d) Python Comparison

### Python Threading
```python
import threading
import time

def task():
    time.sleep(1)
    print("Task done")

# Create thread
thread = threading.Thread(target=task)
thread.start()
thread.join()  # Wait for completion

# Thread pool
from concurrent.futures import ThreadPoolExecutor
with ThreadPoolExecutor(max_workers=5) as executor:
    executor.submit(task)
```

### Go Goroutines
```go
func task() {
    time.Sleep(1 * time.Second)
    fmt.Println("Task done")
}

// Start goroutine
go task()

// Wait for completion
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    task()
}()
wg.Wait()
```

**Key Differences:**

| Feature | Python | Go |
|---------|--------|-----|
| **Unit** | Thread | Goroutine |
| **Cost** | Heavy (~1MB) | Light (~2KB) |
| **Max concurrent** | Limited (~1000s) | Millions |
| **Communication** | Queues, locks | Channels (built-in) |
| **GIL** | Yes (limits parallelism) | No (true parallelism) |
| **Complexity** | High | Low |

**Thinking Difference:**
- Python: Threads are expensive, use sparingly
- Go: Goroutines are cheap, use liberally
- Python: Shared memory with locks
- Go: Channels for communication (safer)

---

## e) Visual Flow / Mental Model

### Goroutine Execution

```
Main goroutine:
1. Start: go task()
   ↓
2. Continue immediately (doesn't wait)
   ↓
3. Main and task() run concurrently
   ↓
4. Both execute simultaneously

Goroutine:
1. Start executing task()
   ↓
2. Run independently
   ↓
3. Complete when done
```

### Channel Communication

```
Goroutine 1          Channel          Goroutine 2
    |                  |                  |
    |-- value ----->   |                  |
    |                  |                  |
    |                  |   <-- value ----|
    |                  |                  |
```

**Unbuffered channel:**
- Sender blocks until receiver ready
- Receiver blocks until sender ready
- Synchronizes goroutines

**Buffered channel:**
- Sender only blocks if buffer full
- Receiver only blocks if buffer empty
- Asynchronous (up to capacity)

---

## f) Demo Example

### Complete Example

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// 1. Simple goroutine
func sayHello() {
    fmt.Println("Hello from goroutine!")
}

// 2. Goroutine with WaitGroup
func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

// 3. Channel communication
func sender(ch chan<- string) {
    ch <- "Hello"
    ch <- "World"
    close(ch)
}

func receiver(ch <-chan string) {
    for msg := range ch {
        fmt.Println("Received:", msg)
    }
}

// 4. Buffered channel
func bufferedExample() {
    ch := make(chan int, 3)
    
    ch <- 1
    ch <- 2
    ch <- 3
    
    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}

// 5. Select statement
func selectExample() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "Message from ch1"
    }()
    
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "Message from ch2"
    }()
    
    select {
    case msg := <-ch1:
        fmt.Println("Received from ch1:", msg)
    case msg := <-ch2:
        fmt.Println("Received from ch2:", msg)
    case <-time.After(3 * time.Second):
        fmt.Println("Timeout")
    }
}

// 6. Producer-Consumer pattern
func producer(ch chan<- int) {
    for i := 0; i < 5; i++ {
        ch <- i
        fmt.Printf("Sent: %d\n", i)
    }
    close(ch)
}

func consumer(ch <-chan int) {
    for value := range ch {
        fmt.Printf("Received: %d\n", value)
        time.Sleep(100 * time.Millisecond)
    }
}

// 7. Worker pool
func workerPool(jobs <-chan int, results chan<- int) {
    for job := range jobs {
        result := job * 2
        results <- result
    }
}

func main() {
    fmt.Println("=== 1. Simple Goroutine ===")
    go sayHello()
    time.Sleep(100 * time.Millisecond)  // Wait for goroutine
    
    fmt.Println("\n=== 2. WaitGroup ===")
    var wg sync.WaitGroup
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }
    wg.Wait()
    
    fmt.Println("\n=== 3. Channel Communication ===")
    ch := make(chan string)
    go sender(ch)
    receiver(ch)
    
    fmt.Println("\n=== 4. Buffered Channel ===")
    bufferedExample()
    
    fmt.Println("\n=== 5. Select Statement ===")
    selectExample()
    
    fmt.Println("\n=== 6. Producer-Consumer ===")
    dataCh := make(chan int)
    go producer(dataCh)
    consumer(dataCh)
    
    fmt.Println("\n=== 7. Worker Pool ===")
    jobs := make(chan int, 10)
    results := make(chan int, 10)
    
    // Start 3 workers
    for w := 1; w <= 3; w++ {
        go workerPool(jobs, results)
    }
    
    // Send jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Collect results
    for r := 1; r <= 5; r++ {
        fmt.Printf("Result: %d\n", <-results)
    }
}
```

**Line-by-line explanation:**

1. **Simple goroutine**: Just add `go` keyword
2. **WaitGroup**: Wait for multiple goroutines to finish
3. **Channels**: Send/receive between goroutines
4. **Buffered channels**: Can hold multiple values
5. **Select**: Wait on multiple channels
6. **Producer-Consumer**: Classic concurrency pattern
7. **Worker pool**: Multiple workers process jobs

**Output:**
```
=== 1. Simple Goroutine ===
Hello from goroutine!

=== 2. WaitGroup ===
Worker 1 starting
Worker 2 starting
Worker 3 starting
Worker 1 done
Worker 2 done
Worker 3 done

=== 3. Channel Communication ===
Received: Hello
Received: World

=== 4. Buffered Channel ===
1
2
3

=== 5. Select Statement ===
Received from ch1: Message from ch1

=== 6. Producer-Consumer ===
Sent: 0
Received: 0
Sent: 1
Received: 1
Sent: 2
Received: 2
Sent: 3
Received: 3
Sent: 4
Received: 4

=== 7. Worker Pool ===
Result: 2
Result: 4
Result: 6
Result: 8
Result: 10
```

---

## g) Use Cases

### When to use goroutines

**1. I/O operations**
```go
go fetchURL(url)  // Don't block on network
```

**2. Parallel processing**
```go
for _, item := range items {
    go process(item)  // Process in parallel
}
```

**3. Background tasks**
```go
go cleanup()  // Run in background
```

### When to use channels

**1. Communication between goroutines**
```go
ch <- result  // Send result
```

**2. Synchronization**
```go
done := make(chan bool)
go func() {
    // work
    done <- true
}()
<-done  // Wait for completion
```

**3. Pipeline processing**
```go
stage1 := make(chan int)
stage2 := make(chan int)
// Connect stages with channels
```

---

## h) Do's and Don'ts / Best Practices

### ✅ Do's

1. **Use WaitGroup to wait for goroutines**
   ```go
   var wg sync.WaitGroup
   wg.Add(1)
   go func() {
       defer wg.Done()
       // work
   }()
   wg.Wait()
   ```

2. **Close channels when done sending**
   ```go
   close(ch)  // ✅ Signal no more values
   ```

3. **Use buffered channels when appropriate**
   ```go
   ch := make(chan int, 10)  // ✅ If you know capacity
   ```

4. **Check channel closure**
   ```go
   value, ok := <-ch
   if !ok {
       // channel closed
   }
   ```

5. **Use select for timeouts**
   ```go
   select {
   case msg := <-ch:
       // got message
   case <-time.After(1 * time.Second):
       // timeout
   }
   ```

### ❌ Don'ts

1. **Don't forget to wait for goroutines**
   ```go
   go task()  // ❌ Program might exit before completion
   ```

2. **Don't send on closed channel**
   ```go
   close(ch)
   ch <- value  // ❌ PANIC!
   ```

3. **Don't create goroutine leaks**
   ```go
   // ❌ Goroutine blocked forever
   go func() {
       <-ch  // Blocked if no sender
   }()
   ```

4. **Don't share memory, share by communicating**
   ```go
   // ❌ Shared memory with mutex
   var counter int
   var mu sync.Mutex
   
   // ✅ Use channel
   ch := make(chan int)
   ```

---

## i) Solved Practice Examples

### Example 1: Parallel Sum

**Task:** Sum numbers in parallel using goroutines.

**Solution:**
```go
package main

import (
    "fmt"
    "sync"
)

func sumPart(numbers []int, ch chan<- int) {
    total := 0
    for _, num := range numbers {
        total += num
    }
    ch <- total
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    ch := make(chan int, 2)
    
    // Split work
    go sumPart(numbers[:5], ch)
    go sumPart(numbers[5:], ch)
    
    // Collect results
    sum1 := <-ch
    sum2 := <-ch
    
    fmt.Printf("Sum: %d\n", sum1+sum2)
}
```

### Example 2: Rate Limiter

**Task:** Process items with rate limiting using channels.

**Solution:**
```go
package main

import (
    "fmt"
    "time"
)

func processWithRateLimit(items []string, limit int) {
    limiter := make(chan struct{}, limit)
    
    for _, item := range items {
        limiter <- struct{}{}  // Acquire
        
        go func(i string) {
            defer func() { <-limiter }()  // Release
            fmt.Printf("Processing: %s\n", i)
            time.Sleep(100 * time.Millisecond)
        }(item)
    }
}

func main() {
    items := []string{"item1", "item2", "item3", "item4", "item5"}
    processWithRateLimit(items, 2)  // Max 2 concurrent
    time.Sleep(1 * time.Second)
}
```

---

## j) Quiz / Practice for Me

### Quiz Questions

1. **What is a goroutine?**
   - [ ] A function
   - [ ] A lightweight thread
   - [ ] A channel
   - [ ] A mutex

2. **What happens if you send on a closed channel?**
   - [ ] Nothing
   - [ ] Panic
   - [ ] Returns error
   - [ ] Blocks forever

3. **How do you wait for multiple goroutines?**
   - [ ] Use channel
   - [ ] Use WaitGroup
   - [ ] Use mutex
   - [ ] Use select

### Practice Tasks

**Task 1: Simple Goroutine**
- Create function that prints numbers 1-5
- Run it in a goroutine
- Wait for it to complete

**Task 2: Channel Communication**
- Create producer that sends numbers 1-10
- Create consumer that receives and prints
- Use channels for communication

### Answers

**Quiz Answers:**
1. A lightweight thread
2. Panic
3. Use WaitGroup

**Practice Solutions:**

**Task 1 Solution:**
```go
package main

import (
    "fmt"
    "sync"
)

func printNumbers(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go printNumbers(&wg)
    wg.Wait()
}
```

**Task 2 Solution:**
```go
package main

import "fmt"

func producer(ch chan<- int) {
    for i := 1; i <= 10; i++ {
        ch <- i
    }
    close(ch)
}

func consumer(ch <-chan int) {
    for value := range ch {
        fmt.Println("Received:", value)
    }
}

func main() {
    ch := make(chan int)
    go producer(ch)
    consumer(ch)
}
```

---

## Key Takeaways

1. **Goroutines are cheap** - Can have millions
2. **Channels for communication** - Share by communicating
3. **WaitGroup for synchronization** - Wait for goroutines
4. **Select for multiple channels** - Like switch for channels
5. **Buffered vs unbuffered** - Choose based on needs
6. **Close channels** - Signal completion

---

## Must Remember Forever

- `go function()` - Start goroutine
- `ch := make(chan Type)` - Create channel
- `ch <- value` - Send
- `value := <-ch` - Receive
- `close(ch)` - Close channel
- `select { case ... }` - Wait on multiple channels
- Use WaitGroup to wait for goroutines

---

---

## k) Additional Practice Problems (Build Strong Logic)

### 10 More Solved Examples

### Solved Problem 1: Parallel File Processing

**Task:** Process multiple files concurrently.

**Solution:**
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func processFile(filename string, wg *sync.WaitGroup, results chan<- string) {
    defer wg.Done()
    // Simulate file processing
    time.Sleep(100 * time.Millisecond)
    results <- fmt.Sprintf("Processed %s", filename)
}

func main() {
    files := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt"}
    results := make(chan string, len(files))
    var wg sync.WaitGroup
    
    for _, file := range files {
        wg.Add(1)
        go processFile(file, &wg, results)
    }
    
    wg.Wait()
    close(results)
    
    for result := range results {
        fmt.Println(result)
    }
}
```

### Solved Problem 2: Worker Pool Pattern

**Task:** Implement worker pool to process jobs concurrently.

**Solution:**
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(100 * time.Millisecond)
        results <- job * 2
    }
}

func main() {
    jobs := make(chan int, 10)
    results := make(chan int, 10)
    var wg sync.WaitGroup
    
    // Start 3 workers
    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }
    
    // Send jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
    
    wg.Wait()
    close(results)
    
    for result := range results {
        fmt.Printf("Result: %d\n", result)
    }
}
```

### Solved Problem 3: Fan-Out Fan-In Pattern

**Task:** Distribute work to multiple workers and collect results.

**Solution:**
```go
package main

import (
    "fmt"
    "sync"
)

func producer(numbers []int, out chan<- int) {
    defer close(out)
    for _, num := range numbers {
        out <- num
    }
}

func worker(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for num := range in {
        out <- num * num
    }
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    producerChan := make(chan int)
    resultChan := make(chan int)
    var wg sync.WaitGroup
    
    go producer(numbers, producerChan)
    
    // Fan-out: Multiple workers
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go worker(producerChan, resultChan, &wg)
    }
    
    go func() {
        wg.Wait()
        close(resultChan)
    }()
    
    // Fan-in: Collect results
    for result := range resultChan {
        fmt.Printf("Squared: %d\n", result)
    }
}
```

### Solved Problem 4: Rate Limiter with Channels

**Task:** Implement rate limiter using channels.

**Solution:**
```go
package main

import (
    "fmt"
    "time"
)

func rateLimiter(limit int, interval time.Duration) chan struct{} {
    ticker := time.NewTicker(interval)
    limiter := make(chan struct{}, limit)
    
    go func() {
        for range ticker.C {
            for i := 0; i < limit; i++ {
                select {
                case limiter <- struct{}{}:
                default:
                }
            }
        }
    }()
    
    return limiter
}

func main() {
    limiter := rateLimiter(5, 1*time.Second)
    
    for i := 0; i < 20; i++ {
        <-limiter
        fmt.Printf("Request %d processed\n", i+1)
    }
}
```

### Solved Problem 5: Pipeline Pattern

**Task:** Create processing pipeline with multiple stages.

**Solution:**
```go
package main

import (
    "fmt"
    "strings"
)

func stage1(in <-chan string, out chan<- string) {
    defer close(out)
    for s := range in {
        out <- strings.ToUpper(s)
    }
}

func stage2(in <-chan string, out chan<- string) {
    defer close(out)
    for s := range in {
        out <- strings.ReplaceAll(s, " ", "-")
    }
}

func stage3(in <-chan string, out chan<- string) {
    defer close(out)
    for s := range in {
        out <- "[" + s + "]"
    }
}

func main() {
    input := make(chan string)
    stage1Out := make(chan string)
    stage2Out := make(chan string)
    output := make(chan string)
    
    go stage1(input, stage1Out)
    go stage2(stage1Out, stage2Out)
    go stage3(stage2Out, output)
    
    go func() {
        input <- "hello world"
        input <- "go programming"
        close(input)
    }()
    
    for result := range output {
        fmt.Println(result)
    }
}
```

### Solved Problem 6: Timeout Pattern

**Task:** Implement operation with timeout using select.

**Solution:**
```go
package main

import (
    "fmt"
    "time"
)

func longOperation(duration time.Duration) <-chan string {
    result := make(chan string)
    go func() {
        time.Sleep(duration)
        result <- "Operation completed"
    }()
    return result
}

func main() {
    timeout := 2 * time.Second
    operation := longOperation(3 * time.Second)
    
    select {
    case result := <-operation:
        fmt.Println(result)
    case <-time.After(timeout):
        fmt.Println("Operation timed out")
    }
}
```

### Solved Problem 7: Concurrent Counter

**Task:** Implement thread-safe counter using channels.

**Solution:**
```go
package main

import (
    "fmt"
    "sync"
)

type SafeCounter struct {
    value int
    inc   chan int
    get   chan int
    done  chan struct{}
}

func NewSafeCounter() *SafeCounter {
    c := &SafeCounter{
        inc:  make(chan int),
        get:  make(chan int),
        done: make(chan struct{}),
    }
    
    go c.run()
    return c
}

func (c *SafeCounter) run() {
    for {
        select {
        case amount := <-c.inc:
            c.value += amount
        case c.get <- c.value:
        case <-c.done:
            return
        }
    }
}

func (c *SafeCounter) Increment() {
    c.inc <- 1
}

func (c *SafeCounter) Get() int {
    return <-c.get
}

func (c *SafeCounter) Close() {
    close(c.done)
}

func main() {
    counter := NewSafeCounter()
    var wg sync.WaitGroup
    
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }
    
    wg.Wait()
    fmt.Printf("Final count: %d\n", counter.Get())
    counter.Close()
}
```

### Solved Problem 8: Concurrent Map Writer

**Task:** Write to map concurrently using channels.

**Solution:**
```go
package main

import (
    "fmt"
    "sync"
)

type ConcurrentMap struct {
    data map[string]int
    ops  chan func()
    done chan struct{}
}

func NewConcurrentMap() *ConcurrentMap {
    cm := &ConcurrentMap{
        data: make(map[string]int),
        ops:  make(chan func()),
        done: make(chan struct{}),
    }
    go cm.run()
    return cm
}

func (cm *ConcurrentMap) run() {
    for {
        select {
        case op := <-cm.ops:
            op()
        case <-cm.done:
            return
        }
    }
}

func (cm *ConcurrentMap) Set(key string, value int) {
    cm.ops <- func() {
        cm.data[key] = value
    }
}

func (cm *ConcurrentMap) Get(key string) int {
    result := make(chan int)
    cm.ops <- func() {
        result <- cm.data[key]
    }
    return <-result
}

func (cm *ConcurrentMap) Close() {
    close(cm.done)
}

func main() {
    cm := NewConcurrentMap()
    var wg sync.WaitGroup
    
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            cm.Set(fmt.Sprintf("key%d", id), id*10)
        }(i)
    }
    
    wg.Wait()
    fmt.Printf("Value: %d\n", cm.Get("key5"))
    cm.Close()
}
```

### Solved Problem 9: Semaphore Pattern

**Task:** Implement semaphore using buffered channel.

**Solution:**
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Semaphore struct {
    ch chan struct{}
}

func NewSemaphore(capacity int) *Semaphore {
    return &Semaphore{
        ch: make(chan struct{}, capacity),
    }
}

func (s *Semaphore) Acquire() {
    s.ch <- struct{}{}
}

func (s *Semaphore) Release() {
    <-s.ch
}

func worker(id int, sem *Semaphore, wg *sync.WaitGroup) {
    defer wg.Done()
    sem.Acquire()
    defer sem.Release()
    
    fmt.Printf("Worker %d started\n", id)
    time.Sleep(1 * time.Second)
    fmt.Printf("Worker %d finished\n", id)
}

func main() {
    sem := NewSemaphore(3) // Allow 3 concurrent workers
    var wg sync.WaitGroup
    
    for i := 1; i <= 10; i++ {
        wg.Add(1)
        go worker(i, sem, &wg)
    }
    
    wg.Wait()
}
```

### Solved Problem 10: Context Cancellation

**Task:** Use context to cancel multiple goroutines.

**Solution:**
```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %d cancelled\n", id)
            return
        default:
            fmt.Printf("Worker %d working...\n", id)
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    var wg sync.WaitGroup
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(ctx, i, &wg)
    }
    
    wg.Wait()
    fmt.Println("All workers stopped")
}
```

---

### 10 More Practice Problems (Solve These!)

**Problem 1: Concurrent Web Scraper**
- Scrape multiple URLs concurrently with rate limiting
- Collect all results into single slice

**Problem 2: Producer-Consumer with Multiple Consumers**
- One producer, multiple consumers
- Ensure all items are processed

**Problem 3: Concurrent Cache**
- Implement cache with Get, Set, Delete
- Use channels for thread-safety

**Problem 4: Parallel Merge Sort**
- Implement merge sort using goroutines
- Merge results from parallel sorts

**Problem 5: Task Scheduler**
- Schedule tasks with priorities
- Execute high priority tasks first

**Problem 6: Concurrent File Downloader**
- Download multiple files concurrently
- Show progress for each download

**Problem 7: Pub-Sub Pattern**
- Implement publisher-subscriber pattern
- Multiple subscribers, one publisher

**Problem 8: Barrier Synchronization**
- Wait for all goroutines to reach barrier before continuing

**Problem 9: Concurrent Queue**
- Implement thread-safe queue using channels
- Support multiple producers and consumers

**Problem 10: Graceful Shutdown**
- Implement graceful shutdown for multiple goroutines
- Clean up resources properly

---

## Next Steps

Now you understand concurrency. Next:
- Packages and Modules
- Code organization
- Dependency management

**Ready? → [14_PACKAGES_AND_MODULES.md](./14_PACKAGES_AND_MODULES.md)**
