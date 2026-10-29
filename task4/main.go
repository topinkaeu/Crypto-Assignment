package main

import (
	"fmt"
)

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func mul(a, b float64) float64 {
	return a * b
}

// Function for Division with error handling
func div(a, b float64) (float64, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

// Function for Modulus (using integers since % works on integers in Go)
func mod(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a % b, true
}
func main() {
	var choice int
	var a, b float64

	for {
		fmt.Println("\n===== Mini Calculator =====")
		fmt.Println("1) Add  2) Sub  3) Mul  4) Div  5) Mod  6) Exit")
		fmt.Print("Choose: ")
		fmt.Scan(&choice)
		if choice == 6 {
			fmt.Println("Exiting, bye bye!")
			break
		}
		if choice < 1 || choice > 6 {
			fmt.Println("Error: invalid choice, please try again.")
			continue
		}

		fmt.Print("Enter a: ")
		fmt.Scan(&a)
		fmt.Print("Enter b: ")
		fmt.Scan(&b)
		switch choice {
		case 1:
			fmt.Printf("Result: %.2f\n", add(a, b))
		case 2:
			fmt.Printf("Result: %.2f\n", sub(a, b))
		case 3:
			fmt.Printf("Result: %.2f\n", mul(a, b))
		case 4:
			result, ok := div(a, b)
			if !ok {
				fmt.Println("Error: division by zero")
			} else {
				fmt.Printf("Result: %.2f\n", result)
			}
		case 5:
			resMod, ok := mod(int(a), int(b))
			if !ok {
				fmt.Println("Error: division by zero")
			} else {
				fmt.Printf("Result: %d\n", resMod)
			}
		}
	}
}
