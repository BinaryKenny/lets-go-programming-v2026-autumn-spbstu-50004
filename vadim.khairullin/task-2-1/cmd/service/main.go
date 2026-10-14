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
		fmt.Println("Error: wrong number of Departments")

		return
	}

	for range Departments {
		_, err = fmt.Scan(&Employees)
		if err != nil || Employees < 1 || Employees > 1000 {
			fmt.Println("Error: wrong number of Employees")

			return
		}

		minimum := 15
		maximum := 30

		for range Employees {
			_, err = fmt.Scan(&Temperature)
			if err != nil || len(Temperature) != 4 {
				fmt.Println("Error: wrong format of temperature")

				return
			}

			sign := Temperature[0:2]
			temp, _ := strconv.Atoi(Temperature[2:4])

			if temp < 15 || temp > 30 {
				fmt.Println("Error: wrong temperature")

				return
			}

			if sign == ">=" && minimum < temp {
				minimum = temp
			} else if sign == "<=" && maximum > temp {
				maximum = temp
			}

			if minimum <= maximum {
				fmt.Println(minimum)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
