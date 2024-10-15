package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("Failed to open file!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner .Scan() {
		lines = append(lines, scanner.Text())
	}

	error := scanner.Err()

	if error != nil {
		// file.Close()

		return nil, errors.New("Failed to read from file!")
	}

	// file.Close()

	return lines, nil

}

func (fm FileManager) WriteJSONToFile(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("Failed to create file!")
	}

	defer file.Close()

	time.Sleep(3 * time.Second)
	
	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)

	if err != nil {
		// file.Close()

		return errors.New("Failed to convert data to JSON!")
	}

	// file.Close()

	return nil

}

func New(inFilePath, outFilePath string) FileManager {
	return FileManager {
		InputFilePath: inFilePath,
		OutputFilePath: outFilePath,
	}
}


