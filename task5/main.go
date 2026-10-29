package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Use bufio.NewReader so the user can type sentences with spaces
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter text to encode: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input = strings.TrimSpace(input)
	// Convert the string into a slice of bytes (raw data)
	data := []byte(input)

	fmt.Printf("\n--- Results for: \"%s\" ---\n", input)

	fmt.Print("1. Binary:     ")
	for _, b := range data {
		fmt.Printf("%08b ", b)
	}
	fmt.Println()
	hexEncoded := hex.EncodeToString(data)
	fmt.Printf("2. Hexadecimal: %s\n", hexEncoded)

	base64Encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Printf("3. Base64:      %s\n", base64Encoded)
}
