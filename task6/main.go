package main

import (
	"encoding/hex"
	"fmt"
)

func xorEncrypt(text string, key byte) string {
	// Convert the string into a slice of bytes
	textBytes := []byte(text)
	resultBytes := make([]byte, len(textBytes))

	// Loop through each byte and XOR it with the key
	for i := 0; i < len(textBytes); i++ {
		resultBytes[i] = textBytes[i] ^ key
	}

	// Convert the resulting bytes back to a string
	return string(resultBytes)
}

func main() {
	// 1. Define our original plaintext message and a single-byte key
	plaintext := "Hello, Crypto Course T-T!"
	var key byte = 'K' // Using the character 'K' as our key (ASCII value)

	fmt.Printf("Original Plaintext: %s\n", plaintext)
	fmt.Printf("Key (Character):    %c (ASCII: %d)\n\n", key, key)

	// 2. Encrypt the plaintext
	ciphertext := xorEncrypt(plaintext, key)

	// printing it as Hex makes it safe and readable to look at!
	fmt.Printf("Ciphertext (Raw):   %s\n", ciphertext)
	fmt.Printf("Ciphertext (Hex):   %s\n", hex.EncodeToString([]byte(ciphertext)))

	// 3. Decrypt the ciphertext using the EXACT SAME function and key
	decryptedText := xorEncrypt(ciphertext, key)
	fmt.Printf("\nDecrypted Message:  %s\n", decryptedText)
}
