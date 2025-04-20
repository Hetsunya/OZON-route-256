package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func processTest(input string) []float64 {
	lines := strings.Split(input, "\n")
	var t int
	fmt.Sscanf(lines[0], "%d", &t)
	lineIdx := 1
	results := make([]float64, 0, t)

	for i := 0; i < t; i++ {
		// Читаем 3 банка
		banks := [3][6]exchange{}
		for j := 0; j < 3; j++ {
			for k, fromTo := range [][2]int{{0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}} {
				var n, m int
				fmt.Sscanf(lines[lineIdx], "%d %d", &n, &m)
				banks[j][k] = exchange{from: fromTo[0], to: fromTo[1], rate: float64(m) / float64(n)}
				lineIdx++
			}
		}
		results = append(results, solve(banks))
	}
	return results
}

func readFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func TestSolution(t *testing.T) {
	testDir := `tests`
	files, err := os.ReadDir(testDir)
	if err != nil {
		t.Fatalf("Failed to read test directory: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() && !strings.HasSuffix(file.Name(), ".a") {
			inputFile := filepath.Join(testDir, file.Name())
			outputFile := filepath.Join(testDir, file.Name()+".a")

			inputData, err := readFile(inputFile)
			if err != nil {
				t.Errorf("Failed to read input file %s: %v", inputFile, err)
				continue
			}

			expectedData, err := readFile(outputFile)
			if err != nil {
				t.Errorf("Failed to read output file %s: %v", outputFile, err)
				continue
			}

			got := processTest(inputData)
			expectedLines := strings.Split(strings.TrimSpace(expectedData), "\n")
			if len(got) != len(expectedLines) {
				t.Errorf("Test %s: expected %d results, got %d", file.Name(), len(expectedLines), len(got))
				continue
			}

			for i := 0; i < len(got); i++ {
				expected, err := strconv.ParseFloat(strings.TrimSpace(expectedLines[i]), 64)
				if err != nil {
					t.Errorf("Test %s, case %d: failed to parse expected value: %v", file.Name(), i+1, err)
					continue
				}
				if abs(got[i]-expected) > 1e-6 && abs((got[i]-expected)/expected) > 1e-6 {
					t.Errorf("Test %s, case %d: expected %.6f, got %.6f", file.Name(), i+1, expected, got[i])
				}
			}
		}
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
