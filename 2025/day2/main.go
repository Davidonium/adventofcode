package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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
				n := strconv.Itoa(i)
				if len(n) % 2 != 0 {
					continue
				}

				mid := len(n)/2
				left := n[:mid]
				right := n[mid:]

				// fmt.Printf("cmp: %d - %s == %s\n", i, left, right)
				if left == right {
					invalidIDs = append(invalidIDs, i)
				}
			}
		}
	}

	n := util.SumSlice(invalidIDs)

	fmt.Printf("Invalid IDs sum: %d\n", n)

	return nil
}
