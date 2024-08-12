package main

import (
	"fmt"
	"strings"
)

const (
	minLength              int = 12
	minLowerCase           int = 1
	minUpperCase           int = 1
	minNumber              int = 1
	minSpecialChar         int = 1
	deductionForDictWord   int = 2
	deductionForRepeatChar int = 1
)

var disallowedWords = []string{"password", "123456", "qwerty", "abc123"} // Add more common passwords

func checkPasswordStrength(password string) (int, string, bool, bool, bool, bool) {
	// Declare variables within the function for proper scope
	var (
		hasLowerCase   bool = strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz")
		hasUpperCase   bool = strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		hasNumber      bool = strings.ContainsAny(password, "0123456789")
		hasSpecialChar bool = strings.ContainsAny(password, "!@#$%^&*()-_=+[]{};':\",<.>/?|\\")
		isLongEnough   bool = len(password) >= minLength
		score          int  = 0
	)

	// Use all the declared variables to avoid "declared and not used" errors
	if isLongEnough {
		score++
	}
	if hasLowerCase {
		score++
	}
	if hasUpperCase {
		score++
	}
	if hasNumber {
		score++
	}
	if hasSpecialChar {
		score++
	}

	// Apply deductions for weaknesses
	if isDictWord(password) {
		score -= deductionForDictWord
	}
	if hasRepeatingChars(password) {
		score -= deductionForRepeatChar
	}

	message := ""
	switch {
	case score < 1:
		message = "Very Weak (password is too short)"
	case score < 3:
		message = "Weak (needs more character types)"
	case score < 4:
		message = "Moderate (consider adding more complexity)"
	default:
		message = "Strong"
	}

	return score, message, hasLowerCase, hasUpperCase, hasNumber, hasSpecialChar
}

func isDictWord(password string) bool {
	for _, word := range disallowedWords {
		if strings.ToLower(password) == word {
			return true
		}
	}
	return false
}

func hasRepeatingChars(password string) bool {
	for i := 1; i < len(password); i++ {
		if password[i] == password[i-1] {
			return true
		}
	}
	return false
}

func main() {
	for {
		fmt.Println("SIMPLE PASSWORD STRENGTH TEST (CLI)")
		fmt.Println("TYPE 'exit' TO END THIS PROGRAM OR INPUT YOUR PASSWORD TO TEST")
		fmt.Println("Enter your password or exit: ")
		var password string
		fmt.Scanln(&password)

		if strings.ToLower(password) == "exit" || strings.ToLower(password) == "y" {
			break
		}
		score, message, hasLowerCase, hasUpperCase, hasNumber, hasSpecialChar := checkPasswordStrength(password)
		fmt.Println(message)

		maxScore := minLength + minLowerCase + minUpperCase + minNumber + minSpecialChar
		fmt.Printf("Your password score is %d (out of a maximum of %d)\n", score, maxScore)

		if score < maxScore {
			fmt.Println("Here are some suggestions for improvement:")
			if !hasLowerCase {
				fmt.Println("- Include at least one lowercase letter")
			}
			if !hasUpperCase {
				fmt.Println("- Include at least one uppercase letter")
			}
			if !hasNumber {
				fmt.Println("- Include at least one number")
			}
			if !hasSpecialChar {
				fmt.Println("- Include at least one special character")
			}
			if len(password) < minLength {
				fmt.Printf("- Increase password length to at least %d characters\n", minLength)
			}
		}
	}
}
