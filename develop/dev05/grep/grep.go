package grep

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Grep — struct to store all fields and methods for the grep functionality.
type Grep struct {
	pattern     string
	filename    string
	optionsInt  map[string]*int
	optionsBool map[string]*bool
	text        string
	result      []string
}

// New — constructor for the Grep struct.
func New(pattern, filename string, optionsInt map[string]*int, optionsBool map[string]*bool) *Grep {
	return &Grep{
		pattern:     pattern,
		filename:    filename,
		optionsInt:  optionsInt,
		optionsBool: optionsBool,
	}
}

// Start — main method that starts all the processes for the grep method.
func (g *Grep) Start() error {
	// Read the file.
	if err := g.readFile(); err != nil {
		return err
	}

	// Check options.
	if *g.optionsBool["c"] {
		g.count()
		return nil
	}

	// Check options.
	if *g.optionsBool["i"] {
		g.ignoreCase()
	}

	// Looking for matches.
	if err := g.search(); err != nil {
		return err
	}

	// Output the results.
	for _, l := range g.result {
		fmt.Println(l)
	}

	return nil
}

func (g *Grep) search() error {
	// Compile pattern.
	expr, err := regexp.Compile(g.pattern)
	if err != nil {
		return err
	}

	// Read int options A, B and C.
	before, after := *g.optionsInt["B"], *g.optionsInt["A"]
	if *g.optionsInt["C"] > 0 {
		before, after = *g.optionsInt["C"], *g.optionsInt["C"]
	}

	// Split text by lines.
	lines := strings.Split(g.text, "\n")

	// Iterate through lines and search for pattern match.
	for i, l := range lines {
		if expr.MatchString(l) && !(*g.optionsBool["F"]) ||
			strings.Contains(l, g.pattern) && *g.optionsBool["F"] {
			if *g.optionsBool["v"] {
				continue
			}

			temp := make([]string, 0, before)
			for j := i - 1; j+1 > i-before && j >= 0; j-- {
				if *g.optionsBool["n"] {
					temp = append(temp, fmt.Sprintf("%v:%v", j, lines[j]))
				} else {
					temp = append(temp, lines[j])
				}
			}

			for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
				temp[i], temp[j] = temp[j], temp[i]
			}

			g.result = append(g.result, temp...)

			if *g.optionsBool["n"] {
				g.result = append(g.result, fmt.Sprintf("%v:%v", i, l))
			} else {
				g.result = append(g.result, l)
			}

			for j := i + 1; j-i-1 < after && j < len(lines); j++ {
				if *g.optionsBool["n"] {
					g.result = append(g.result, fmt.Sprintf("%v:%v", j, lines[j]))
				} else {
					g.result = append(g.result, lines[j])
				}
			}
		}

		if *g.optionsBool["v"] {
			if *g.optionsBool["n"] {
				g.result = append(g.result, fmt.Sprintf("%v:%v", i, l))
			} else {
				g.result = append(g.result, l)
			}
		}
	}

	return nil
}

func (g *Grep) count() {
	lines := strings.Split(g.text, "\n")
	fmt.Printf("Total number of lines: %v\n", len(lines))
}

func (g *Grep) ignoreCase() {
	g.pattern = strings.ToLower(g.pattern)
	g.text = strings.ToLower(g.text)
}

func (g *Grep) readFile() error {
	content, err := os.ReadFile(g.filename)
	if err != nil {
		return err
	}

	g.text = string(content)

	return nil
}
