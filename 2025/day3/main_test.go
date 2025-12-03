package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestChallenge(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			name: "example",
			input: `
987654321111111
811111111111119
234234234234278
818181911112111
			`,
			expect: "r=357",
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

			if !strings.Contains(string(output), tt.expect) {
				t.Errorf("wanted %s in the output, got\n%s\n", tt.expect, string(output))
			}
		})
	}
}
