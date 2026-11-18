package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func main() {
	configPath := flag.String("config", "./cmd/source/config.yaml", "Path to config file")
	flag.Parse()

	file, err := os.Open(*configPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to open config file: %s", err))
	}

	defer file.Close()

	var conf Config

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse config file: %s", err))
	}
}
