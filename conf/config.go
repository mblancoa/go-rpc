package conf

import (
	"encoding/json"
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/mblancoa/go-rpc/internal/errors"
	"io"
	"os"
)

const (
	RunMode = "RUN_MODE"
)

var configFile string

func GetConfigFile() string {
	if configFile == "" {
		mode := os.Getenv(RunMode)
		if mode != "" {
			configFile = fmt.Sprintf("conf/%s.application.yml", mode)
		} else {
			configFile = "conf/application.yml"
		}
	}
	return configFile
}

func LoadJsonConfiguration(fileName string, configObj interface{}) {
	bts := loadFile(fileName)
	err := json.Unmarshal(bts, configObj)
	errors.ManageErrorPanic(err)
}

func LoadYamlConfiguration(fileName string, configObj interface{}) {
	bts := loadFile(fileName)
	err := yaml.Unmarshal(bts, configObj)
	errors.ManageErrorPanic(err)
}

func loadFile(file string) []byte {
	confFile, err := os.Open(file)
	errors.ManageErrorPanic(err)
	defer func() {
		err := confFile.Close()
		errors.ManageErrorPanic(err)
	}()

	bts, err := io.ReadAll(confFile)
	errors.ManageErrorPanic(err)
	return bts
}
