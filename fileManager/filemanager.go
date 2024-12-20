package fileManager

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

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

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
		err := file.Close()
		if err != nil {
			return nil, err
		}
		return nil, errors.New("could not Read file")
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	return lines, nil

}

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("could not open file")
	}

	time.Sleep(3 * time.Second) //일부러 고루틴 적용하려고 의도적으로 지체함

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		err := file.Close()
		if err != nil {
			return err
		}
		return errors.New("could not WriteJson")
	}

	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

func New(inputPath string, outputPath string) FileManager {
	return FileManager{inputPath, outputPath}
}
