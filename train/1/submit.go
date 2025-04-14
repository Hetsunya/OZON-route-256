package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(s string) bool {
	if len(s) <= 1 {
		return true
	}

	runes := []rune(s)
	n := len(runes)

	// Массив для пометки удалённых символов
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
	return hasFirst
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var t int
	fmt.Sscanf(scanner.Text(), "%d", &t)

	for i := 0; i < t; i++ {
		scanner.Scan()
		s := scanner.Text()
		if solve(s) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
