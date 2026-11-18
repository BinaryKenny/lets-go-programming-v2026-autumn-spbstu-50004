package main

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func main() {
	configPath := flag.String("config", "config.yaml", "Path to config file")
	flag.Parse()

	file, err := os.Open(*configPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to open config file: %s", err))
	}

	defer file.Close()

}
