package day4

import (
	"bufio"
	"fmt"
	"os"
)

func Solver(f *os.File) error {
	m, err := readIntoMatrix(f)
	if err != nil {
		return err
	}
	words := countWord(m, []rune("XMAS"))
	fmt.Printf("Total XMAS: %d\n", words)
	mas := countX_MAS(m)
	fmt.Printf("Total X-MAS: %d\n", mas)
	return err
}

func readIntoMatrix(f *os.File) ([][]rune, error) {
	reader := bufio.NewReader(f)
	var m [][]rune
	var row []rune
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if r == '\n' {
			m = append(m, row)
			row = []rune{}
		} else {
			row = append(row, r)
		}
	}
	if len(row) > 0 {
		m = append(m, row)
	}
	fmt.Printf("Total Lines: %d\n", len(m))
	return m, nil
}

func countWord(m [][]rune, word []rune) int {
	total := 0
	for i := 0; i < len(m); i++ {
		lineTotal := 0
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == 'X' {
				lineTotal += countFrom(m, i, j, word)
			}
		}
		// fmt.Printf("%d:%d\n", i+1, lineTotal)
		total += lineTotal
	}
	return total
}

func countFrom(m [][]rune, i, j int, word []rune) int {
	total := 0
	if searchDirection(m, i, j, 0, 1, word) { // right
		total++
	}
	if searchDirection(m, i, j, 1, 1, word) { // down-right
		total++
	}
	if searchDirection(m, i, j, 1, 0, word) { // down
		total++
	}
	if searchDirection(m, i, j, 1, -1, word) { // down-left
		total++
	}
	if searchDirection(m, i, j, 0, -1, word) { // left
		total++
	}
	if searchDirection(m, i, j, -1, -1, word) { // up-left
		total++
	}
	if searchDirection(m, i, j, -1, 0, word) { // up
		total++
	}
	if searchDirection(m, i, j, -1, 1, word) { // up-right
		total++
	}
	return total
}

func searchDirection(m [][]rune, i, j, dI, dJ int, word []rune) bool {
	for _, c := range word {
		if i < 0 || i >= len(m) || j < 0 || j >= len(m[i]) {
			return false
		}
		s := m[i][j]
		if s != c {
			return false
		}
		i += dI
		j += dJ
	}
	return true
}

func countX_MAS(m [][]rune) int {
	word := []rune("MAS")
	total := 0
	for i := 1; i < len(m)-1; i++ {
		for j := 1; j < len(m[i])-1; j++ {
			if m[i][j] == 'A' {
				if (searchDirection(m, i-1, j-1, 1, 1, word) || searchDirection(m, i+1, j+1, -1, -1, word)) && (searchDirection(m, i-1, j+1, 1, -1, word) || searchDirection(m, i+1, j-1, -1, 1, word)) {
					total++
				}
			}
		}
	}
	return total
}
