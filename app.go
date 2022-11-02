package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	lowerCharSet   = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*/\\?"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func main() {
	passwordLength := 10
	password := generatePassword(passwordLength)
	fmt.Println("Your password: ", password)
}

func generatePassword(length int) string {
	rand.Seed(time.Now().Unix())
	inRune := []rune(allCharSet)
	var password strings.Builder
	for index := 0; index < length; index++ {
		password.WriteRune(inRune[rand.Intn(len(inRune))])
	}
	return password.String()
}
