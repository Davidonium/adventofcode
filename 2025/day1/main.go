package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/davidonium/adventofcode/util"
)

const safeMaxNumber = 100

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

	pos := 50
	safePwd := 0

	fmt.Printf("at %d\n", pos)

	for scanner.Scan() {
		t := scanner.Text()
		dir := t[0]
		rc := t[1:]
		c := util.ParseInt(rc)

		clicked := false
		switch dir {
		case 'R':
			fmt.Printf("R%d", c)
			for range c {
				pos++
				if pos > safeMaxNumber-1 {
					pos = 0
					safePwd++
					clicked = true
				}
			}
		case 'L':
			fmt.Printf("L%d", c)
			for range c {
				pos--
				if pos < 0 {
					pos = safeMaxNumber - 1
				} else if pos == 0 {
					safePwd++
					clicked = true
				}
			}

		default:
			panic(fmt.Sprintf("unknown direction '%s', expected 'R' or 'L'", string(dir)))
		}

		fmt.Printf(" (clicked: %t) => %d\n", clicked, pos)
	}

	fmt.Printf("\n\n\tSafe password is %d\n", safePwd)

	return nil
}
