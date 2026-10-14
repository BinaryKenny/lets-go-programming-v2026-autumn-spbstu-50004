package main

import "fmt"

func main() {
	var (
		Departments int
		Employees   int
		Temperature string
	)
	minimum := 15
	maximum := 30
	_, err := fmt.Scan(&Departments)
	if err != nil  Departments < 1  Departments > 1000 {
		fmt.Println("Error: wrong number of Departments", err)
		return
	}
}
