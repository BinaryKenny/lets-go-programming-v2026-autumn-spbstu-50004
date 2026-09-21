package main

import (
	"fmt"
)

func main() {
	var (
		first    int
		second   int
		operator string
	)
	_, err_first := fmt.Scan(&first)
	if err_first != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err_second := fmt.Scan(&second)
	if err_second != nil {
		fmt.Println("Invalid second operand")
		return
	}
	_, err_operator := fmt.Scan(&operator)
	if err_operator != nil {
		fmt.Println("Invalid operation")
		return
	}
	switch operator {
	case "+":
		fmt.Println(first + second)
	case "-":
		fmt.Println(first - second)
	case "*":
		fmt.Println(first * second)
	case "/":
		if second == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(first / second)
	default:
		fmt.Println("Invalid operation")
	}
}
