package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		Departments int
		Employees   int
		Temperature string
	)
	_, err := fmt.Scan(&Departments)
	if err != nil || Departments < 1 || Departments > 1000 {
		fmt.Println("Error: wrong number of Departments", err)
		return
	}
	for i := 0; i < Departments; i++ {
		_, err = fmt.Scan(&Employees)
		if err != nil || Employees < 1 || Employees > 1000 {
			fmt.Println("Error: wrong number of Employees", err)
			return
		}
		minimum := 15
		maximum := 30
		for j := 0; j < Employees; j++ {
			_, err = fmt.Scan(&Temperature)
			if err != nil || len(Temperature) != 4 {
				fmt.Println("Error: wrong number of Temperature", err)
				return
			}
			sign := Temperature[0:2]
			temp, _ := strconv.Atoi(Temperature[2:4])
			if temp < 15 || temp > 30 {
				fmt.Println("Error: wrong temperature")
				return
			}
			if sign == ">=" {
				minimum = temp
			} else if sign == "<=" {
				maximum = temp
			} else {
				fmt.Println("Error: wrong temperature sign")
				return
			}
			if minimum <= maximum {
				fmt.Println(minimum)
			} else {
				fmt.Println(-1)
				return
			}
		}
	}
}
