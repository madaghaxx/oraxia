package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: oraxia <word>")
		fmt.Println("Generates a deterministic secure password based on the provided word.")
		os.Exit(1)
	}

	secret, err := os.ReadFile("secret.txt")
	if err != nil {
		fmt.Println("Error reading secret.txt:", err)
		fmt.Println("Please ensure secret.txt exists with your secret token.")
		os.Exit(1)
	}

	// Trim whitespace from secret
	secret = []byte(strings.TrimSpace(string(secret)))

	if len(secret) == 0 {
		fmt.Println("Error: secret.txt is empty. Please add a secret token.")
		os.Exit(1)
	}

	// Use the single input word
	input := os.Args[1]

	// Compute HMAC-SHA256
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(input))
	hash := h.Sum(nil)

	// Take first 16 bytes for 128-bit security (shorter password)
	shortHash := hash[:16]

	// Encode to base64 for a printable password
	password := base64.StdEncoding.EncodeToString(shortHash)

	fmt.Println(password)
}
