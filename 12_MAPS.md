# Maps

## a) Overview

### What this topic is
Maps are key-value pairs - like dictionaries in Python. They let you store and retrieve values by key.

### Why it exists in Go
Maps provide efficient lookup, storage, and retrieval of key-value data. They're essential for many algorithms and data structures.

---

## b) Syntax

### Basic Map Operations
```go
// Create
var m map[string]int           // nil map
m := make(map[string]int)      // Empty map
m := map[string]int{           // With initial values
    "apple": 5,
    "banana": 3,
}

// Access
value := m["apple"]            // Get value
value, exists := m["apple"]    // Check existence

// Set
m["orange"] = 10

// Delete
delete(m, "apple")

// Iterate
for key, value := range m {
    // code
}

// Length
len(m)
```

---

## c) Explanation

### Step-by-Step Map Operations

**1. Create map**
```go
m := make(map[string]int)
```
- `make` creates empty map
- Type: `map[keyType]valueType`
- Zero value is `nil` (can't use until initialized)

**2. Set values**
```go
m["apple"] = 5
```
- Key: `"apple"` (string)
- Value: `5` (int)
- Creates entry if key doesn't exist
- Updates if key exists

**3. Get values**
```go
value := m["apple"]
```
- Returns value for key
- Returns zero value if key doesn't exist
- Can't tell if key exists or value is zero!

**4. Check existence**
```go
value, exists := m["apple"]
```
- `exists` is `true` if key exists
- `exists` is `false` if key doesn't exist
- Always use this pattern!

**5. Delete**
```go
delete(m, "apple")
```
- Removes key-value pair
- Safe to delete non-existent key (no error)

### Characteristics

- **Zero value**: `nil` (can't use until initialized)
- **Reference type**: Maps are references (like slices)
- **Keys must be comparable**: Numbers, strings, arrays, structs (not slices, maps, functions)
- **Unordered**: Iteration order is random
- **Dynamic**: Grows automatically

---

## d) Python Comparison

### Python Dictionary
```python
# Create
d = {}
d = {"apple": 5, "banana": 3}

# Access
value = d["apple"]              # Raises KeyError if missing
value = d.get("apple", 0)        # Returns default if missing
value = d.get("apple")           # Returns None if missing

# Set
d["orange"] = 10

# Delete
del d["apple"]
d.pop("apple", None)

# Check existence
if "apple" in d:
    value = d["apple"]

# Iterate
for key, value in d.items():
    print(key, value)
```

### Go Map
```go
// Create
m := make(map[string]int)
m := map[string]int{"apple": 5, "banana": 3}

// Access
value := m["apple"]             // Returns zero value if missing
value, exists := m["apple"]     // Check existence

// Set
m["orange"] = 10

// Delete
delete(m, "apple")

// Check existence
if value, exists := m["apple"]; exists {
    // use value
}

// Iterate
for key, value := range m {
    // code
}
```

**Key Differences:**

| Feature | Python | Go |
|---------|--------|-----|
| **Access missing key** | Raises KeyError | Returns zero value |
| **Check existence** | `key in dict` | `value, ok := m[key]` |
| **Delete** | `del dict[key]` | `delete(m, key)` |
| **Order** | Preserved (Python 3.7+) | Random (unordered) |
| **Zero value** | `{}` (empty dict) | `nil` (can't use) |

**Thinking Difference:**
- Python: Raises error for missing key
- Go: Returns zero value (must check existence)
- Python: Ordered (recent versions)
- Go: Unordered (random iteration)

---

## e) Visual Flow / Mental Model

### Map Structure

```
m := map[string]int{
    "apple": 5,
    "banana": 3,
}

Memory:
┌─────────┬────────┐
│  Key    │ Value  │
├─────────┼────────┤
│ "apple" │   5    │
│ "banana"│   3    │
└─────────┴────────┘
```

### Access Operation

```
value := m["apple"]
  ↓
1. Hash "apple" to get bucket
  ↓
2. Look up in bucket
  ↓
3. If found: return value (5)
  ↓
4. If not found: return zero value (0)
```

### Check Existence

```
value, exists := m["apple"]
  ↓
1. Look up "apple"
  ↓
2. If found:
     value = 5, exists = true
  ↓
3. If not found:
     value = 0, exists = false
```

---

## f) Demo Example

### Complete Example

```go
package main

import "fmt"

func main() {
    // 1. Create maps
    var m1 map[string]int
    fmt.Println("Nil map:", m1, "is nil:", m1 == nil)
    // m1["key"] = 1  // ❌ PANIC! Can't use nil map
    
    m2 := make(map[string]int)
    fmt.Println("Empty map:", m2)
    
    m3 := map[string]int{
        "apple":  5,
        "banana": 3,
        "cherry": 8,
    }
    fmt.Println("Map with values:", m3)
    
    // 2. Set values
    m2["orange"] = 10
    m2["grape"] = 7
    fmt.Println("After setting:", m2)
    
    // 3. Get values
    appleCount := m3["apple"]
    fmt.Println("Apple count:", appleCount)
    
    // 4. Get with existence check
    value, exists := m3["apple"]
    if exists {
        fmt.Printf("Apple exists: %d\n", value)
    } else {
        fmt.Println("Apple doesn't exist")
    }
    
    // 5. Missing key (returns zero value)
    missing := m3["mango"]
    fmt.Println("Missing key (zero value):", missing)
    
    missingValue, exists := m3["mango"]
    if !exists {
        fmt.Printf("Mango doesn't exist (value: %d)\n", missingValue)
    }
    
    // 6. Update value
    m3["apple"] = 10
    fmt.Println("After updating apple:", m3)
    
    // 7. Delete
    fmt.Println("Before delete:", m3)
    delete(m3, "banana")
    fmt.Println("After delete banana:", m3)
    
    delete(m3, "nonexistent")  // Safe, no error
    fmt.Println("After delete nonexistent:", m3)
    
    // 8. Iterate
    fmt.Println("\nIterating map:")
    for key, value := range m3 {
        fmt.Printf("  %s: %d\n", key, value)
    }
    
    // 9. Iterate keys only
    fmt.Println("\nKeys only:")
    for key := range m3 {
        fmt.Println("  Key:", key)
    }
    
    // 10. Iterate values only
    fmt.Println("\nValues only:")
    for _, value := range m3 {
        fmt.Println("  Value:", value)
    }
    
    // 11. Length
    fmt.Printf("\nMap length: %d\n", len(m3))
    
    // 12. Check if empty
    emptyMap := make(map[string]int)
    if len(emptyMap) == 0 {
        fmt.Println("Map is empty")
    }
    
    // 13. Map of maps
    students := map[string]map[string]int{
        "Alice": {
            "math":    90,
            "science": 85,
        },
        "Bob": {
            "math":    80,
            "science": 95,
        },
    }
    
    fmt.Println("\nNested map:")
    for name, grades := range students {
        fmt.Printf("  %s: %v\n", name, grades)
    }
    
    // 14. Map as set (using bool values)
    set := make(map[string]bool)
    set["item1"] = true
    set["item2"] = true
    set["item1"] = true  // Duplicate (no effect)
    
    if set["item1"] {
        fmt.Println("\nitem1 is in set")
    }
}
```

**Line-by-line explanation:**

1. **Nil map**: Can't use until initialized
2. **Empty map**: Created with `make`
3. **Initial values**: Can initialize with values
4. **Set values**: Add or update entries
5. **Get values**: Returns value or zero value
6. **Existence check**: Always use `value, ok` pattern
7. **Update**: Same as set (overwrites)
8. **Delete**: Removes key-value pair
9. **Iterate**: `for key, value := range map`
10. **Length**: Number of key-value pairs
11. **Nested maps**: Maps can contain maps
12. **Set simulation**: Use map with bool values

**Output:**
```
Nil map: map[] is nil: true
Empty map: map[]
Map with values: map[apple:5 banana:3 cherry:8]
After setting: map[grape:7 orange:10]
Apple count: 5
Apple exists: 5
Missing key (zero value): 0
Mango doesn't exist (value: 0)
After updating apple: map[apple:10 banana:3 cherry:8]
Before delete: map[apple:10 banana:3 cherry:8]
After delete banana: map[apple:10 cherry:8]
After delete nonexistent: map[apple:10 cherry:8]

Iterating map:
  apple: 10
  cherry: 8

Keys only:
  Key: apple
  Key: cherry

Values only:
  Value: 10
  Value: 8

Map length: 2
Map is empty

Nested map:
  Alice: map[math:90 science:85]
  Bob: map[math:80 science:95]

item1 is in set
```

---

## g) Use Cases

### When to use maps

**1. Key-value lookups**
```go
userScores := map[string]int{
    "alice": 100,
    "bob":   85,
}
```

**2. Counting**
```go
counts := make(map[string]int)
for _, item := range items {
    counts[item]++
}
```

**3. Grouping**
```go
groups := make(map[string][]string)
for _, item := range items {
    groups[item.Category] = append(groups[item.Category], item.Name)
}
```

**4. Caching**
```go
cache := make(map[string]Result)
if result, exists := cache[key]; exists {
    return result
}
```

**5. Sets (using bool)**
```go
seen := make(map[string]bool)
if seen[item] {
    // already seen
}
```

---

## h) Do's and Don'ts / Best Practices

### ✅ Do's

1. **Always check existence when zero value matters**
   ```go
   value, exists := m[key]
   if exists {
       // use value
   }
   ```

2. **Initialize with make**
   ```go
   m := make(map[string]int)  // ✅
   ```

3. **Use delete safely**
   ```go
   delete(m, key)  // ✅ Safe even if key doesn't exist
   ```

4. **Pre-allocate capacity if known**
   ```go
   m := make(map[string]int, 100)  // ✅ Pre-allocate
   ```

5. **Use maps for fast lookups**
   ```go
   // ✅ O(1) lookup
   if value, exists := m[key]; exists {
       // found
   }
   ```

### ❌ Don'ts

1. **Don't use nil map**
   ```go
   var m map[string]int
   m["key"] = 1  // ❌ PANIC!
   ```

2. **Don't assume zero value means missing**
   ```go
   value := m[key]  // ❌ Can't tell if key exists
   value, ok := m[key]  // ✅ Check existence
   ```

3. **Don't rely on iteration order**
   ```go
   // ❌ Order is random
   for key, value := range m {
       // process
   }
   ```

4. **Don't use slices/maps as keys**
   ```go
   // ❌ Can't use slice as key
   m := make(map[[]int]string)
   
   // ✅ Use comparable types
   m := make(map[string]int)
   ```

---

## i) Solved Practice Examples

### Example 1: Word Counter

**Task:** Count occurrences of each word in a slice.

**Solution:**
```go
package main

import "fmt"

func countWords(words []string) map[string]int {
    counts := make(map[string]int)
    for _, word := range words {
        counts[word]++
    }
    return counts
}

func main() {
    words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
    counts := countWords(words)
    fmt.Println("Word counts:", counts)
}
```

### Example 2: Group by Category

**Task:** Group items by their category.

**Solution:**
```go
package main

import "fmt"

type Item struct {
    Name     string
    Category string
}

func groupByCategory(items []Item) map[string][]string {
    groups := make(map[string][]string)
    for _, item := range items {
        groups[item.Category] = append(groups[item.Category], item.Name)
    }
    return groups
}

func main() {
    items := []Item{
        {"Apple", "Fruit"},
        {"Banana", "Fruit"},
        {"Carrot", "Vegetable"},
        {"Orange", "Fruit"},
        {"Broccoli", "Vegetable"},
    }
    
    groups := groupByCategory(items)
    for category, names := range groups {
        fmt.Printf("%s: %v\n", category, names)
    }
}
```

---

## j) Quiz / Practice for Me

### Quiz Questions

1. **What is the zero value of a map?**
   - [ ] Empty map
   - [ ] nil
   - [ ] {}
   - [ ] 0

2. **What happens when you access a non-existent key?**
   - [ ] Panic
   - [ ] Error
   - [ ] Returns zero value
   - [ ] Returns nil

3. **How do you check if a key exists?**
   - [ ] `key in map`
   - [ ] `value, ok := map[key]`
   - [ ] `map.hasKey(key)`
   - [ ] `map.contains(key)`

### Practice Tasks

**Task 1: Character Counter**
- Create function `countChars(s string) map[rune]int`
- Count occurrences of each character
- Return map of character to count

**Task 2: Remove Duplicates Using Map**
- Create function `unique(items []string) []string`
- Use map to track seen items
- Return slice with duplicates removed

### Answers

**Quiz Answers:**
1. nil
2. Returns zero value
3. `value, ok := map[key]`

**Practice Solutions:**

**Task 1 Solution:**
```go
package main

import "fmt"

func countChars(s string) map[rune]int {
    counts := make(map[rune]int)
    for _, char := range s {
        counts[char]++
    }
    return counts
}

func main() {
    text := "hello"
    counts := countChars(text)
    for char, count := range counts {
        fmt.Printf("%c: %d\n", char, count)
    }
}
```

**Task 2 Solution:**
```go
package main

import "fmt"

func unique(items []string) []string {
    seen := make(map[string]bool)
    result := []string{}
    
    for _, item := range items {
        if !seen[item] {
            seen[item] = true
            result = append(result, item)
        }
    }
    return result
}

func main() {
    items := []string{"apple", "banana", "apple", "cherry", "banana"}
    uniqueItems := unique(items)
    fmt.Println("Original:", items)
    fmt.Println("Unique:", uniqueItems)
}
```

---

## Key Takeaways

1. **Maps are key-value pairs** - Like Python dictionaries
2. **Zero value is nil** - Must initialize with `make`
3. **Check existence** - Always use `value, ok := m[key]`
4. **Unordered** - Iteration order is random
5. **Dynamic** - Grows automatically
6. **Keys must be comparable** - No slices/maps as keys

---

## Must Remember Forever

- `make(map[keyType]valueType)` - Create map
- `value, ok := m[key]` - Get with existence check
- `delete(m, key)` - Delete key-value pair
- `len(m)` - Number of key-value pairs
- Maps are unordered - don't rely on order
- Zero value is `nil` - can't use until initialized

---

---

## k) Additional Practice Problems (Build Strong Logic)

### 10 More Solved Examples

### Solved Problem 1: Frequency Counter

**Task:** Count frequency of each element in slice.

**Solution:**
```go
package main

import "fmt"

func frequencyCounter(arr []int) map[int]int {
    freq := make(map[int]int)
    for _, num := range arr {
        freq[num]++
    }
    return freq
}

func main() {
    arr := []int{1, 2, 3, 2, 1, 3, 3, 4, 5, 4}
    freq := frequencyCounter(arr)
    fmt.Printf("Frequencies: %v\n", freq)
}
```

### Solved Problem 2: Group Anagrams

**Task:** Group strings that are anagrams of each other.

**Solution:**
```go
package main

import (
    "fmt"
    "sort"
    "strings"
)

func groupAnagrams(strs []string) map[string][]string {
    groups := make(map[string][]string)
    
    for _, str := range strs {
        key := sortString(str)
        groups[key] = append(groups[key], str)
    }
    
    return groups
}

func sortString(s string) string {
    chars := strings.Split(s, "")
    sort.Strings(chars)
    return strings.Join(chars, "")
}

func main() {
    words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
    groups := groupAnagrams(words)
    fmt.Printf("Anagram groups: %v\n", groups)
}
```

### Solved Problem 3: Two Sum with Map

**Task:** Find two numbers that sum to target using map.

**Solution:**
```go
package main

import "fmt"

func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    
    for i, num := range nums {
        complement := target - num
        if idx, exists := seen[complement]; exists {
            return []int{idx, i}
        }
        seen[num] = i
    }
    return nil
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    result := twoSum(nums, target)
    fmt.Printf("Indices: %v\n", result)
}
```

### Solved Problem 4: First Non-Repeating Character

**Task:** Find first character that appears only once.

**Solution:**
```go
package main

import "fmt"

func firstNonRepeating(s string) rune {
    count := make(map[rune]int)
    
    for _, char := range s {
        count[char]++
    }
    
    for _, char := range s {
        if count[char] == 1 {
            return char
        }
    }
    return 0
}

func main() {
    str := "leetcode"
    result := firstNonRepeating(str)
    fmt.Printf("First non-repeating: %c\n", result)
}
```

### Solved Problem 5: Map Inversion

**Task:** Invert a map (swap keys and values).

**Solution:**
```go
package main

import "fmt"

func invertMap(m map[string]int) map[int]string {
    inverted := make(map[int]string)
    for key, value := range m {
        inverted[value] = key
    }
    return inverted
}

func main() {
    original := map[string]int{"a": 1, "b": 2, "c": 3}
    inverted := invertMap(original)
    fmt.Printf("Original: %v\n", original)
    fmt.Printf("Inverted: %v\n", inverted)
}
```

### Solved Problem 6: Map Merging

**Task:** Merge two maps, handling conflicts.

**Solution:**
```go
package main

import "fmt"

func mergeMaps(m1, m2 map[string]int) map[string]int {
    merged := make(map[string]int)
    
    for k, v := range m1 {
        merged[k] = v
    }
    
    for k, v := range m2 {
        if existing, exists := merged[k]; exists {
            merged[k] = existing + v // Sum on conflict
        } else {
            merged[k] = v
        }
    }
    return merged
}

func main() {
    m1 := map[string]int{"a": 1, "b": 2}
    m2 := map[string]int{"b": 3, "c": 4}
    merged := mergeMaps(m1, m2)
    fmt.Printf("Merged: %v\n", merged)
}
```

### Solved Problem 7: Substring with Map

**Task:** Find longest substring without repeating characters.

**Solution:**
```go
package main

import "fmt"

func longestSubstring(s string) int {
    charMap := make(map[byte]int)
    maxLen := 0
    start := 0
    
    for end := 0; end < len(s); end++ {
        if idx, exists := charMap[s[end]]; exists && idx >= start {
            start = idx + 1
        }
        charMap[s[end]] = end
        if end-start+1 > maxLen {
            maxLen = end - start + 1
        }
    }
    return maxLen
}

func main() {
    str := "abcabcbb"
    length := longestSubstring(str)
    fmt.Printf("Longest substring length: %d\n", length)
}
```

### Solved Problem 8: Map Filtering

**Task:** Filter map based on condition.

**Solution:**
```go
package main

import "fmt"

func filterMap(m map[string]int, minValue int) map[string]int {
    filtered := make(map[string]int)
    for k, v := range m {
        if v >= minValue {
            filtered[k] = v
        }
    }
    return filtered
}

func main() {
    scores := map[string]int{"Alice": 85, "Bob": 60, "Charlie": 95, "David": 45}
    passed := filterMap(scores, 70)
    fmt.Printf("Passed (>=70): %v\n", passed)
}
```

### Solved Problem 9: Map Sorting by Value

**Task:** Get keys sorted by their values.

**Solution:**
```go
package main

import (
    "fmt"
    "sort"
)

type Pair struct {
    Key   string
    Value int
}

func sortMapByValue(m map[string]int) []string {
    pairs := make([]Pair, 0, len(m))
    for k, v := range m {
        pairs = append(pairs, Pair{k, v})
    }
    
    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i].Value > pairs[j].Value
    })
    
    result := make([]string, len(pairs))
    for i, p := range pairs {
        result[i] = p.Key
    }
    return result
}

func main() {
    scores := map[string]int{"Alice": 85, "Bob": 60, "Charlie": 95}
    sorted := sortMapByValue(scores)
    fmt.Printf("Sorted by value: %v\n", sorted)
}
```

### Solved Problem 10: Nested Map Operations

**Task:** Work with nested maps (map of maps).

**Solution:**
```go
package main

import "fmt"

type NestedMap map[string]map[string]int

func NewNestedMap() NestedMap {
    return make(NestedMap)
}

func (nm NestedMap) Set(category, item string, value int) {
    if nm[category] == nil {
        nm[category] = make(map[string]int)
    }
    nm[category][item] = value
}

func (nm NestedMap) Get(category, item string) (int, bool) {
    if cat, exists := nm[category]; exists {
        if val, exists := cat[item]; exists {
            return val, true
        }
    }
    return 0, false
}

func main() {
    nm := NewNestedMap()
    nm.Set("fruits", "apple", 5)
    nm.Set("fruits", "banana", 3)
    nm.Set("vegetables", "carrot", 10)
    
    val, _ := nm.Get("fruits", "apple")
    fmt.Printf("Apples: %d\n", val)
}
```

---

### 10 More Practice Problems (Solve These!)

**Problem 1: Word Pattern Matching**
- Check if string follows pattern: "abba" matches "dog cat cat dog"

**Problem 2: Isomorphic Strings**
- Check if two strings are isomorphic (character mapping)

**Problem 3: Valid Anagram**
- Check if two strings are anagrams using map

**Problem 4: Group Shifted Strings**
- Group strings that can be shifted to match each other

**Problem 5: Top K Frequent Elements**
- Find k most frequent elements in array

**Problem 6: Design HashMap**
- Implement basic hash map operations: Put, Get, Remove

**Problem 7: LRU Cache**
- Implement Least Recently Used cache using map

**Problem 8: Word Dictionary**
- Create dictionary with add, search, delete operations

**Problem 9: Map with Default Value**
- Create map wrapper that returns default value for missing keys

**Problem 10: Concurrent Safe Map**
- Create thread-safe map using channels or mutex

---

## Next Steps

Now you understand maps. Next:
- Concurrency
- Goroutines and channels
- Go's powerful concurrency features

**Ready? → [13_CONCURRENCY.md](./13_CONCURRENCY.md)**
