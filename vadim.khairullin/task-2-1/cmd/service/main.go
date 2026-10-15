package main

import (
	"errors"
	"fmt"
)

var (
	errEmployees   = errors.New("Error: wrong number of Employees")
	errTemperature = errors.New("Error: wrong format of temperature")
	errSign        = errors.New("Error: wrong format of sign")
)

func main() {
	var Departments int

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
		return errEmployees
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

func defineTemperatureRegime(minTemp, maxTemp int) (int, int, error) {
	var (
		sign string
		temp int
	)

	_, err := fmt.Scan(&sign, &temp)

	if err != nil {
		return 0, 0, errTemperature
	}

	if temp < 15 || temp > 30 {
		return 0, 0, errTemperature
	}

	switch sign {
	case ">=":
		if minTemp < temp {
			minTemp = temp
		}
	case "<=":
		if maxTemp > temp {
			maxTemp = temp
		}
	default:
		return 0, 0, errSign
	}

	return minTemp, maxTemp, nil
}
