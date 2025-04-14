package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func solve(s string) bool {
	if len(s) <= 1 {
		return true
	}

	runes := []rune(s)
	n := len(runes)

	// Массив для пометки удалённых символов (0 — удалён)
	active := make([]bool, n)
	for i := range active {
		active[i] = true
	}

	// Итеративно удаляем символы между одинаковыми
	for {
		changed := false
		for i := 0; i < n-2; i++ {
			if active[i] && active[i+2] && runes[i] == runes[i+2] && active[i+1] {
				// Найдена тройка x_y_x, удаляем y
				active[i+1] = false
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	// Проверяем, что остались только одинаковые символы
	var first rune
	hasFirst := false
	for i := 0; i < n; i++ {
		if active[i] {
			if !hasFirst {
				first = runes[i]
				hasFirst = true
			} else if runes[i] != first {
				return false
			}
		}
	}
	return hasFirst // Если ничего не осталось, строка валидна
}

func processTest(input string) []string {
	var results []string
	lines := strings.Split(input, "\n")
	var t int
	fmt.Sscanf(lines[0], "%d", &t)
	lineIdx := 1

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
				if got[i] != expectedLines[i] {
					t.Errorf("Test %s, case %d: expected %s, got %s", file.Name(), i+1, expectedLines[i], got[i])
				}
			}
		}
	}
}

func main() {
	testing.Main(func(pat, str string) (bool, error) { return true, nil }, []testing.InternalTest{
		{Name: "TestSolution", F: TestSolution},
	}, nil, nil)
}
