package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename, err := getFilenameFromArgs()
	if err != nil {
		panic(err)
	}
	reports, err := readReportsFromFile(filename)
	if err != nil {
		panic(err)
	}
	safeReports := countSafeReports(reports, isSafe)
	fmt.Printf("Number of safe reports: %d\n", safeReports)
	safeWithTolerance := countSafeReports(reports, isSafeWithTolerance)
	fmt.Printf("Number of safe reports with tolerance: %d\n", safeWithTolerance)
}

func getFilenameFromArgs() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("filename not provided")
	}
	filename := os.Args[1]
	return filename, nil
}

type Level = int
type Report []Level

func parseReport(line string) (Report, error) {
	sLevels := strings.Split(line, " ")
	var report Report
	for _, sLevel := range sLevels {
		level, err := strconv.Atoi(sLevel)
		if err != nil {
			return nil, err
		}
		report = append(report, level)
	}
	return report, nil
}

func readReportsFromFile(filename string) ([]Report, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var reports []Report
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

type CheckSafetyFunc func(Report) bool

func countSafeReports(reports []Report, checkSafety CheckSafetyFunc) int {
	var safeReports int
	for _, report := range reports {
		if checkSafety(report) {
			safeReports++
		}
	}
	return safeReports
}

func isSafe(r Report) bool {
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

func isSafeWithTolerance(report Report) bool {
	if len(report) < 2 {
		return true
	}
	if isSafe(report) {
		return true
	}
	for i := 0; i < len(report); i++ {
		r := make(Report, len(report))
		copy(r, report)
		r = append(r[:i], r[i+1:]...)
		if isSafe(r) {
			return true
		}
	}
	return false
}
