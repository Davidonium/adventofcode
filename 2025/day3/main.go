package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

const totalBatteriesToActivate = 12

func run(in io.Reader) error {
	scanner := bufio.NewScanner(in)

	var banks []Bank
	for scanner.Scan() {
		t := scanner.Text()
		bank := Bank{}
		for _, b := range t {
			bank.Values = append(bank.Values, util.ParseInt(string(b)))
		}

		banks = append(banks, bank)
	}
	var values []int

	for _, b := range banks {
		value := 0
		cursor := 0
		room := len(b.Values) - totalBatteriesToActivate
		for i := range totalBatteriesToActivate {
			from := cursor
			until := from + room + 1

			max := b.Values[from]
			for j := from; j < until; j++ {
				if b.Values[j] > max {
					max = b.Values[j]
					cursor = j
				}
			}
			room -= cursor - from
			if cursor < len(b.Values)-1 {
				cursor++
			}

			value += max * int(math.Pow10(totalBatteriesToActivate-i-1))
		}

		values = append(values, value)
	}

	fmt.Printf("r=%d\n", util.SumSlice(values))

	return nil
}

type Bank struct {
	Values []int
}
