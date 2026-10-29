package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Print("Enter the first integer: X = ")
	fmt.Scan(&x)
	fmt.Print("Enter the second integer: Y =  ")
	fmt.Scan(&y)

	fmt.Printf("\n--- Your Values: x = %d, y = %d ---\n\n", x, y)

	bothPositive := (x > 0) && (y > 0)
	fmt.Printf("Both positive (x > 0 && y > 0): %t\n", bothPositive)

	oneGreater := (x > y) || (y > x)
	fmt.Printf("One greater than the other (x > y || y > x): %t\n", oneGreater)

	notEqual := !(x == y)
	fmt.Printf("Not equal (!(x == y)): %t\n", notEqual)
}
