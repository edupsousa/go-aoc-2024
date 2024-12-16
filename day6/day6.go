package day6

import (
	"bufio"
	"fmt"
	"os"
)

type direction uint8

const (
	up direction = iota
	right
	down
	left
)

func (d direction) turnRight() direction {
	return (d + 1) % 4
}

type movement struct {
	dX, dY int
}

var movements = map[direction]movement{
	up:    {0, -1},
	right: {1, 0},
	down:  {0, 1},
	left:  {-1, 0},
}

type tileStatus uint8

func (ts tileStatus) String() string {
	if ts == emptyTile {
		return "."
	}
	if ts == obstacleTile {
		return "#"
	}
	if ts == upTile {
		return "^"
	}
	if ts == rightTile {
		return ">"
	}
	if ts == downTile {
		return "v"
	}
	if ts == leftTile {
		return "<"
	}
	if (ts&upTile == upTile || ts&downTile == downTile) && (ts&rightTile == rightTile || ts&leftTile == leftTile) {
		return "+"
	}
	if ts&upTile == upTile || ts&downTile == downTile {
		return "|"
	}
	if ts&rightTile == rightTile || ts&leftTile == leftTile {
		return "-"
	}
	return "?"
}

func (ts *tileStatus) setMovement(d direction) {
	switch d {
	case up:
		*ts |= upTile
	case right:
		*ts |= rightTile
	case down:
		*ts |= downTile
	case left:
		*ts |= leftTile
	}
}

func (ts *tileStatus) setObstacle() {
	*ts |= obstacleTile
}

func (ts *tileStatus) clearObstacle() {
	*ts &= ^obstacleTile
}

const emptyTile tileStatus = 0

const (
	obstacleTile tileStatus = 1 << iota
	upTile
	rightTile
	downTile
	leftTile
)

type tileMap [][]tileStatus

func (tm tileMap) isInBounds(p position) bool {
	return p.y >= 0 && p.y < len(tm) && p.x >= 0 && p.x < len(tm[p.y])
}

func (tm *tileMap) setMovement(p position, d direction) {
	(*tm)[p.y][p.x].setMovement(d)
}

func (tm *tileMap) setObstacle(p position) {
	(*tm)[p.y][p.x].setObstacle()
}

func (tm *tileMap) clearObstacle(p position) {
	(*tm)[p.y][p.x].clearObstacle()
}

func (tm tileMap) hasObstacle(p position) bool {
	return tm[p.y][p.x]&obstacleTile == obstacleTile
}

func (tm tileMap) countMovementTiles() int {
	count := 0
	for _, row := range tm {
		for _, tile := range row {
			if tile != emptyTile && tile != obstacleTile {
				count++
			}
		}
	}
	return count
}

func (tm tileMap) String() string {
	var s string
	for _, row := range tm {
		for _, tile := range row {
			s += tile.String()
		}
		s += "\n"
	}
	return s
}

type position struct {
	x, y int
}

func (p position) move(d direction) position {
	return position{p.x + movements[d].dX, p.y + movements[d].dY}
}

type player struct {
	pos position
	dir direction
}

type game struct {
	tm    tileMap
	guard player
}

func (g *game) moveGuard() bool {
	newPos := g.guard.pos.move(g.guard.dir)
	fmt.Printf("Moving guard from %v to %v\n", g.guard.pos, newPos)
	if !g.tm.isInBounds(newPos) {
		return false
	}
	if g.tm.hasObstacle(newPos) {
		g.guard.dir = g.guard.dir.turnRight()
		return g.moveGuard()
	}
	g.tm.setMovement(newPos, g.guard.dir)
	g.guard.pos = newPos
	return true
}

func createGameFromInput(f *os.File) game {
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	var tm tileMap
	var start position
	for s.Scan() {
		line := s.Text()
		var row []tileStatus
		for _, c := range line {
			switch c {
			case '.':
				row = append(row, emptyTile)
			case '#':
				row = append(row, obstacleTile)
			case '^':
				start = position{len(row), len(tm)}
				row = append(row, upTile)
			}
		}
		tm = append(tm, row)
	}
	return game{
		tm: tm,
		guard: player{
			pos: start,
			dir: up,
		},
	}
}

func Solver(f *os.File) error {
	game := createGameFromInput(f)
	fmt.Println(game.tm.String())
	for game.moveGuard() {
	}
	fmt.Println(game.tm.String())
	fmt.Printf("Number of movement tiles: %d\n", game.tm.countMovementTiles())
	return nil
}
