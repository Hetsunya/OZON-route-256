package main

import (
	"bufio"
	"fmt"
	"os"
)

// Структура для курса обмена
type exchange struct {
	from, to int
	rate     float64
}

func solve(banks [3][6]exchange) float64 {
	maxUSD := 0.0

	// 1 обмен: RUB -> USD в любом банке
	for i := 0; i < 3; i++ {
		for _, ex := range banks[i] {
			if ex.from == 0 && ex.to == 1 {
				maxUSD = max(maxUSD, ex.rate)
			}
		}
	}

	// 2 обмена: RUB -> X -> USD в парах банков
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				continue
			}
			for _, ex1 := range banks[i] {
				if ex1.from == 0 { // RUB -> X
					for _, ex2 := range banks[j] {
						if ex2.from == ex1.to && ex2.to == 1 { // X -> USD
							maxUSD = max(maxUSD, ex1.rate*ex2.rate)
						}
					}
				}
			}
		}
	}

	// 3 обмена: RUB -> X -> Y -> USD в трёх банках
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				continue
			}
			for k := 0; k < 3; k++ {
				if k == i || k == j {
					continue
				}
				for _, ex1 := range banks[i] {
					if ex1.from == 0 { // RUB -> X
						for _, ex2 := range banks[j] {
							if ex2.from == ex1.to { // X -> Y
								for _, ex3 := range banks[k] {
									if ex3.from == ex2.to && ex3.to == 1 { // Y -> USD
										maxUSD = max(maxUSD, ex1.rate*ex2.rate*ex3.rate)
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return maxUSD
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		// Читаем 3 банка
		banks := [3][6]exchange{}
		for j := 0; j < 3; j++ {
			for k, fromTo := range [][2]int{{0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}} {
				var n, m int
				fmt.Fscan(in, &n, &m)
				banks[j][k] = exchange{from: fromTo[0], to: fromTo[1], rate: float64(m) / float64(n)}
			}
		}
		fmt.Fprintf(out, "%.6f\n", solve(banks))
	}
}
