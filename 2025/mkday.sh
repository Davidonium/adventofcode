#!/usr/bin/env bash

set -euo pipefail


if [[ -z $1 ]] 
then
    echo "please provide a day number"
    exit
fi

DAY_DIR=day$1

mkdir $DAY_DIR

touch $DAY_DIR/input.txt

tee -a $DAY_DIR/main.go << END
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

	for scanner.Scan() {
		t := scanner.Text()
        // TODO start coding!
	}

    return nil
}
END

tee -a $DAY_DIR/main_test.go << END
package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestChallenge(t *testing.T) {
	tests := []struct {
		name        string
		input       string
	}{
        {
            // TODO test case
        },
    }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := os.Stdout

			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("failed to create os.Pipe: %v", err)
			}

			os.Stdout = w

			in := strings.NewReader(strings.TrimSpace(tt.input))

			if err := run(in); err != nil {
				w.Close()
				os.Stdout = stdout
				t.Fatalf("failed to run the program")
			}

			w.Close()
			os.Stdout = stdout

			output, err := io.ReadAll(r)
			if err != nil {
				t.Fatalf("could not read from output: %v", err)
			}

            // TODO test the thing
		})
	}
}

END
