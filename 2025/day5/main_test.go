package main

import (
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestChallenge(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "example",
			input: `
3-5
10-14
16-20
12-18

1
5
8
11
17
32`,
			expected: "r=14",
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

			ob, err := io.ReadAll(r)
			if err != nil {
				t.Fatalf("could not read from output: %v", err)
			}

			o := string(ob)
			if !strings.Contains(o, tt.expected) {
				t.Errorf("expected %s to be in the output:\n%s\n", tt.expected, o)
			}
		})
	}
}

func TestPairCombinations(t *testing.T) {
	i := [][2]uint64{{3, 5}, {10, 14}, {16, 20}, {12, 18}}
	c := PairCombinations(i)

	expect := [][2][2]uint64{
		{{3, 5}, {10, 14}}, {{3, 5}, {16, 20}}, {{3, 5}, {12, 18}},
		{{10, 14}, {16, 20}}, {{10, 14}, {12, 18}},
		{{16, 20}, {12, 18}},
	}
	if !reflect.DeepEqual(c, expect) {
		t.Errorf("input:\n\t%v\ngot:\n\t%v\nwant:\n\t%v\n", i, c, expect)
	}
}
