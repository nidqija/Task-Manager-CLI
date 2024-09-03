package main

import (
	"encoding/json"
	"os"
)

// Generic storage structure for any type T
type Storage[T any] struct {
	FileName string
}

// Constructor for creating a new Storage instance
func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}

// Save method for writing data to the file in JSON format
func (s *Storage[T]) Save(data T) error {
	// Marshal the data into indented JSON format
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Write the JSON data to the file
	return os.WriteFile(s.FileName, fileData, 0644)
}

// Load method for reading data from the file and unmarshalling it into the provided pointer
func (s *Storage[T]) Load(data *T) error {
	// Read the file contents
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		return err
	}

	// Unmarshal the JSON data into the provided pointer
	return json.Unmarshal(fileData, data)
}