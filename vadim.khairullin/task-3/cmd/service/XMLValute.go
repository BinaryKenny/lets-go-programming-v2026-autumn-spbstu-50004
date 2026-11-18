package main

import (
	"strconv"
	"strings"
)

type XMLValute struct {
	ID        string `xml:"id,attr"`
	NumCode   int    `xml:"NumCode"`
	CharCode  string `xml:"CharCode"`
	Nominal   int    `xml:"Nominal"`
	Name      string `xml:"Name"`
	Value     string `xml:"Value"`
	VunitRate string `xml:"VunitRate"`
}

func valuteParseToJson(val *XMLValute) (JsonValute, error) {
	num := val.Value
	num = strings.Replace(num, ",", ".", -1)
	result, err := strconv.ParseFloat(num, 64)
	if err != nil {

		return JsonValute{}, err
	}

	return JsonValute{val.NumCode, val.CharCode, result}, nil
}
