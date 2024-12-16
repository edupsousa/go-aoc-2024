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
	// if ts == upTile {
	// 	return "^"
	// }
	// if ts == rightTile {
	// 	return ">"
	// }
	// if ts == downTile {
	// 	return "v"
	// }
	// if ts == leftTile {
	// 	return "<"
	// }
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

func (ts tileStatus) hasMovement(d direction) bool {
	return (d == up && ts&upTile == upTile) ||
		(d == right && ts&rightTile == rightTile) ||
		(d == down && ts&downTile == downTile) ||
		(d == left && ts&leftTile == leftTile)
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

func (tm tileMap) hasMovement(p position, d direction) bool {
	return tm[p.y][p.x].hasMovement(d)
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

type movementResult uint8

const (
	movementOk movementResult = iota
	movementOutOfBounds
	movementLoop
)

func (g *game) moveGuard() movementResult {
	newPos := g.guard.pos.move(g.guard.dir)
	// fmt.Printf("Moving guard from %v to %v\n", g.guard.pos, newPos)
	if !g.tm.isInBounds(newPos) {
		return movementOutOfBounds
	}
	if g.tm.hasMovement(newPos, g.guard.dir) {
		// fmt.Printf("Guard is in a loop at %v\n", newPos)
		return movementLoop
	}
	if g.tm.hasObstacle(newPos) {
		g.guard.dir = g.guard.dir.turnRight()
		return g.moveGuard()
	}
	g.tm.setMovement(newPos, g.guard.dir)
	g.guard.pos = newPos
	return movementOk
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

type part1Solver struct {
	game game
}

func (s *part1Solver) Solve() error {
	// fmt.Println(s.game.tm.String())
	for {
		result := s.game.moveGuard()
		if result == movementOutOfBounds {
			break
		}
		if result == movementLoop {
			return fmt.Errorf("(part1) guard is in a loop at %v", s.game.guard.pos)
		}
	}
	// fmt.Println(s.game.tm.String())
	fmt.Printf("Number of movement tiles: %d\n", s.game.tm.countMovementTiles())
	return nil
}

func (s part1Solver) getMovementPositions() []position {
	var positions []position
	for y, row := range s.game.tm {
		for x, tile := range row {
			if tile != emptyTile && tile != obstacleTile {
				positions = append(positions, position{x, y})
			}
		}
	}
	return positions
}

type part2Solver struct {
	baseGame   game
	candidates []position
}

func (s *part2Solver) Solve() error {
	var results []position
	startPos := s.baseGame.guard.pos
	for _, candidate := range s.candidates {
		if candidate == startPos {
			continue
		}
		game := copyGame(s.baseGame)
		game.tm.setObstacle(candidate)
		for {
			r := game.moveGuard()
			if r == movementOutOfBounds {
				break
			}
			if r == movementLoop {
				// fmt.Println(game.tm.String())
				results = append(results, candidate)
				break
			}
		}
	}
	fmt.Printf("Number of loops: %d\n", len(results))
	return nil
}

func copyGame(g game) game {
	tm := make(tileMap, len(g.tm))
	for y, row := range g.tm {
		tm[y] = make([]tileStatus, len(row))
		copy(tm[y], row)
	}
	return game{
		tm: tm,
		guard: player{
			pos: g.guard.pos,
			dir: g.guard.dir,
		},
	}
}

func Solver(f *os.File) error {
	baseGame := createGameFromInput(f)
	p1Solver := part1Solver{game: copyGame(baseGame)}
	err := p1Solver.Solve()
	if err != nil {
		return err
	}
	candidates := p1Solver.getMovementPositions()
	p2Solver := part2Solver{
		baseGame:   copyGame(baseGame),
		candidates: candidates,
	}
	return p2Solver.Solve()

}
