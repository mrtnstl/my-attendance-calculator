package cmd

import (
	"fmt"
	"os"
)

func ReadFile(filePath string) (string, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("err: %s\n", err)
	}
	return string(b), nil
}
