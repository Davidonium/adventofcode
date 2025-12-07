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
		r := int(a[0] - b[0])
		if r == 0 {
			r += int(a[1] - b[1])
		}
		return r
	})

	final := [][2]uint64{}
	i := 0
	cursor := ranges[i]
	endWithIntersect := false
	for i < len(ranges)-1 {
		fmt.Printf("(i: %d, cursor: %v, next: %v)", i, cursor, ranges[i+1])
		n, ok := intersect(cursor, ranges[i+1])
		if ok {
			fmt.Printf(" intersect => n: %v ", n)
			cursor = n
			i++
			endWithIntersect = true
		} else {
			fmt.Printf(" ()")
			cursor = ranges[i+1]
			i++
			final = append(final, cursor)
			endWithIntersect = false
		}

		fmt.Print("\n")
	}

	if endWithIntersect {
		// add remaining cursor
		final = append(final, cursor)
	}

	freshCount := uint64(0)
	for _, f := range final {
		freshCount += (f[1] - f[0]) + 1
	}
	fmt.Printf("r=%d\n", freshCount)

	return nil
}

func intersect(a, b [2]uint64) ([2]uint64, bool) {
	if a[1] < b[0] {
		return [2]uint64{}, false
	}
	if a[0] > b[1] {
		return [2]uint64{}, false
	}

	left := min(a[0], b[0])
	right := max(a[1], b[1])
	return [2]uint64{left, right}, true
}

func PairCombinations(slice [][2]uint64) [][2][2]uint64 {
	res := [][2][2]uint64{}
	for i, s := range slice {
		for _, s2 := range slice[i+1:] {
			res = append(res, [2][2]uint64{s, s2})
		}
	}

	return res

}
