package filemanager

import (
	"bufio"
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
