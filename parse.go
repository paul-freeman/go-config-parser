package parse

import "fmt"

// Config is a string representing a configuration.
//
// It will be parsed on creation to ensure validity.
type Config struct {
	str string
}

// New creates a new Config from a string.
//
// If the input string fails to parse, an error is returned.
func New(in string) (Config, error) {
	err := validate(in)
	if err != nil {
		return Config{}, err
	}
	return Config{str: in}, nil
}

// String returns the string representation of the configuration.
func (c Config) String() string {
	return c.str
}

// validate is a helper function to ensure the input string is valid.
//
// It will iterate over the string and put open bracket characters onto a stack.
// When closed brackets are encountered, the stack is checked to see if the
// expected bracket is on the top. If not, an error is returned.
//
// If the stack is empty at the end, the input string is valid.
func validate(in string) error {
	var stack []rune
	for _, r := range in {
		switch r {
		case '(', '[', '{':
			stack = append(stack, r)
		case ')', ']', '}':
			if len(stack) == 0 {
				return fmt.Errorf("unexpected %c", r)
			}
			expected := stack[len(stack)-1]
			if r == ')' && expected != '(' {
				return fmt.Errorf("expected %c, got %c", expected, r)
			}
			if r == ']' && expected != '[' {
				return fmt.Errorf("expected %c, got %c", expected, r)
			}
			if r == '}' && expected != '{' {
				return fmt.Errorf("expected %c, got %c", expected, r)
			}
			stack = stack[:len(stack)-1]
		default:
			// any other runes are ignored
			continue
		}
	}
	if len(stack) > 0 {
		return fmt.Errorf("unclosed %c", stack[len(stack)-1])
	}
	return nil
}
