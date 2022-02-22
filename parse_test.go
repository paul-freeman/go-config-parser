package parse_test

import (
	"fmt"
	"testing"

	parser "github.com/paul-freeman/go-config-parser"
)

func TestParser(t *testing.T) {
	tests := []struct {
		in          string
		shouldError bool
	}{
		// basic tests
		{"([)]", true},
		{"{[}", true},
		{"([])", false},
		{"([]{}())", false},
		{"(){}()", false},
		// extra tests
		{"", false},
		{"{", true},
		{"{]", true},
		{"{{{{{}}}}", true},
		{"{{{{}}}}}", true},
		{"(letters[can]be{inside}the(configuration))", false},
		{"(letters[can]be{inside}\n\nthe(configuration))", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("test %s", test.in), func(t *testing.T) {
			c, err := parser.New(test.in)
			if test.shouldError && err == nil {
				t.Fatalf("Expected error for %s", test.in)
			}
			if !test.shouldError && test.in != c.String() {
				t.Fatalf("Configuration lost during parsing: expected %s, got %s", test.in, c.String())
			}
		})
	}
}
