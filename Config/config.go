package Config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Server struct {
		Port string
	}
}

var AppConfig Config

func GetConfig() error {

	file, error := os.Open("config.yml")
	if error != nil {
		return error
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err := decoder.Decode(&AppConfig)
	if err != nil {
		return err
	}
	return nil
}
