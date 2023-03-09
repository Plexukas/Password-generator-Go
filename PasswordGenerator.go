package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// asks the user for the password length
	var length int
	fmt.Println("Enter password length: ")
	fmt.Scanln(&length)

	// possible characters in the password
	const (
		letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numbers = "0123456789"
		symbols = "!@#$%^&*()_+-=[]{}|;:,./<>?"
	)

	// combines all possible characters into a single string
	allChars := letters + numbers + symbols

	// generates the password using a random index from the allChars string
	password := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(allChars))
		password += string(allChars[randomIndex])
	}
	fmt.Println("Your password is:", password)
}
