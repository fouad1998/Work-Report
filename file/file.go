package file

import (
	"encoding/base64"
	"os"
)

func ReadFile(filename string) ([]byte, error) {
	nameCrypt := base64.StdEncoding.EncodeToString([]byte(filename))
	content, err := os.ReadFile(nameCrypt)
	if err != nil {
		return nil, err
	}

	d := decrypt((content))

	return d, nil
}

func WriteFile(filename string, content []byte) error {
	nameCrypt := base64.StdEncoding.EncodeToString([]byte(filename))
	c := encrypt(content)

	return os.WriteFile(nameCrypt, c, 0644)
}
