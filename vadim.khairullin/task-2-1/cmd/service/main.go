package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	var (
		Departments int
	)

	_, err := fmt.Scan(&Departments)
	if err != nil || Departments < 1 || Departments > 1000 {
		fmt.Println("Error: wrong number of Departments")

		return
	}

	for range Departments {
		err = setTemperatureByEmployees()
		if err != nil {
			fmt.Println(err)

			return
		}
	}
}

func setTemperatureByEmployees() error {
	var Employees int
	_, err := fmt.Scan(&Employees)
	if err != nil || Employees < 1 || Employees > 1000 {
		return errors.New("Error: wrong number of Employees")
	}

	minimum := 15
	maximum := 30

	for range Employees {
		minimum, maximum, err = defineTemperatureRegime(minimum, maximum)
		if err != nil {
			return err
		}

		if minimum <= maximum {
			fmt.Println(minimum)
		} else {
			fmt.Println(-1)
		}
	}
	return nil
}

func defineTemperatureRegime(min int, max int) (int, int, error) {
	var Temperature string

	_, err := fmt.Scan(&Temperature)
	if err != nil || len(Temperature) != 4 {
		return 0, 0, errors.New("Error: wrong format of temperature")
	}

	sign := Temperature[0:2]
	temp, _ := strconv.Atoi(Temperature[2:4])

	if temp < 15 || temp > 30 {
		return 0, 0, errors.New("Error: temperature out of range")
	}

	if sign == ">=" && min < temp {
		min = temp
	} else if sign == "<=" && max > temp {
		max = temp
	}

	return min, max, nil
}
