package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	filename, err := readFilename()
	if err != nil {
		panic(err)
	}
	fmt.Println("Reading file", filename)
	lHeap, rHeap, rMap, err := readInput(filename)
	if err != nil {
		panic(err)
	}
	sScore := similarityScore(lHeap, rMap)
	fmt.Println("Similarity score is", sScore)
	diffSum := sumDifferences(lHeap, rHeap)
	fmt.Println("Sum of differences is", diffSum)
}

func similarityScore(lHeap *MinHeap, rMap *map[int]int) int {
	var score int
	for _, l := range *lHeap {
		score += l * (*rMap)[l]
	}
	return score
}

func sumDifferences(lHeap *MinHeap, rHeap *MinHeap) int {
	var sum int
	var l int
	var r int
	for lHeap.Len() > 0 {
		l = heap.Pop(lHeap).(int)
		r = heap.Pop(rHeap).(int)
		sum += int(math.Abs(float64(l - r)))
	}
	return sum
}

func readFilename() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("filename not provided")
	}
	filename := os.Args[1]
	return filename, nil
}

func readInput(filename string) (*MinHeap, *MinHeap, *map[int]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, nil, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	lHeap := make(MinHeap, 0)
	rHeap := make(MinHeap, 0)
	rMap := make(map[int]int)

	for fileScanner.Scan() {
		var l int
		var r int
		_, err := fmt.Sscanf(fileScanner.Text(), "%d   %d", &l, &r)
		if err != nil {
			return nil, nil, nil, err
		}
		lHeap = append(lHeap, l)
		rHeap = append(rHeap, r)
		rMap[r] = rMap[r] + 1
	}
	heap.Init(&lHeap)
	heap.Init(&rHeap)
	return &lHeap, &rHeap, &rMap, nil
}
