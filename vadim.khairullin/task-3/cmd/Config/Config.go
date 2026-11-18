package Config

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func ParseConfig() (Config, error) {
	configPath := flag.String("config", "config.yaml", "Path to config file")
	flag.Parse()

	file, err := os.Open(*configPath)
	if err != nil {
		return Config{}, err
	}

	defer file.Close()

	var conf Config

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
