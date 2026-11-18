package XML

import (
	"strconv"
	"strings"

	"github.com/BinaryKenny/task-3/cmd/Json"
)

type XmlValute struct {
	ID        string `xml:"id,attr"`
	NumCode   int    `xml:"NumCode"`
	CharCode  string `xml:"CharCode"`
	Nominal   int    `xml:"Nominal"`
	Name      string `xml:"Name"`
	Value     string `xml:"Value"`
	VunitRate string `xml:"VunitRate"`
}

func ValuteParseToJson(val XmlValute) (Json.JsonValute, error) {
	num := val.Value
	num = strings.Replace(num, ",", ".", -1)
	result, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return Json.JsonValute{}, err
	}

	return Json.JsonValute{
		NumCode:  val.NumCode,
		CharCode: val.CharCode,
		Value:    result,
	}, nil
}
