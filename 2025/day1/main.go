package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/davidonium/adventofcode/util"
)

const safeCount = 100

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

		switch dir {
		case 'R':
			fmt.Printf("R%d", c)
			new := pos + c
			if new+1 > safeCount {
				pos = new % safeCount
				times := (c / safeCount)
				if times == 0 {
					times = 1
				}
				// safePwd += times

				fmt.Printf(" (new: %d, times: %d) ", new, times)
			} else {
				pos = new
			}
		case 'L':
			fmt.Printf("L%d", c)
			new := pos - c
			if new < 0 {
				pos = (new % safeCount)
				if pos != 0 {
					pos += safeCount
				}
				times := (c / safeCount)
				if times == 0 {
					times = 1
				}
				// safePwd += times

				fmt.Printf(" (new: %d, times: %d) ", new, times)
			} else {
				pos = new
			}
		default:
			panic(fmt.Sprintf("unknown direction '%s', expected 'R' or 'L'", string(dir)))
		}

		if pos == 0 {
			safePwd++
		}

		fmt.Printf(" => %d\n", pos)
	}

	fmt.Printf("\n\n\tSafe password is %d\n", safePwd)

	return nil
}
