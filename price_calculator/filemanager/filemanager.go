package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("Failed to open file")
	}

	defer file.Close() // will be called when the surrounding function will be closed

	scanner := bufio.NewScanner(file)

	var lines []string // ['10', '20', '30'] each number will become a line

	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // that's why append
	}

	err = scanner.Err()
	if err != nil {
		// file.Close()
		return nil, errors.New("Failed to read file")
	}

	// file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("Failed to create file")
	}

	defer file.Close()

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		// file.Close()
		return errors.New("Failed to convert data to json")
	}

	// file.Close()
	return nil
}
