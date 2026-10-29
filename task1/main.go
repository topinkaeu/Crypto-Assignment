package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Print("Enter x: ")
	fmt.Scan(&x)
	fmt.Print("Enter y: ")
	fmt.Scan(&y)

	fmt.Printf("\nInitial: x = %d, y = %d\n\n", x, y)

	x = x
	fmt.Printf("After x = x (Simple Assignment): x = %d\n", x)

	x += y
	fmt.Printf("After x += y: x = %d\n", x)

	x -= y
	fmt.Printf("After x -= y: x = %d\n", x)

	x *= y
	fmt.Printf("After x *= y: x = %d\n", x)

	if y != 0 {
		x /= y // x = x / y
		fmt.Printf("After x /= y: x = %d\n", x)
	} else {
		fmt.Println("After x /= y: Skipped (cannot divide by zero)")
	}

	if y != 0 {
		x %= y // x = x % y
		fmt.Printf("After x %%= y: x = %d\n", x)
	} else {
		fmt.Println("After x %%= y: Skipped (cannot divide by zero)")
	}
}
