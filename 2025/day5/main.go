package main

import (
	"bufio"
	"fmt"
	"io"
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

	ranges := [][2]uint64{}
	ingredients := []uint64{}

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
			ingredients = append(ingredients, util.ParseUInt64(t))
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	freshCount := 0
	for _, ing := range ingredients {
		for _, r := range ranges {
			if ing >= r[0] && ing <= r[1] {
				freshCount++
				break
			}
		}
	}

	fmt.Printf("r=%d\n", freshCount)

	return nil
}

