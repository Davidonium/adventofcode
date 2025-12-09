package main

import (
	"io"
	"os"
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
.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`,
			expected: "r=21",
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
				t.Errorf("failure\nwant:\n%s\ngot:\n%s\n", tt.expected, o)
			}
		})
	}
}
