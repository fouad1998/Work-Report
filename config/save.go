package config

import (
	"encoding/json"

	"github.com/fouad1998/work-reporter/file"
)

func (c *Config) Save() error {
	content, err := json.Marshal(c)
	if err != nil {
		return err
	}

	return file.WriteFile(filename, content)
}
