package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fd, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf("failed to read input file: %v\n", err)
		os.Exit(1)
	}

	defer fd.Close()

	if err := run(fd); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run the program: %v", err)
		os.Exit(1)
	}
}

func run(in io.Reader) error {
	scanner := bufio.NewScanner(in)

	g := Grid{}

	i := 0
	for scanner.Scan() {
		t := scanner.Text()

		if len(g.Positions) == 0 {
			g.Positions = [][]Cell{}
		}
		cols := make([]Cell, 0, len(t))

		g.Positions = append(g.Positions, cols)

		for _, v := range t {
			var cell Cell
			switch v {
			case '.':
				cell = CellEmpty
			case '@':
				cell = CellWithPaper
			}

			g.Positions[i] = append(g.Positions[i], cell)
		}
		i++
	}

	total := 0
	for {
		total += g.ReachablePaperCount()
		g.RemoveReachablePaper()
		after := g.ReachablePaperCount()
		if after == 0 {
			break
		}
	}
	fmt.Printf("r=%d\n", total)

	return nil
}

type Cell int

const (
	CellWithPaper Cell = iota + 1
	CellEmpty
)

type Point struct {
	X int
	Y int
}

type Grid struct {
	Positions [][]Cell
}

func (g Grid) AdjacentPaperCount(x, y int) (int, error) {
	if x < 0 || x > len(g.Positions[0]) {
		return 0, fmt.Errorf("out of bounds, got (%d, %d) bounds: (%d, %d)", x, y, len(g.Positions[0]), len(g.Positions))
	}
	if y < 0 || y > len(g.Positions) {
		return 0, fmt.Errorf("out of bounds, got (%d, %d) bounds: (%d, %d)", x, y, len(g.Positions[0]), len(g.Positions))
	}

	var positions []Point

	if y+1 < len(g.Positions) {
		positions = append(positions, Point{X: x, Y: y + 1})
	}
	if y-1 >= 0 {
		positions = append(positions, Point{X: x, Y: y - 1})
	}

	if x+1 < len(g.Positions[0]) {
		positions = append(positions, Point{X: x + 1, Y: y})
		if y+1 < len(g.Positions) {
			positions = append(positions, Point{X: x + 1, Y: y + 1})
		}
		if y-1 >= 0 {
			positions = append(positions, Point{X: x + 1, Y: y - 1})
		}
	}

	if x-1 >= 0 {
		positions = append(positions, Point{X: x - 1, Y: y})
		if y-1 >= 0 {
			positions = append(positions, Point{X: x - 1, Y: y - 1})
		}
		if y+1 < len(g.Positions) {
			positions = append(positions, Point{X: x - 1, Y: y + 1})
		}
	}

	c := 0
	for _, p := range positions {
		if g.Positions[p.Y][p.X] == CellWithPaper {
			c++
		}
	}
	return c, nil
}

func (g Grid) String() string {
	sb := &strings.Builder{}
	for i, row := range g.Positions {
		for j, col := range row {
			switch col {
			case CellEmpty:
				sb.WriteByte('.')
			case CellWithPaper:
				c, _ := g.AdjacentPaperCount(j, i)
				if c < 4 {
					sb.WriteByte('x')
				} else {
					sb.WriteByte('@')
				}
			}
		}

		sb.WriteByte('\n')
	}

	return sb.String()
}

func (g Grid) ReachablePaperPoints() []Point {
	var points []Point
	for i, row := range g.Positions {
		for j, col := range row {
			if col == CellWithPaper {
				c, _ := g.AdjacentPaperCount(j, i)
				if c < 4 {
					points = append(points, Point{X: j, Y: i})
				}
			}
		}
	}

	return points
}

func (g Grid) ReachablePaperCount() int {
	p := g.ReachablePaperPoints()
	return len(p)
}

func (g *Grid) RemoveReachablePaper() {
	pts := g.ReachablePaperPoints()

	for _, p := range pts {
		g.Positions[p.Y][p.X] = CellEmpty
	}
}
