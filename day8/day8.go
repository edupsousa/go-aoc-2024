package day8

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func Solver(f *os.File) error {
	am := parseInputMap(f)
	_, distinct := am.getAntinodes()
	// fmt.Printf("%+v\n", an)
	fmt.Printf("Distinct antinodes: %d\n", distinct)
	return nil
}

type position struct {
	x, y float64
}

type line struct {
	p1, p2 position
}

func (l line) getPointEquiDistantFromStart() position {
	return position{
		l.p1.x + (l.p2.x-l.p1.x)*-1,
		l.p1.y + (l.p2.y-l.p1.y)*-1,
	}
}

func (l line) getPointEquiDistantFromEnd() position {
	return position{
		l.p2.x + (l.p1.x-l.p2.x)*-1,
		l.p2.y + (l.p1.y-l.p2.y)*-1,
	}
}

type antennasMap struct {
	width    int
	height   int
	antennas map[rune][]position
}

func (am antennasMap) getAntinodes() (map[rune][]position, int) {
	posSet := make(map[position]bool)
	an := make(map[rune][]position)
	for k, v := range am.antennas {
		an[k] = make([]position, 0)
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				ans := getAntinodes(v[i], v[j])
				for _, a := range ans {
					if a.x >= 0 && a.x < float64(am.width) && a.y >= 0 && a.y < float64(am.height) {
						an[k] = append(an[k], a)
						posSet[a] = true
					}
				}
			}
		}
	}
	return an, len(posSet)
}

func getAntinodes(a position, b position) []position {
	l := line{a, b}
	an1 := l.getPointEquiDistantFromStart()
	an2 := l.getPointEquiDistantFromEnd()
	return []position{an1, an2}
}

func parseInputMap(f *os.File) antennasMap {
	r := bufio.NewReader(f)
	antennas := make(map[rune][]position)
	y := 0
	x := 0
	width := 0
	height := 0
	for {
		if c, _, err := r.ReadRune(); err != nil {
			break
		} else {
			if c == '\n' {
				y++
				x = 0
			} else {
				if c != '.' {
					if _, ok := antennas[c]; !ok {
						antennas[c] = make([]position, 0)
					}
					antennas[c] = append(antennas[c], position{float64(x), float64(y)})
				}
				height = int(math.Max(float64(height), float64(y+1)))
				width = int(math.Max(float64(width), float64(x+1)))
				x++
			}
		}
	}
	return antennasMap{width, height, antennas}
}
