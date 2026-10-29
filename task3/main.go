package main

import (
	"fmt"
)

// Function for Bitwise AND
func myAND(a, b int) {
	result := a & b
	fmt.Printf("Bitwise AND (%d & %d) = %d\n", a, b, result)
}

func myOR(a, b int) {
	result := a | b
	fmt.Printf("Bitwise OR (%d | %d) = %d\n", a, b, result)
}

func myXor(a, b int) {
	result := a ^ b
	fmt.Printf("Bitwise XOR (%d ^ %d) = %d\n", a, b, result)
}

func myNOT(a, b int) {
	fmt.Printf("Bitwise NOT (^%d) = %d\n", a, ^a)
	fmt.Printf("Bitwise NOT (^%d) = %d\n", b, ^b)
}

func myShifts(a, b int) {
	// Using b as the shift amount (handling positive safe values)
	shiftAmt := uint(b)
	if b < 0 {
		shiftAmt = 0
	}
	fmt.Printf("Left Shift (%d << %d) = %d\n", a, b, a<<shiftAmt)
	fmt.Printf("Right Shift (%d >> %d) = %d\n", a, b, a>>shiftAmt)
}

// Function to demonstrate bitwise assignment operators
func demonstrateAssignments(a, b int) {
	fmt.Println("\n--- Bitwise Assignment Operators ---")
	x := a
	// AND Assignment
	x &= b
	fmt.Printf("After x &= b: x = %d\n", x)

	// Reset x and do OR Assignment
	x = a
	x |= b
	fmt.Printf("After x |= b: x = %d\n", x)

	// Reset x and do XOR Assignment
	x = a
	x ^= b
	fmt.Printf("After x ^= b: x = %d\n", x)

	// Reset x and do Left/Right Shift Assignments
	if b >= 0 && b < 5 {
		x = a
		x <<= uint(b)
		fmt.Printf("After x <<= b: x = %d\n", x)

		x = a
		x >>= uint(b)
		fmt.Printf("After x >>= b: x = %d\n", x)
	}
}

func main() {
	var a, b int
	fmt.Print("Enter the first integer (a): ")
	fmt.Scan(&a)
	fmt.Print("Enter the second integer (b): ")
	fmt.Scan(&b)

	fmt.Printf("\n--- Executing Functions (Lab #3) ---\n")

	// Calling functions exactly as requested in the lab pseudocode
	myXor(a, b)
	myNOT(a, b)
	myOR(a, b)
	myAND(a, b)
	myShifts(a, b)

	// Show assignment operators effect
	demonstrateAssignments(a, b)
}
