package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

)

func processTest(input string) []string {
	lines := strings.Split(input, "\n")
	var t int
	fmt.Sscanf(lines[0], "%d", &t)
	lineIdx := 1
	results := make([]string, 0, t)

	for i := 0; i < t; i++ {
		s := strings.TrimSpace(lines[lineIdx])
		if solve(s) {
			results = append(results, "YES")
		} else {
			results = append(results, "NO")
		}
		lineIdx++
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
				expected := strings.TrimSpace(expectedLines[i])
				if got[i] != expected {
					t.Errorf("Test %s, case %d: expected %s, got %s", file.Name(), i+1, expected, got[i])
				}
			}
		}
	}
}
