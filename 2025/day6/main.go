package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
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

	var numLines []string
	var ops []struct {
		op  string
		idx int
	}

	maxLineLength := 0
	r := regexp.MustCompile(`\d+`)
	opR := regexp.MustCompile(`(\+|\*)`)
	for scanner.Scan() {
		t := scanner.Text()
		ln := r.FindAllString(t, -1)

		if len(ln) > 0 {
			numLines = append(numLines, t)
		} else {
			idxs := opR.FindAllStringIndex(t, -1)
			for _, a := range idxs {
				ops = append(ops, struct {
					op  string
					idx int
				}{
					op:  t[a[0]:a[1]],
					idx: a[0],
				})
			}
			maxLineLength = len(t)
		}
	}

	sum := 0
	for i := range ops {
		finalIdx := maxLineLength
		if i < len(ops)-1 {
			// subtract 1 because that's the separation between columns
			// each operation index marks the start of a column
			finalIdx = ops[i+1].idx - 1
		}

		colLen := finalIdx - ops[i].idx

		colNums := make([]string, 0, len(numLines))

		for _, nl := range numLines {
			colNums = append(colNums, nl[ops[i].idx:finalIdx])
		}

		colSum := 0
		// since the number is multiplying itself, it's easier to start at 1
		if ops[i].op == "*" {
			colSum = 1
		}

		for j := range colLen {
			num := 0
			factor := 0
			for z := len(colNums) - 1; z >= 0; z-- {
				n := colNums[z]
				if n[j] == ' ' {
					continue
				}
				// transform ascii byte to integer
				d := int(n[j]-'0')

				factored := d * int(math.Pow10(factor))
				// factor only increases when encountering an actual number
				factor++

				num += factored
			}


			switch ops[i].op {
			case "+":
				colSum += num
			case "*":
				colSum *= num
			}
		}

		sum += colSum
	}

	fmt.Printf("r=%d\n", sum)

	return nil
}
