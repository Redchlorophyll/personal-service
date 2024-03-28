package environment_variable

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func GetEnvironmentVariables() Config {
	appConfig := Config{}
	configFilePath := os.Getenv("CONFIG_FILE_PATH")
	if configFilePath == "" {
		configFilePath = ".config.yaml"
	}

	err := readConfig(&appConfig, configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	return appConfig
}

func readConfig(config *Config, fileName string) error {
	fname, err := filepath.Abs(fileName)
	if err != nil {
		return err
	}

	yamlFile, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return err
	}

	return nil
}
