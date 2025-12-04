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

	fmt.Printf("The parsed grid:\n%s\n", g)

	return nil
}

type Cell int

const (
	CellWithPaper Cell = iota + 1
	CellEmpty
)

type Grid struct {
	Positions [][]Cell
}

func (g Grid) String() string {
	sb := &strings.Builder{}
	for _, row := range g.Positions {
		for _, col := range row {

			switch col {
			case CellEmpty:
				sb.WriteString(".")
			case CellWithPaper:
				sb.WriteString("@")
			}
		}

		sb.WriteString("\n")
	}

	return sb.String()
}
