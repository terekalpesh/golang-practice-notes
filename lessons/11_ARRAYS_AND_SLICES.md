# Arrays and Slices

## a) Overview

### What this topic is
Arrays are fixed-size collections, and slices are dynamic arrays. Slices are what you'll use 99% of the time in Go.

### Why it exists in Go
Arrays provide fixed-size collections, but slices are more flexible and are the primary way to work with sequences of data in Go.

---

## b) Syntax

### Arrays
```go
var arr [5]int                    // [0, 0, 0, 0, 0]
arr := [5]int{1, 2, 3, 4, 5}      // [1, 2, 3, 4, 5]
arr := [...]int{1, 2, 3}         // Size inferred: [1, 2, 3]

arr[0] = 10                       // Access/modify
value := arr[0]
len(arr)                          // Length
```

### Slices
```go
var s []int                       // nil slice
s := []int{1, 2, 3}               // [1, 2, 3]
s := make([]int, 5)                // [0, 0, 0, 0, 0]
s := make([]int, 0, 10)           // Length 0, capacity 10

s = append(s, 4)                   // Add element
s = append(s, 5, 6, 7)            // Add multiple

s[1:3]                            // Slice from index 1 to 3
len(s)                            // Length
cap(s)                            // Capacity
```

---

## c) Explanation

### Step-by-Step Understanding

**1. Arrays (fixed size)**
```go
var arr [5]int
```
- Size is part of type: `[5]int` is different from `[10]int`
- Size cannot change
- Rarely used directly (slices are preferred)

**2. Slices (dynamic size)**
```go
s := []int{1, 2, 3}
```
- No size in type (just `[]int`)
- Can grow and shrink
- Built on top of arrays
- This is what you'll use!

**3. Slice internals**
```
Slice = [Pointer | Length | Capacity]
         ↓
    Underlying Array
```
- Pointer: Points to underlying array
- Length: Number of elements
- Capacity: Total space available

**4. Append operation**
```go
s = append(s, 4)
```
- Adds element to end
- Grows slice if needed
- May create new underlying array if capacity exceeded

### Characteristics

- **Arrays**: Fixed size, value type (copied)
- **Slices**: Dynamic size, reference type (shared)
- **Zero value**: Array = zeroed elements, Slice = `nil`
- **Indexing**: Both use `[index]`
- **Range**: Both work with `for range`

---

## d) Python Comparison

### Python Lists
```python
# Python list (dynamic)
items = [1, 2, 3]
items.append(4)
items.extend([5, 6])
items[0] = 10

# Slicing
items[1:3]  # [2, 3]
items[:3]   # [1, 2, 3]
items[2:]   # [3, 4, 5, 6]
```

### Go Slices
```go
// Go slice (dynamic)
items := []int{1, 2, 3}
items = append(items, 4)
items = append(items, 5, 6)
items[0] = 10

// Slicing
items[1:3]  // [2, 3]
items[:3]   // [1, 2, 3]
items[2:]   // [3, 4, 5, 6]
```

**Key Differences:**

| Feature | Python | Go |
|---------|--------|-----|
| **Type** | List (always dynamic) | Array (fixed) or Slice (dynamic) |
| **Append** | `list.append(item)` | `slice = append(slice, item)` |
| **Slicing** | `list[start:end]` | `slice[start:end]` (similar) |
| **Copy behavior** | Reference (shared) | Reference (shared for slices) |
| **Size** | Always dynamic | Array fixed, slice dynamic |

**Thinking Difference:**
- Python: Only lists (always dynamic)
- Go: Arrays (fixed) and Slices (dynamic)
- Python: `append()` modifies in place
- Go: `append()` returns new slice (may need to reassign)

---

## e) Visual Flow / Mental Model

### Array in Memory

```
arr := [5]int{1, 2, 3, 4, 5}

Memory:
Index:  0   1   2   3   4
Value:  1   2   3   4   5
```

### Slice in Memory

```
s := []int{1, 2, 3, 4, 5}

Slice structure:
┌─────────┬────────┬──────────┐
│ Pointer │ Length │ Capacity │
│ 0x1000  │   5    │    5     │
└─────────┴────────┴──────────┘
     ↓
Underlying Array (at 0x1000):
[1, 2, 3, 4, 5]
```

### Append Operation

```
s := []int{1, 2, 3}  // len=3, cap=3
s = append(s, 4)

If capacity available:
  Add 4 to existing array
  s = [1, 2, 3, 4]  // len=4, cap=3

If capacity exceeded:
  Create new array (usually 2x capacity)
  Copy old elements
  Add new element
  s = [1, 2, 3, 4]  // len=4, cap=6 (or more)
```

### Slice Slicing

```
original := []int{1, 2, 3, 4, 5}
slice := original[1:3]  // [2, 3]

Memory:
original: [1, 2, 3, 4, 5]
            ↑     ↑
            |     |
slice:      [2, 3]
         (shares same underlying array!)
```

**Important**: Slicing creates new slice but shares underlying array!

---

## f) Demo Example

### Complete Example

```go
package main

import "fmt"

func main() {
    // 1. Arrays (fixed size)
    var arr1 [5]int
    fmt.Println("Empty array:", arr1)
    
    arr2 := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Initialized array:", arr2)
    
    arr3 := [...]int{1, 2, 3}  // Size inferred
    fmt.Println("Inferred size:", arr3)
    fmt.Println("Length:", len(arr3))
    
    // 2. Slices (dynamic)
    var s1 []int
    fmt.Println("Nil slice:", s1, "is nil:", s1 == nil)
    
    s2 := []int{1, 2, 3}
    fmt.Println("Slice literal:", s2)
    
    s3 := make([]int, 5)  // Length 5, capacity 5
    fmt.Println("Make slice:", s3, "len:", len(s3), "cap:", cap(s3))
    
    s4 := make([]int, 0, 10)  // Length 0, capacity 10
    fmt.Println("Empty with capacity:", s4, "len:", len(s4), "cap:", cap(s4))
    
    // 3. Append
    s5 := []int{1, 2, 3}
    fmt.Println("Before append:", s5)
    s5 = append(s5, 4)
    fmt.Println("After append 4:", s5)
    s5 = append(s5, 5, 6, 7)
    fmt.Println("After append multiple:", s5)
    
    // 4. Slicing
    original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Println("Original:", original)
    
    slice1 := original[1:4]  // Index 1 to 4 (exclusive)
    fmt.Println("original[1:4]:", slice1)
    
    slice2 := original[:5]  // Start to index 5
    fmt.Println("original[:5]:", slice2)
    
    slice3 := original[5:]  // Index 5 to end
    fmt.Println("original[5:]:", slice3)
    
    slice4 := original[:]  // Copy of entire slice
    fmt.Println("original[:]:", slice4)
    
    // 5. Modifying slices (they share underlying array!)
    fmt.Println("\n--- Modifying shared slice ---")
    shared := []int{1, 2, 3, 4, 5}
    part := shared[1:4]  // [2, 3, 4]
    fmt.Println("shared:", shared)
    fmt.Println("part:", part)
    
    part[0] = 99  // Modify part
    fmt.Println("After modifying part[0] = 99:")
    fmt.Println("shared:", shared)  // Also changed!
    fmt.Println("part:", part)
    
    // 6. Copy (to avoid sharing)
    fmt.Println("\n--- Copying slice ---")
    src := []int{1, 2, 3, 4, 5}
    dst := make([]int, len(src))
    copy(dst, src)
    
    dst[0] = 99
    fmt.Println("src:", src)  // Unchanged
    fmt.Println("dst:", dst)  // Changed
    
    // 7. Range loop
    fmt.Println("\n--- Range loop ---")
    numbers := []int{10, 20, 30, 40, 50}
    for index, value := range numbers {
        fmt.Printf("Index %d: %d\n", index, value)
    }
    
    // Ignore index
    for _, value := range numbers {
        fmt.Printf("Value: %d\n", value)
    }
    
    // 8. Length and capacity
    fmt.Println("\n--- Length and Capacity ---")
    growable := make([]int, 3, 10)
    fmt.Printf("len=%d, cap=%d, slice=%v\n", len(growable), cap(growable), growable)
    
    growable = append(growable, 1, 2, 3, 4, 5, 6, 7, 8)
    fmt.Printf("After append 8 elements:\n")
    fmt.Printf("len=%d, cap=%d, slice=%v\n", len(growable), cap(growable), growable)
}
```

**Line-by-line explanation:**

1. **Arrays**: Fixed size, rarely used
2. **Slices**: Dynamic, what you'll use
3. **Append**: Adds elements, may grow capacity
4. **Slicing**: Creates new slice sharing underlying array
5. **Shared memory**: Modifying slice affects original (if sharing)
6. **Copy**: Creates independent copy
7. **Range**: Iterate over elements
8. **Capacity**: Pre-allocated space for growth

**Output:**
```
Empty array: [0 0 0 0 0]
Initialized array: [1 2 3 4 5]
Inferred size: [1 2 3]
Length: 3
Nil slice: [] is nil: true
Slice literal: [1 2 3]
Make slice: [0 0 0 0 0] len: 5 cap: 5
Empty with capacity: [] len: 0 cap: 10
Before append: [1 2 3]
After append 4: [1 2 3 4]
After append multiple: [1 2 3 4 5 6 7]
Original: [1 2 3 4 5 6 7 8 9 10]
original[1:4]: [2 3 4]
original[:5]: [1 2 3 4 5]
original[5:]: [6 7 8 9 10]
original[:]: [1 2 3 4 5 6 7 8 9 10]

--- Modifying shared slice ---
shared: [1 2 3 4 5]
part: [2 3 4]
After modifying part[0] = 99:
shared: [1 99 3 4 5]
part: [99 3 4]

--- Copying slice ---
src: [1 2 3 4 5]
dst: [99 2 3 4 5]

--- Range loop ---
Index 0: 10
Index 1: 20
Index 2: 30
Index 3: 40
Index 4: 50
Value: 10
Value: 20
Value: 30
Value: 40
Value: 50

--- Length and Capacity ---
len=3, cap=10, slice=[0 0 0]
After append 8 elements:
len=11, cap=12, slice=[0 0 0 1 2 3 4 5 6 7 8]
```

---

## g) Use Cases

### When to use arrays

**1. Fixed-size collections**
```go
var buffer [1024]byte  // Fixed size buffer
```

**2. When size is part of type**
```go
type Point [3]float64  // 3D point
```

### When to use slices

**1. Dynamic collections (99% of cases)**
```go
items := []string{"apple", "banana"}
items = append(items, "cherry")
```

**2. Working with variable data**
```go
func processItems(items []int) {
    // Works with any size
}
```

**3. Building collections**
```go
var results []Result
for item := range items {
    results = append(results, process(item))
}
```

---

## h) Do's and Don'ts / Best Practices

### ✅ Do's

1. **Use slices, not arrays**
   ```go
   items := []int{1, 2, 3}  // ✅ Slice
   ```

2. **Pre-allocate capacity when you know size**
   ```go
   items := make([]int, 0, 100)  // ✅ Pre-allocate
   ```

3. **Check slice before indexing**
   ```go
   if len(s) > 0 {
       first := s[0]  // ✅ Safe
   }
   ```

4. **Use copy when you need independent slice**
   ```go
   dst := make([]int, len(src))
   copy(dst, src)  // ✅ Independent copy
   ```

5. **Use range for iteration**
   ```go
   for i, v := range items {  // ✅
       // process
   }
   ```

### ❌ Don'ts

1. **Don't use arrays unless you really need fixed size**
   ```go
   var arr [5]int  // ❌ Usually not needed
   slice := []int{}  // ✅ Use slice
   ```

2. **Don't forget to reassign with append**
   ```go
   append(s, 4)  // ❌ Doesn't modify s!
   s = append(s, 4)  // ✅ Correct
   ```

3. **Don't assume slice is independent after slicing**
   ```go
   part := original[1:3]
   part[0] = 99  // ❌ Also modifies original!
   ```

4. **Don't index without checking length**
   ```go
   value := s[0]  // ❌ Can panic if empty
   if len(s) > 0 {
       value := s[0]  // ✅ Safe
   }
   ```

---

## i) Solved Practice Examples

### Example 1: Sum Function

**Task:** Create function that sums all numbers in a slice.

**Solution:**
```go
package main

import "fmt"

func sum(numbers []int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println("Sum:", sum(nums))
}
```

### Example 2: Filter Function

**Task:** Create function that filters even numbers from slice.

**Solution:**
```go
package main

import "fmt"

func filterEven(numbers []int) []int {
    result := []int{}
    for _, num := range numbers {
        if num%2 == 0 {
            result = append(result, num)
        }
    }
    return result
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    evens := filterEven(nums)
    fmt.Println("Original:", nums)
    fmt.Println("Evens:", evens)
}
```

### Example 3: Reverse Slice

**Task:** Create function that reverses a slice in place.

**Solution:**
```go
package main

import "fmt"

func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println("Before:", nums)
    reverse(nums)
    fmt.Println("After:", nums)
}
```

---

## j) Quiz / Practice for Me

### Quiz Questions

1. **What is the difference between array and slice?**
   - [ ] No difference
   - [ ] Array is fixed size, slice is dynamic
   - [ ] Array is dynamic, slice is fixed
   - [ ] Both are the same

2. **What does append return?**
   - [ ] Nothing (modifies in place)
   - [ ] New slice
   - [ ] Boolean
   - [ ] Error

3. **Do slices share underlying array when sliced?**
   - [ ] No, always independent
   - [ ] Yes, they share
   - [ ] Sometimes
   - [ ] Never

### Practice Tasks

**Task 1: Find Maximum**
- Create function `findMax(numbers []int) int`
- Return maximum value in slice
- Handle empty slice (return 0)

**Task 2: Remove Duplicates**
- Create function `removeDuplicates(numbers []int) []int`
- Return new slice with duplicates removed
- Preserve order

### Answers

**Quiz Answers:**
1. Array is fixed size, slice is dynamic
2. New slice
3. Yes, they share

**Practice Solutions:**

**Task 1 Solution:**
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
    fmt.Println("Max:", findMax(nums))
}
```

**Task 2 Solution:**
```go
package main

import "fmt"

func removeDuplicates(numbers []int) []int {
    seen := make(map[int]bool)
    result := []int{}
    
    for _, num := range numbers {
        if !seen[num] {
            seen[num] = true
            result = append(result, num)
        }
    }
    return result
}

func main() {
    nums := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
    unique := removeDuplicates(nums)
    fmt.Println("Original:", nums)
    fmt.Println("Unique:", unique)
}
```

---

## Key Takeaways

1. **Arrays are fixed** - Size is part of type, rarely used
2. **Slices are dynamic** - What you'll use 99% of the time
3. **Append returns new slice** - Must reassign: `s = append(s, item)`
4. **Slicing shares memory** - Be careful when modifying
5. **Use copy for independence** - When you need separate slice
6. **Range for iteration** - Clean way to iterate

---

## Must Remember Forever

- `[5]int` = array (fixed size)
- `[]int` = slice (dynamic)
- `append(slice, item)` returns new slice
- Slicing shares underlying array
- `copy(dst, src)` for independent copy
- `len(slice)` and `cap(slice)` for size info

---

---

## k) Additional Practice Problems (Build Strong Logic)

### 10 More Solved Examples

### Solved Problem 1: Rotate Array

**Task:** Rotate array to the right by k positions.

**Solution:**
```go
package main

import "fmt"

func rotateRight(arr []int, k int) []int {
    if len(arr) == 0 {
        return arr
    }
    k = k % len(arr)
    result := make([]int, len(arr))
    for i := 0; i < len(arr); i++ {
        result[(i+k)%len(arr)] = arr[i]
    }
    return result
}

func main() {
    arr := []int{1, 2, 3, 4, 5}
    rotated := rotateRight(arr, 2)
    fmt.Printf("Original: %v\n", arr)
    fmt.Printf("Rotated: %v\n", rotated)
}
```

### Solved Problem 2: Merge Sorted Arrays

**Task:** Merge two sorted arrays into one sorted array.

**Solution:**
```go
package main

import "fmt"

func mergeSorted(arr1, arr2 []int) []int {
    result := make([]int, 0, len(arr1)+len(arr2))
    i, j := 0, 0
    
    for i < len(arr1) && j < len(arr2) {
        if arr1[i] <= arr2[j] {
            result = append(result, arr1[i])
            i++
        } else {
            result = append(result, arr2[j])
            j++
        }
    }
    
    result = append(result, arr1[i:]...)
    result = append(result, arr2[j:]...)
    return result
}

func main() {
    arr1 := []int{1, 3, 5, 7}
    arr2 := []int{2, 4, 6, 8}
    merged := mergeSorted(arr1, arr2)
    fmt.Printf("Merged: %v\n", merged)
}
```

### Solved Problem 3: Find All Duplicates

**Task:** Find all duplicate elements in an array.

**Solution:**
```go
package main

import "fmt"

func findDuplicates(arr []int) []int {
    seen := make(map[int]int)
    duplicates := []int{}
    
    for _, num := range arr {
        seen[num]++
    }
    
    for num, count := range seen {
        if count > 1 {
            duplicates = append(duplicates, num)
        }
    }
    return duplicates
}

func main() {
    arr := []int{1, 2, 3, 2, 4, 5, 3, 6, 3}
    dupes := findDuplicates(arr)
    fmt.Printf("Duplicates: %v\n", dupes)
}
```

### Solved Problem 4: Slice Partition

**Task:** Partition slice around a pivot value.

**Solution:**
```go
package main

import "fmt"

func partition(arr []int, pivot int) ([]int, []int) {
    smaller := []int{}
    larger := []int{}
    
    for _, num := range arr {
        if num < pivot {
            smaller = append(smaller, num)
        } else {
            larger = append(larger, num)
        }
    }
    return smaller, larger
}

func main() {
    arr := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
    smaller, larger := partition(arr, 5)
    fmt.Printf("Smaller than 5: %v\n", smaller)
    fmt.Printf("Larger or equal: %v\n", larger)
}
```

### Solved Problem 5: Slice Intersection

**Task:** Find common elements between two slices.

**Solution:**
```go
package main

import "fmt"

func intersection(arr1, arr2 []int) []int {
    set := make(map[int]bool)
    result := []int{}
    
    for _, num := range arr1 {
        set[num] = true
    }
    
    for _, num := range arr2 {
        if set[num] {
            result = append(result, num)
            set[num] = false // Avoid duplicates
        }
    }
    return result
}

func main() {
    arr1 := []int{1, 2, 3, 4, 5}
    arr2 := []int{3, 4, 5, 6, 7}
    common := intersection(arr1, arr2)
    fmt.Printf("Common elements: %v\n", common)
}
```

### Solved Problem 6: Slice Difference

**Task:** Find elements in first slice but not in second.

**Solution:**
```go
package main

import "fmt"

func difference(arr1, arr2 []int) []int {
    set := make(map[int]bool)
    result := []int{}
    
    for _, num := range arr2 {
        set[num] = true
    }
    
    for _, num := range arr1 {
        if !set[num] {
            result = append(result, num)
        }
    }
    return result
}

func main() {
    arr1 := []int{1, 2, 3, 4, 5}
    arr2 := []int{3, 4, 5}
    diff := difference(arr1, arr2)
    fmt.Printf("Elements only in arr1: %v\n", diff)
}
```

### Solved Problem 7: Slice Chunking

**Task:** Split slice into chunks of specified size.

**Solution:**
```go
package main

import "fmt"

func chunkSlice(arr []int, size int) [][]int {
    chunks := [][]int{}
    for i := 0; i < len(arr); i += size {
        end := i + size
        if end > len(arr) {
            end = len(arr)
        }
        chunks = append(chunks, arr[i:end])
    }
    return chunks
}

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    chunks := chunkSlice(arr, 3)
    fmt.Printf("Chunks: %v\n", chunks)
}
```

### Solved Problem 8: Slice Flattening

**Task:** Flatten a 2D slice into 1D.

**Solution:**
```go
package main

import "fmt"

func flatten(arr [][]int) []int {
    result := []int{}
    for _, subarr := range arr {
        result = append(result, subarr...)
    }
    return result
}

func main() {
    arr := [][]int{{1, 2, 3}, {4, 5}, {6, 7, 8, 9}}
    flattened := flatten(arr)
    fmt.Printf("Flattened: %v\n", flattened)
}
```

### Solved Problem 9: Slice Shuffling

**Task:** Randomly shuffle elements in a slice.

**Solution:**
```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func shuffle(arr []int) []int {
    result := make([]int, len(arr))
    copy(result, arr)
    
    rand.Seed(time.Now().UnixNano())
    for i := len(result) - 1; i > 0; i-- {
        j := rand.Intn(i + 1)
        result[i], result[j] = result[j], result[i]
    }
    return result
}

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    shuffled := shuffle(arr)
    fmt.Printf("Shuffled: %v\n", shuffled)
}
```

### Solved Problem 10: Slice Window

**Task:** Get all subarrays of specified size (sliding window).

**Solution:**
```go
package main

import "fmt"

func slidingWindow(arr []int, size int) [][]int {
    if size > len(arr) {
        return [][]int{}
    }
    
    windows := [][]int{}
    for i := 0; i <= len(arr)-size; i++ {
        windows = append(windows, arr[i:i+size])
    }
    return windows
}

func main() {
    arr := []int{1, 2, 3, 4, 5, 6}
    windows := slidingWindow(arr, 3)
    fmt.Printf("Windows of size 3: %v\n", windows)
}
```

---

### 10 More Practice Problems (Solve These!)

**Problem 1: Two Sum**
- Find two numbers in array that add up to target
- Return indices of the two numbers

**Problem 2: Three Sum**
- Find all unique triplets that sum to zero

**Problem 3: Maximum Subarray**
- Find contiguous subarray with largest sum (Kadane's algorithm)

**Problem 4: Product of Array Except Self**
- Return array where each element is product of all other elements

**Problem 5: Rotate Matrix**
- Rotate 2D matrix 90 degrees clockwise in-place

**Problem 6: Spiral Matrix**
- Return elements of matrix in spiral order

**Problem 7: Merge Intervals**
- Merge overlapping intervals in slice of [start, end] pairs

**Problem 8: Next Permutation**
- Find next lexicographically greater permutation

**Problem 9: Container With Most Water**
- Find two lines that together form container with most water

**Problem 10: Longest Increasing Subsequence**
- Find length of longest increasing subsequence

---

## Next Steps

Now you understand arrays and slices. Next:
- Maps
- Key-value pairs
- Working with dictionaries

**Ready? → [12_MAPS.md](./12_MAPS.md)**
