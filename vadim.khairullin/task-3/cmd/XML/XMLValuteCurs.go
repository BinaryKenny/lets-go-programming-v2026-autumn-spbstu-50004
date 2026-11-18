package XML

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/BinaryKenny/task-3/cmd/Json"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

type ValCurs struct {
	Date    string      `xml:"Date,attr"`
	Name    string      `xml:"name,attr"`
	Valutes []XmlValute `xml:"Valute"`
}

func ParseXMLFile(path string) ValCurs {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintln("XML file is uncorrect " + err.Error()))
	}

	defer file.Close()

	decoder := xml.NewDecoder(file)

	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		return transform.NewReader(input, charmap.Windows1251.NewDecoder()), nil
	}

	var valCurs ValCurs
	err = decoder.Decode(&valCurs)
	if err != nil {
		panic(fmt.Sprintln("The structure of XML file is wrong" + err.Error()))
	}

	return valCurs
}

func ValuteSliceParseToJson(curs ValCurs) ([]Json.JsonValute, error) {
	var valutes []Json.JsonValute
	for _, val := range curs.Valutes {
		valute, err := ValuteParseToJson(val)
		if err != nil {
			return nil, err
		}
		valutes = append(valutes, valute)
	}

	return valutes, nil
}
