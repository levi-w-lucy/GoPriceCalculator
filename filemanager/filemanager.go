package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, fmt.Errorf("could not open file %s", fm.InputFilePath)
	}

	scanner := bufio.NewScanner(file)
	var prices []string
	for scanner.Scan() {
		err = scanner.Err()
		if err != nil {
			file.Close()
			return nil, fmt.Errorf("reading content from %s failed", fm.InputFilePath)
		}

		prices = append(prices, scanner.Text())
	}

	return prices, nil
}

func (fm FileManager) WriteResult(data any) error {

	var newFile, err = os.Create(fm.OutputFilePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s with error %s", fm.OutputFilePath, err)
	}

	defer newFile.Close()

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(newFile)
	err = encoder.Encode(data)

	if err != nil {
		newFile.Close()
		return errors.New("failed to convert data to JSON format")
	}

	newFile.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
