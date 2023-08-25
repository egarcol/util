package readfile

import (
	"bufio"
	"os"
)

func GetLines(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}
