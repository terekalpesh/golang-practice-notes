package main

// ARRAYS AND SLICES

/*
-	Arrays are fixed-size collections, and slices are dynamic arrays. Slices are what you'll use 99% of the time in Go.
-	Arrays provide fixed-size collections, but slices are more flexible and are the primary way to work with sequences of data in Go.
*/

// SYNTAX

// ARRAYS
/*
var arr [5]int
arr := [5]int{1, 2, 3, 4, 5}
arr := [...]int{1, 2, 3}	// Size inferred

arr[0] = 10	// Access/modify
value := arr[0]
len(arr)	// Length
*/


// SLICES
/*
var s []int	// nil slice
s := []int{1,  2, 3}
s := make([]int, 5)
s := make([]int, 0, 10)	// Length 0, capacity 10

s = append(s, 4)	// Add element
s = append(s, 5, 6, 7)	// Add multiple

s[1:3] 	// Slice from index 1 to 3
len(s)	// Length
cap(s)	// Capacity
*/


import "fmt"

func main() {

	// 1. Arrays (fix size)

	var arr1 [5]int
	fmt.Println("Empty array: ", arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Initialized array: ", arr2)

	arr3 := [...]int{1, 2, 3}	// Size inferred
	fmt.Println("Inferred size:", arr3)
	fmt.Println("Length:", len(arr3))



	// 2. Slices (dynamic)

	var s1 []int
	fmt.Println("Nil slice:", s1, "is nil:", s1 == nil)

	s2 := make([]int, 5)	// Length 5, capacity 5
	fmt.Println("Make slice:", s2, "len:", len(s2), "cap", cap(s2))

	s3 := make([]int, 0, 10)	// Length 0, capacity 10
	fmt.Println("Empty with capacity:", s3, "len:", len(s3), "cap:", cap(s3))

	s4 := []int{1, 2, 3}
	fmt.Println("Slice literal:", s4)



	// 3. Append

	s5 := []int{1, 2, 3}
	fmt.Println("Before append:", s5)

	s5 = append(s5, 4)
	fmt.Println("After append:", s5)

	s5 = append(s5, 5, 6, 7)
	fmt.Println("After append multiple:", s5, "len:", len(s5), "cap:", cap(s5))



	// 4. Slicing

	original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("original:", original)

	slice1 := original[1:4]	// Index 1 to 4 (exlusive)
	fmt.Println("original[1:4]:", slice1)

	slice2 := original[:5]	// Start to index 5
	fmt.Println("original[:5]:", slice2)

	slice3 := original[5:]	// Index 5 to end
	fmt.Println("original[5:]:", slice3)

	slice4 := original[:]	// Copy of entire slice
	fmt.Println("original[:]:", slice4)



	// 5. Modifying slices (they share underlying array!)

	fmt.Println("\n---Modifying shared slice---")
	shared := []int{1, 2, 3, 4, 5}
	part := shared[1:4]
	fmt.Println("shared:", shared)
	fmt.Println("part:", part)

	part[0] = 99	// Modify part
	fmt.Println("\nAfter modifying part[0] = 99:")
	fmt.Println("shared:", shared)
	fmt.Println("part:", part)



	// 6. Copy (to avoid sharing)
	fmt.Println("\n---Copying slice---")
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	copy(dst, src)

	dst[0] = 99
	fmt.Println("src:", src)	// Unchanged
	fmt.Println("dst:", dst)	// Changed



	// 7. Range loop
	fmt.Println("\n---Range loop---")
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range(numbers) {
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
	fmt.Printf("len=%d, cap=%d, slice=%d\n", len(growable), cap(growable), growable)
	
}