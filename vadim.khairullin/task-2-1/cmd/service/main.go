package main

import (
	"errors"
	"fmt"
)

var (
	errEmployees   = errors.New("wrong number of employees")
	errTemperature = errors.New("wrong format of temperature")
	errSign        = errors.New("wrong format of sign")
)

func main() {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil || departments < 1 || departments > 1000 {
		fmt.Println("Error: wrong number of Departments")

		return
	}

	for range departments {
		err = setTemperatureByEmployees()
		if err != nil {
			fmt.Println(err)

			return
		}
	}
}

func setTemperatureByEmployees() error {
	var employees int

	_, err := fmt.Scan(&employees)
	if err != nil || employees < 1 || employees > 1000 {
		return errEmployees
	}

	minimum := 15
	maximum := 30

	for range employees {
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
