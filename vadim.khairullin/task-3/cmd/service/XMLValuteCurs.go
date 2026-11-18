package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

type ValCurs struct {
	Date   string      `xml:"Date,attr"`
	Name   string      `xml:"name,attr"`
	Valute []XMLValute `xml:"Valute"`
}

func parseXMLFile(path string) *ValCurs {
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

	return &valCurs
}

func 