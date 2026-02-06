package main

// IF/ELSE

// if condition {
// 	// code
// }

// if condition {
// 	// code
// } else {
// 	// code
// }

// if condition {
// 	// code
// } else if condition {
// 	// code
// } else {
// 	// code
// }

// If with initialization
// if x := getValue; x > 0 {
// 	// x is available here
// }


// FOR LOOPS

// Traditional for loop
// for i := 0; i < 10; i++ {
// 	// code
// }

// While-style loop
// for condition {
// 	// code
// }

// Infinite loop
// for {
// 	// code
// 	break	// Exit loop
// }

// Range loop (for slices, maps, etc.)
// for index, value := range slice {
// 	// code
// }


// SWITCH

// switch value {
// case option1:
// 	// code
// case option2:
// 	// code
// default:
// 	// code
// }

// switch with no value (like if/else chain)
// switch {
// case condition1:
// 	// code
// case condition2:
// 	// code
// }


import "fmt"

func main() {
	// 1. Basic if
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	// 2. If/else
	if x%2 == 0 {
		fmt.Println("x is even")
	} else {
		fmt.Println("x is odd")
	}

	// 3. If/else if/else
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >=  80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// 4. If with initialization
	if y := 20; y > 10 {
		fmt.Printf("y (%d) is greater than 10\n", y)
		// y is only available in this if block
	}

	// ***************************************************
	// 5. Traditonal for loop
	fmt.Println("Counting 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 6. While-style loop (using for)
	fmt.Println("Countdown: ")
	count := 5
	for count > 0 {
		fmt.Println(count)
		count--
	}
	fmt.Println("let's go...")

	// 7. Infinite loop with break
	fmt.Println("Finding first even numbers > 10: ")
	num := 11
	for {
		if num%2 == 0 {
			fmt.Println("Found: ", num)
			// break
		}
		num++
		if num > 20 {
			break
		}
	}

	// 8. Continue statement
	fmt.Println("Odd numbers 1 to 10: ")
	for i := 1;  i <= 10; i++ {
		if i%2 == 0 {
			continue	// Skip even numbers
		}
		fmt.Printf("%d", i)
	}
	fmt.Println()

	// 9. Range loop (slices)
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("Fruits: ")
	for index, value := range fruits {
		fmt.Printf("%d:%s\n", index, value)
	}

	// 10. Range loop (ignore index)
	fmt.Println("Fruite (value only): ")
	for _, fruit := range fruits {
		fmt.Printf("%s\n", fruit)
	}

	// 11. Switch statement
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of work week")
	case "Friday":
		fmt.Println("TGIF!")
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Regular day")
	}

	// 12. Switch with no value (like if/else)
	hour := 14
	switch {
	case hour < 12:
		fmt.Println("Good Morning!")
	case hour < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}