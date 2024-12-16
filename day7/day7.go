package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solver(f *os.File) error {
	eqs, err := parseInputFile(f)
	if err != nil {
		return err
	}
	total := 0
	for _, eq := range eqs {
		if eq.isSolvable() {
			total += eq.result
			// fmt.Printf("Equation %v is solvable\n", eq)
		} else {
			// fmt.Printf("Equation %v is not solvable\n", eq)
		}
	}
	fmt.Printf("Total solvable equations: %d\n", total)
	return nil
}

type equation struct {
	result int
	parts  []int
}

func (e equation) isSolvable() bool {
	return checkSolvable(e.parts, e.result)
}

func checkSolvable(parts []int, result int) bool {
	if len(parts) == 1 {
		return parts[0] == result
	}
	sumParts := make([]int, len(parts)-1)
	mulParts := make([]int, len(parts)-1)
	copy(sumParts, parts[1:])
	copy(mulParts, parts[1:])
	sumParts[0] = parts[0] + parts[1]
	mulParts[0] = parts[0] * parts[1]
	return (sumParts[0] <= result && checkSolvable(sumParts, result)) || (mulParts[0] <= result && checkSolvable(mulParts, result))
}

func parseEquation(line string) (equation, error) {
	mainParts := strings.Split(line, ":")
	result, err := strconv.Atoi(mainParts[0])
	if err != nil {
		return equation{}, err
	}
	sParts := strings.Split(strings.Trim(mainParts[1], " "), " ")
	parts := make([]int, len(sParts))
	for i, sPart := range sParts {
		part, err := strconv.Atoi(sPart)
		if err != nil {
			return equation{}, err
		}
		parts[i] = part
	}
	return equation{result, parts}, nil
}

func parseInputFile(f *os.File) ([]equation, error) {
	var equations []equation
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		eq, err := parseEquation(s.Text())
		if err != nil {
			return nil, err
		}
		equations = append(equations, eq)
	}
	return equations, nil
}
