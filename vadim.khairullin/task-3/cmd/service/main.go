package main

import (
	"github.com/BinaryKenny/task-3/Config"
	"github.com/BinaryKenny/task-3/Json"
	"github.com/BinaryKenny/task-3/XML"
)

func main() {
	conf, err := Config.ParseConfig()
	if err != nil {
		panic("Failed to parse config: " + err.Error())
	}

	ValCurs := XML.ParseXMLFile(conf.InputFile)

	valutes, err := XML.ValuteSliceParseToJson(ValCurs)
	if err != nil {
		panic(err)
	}
	Json.EncodeToJson(valutes, conf.OutputFile)
}
