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

	i := 0
	g := Grid{}
	for scanner.Scan() {
		t := scanner.Text()

		g.Cells = append(g.Cells, make([]Cell, 0, len(t)))

		for j, c := range t {
			g.Cells[i] = append(g.Cells[i], Cell(c))
			if Cell(c) == CellStart {
				g.Beams = append(g.Beams, Point{X: j, Y: i})
			}
		}
		i++
	}

	i = 0
	for g.Advance() {
		fmt.Printf("Iteration: %d\n", i)
		fmt.Print(g)
		fmt.Println("")
		i++
	}

	fmt.Printf("r=%d\n", g.Splits)

	return nil
}

type Cell rune

const (
	CellEmpty = '.'
	CellSplit = '^'
	CellStart = 'S'
	CellBeam  = '|'
)

type Point struct {
	X int
	Y int
}

type Grid struct {
	Cells [][]Cell
	Beams []Point
	Splits int
}

func (g *Grid) Advance() bool {
	keepGoing := true
	newBeams := []Point{}
	for _, b := range g.Beams {
		nextY := b.Y + 1
		if nextY < len(g.Cells) {
			switch g.Cells[nextY][b.X] {
			case CellSplit:
				if b.X+1 < len(g.Cells[nextY]) && g.Cells[nextY][b.X+1] != CellBeam {
					newBeams = append(newBeams, Point{Y: nextY, X: b.X + 1})
					g.Cells[nextY][b.X+1] = CellBeam
				}
				if b.X-1 >= 0 && g.Cells[nextY][b.X-1] != CellBeam {
					newBeams = append(newBeams, Point{Y: nextY, X: b.X - 1})
					g.Cells[nextY][b.X-1] = CellBeam
				}

				g.Splits++
			case CellEmpty:
				g.Cells[nextY][b.X] = CellBeam
				b.Y = nextY
				newBeams = append(newBeams, b)
			}
		}

		if nextY == len(g.Cells) {
			keepGoing = false
		}
	}

	g.Beams = newBeams

	return keepGoing
}

func (g Grid) String() string {
	sb := strings.Builder{}
	for _, row := range g.Cells {
		for _, col := range row {
			sb.WriteRune(rune(col))
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}
