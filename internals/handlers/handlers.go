package handlers

import (
	"encoding/json"
	"fmt"
	"os"
)

// Function to read data from a JSON file
func ReadJSON(filePath string, result any) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return err
	}

	return nil
}

// Function to write data to a JSON file
func WriteJSON(filePath string, data any) error {
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Data written successfully to JSON file.")
	return nil
}
