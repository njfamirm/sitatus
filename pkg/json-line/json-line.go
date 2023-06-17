package jsonLine

import (
	"encoding/json"
	"os"
)

// WriteJSONLine writes a JSON object to a file as a single line
func WriteJSONLine(filepath string, data interface{}) error {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = file.Write(append(jsonData, '\n'))
	if err != nil {
		return err
	}

	return nil
}
