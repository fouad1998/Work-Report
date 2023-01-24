package config

import (
	"encoding/json"
	"os"

	"github.com/fouad1998/work-reporter/file"
)

var filename = "config.json"

func (c *Config) Read() error {
	content, err := file.ReadFile(filename)
	if os.IsNotExist(err) {
		return nil
	}

	if err != nil {
		return err
	}

	return json.Unmarshal(content, c)
}
