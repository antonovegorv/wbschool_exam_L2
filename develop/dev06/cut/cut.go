package cut

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

// Cut — struct to store all fields and methods for a cut util.
type Cut struct {
	filename string
	f        int
	d        string
	s        bool
}

// New — constructor for a Cut struct.
func New(filename string, f int, d string, s bool) *Cut {
	return &Cut{filename: filename, f: f, d: d, s: s}
}

// Run — main method to start cut util.
func (c *Cut) Run() error {
	// If we got the file name, then we read the file. Otherwise, we read stdin.
	if c.filename == "" {
		// Create a reader.
		reader := bufio.NewReader(os.Stdin)

		// Read line by line until ^C is pressed.
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			if res, ok := cut(line, c.f, c.d, c.s); ok {
				fmt.Println(strings.TrimSpace(res))
			}
		}
	} else {
		// Read file.
		content, err := os.ReadFile(c.filename)
		if err != nil {
			return err
		}

		lines := bytes.Split(content, []byte{'\n'})
		for _, l := range lines {
			if res, ok := cut(string(l), c.f, c.d, c.s); ok {
				fmt.Println(strings.TrimSpace(res))
			}
		}
	}

	return nil
}

func cut(str string, f int, d string, s bool) (string, bool) {
	// Check for the "-s" flag.
	if s {
		if !strings.Contains(str, d) {
			return "", false
		}
	}

	// Split string by delimiter.
	columns := strings.Split(str, d)
	if f < len(columns) {
		return columns[f], true
	}

	return "", false
}
