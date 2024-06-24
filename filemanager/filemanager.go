package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func ReadLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, fmt.Errorf("could not open file %s", fileName)
	}

	scanner := bufio.NewScanner(file)
	var prices []string
	for scanner.Scan() {
		err = scanner.Err()
		if err != nil {
			file.Close()
			return nil, fmt.Errorf("reading content from %s failed", fileName)
		}

		prices = append(prices, scanner.Text())
	}

	return prices, nil
}

func WriteJSON(path string, data any) error {
	var newFile, err = os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file %s with error %s", path, err)
	}

	encoder := json.NewEncoder(newFile)
	err = encoder.Encode(data)

	if err != nil {
		newFile.Close()
		return errors.New("failed to convert data to JSON format")
	}

	newFile.Close()
	return nil
}
