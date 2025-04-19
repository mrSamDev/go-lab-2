package utils

import (
	"errors"
	"os"
)

func ConvertToByteSize(s string) []byte {
	return []byte(s)
}

func ConvertToByteToString(b []byte) string {
	return string(b)
}

func SaveFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func ReadFile(path string) (string, error) {
	file, error := os.ReadFile(path)

	if error != nil {
		return "", errors.New("unable to open file")
	}

	return ConvertToByteToString(file), nil
}

func RemoveFile(path string) error {
	return os.Remove(path)
}
