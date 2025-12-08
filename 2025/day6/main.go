package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"

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

	var nums [][]int
	var ops []string

	r := regexp.MustCompile(`(\d+)`)
	opR := regexp.MustCompile(`(\+|\*)`)
	for scanner.Scan() {
		t := scanner.Text()
		ln := r.FindAllString(t, -1)

		if len(ln) > 0 {
			var all []int
			for _, raw := range ln {
				all = append(all, util.ParseInt(raw))
			}
			nums = append(nums, all)
		} else {
			lo := opR.FindAllString(t, -1)
			for _, o := range lo {
				ops = append(ops, o)
			}
		}
	}

	sum := 0
	for i := 0; i < len(ops); i++ {
		col := 0
		if ops[i] == "*" {
			col = 1
		}
		for _, n := range nums {
			switch ops[i] {
			case "+":
				col += n[i]
			case "*":
				col *= n[i]
			}
		}

		sum += col
	}

	fmt.Printf("r=%d\n", sum)

	return nil
}
