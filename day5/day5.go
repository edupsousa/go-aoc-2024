package day5

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solver(f *os.File) error {
	r, u := readRulesAndUpdates(f)
	ruleList := rulesToAdjList(r)
	// fmt.Println("Rule list:", ruleList)
	valid, invalid := getValidUpdates(u, ruleList)
	// fmt.Println("Valid updates:", valid)
	// fmt.Println("Invalid updates:", invalid)
	reorderedInvalid := reorderInvalidUpdates(invalid, ruleList)
	// fmt.Println("Reordered invalid updates:", reorderedInvalid)
	sumValid := sumMiddlePages(valid)
	fmt.Println("Middle page sum for valid updates:", sumValid)
	sumReordered := sumMiddlePages(reorderedInvalid)
	fmt.Println("Middle page sum for reordered invalid updates:", sumReordered)
	return nil
}

type ruleAdjList = map[string]map[string]bool

type sortableUpdate struct {
	update []string
	rules  ruleAdjList
}

func (su sortableUpdate) Len() int {
	return len(su.update)
}

func (su sortableUpdate) Swap(i, j int) {
	su.update[i], su.update[j] = su.update[j], su.update[i]
}

func (su sortableUpdate) Less(i, j int) bool {
	return su.rules[su.update[i]][su.update[j]]
}

func reorderInvalidUpdates(updates [][]string, rules ruleAdjList) [][]string {
	for _, update := range updates {
		sort.Sort(sortableUpdate{update, rules})
	}
	return updates
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

func rulesToAdjList(rules []string) ruleAdjList {
	adjList := make(ruleAdjList)
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

func getValidUpdates(updates [][]string, rules ruleAdjList) ([][]string, [][]string) {
	var valid [][]string
	var invalid [][]string
	for _, u := range updates {
		if isValidUpdate(u, rules) {
			valid = append(valid, u)
		} else {
			invalid = append(invalid, u)
		}
	}
	return valid, invalid
}

func isValidUpdate(update []string, rules ruleAdjList) bool {
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
