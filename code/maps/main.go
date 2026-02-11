package main

// MAPS
/*
-	Maps are key-value pairs. They let you store and retrieve values by key.
*/

// BASIC MAP OPERATION
/*
// Create
var m map[string]int	// nil map
m := make(map[string]int)	// Empty map
m := map[string]int{	// With initial values
	"apple": 5,
	"banana": 3
}


// 	Access
value := m["apple"]	// Get value
value, exist := m["apple"]	// Check existence


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
*/

import "fmt"

func main() {
	// // 1. Create maps
	// var m1 map[string]int
	// fmt.Println("Nil map:", m1, "is nill:", m1 == nil)
	// // m1["key"] = 1	// ❌ PANIC! Can't use nil map

	// m2 := make(map[string]int)
	// fmt.Println("Empty map:", m2, "is nill:", m2 == nil)

	m3 := map[string]int {
		"apple": 5,
		"banana": 3,
		"cherry": 8,
	}
	fmt.Println("Map with values:", m3)


	// // 2. Set values
	// m2["orange"] = 10
	// m2["grape"] = 7
	// fmt.Println("After setting m2:", m2)


	// // 3. Get values
	// appleCount := m3["apple"]
	// fmt.Println("Apple count:", appleCount)


	// // 4. Get with existence check
	// value, exists := m3["apple"]
	// if exists {
	// 	fmt.Printf("Apple exists: %d\n", value)
	// } else {
	// 	fmt.Println("Apple doesn't exist")
	// }


	// // 5. Missing key (return zero value)
	// missing := m3["mango"]
	// fmt.Println("Missing key (zero value):", missing)

	// missingValue, exists := m3["mango"]
	// if !exists {
	// 	fmt.Printf("Mango doesn't exist (value: %d)\n", missingValue)
	// }


	// // 6. Update value
	// m3["apple"] = 10
	// fmt.Println("After updating apple:", m3)


	// // 7. Delete
	// fmt.Println("Before delete:", m3)
	// delete(m3, "banana")
	// fmt.Println("After delete banana:", m3)

	// delete(m3, "nonexistent")	// Safe, no error
	// fmt.Println("After deleting nonexistent:", m3)


	// // 8. Iterate
	// fmt.Println("\nIterating map:")
	// for key, value := range m3 {
	// 	fmt.Printf("%s: %d\n", key, value)
	// }


	// // 9. Iterate keys only
	// fmt.Println("\nKeys only:")
	// for key, _ := range m3 {
	// 	fmt.Println("Key:", key)
	// }


	// 10. Iterate values only
	fmt.Println("\nValues only:")
	for _, value := range m3 {
		fmt.Println("Value:", value)
	}


	// 11. Length
	fmt.Printf("\nMap length: %d\n", len(m3))


	// 12. Check if empty
	emptyMap := make(map[string]int)

	if len(emptyMap) == 0 {
		fmt.Println("Map is empty")
	}


	// Map of maps
	students := map[string]map[string]int{
		"Alice" : {
			"math": 90,
			"science": 85,
		},
		"Bob" : {
			"math": 80,
			"science": 95,
		},
	}

	fmt.Println("\nNested map:")
	for name, grades := range students {
		fmt.Printf("%s: %v\n", name, grades)
	}


	// 14. Map as set (using bool values)
	set := make(map[string]bool)
	set["item1"] = true
	set["item2"] = true
	set["item1"] = true	// Duplicate (no effect)

	if set["item1"] {
		fmt.Println("\nitem is in set")
		fmt.Println("Set:", set)
	}
}