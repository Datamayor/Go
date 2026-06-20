package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	FilePath string
}

func NewStorage[T any](filePath string) *Storage[T] {
	return &Storage[T]{
		FilePath: filePath,
	}
}

func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		return err
	}

	return os.WriteFile(s.FilePath, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FilePath)

	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data)
}
