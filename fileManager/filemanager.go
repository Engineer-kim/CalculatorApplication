package fileManager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, errors.New("could not open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("could not Read file")
	}
	file.Close()
	return lines, nil

}

func WriteJson(path string, data interface{}) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("could not open file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("could not WriteJson")
	}

	file.Close()
	return nil
}
