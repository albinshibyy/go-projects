package main

import (
	"fmt"
)

func main() {

	var num1, num2, ans int
	var ope, exit string

	for {
		fmt.Println("SIMPLE CALCULATOR CLI")
		fmt.Println("Enter first value:")
		fmt.Scanln(&num1)
		fmt.Println("Enter second value:")
		fmt.Scanln(&num2)
		fmt.Println("Enter the operator (+, -, *, /, %):")
		fmt.Scanln(&ope)
		if ope == "+" || ope == "-" || ope == "*" || ope == "/" || ope == "%" {
			switch ope {
			case "+":
				ans = num1 + num2
				fmt.Println("Answer for", num1, " ", ope, " ", num2, " = ", ans)
			case "-":
				ans = num1 - num2
				fmt.Println("Answer for", num1, " ", ope, " ", num2, " = ", ans)
			case "*":
				ans = num1 * num2
				fmt.Println("Answer for", num1, " ", ope, " ", num2, " = ", ans)
			case "/":
				if num2 == 0 {
					fmt.Println("Division by zero is not possible")
				} else {
					ans = num1 / num2
					fmt.Println("Answer for", num1, " ", ope, " ", num2, " = ", ans)
				}
			case "%":

				if num2 == 0 {
					fmt.Println("Modulo by zero is not possible")
				} else {
					ans = num1 % num2
					fmt.Println("Answer for", num1, " ", ope, " ", num2, " = ", ans)
				}
			default:
				fmt.Println("Invalid operator")
			}
		} else {
			fmt.Println("Error in entered operator!")
		}
		// LOOPING PERMISSION
		fmt.Println("do u want to continue 'yes/no': ")
		fmt.Scanln(&exit)
		if exit == "no" || exit == "n" || exit == "exit" {
			break
		} else {
			continue
		}

	}
}
