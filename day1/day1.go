package day1

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func similarityScore(lHeap *minHeap, rMap *map[int]int) int {
	var score int
	for _, l := range *lHeap {
		score += l * (*rMap)[l]
	}
	return score
}

func sumDifferences(lHeap *minHeap, rHeap *minHeap) int {
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

func readInput(file *os.File) (*minHeap, *minHeap, *map[int]int, error) {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	lHeap := make(minHeap, 0)
	rHeap := make(minHeap, 0)
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

func Solver(file *os.File) error {
	lHeap, rHeap, rMap, err := readInput(file)
	if err != nil {
		return err
	}
	sScore := similarityScore(lHeap, rMap)
	fmt.Println("Similarity score is", sScore)
	diffSum := sumDifferences(lHeap, rHeap)
	fmt.Println("Sum of differences is", diffSum)
	return nil
}
