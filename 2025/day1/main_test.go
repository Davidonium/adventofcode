package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestDay1Part1(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedPwd string
	}{
		{
			name: "example",
			input: `
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`,
			expectedPwd: "3",
		},
		{
			name: "single overflow right",
			input: `
R88
`,
			expectedPwd: "0",
		},
		{
			name: "single overflow left",
			input: `
L59
`,
			expectedPwd: "0",
		},
		{
			name: "multiple overflow right",
			input: `
R200
R2
`,
			expectedPwd: "0",
		},
		{
			name: "multiple overflow right, land in 0",
			input: `
R750
`,
			expectedPwd: "1",
		},
		{
			name: "multiple overflow left",
			input: `
L100
L20
`,
			expectedPwd: "0",
		},
		{
			name: "multiple overflow left, land in 0",
			input: `
L750
`,
			expectedPwd: "1",
		},
		{
			name: "overflow left 1 by 1",
			input: `
L49
L1
L1
`,
			expectedPwd: "1",
		},
		{
			name: "overflow right 1 by 1",
			input: `
R49
R1
R1
`,
			expectedPwd: "1",
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

			if !strings.Contains(string(output), "Safe password is "+tt.expectedPwd) {
				t.Errorf("bad password - want %q, entire output \n%s\n", tt.expectedPwd, string(output))
			}

		})
	}
}
