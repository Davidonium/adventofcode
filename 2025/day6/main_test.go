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
        expected    string
	}{
        {
            name: "example",
            input: `
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +`,
            expected: "r=4277556",
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
				t.Errorf("FAILURE:\nwant:\n\t%s\ngot:\n\t%s\n", tt.expected, o)
            }
		})
	}
}

