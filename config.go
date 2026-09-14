package commandline

import (
	"encoding/json"
	"github.com/GabeCordo/Commandline/paths"
)

const (
	DefaultJSONPrefix string = ""
	DefaultJSONIndent string = " "
)

type Config struct {
	Version     float32 `json:"version"`
	UserProfile struct {
		FirstName string `json:"first-Name"`
		LastName  string `json:"last-Name"`
		Email     string `json:"email"`
	} `json:"profile"`
}

func NewConfig() *Config {
	config := new(Config)
	return config
}

func (config *Config) ToJson(path paths.Path) error {
	if path.DoesNotExist() {
		panic("the path is not valid, it cannot be converted to JSON")
	}

	bytes, err := json.MarshalIndent(config, DefaultJSONPrefix, DefaultJSONIndent)
	if err != nil {
		panic("there was an issue marshalling the Config to JSON")
	}

	return path.Write(bytes)
}

func (config *Config) FromJson(path paths.Path) *Config {
	if path.DoesNotExist() {
		panic("the path is not valid, the config cannot be updated to match the JSON file")
	}

	bytes, err := path.Read()
	if err != nil {
		panic("there was an issue reading the path")
	}

	err = json.Unmarshal(bytes, config)
	if err != nil {
		panic("there was an issue marshalling from JSON to the Config instance")
	}

	return config
}
