package validator

import (
	"encoding/json"
	"fmt"

	"github.com/Rajeevnita1993/json-parser/fileio"
)

func ValidateJSONFile(filename string) error {
	data, err := fileio.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	err = ParseJSON(data)
	if err != nil {
		return fmt.Errorf("invalid JSON: %v", err)

	}

	return nil

}

func ParseJSON(data []byte) error {
	var obj map[string]interface{}
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}
	fmt.Println("Parsed data: ", obj)
	return nil
}
