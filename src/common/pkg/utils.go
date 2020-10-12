package utils

import (
	"bytes"
	"encoding/json"
)

func StructToJson(data interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := json.NewEncoder(buf).Encode(data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var (
	Greeting      = "Hello, Shadow. It's time to learn GoLang!"
	innerGreeting = "Hidden Message."
)
