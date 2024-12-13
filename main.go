package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/edupsousa/go-aoc-2024/day1"
	"github.com/edupsousa/go-aoc-2024/day2"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <day> <input-file>")
		os.Exit(1)
	}
	day, err := strconv.Atoi(os.Args[1])
	if err != nil || day < 1 || day > 25 {
		fmt.Println("Invalid day")
		os.Exit(1)
	}
	inputPath := os.Args[2]
	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()
	solverFn := getDaySolver(day)
	err = solverFn(file)
	if err != nil {
		fmt.Println("Error solving day:", err)
		os.Exit(1)
	}
}

type SolverFunc func(*os.File) error

func getDaySolver(day int) SolverFunc {
	switch day {
	case 1:
		return day1.Solver
	case 2:
		return day2.Solver
	default:
		return func(_ *os.File) error {
			fmt.Println("Day not implemented yet")
			return nil
		}
	}
}
