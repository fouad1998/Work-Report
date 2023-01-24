package reporter

import (
	"encoding/json"
	"os"

	"github.com/fouad1998/work-reporter/file"
)

func (r *Reporter) read() error {
	content, err := file.ReadFile(filename)
	if os.IsNotExist(err) {
		return nil
	}

	if err != nil {
		return err
	}

	return json.Unmarshal(content, &r.Items)
}
