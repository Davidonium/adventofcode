package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"github.com/davidonium/adventofcode/util"
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

	invalidIDs := []int{}
	for scanner.Scan() {
		t := scanner.Text()

		ranges := strings.SplitSeq(t, ",")
		for r := range ranges {
			parts := strings.Split(r, "-")

			low := util.ParseInt(parts[0])
			high := util.ParseInt(parts[1])

			for i := low; i <= high; i++ {
				d := util.DigitCount(i)
				if d%2 != 0 {
					continue
				}

				mid := d / 2
				div := int(math.Pow10(mid))
				left :=  i / div
				right := i % div
				fmt.Printf("cmp: %d - (mid: %d) %d == %d", i, mid, left, right)
				if left == right {
					fmt.Print(" (!)")
					invalidIDs = append(invalidIDs, i)
				}
				fmt.Print("\n")
			}
		}
	}

	n := util.SumSlice(invalidIDs)

	fmt.Printf("Invalid IDs sum: %d\n", n)

	return nil
}
