package tests

import (
	"bufio"
	"dukecon/evaluator"
	"os"
	"strconv"
	"strings"
	"testing"
)

const SINVALS_FILEPATH = "testdata/sin_vals.txt"

func TestSinFunction(t *testing.T) {
	lines, err := readFile(SINVALS_FILEPATH)
	if err != nil {
		t.Fatalf("Cannot read file %s: %s", SINVALS_FILEPATH, err.Error())
	}

	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		x, err := strconv.ParseInt(splitLine[0], 10, 64)

		if err != nil {
			t.Fatalf("Cannot parse integer %s", splitLine[0])
		}

		expectedValue, err := strconv.ParseInt(splitLine[1], 10, 64)

		if err != nil {
			t.Fatalf("Cannot parse integer %s", splitLine[1])
		}

		value := evaluator.Sin(x)

		if value != expectedValue {
			t.Fatalf("Sin(%d) returned incorrect value. got=%d, expected=%d", x, value, expectedValue)
		}
	}
}

func readFile(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
