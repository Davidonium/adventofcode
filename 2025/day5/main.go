package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
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

	ranges := [][2]uint64{}

	parsingRanges := true
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			parts := strings.Split(t, "-")
			low := util.ParseUInt64(parts[0])
			high := util.ParseUInt64(parts[1])
			ranges = append(ranges, [2]uint64{low, high})
		} else {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	slices.SortFunc(ranges, func(a, b [2]uint64) int {
		return int(a[0] - b[0])
	})

	sum := uint64(0)
	for i := 0; i < len(ranges); i++ {
		left := ranges[i][0]
		right := ranges[i][1]
		for j := i + 1; j < len(ranges) && ranges[j][0] <= right; j++ {
			right = max(ranges[j][1], right)
			i = j
		}
		sum += right - left + 1
	}

	fmt.Printf("r=%d\n", sum)


	return nil
}
