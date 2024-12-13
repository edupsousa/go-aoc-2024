package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solver(f *os.File) error {
	r, u := readRulesAndUpdates(f)
	ruleList := rulesToAdjList(r)
	fmt.Println("Rule list:", ruleList)
	validUpd := getValidUpdates(u, ruleList)
	fmt.Println("Valid updates:", validUpd)
	sum := sumMiddlePages(validUpd)
	fmt.Println("Sum of middle pages:", sum)
	return nil
}

func readRulesAndUpdates(f *os.File) ([]string, [][]string) {
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	var rules []string
	var updates [][]string
	section := 0
	for s.Scan() {
		line := s.Text()
		if line == "" {
			section++
			continue
		}
		if section == 0 {
			rules = append(rules, line)
		} else {
			updates = append(updates, strings.Split(line, ","))
		}
	}
	return rules, updates
}

type PageSet = map[string]bool
type RuleAdjList = map[string]map[string]bool

func rulesToAdjList(rules []string) RuleAdjList {
	adjList := make(RuleAdjList)
	for _, r := range rules {
		rParts := strings.Split(r, "|")
		page := rParts[0]
		if _, ok := adjList[page]; !ok {
			adjList[page] = make(map[string]bool)
		}
		adjList[page][rParts[1]] = true
	}
	return adjList
}

func getValidUpdates(updates [][]string, rules RuleAdjList) [][]string {
	var validUpdates [][]string
	for _, u := range updates {
		if isValidUpdate(u, rules) {
			validUpdates = append(validUpdates, u)
		}
	}
	return validUpdates
}

func isValidUpdate(update []string, rules RuleAdjList) bool {
	seen := make(map[string]bool)
	for _, page := range update {
		seen[page] = true
		succesors, ok := rules[page]
		if !ok {
			continue
		}
		// fmt.Println("Page:", page, "Succesors:", succesors)
		for s := range succesors {
			if seen[s] {
				return false
			}
		}
	}
	return true
}

func sumMiddlePages(updates [][]string) int {
	sum := 0
	for _, u := range updates {
		i := len(u) / 2
		mid, _ := strconv.Atoi(u[i])
		sum += mid
	}
	return sum
}
