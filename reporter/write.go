package reporter

import (
	"encoding/json"

	"github.com/fouad1998/work-reporter/file"
)

var filename = "reporter"

func (r *Reporter) write() error {
	content, _ := json.Marshal(r.Items)

	return file.WriteFile(filename, content)
}
