package day3

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func Solver(file *os.File) error {
	data, err := readFile(file)
	if err != nil {
		return err
	}
	re := regexp.MustCompile(`(mul)\(([0-9]{1,3}),([0-9]{1,3})\)|(do)\(\)|(don)'t\(\)`)
	matches := re.FindAllStringSubmatch(data, -1)
	total := 0
	enabled := true
	for _, match := range matches {
		if match[4] == "do" {
			enabled = true
		} else if match[5] == "don" {
			enabled = false
		} else if enabled {
			a, err := strconv.Atoi(match[2])
			if err != nil {
				return err
			}
			b, err := strconv.Atoi(match[3])
			if err != nil {
				return err
			}
			total += a * b
		}
	}
	fmt.Printf("Total: %d\n", total)
	return nil
}

func readFile(file *os.File) (string, error) {
	var data string
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		data += string(line)
	}
	return data, nil
}
