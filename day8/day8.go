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

func (l line) getPointEquiDistantFromStart(multiplier int) position {
	return position{
		l.p1.x + (l.p2.x-l.p1.x)*-float64(multiplier),
		l.p1.y + (l.p2.y-l.p1.y)*-float64(multiplier),
	}
}

func (l line) getPointEquiDistantFromEnd(multiplier int) position {
	return position{
		l.p2.x + (l.p1.x-l.p2.x)*-float64(multiplier),
		l.p2.y + (l.p1.y-l.p2.y)*-float64(multiplier),
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
				ans := getAntinodes(v[i], v[j], am.width, am.height)
				for _, a := range ans {
					an[k] = append(an[k], a)
					posSet[a] = true
				}
			}
		}
	}
	return an, len(posSet)
}

func isInBounds(p position, w, h int) bool {
	return p.x >= 0 && p.x < float64(w) && p.y >= 0 && p.y < float64(h)
}

func getAntinodes(a position, b position, w, h int) []position {
	var ans []position
	l := line{a, b}
	i := 0
	for {
		p := l.getPointEquiDistantFromStart(i)
		if !isInBounds(p, w, h) {
			break
		}
		ans = append(ans, p)
		i++
	}
	i = 0
	for {
		p := l.getPointEquiDistantFromEnd(i)
		if !isInBounds(p, w, h) {
			break
		}
		ans = append(ans, p)
		i++
	}

	return ans
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
