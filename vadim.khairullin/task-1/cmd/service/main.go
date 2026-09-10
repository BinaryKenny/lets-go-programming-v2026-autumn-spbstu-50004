package main

import "fmt"

func main() {
	var number1, number2 int
	var operator rune

	_, err1 := fmt.Scan(&number1)
	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	}
}
