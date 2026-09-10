package main

import "fmt"

func main() {
	var number1, number2 int
	var operator string

	_, err1 := fmt.Scan(&number1)
	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err2 := fmt.Scan(&number2)
	if err2 != nil {
		fmt.Println("Invalid second operand")
		return
	}
	_, err3 := fmt.Scan(&operator)
	if err3 != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operator {
	case "+":
		fmt.Println(number1 + number2)
	case "-":
		fmt.Println(number1 - number2)
	case "*":
		fmt.Println(number1 * number2)
	case "/":
		if number2 == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(number1 / number2)
	default:
		fmt.Println("Invalid operation")
	}
}
