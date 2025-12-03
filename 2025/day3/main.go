package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

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

	var banks []Bank
	for scanner.Scan() {
		t := scanner.Text()
		bank := Bank{}
		for _, b := range t {
			bank.Values = append(bank.Values, byte(util.ParseInt(string(b))))
		}

		banks = append(banks, bank)
	}
	var values []int

	for _, b := range banks {
		firstIdx := 0
		for j := 0; j < len(b.Values); j++ {
			for z := firstIdx+1; z < len(b.Values); z++ {
				first := b.Values[firstIdx]
				vin := b.Values[z]

				if vin > first && z < len(b.Values)-1 {
					firstIdx = z
					break
				}
			}
		}

		secondIdx := firstIdx+1
		for j := firstIdx+1; j < len(b.Values); j++ {
			v := b.Values[j]
			if v > b.Values[secondIdx] {
				secondIdx = j
			}
		}

		values = append(values, int(b.Values[firstIdx])*10 + int(b.Values[secondIdx]))
	}

	fmt.Printf("r=%d\n", util.SumSlice(values))

	return nil
}

type Bank struct {
	Values []byte
}
