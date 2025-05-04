package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type NoTxtError struct {
	Message string
}

// Custom errors must implement the Error() interface method
func (e *NoTxtError) Error() string {
	return e.Message
}

func loadFile(path string) (string, error) {
	if !strings.HasSuffix(path, ".txt") {
		return "", &NoTxtError{"File must be a .txt file"}
	}

	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	return path, nil
}

func main() {
	p, err := loadFile("test.txt")
	var noTxtError *NoTxtError
	if errors.Is(err, os.ErrNotExist) { // errors.Is() checks if the error is of a specific type
		fmt.Println("File doesn't exist")
	} else if errors.As(err, &noTxtError) { // errors.As() checks if the error is of a specific type (castable) and stores it in the second argument
		fmt.Println("NoTxtError: ", noTxtError.Message)
	}

	fmt.Println("Loaded file: ", p)
}
