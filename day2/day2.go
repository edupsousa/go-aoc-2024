package day2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseReport(line string) ([]int, error) {
	sLevels := strings.Split(line, " ")
	var report []int
	for _, sLevel := range sLevels {
		level, err := strconv.Atoi(sLevel)
		if err != nil {
			return nil, err
		}
		report = append(report, level)
	}
	return report, nil
}

func readReportsFromFile(file *os.File) ([][]int, error) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var reports [][]int
	for scanner.Scan() {
		line := scanner.Text()
		report, err := parseReport(line)
		if err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}
	return reports, nil
}

type checkSafetyFunc func([]int) bool

func countSafeReports(reports [][]int, checkSafety checkSafetyFunc) int {
	var safeReports int
	for _, report := range reports {
		if checkSafety(report) {
			safeReports++
		}
	}
	return safeReports
}

func isSafe(r []int) bool {
	if len(r) < 2 {
		return true
	}
	inc := r[1] > r[0]
	for i := 0; i < len(r)-1; i++ {
		if (r[i+1] > r[i]) != inc {
			return false
		}
		diff := math.Abs(float64(r[i+1] - r[i]))
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func isSafeWithTolerance(report []int) bool {
	if len(report) < 2 {
		return true
	}
	if isSafe(report) {
		return true
	}
	for i := 0; i < len(report); i++ {
		r := make([]int, len(report))
		copy(r, report)
		r = append(r[:i], r[i+1:]...)
		if isSafe(r) {
			return true
		}
	}
	return false
}

func Solver(file *os.File) error {
	reports, err := readReportsFromFile(file)
	if err != nil {
		return err
	}
	safeReports := countSafeReports(reports, isSafe)
	fmt.Printf("Number of safe reports: %d\n", safeReports)
	safeWithTolerance := countSafeReports(reports, isSafeWithTolerance)
	fmt.Printf("Number of safe reports with tolerance: %d\n", safeWithTolerance)
	return nil
}
