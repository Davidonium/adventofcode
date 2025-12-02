package main

import (
	"bufio"
	"fmt"
	"io"
	"maps"
	"math"
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

	invalidIDs := []int{}
	for scanner.Scan() {
		t := scanner.Text()

		ranges := strings.SplitSeq(t, ",")
		for r := range ranges {
			rangeInvalidIDs := map[int]struct{}{}
			parts := strings.Split(r, "-")

			low := util.ParseInt(parts[0])
			high := util.ParseInt(parts[1])

			for i := low; i <= high; i++ {
				d := util.DigitCount(i)

				factors := []int{}
				sq := int(math.Sqrt(float64(d)))
				for j := 1; j <= sq; j++ {
					if d%j == 0 {
						div := d / j

						if div == j {
							// example 9 / 3 = 3, gotta add 3
							factors = append(factors, j)
						} else {
							// example 10 / 2 = 5, gotta add 2 and 5
							factors = append(factors, j)

							// not interested in the number itself on this exercise
							// example 10 / 1 = 10, not adding it (but adding 1 in the previous line)
							if div != d {
								factors = append(factors, div)
							}
						}
					}
				}

				for _, div := range factors {
					nums := map[int]struct{}{}
					for j := d - div; j >= 0; j -= div {
						// extract exact digits,
						// example:
						// 		i = 123456788 where div is 3 d is 9 and j is at 6
						// 		123456789 / 10^6 = 123456
						// 		123456 % 10^3 = 456
						//
						// 		i = 123456788 where div is 3 d is 9 and j is at 3
						// 		123456789 / 10^3 = 123
						// 		123 % 10^3 = 123
						p := int(math.Pow10(j))
						l := i / p
						n := l % int(math.Pow10(div))
						nums[n] = struct{}{}
					}

					if len(nums) == 1 {
						rangeInvalidIDs[i] = struct{}{}
					}
				}
			}

			ids := slices.Collect(maps.Keys(rangeInvalidIDs))
			fmt.Printf("%d-%d: %v\n", low, high, ids)

			invalidIDs = append(invalidIDs, ids...)
		}
	}

	n := util.SumSlice(invalidIDs)

	fmt.Printf("Invalid IDs sum: %d\n", n)

	return nil
}

